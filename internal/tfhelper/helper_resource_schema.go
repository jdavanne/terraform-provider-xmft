package tfhelper

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ModelToSchema(ctx context.Context, modelName string, model interface{}) schema.Schema {
	return schema.Schema{Attributes: _modelToAttributes(ctx, modelName, reflect.TypeOf(reflect.ValueOf(model).Elem().Interface()))}
}

func _modelToAttributes(ctx context.Context, modelName string, value reflect.Type) map[string]schema.Attribute {
	attrs := make(map[string]schema.Attribute)
	switch value.Kind() {
	case reflect.Struct:
		// v := value.Interface()
		// first := true
		// reflectType := reflect.TypeOf(v)
		reflectType := value
		// reflectValue := reflect.ValueOf(v)
		for i := 0; i < reflectType.NumField(); i++ {
			fieldName := reflectType.Field(i).Name
			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			name := FlagsTfsdkGetName(tfsdk)

			mustCheckSupportedAttributes(modelName+"."+fieldName, flags)
			required := FlagsHas(flags, "required")
			computed := FlagsHas(flags, "computed")
			sensitive := FlagsHas(flags, "sensitive")
			elementtype, _ := FlagsGet(flags, "elementtype")

			state := FlagsHas(flags, "state")
			def, defok := FlagsGet(flags, "default")
			var s []planmodifier.String = nil
			if state {
				s = []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				}
			}
			optional := (!required && !computed) || FlagsHas(flags, "optional") || defok
			computed = computed || defok

			// typ := reflectType.Field(i).Type.
			typestr := reflectType.Field(i).Type.String()
			kind := reflectType.Field(i).Type.Kind()
			// tflog.Debug(ctx, "ResourceModelToSchema >> "+modelName+"."+fieldName+" : "+typestr+"/"+kind.String())

			switch kind {
			case reflect.Slice:
				t := reflectType.Field(i).Type.Elem()
				typestr := t.String()
				kind2 := t.Kind()

				switch kind2 {
				case reflect.Struct:
					if strings.HasPrefix(typestr, "basetypes.") {
						switch typestr {
						case "basetypes.StringValue":
							var d defaults.List
							if defok && def == "" {
								f, _ := types.ListValue(types.StringType, make([]attr.Value, 0))
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + modelName + "." + fieldName + ")")
							}
							attrs[name] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(flags, "[]"),
							}
						default:
							panic("unsupported slice type: " + typestr + "(" + modelName + "." + fieldName + ")")
						}
					} else {
						var d defaults.List
						if defok && def == "" {
							f, _ := types.ListValue(types.ObjectType{}, make([]attr.Value, 0))
							d = listdefault.StaticValue(f)
						} else if defok {
							panic("unsupported default value: " + def + "(" + modelName + "." + fieldName + ")")
						}
						attrs[name] = schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: _modelToAttributes(ctx, modelName+"."+fieldName, reflectType.Field(i).Type.Elem()),
							},
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "[]"),
						}
					}
				default:
					panic("unsupported: Slice" + kind2.String() + " (" + modelName + "." + fieldName + ")")
				}

			case reflect.Struct:
				if strings.HasPrefix(typestr, "basetypes.") {
					switch typestr {
					case "basetypes.MapValue":
						if elementtype == "string" {
							attrs[name] = schema.MapAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(flags, "{}"),
							}
						} else {
							panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
						}
					case "basetypes.ObjectValue":
						elementModel := registeredTypes[elementtype]
						if elementModel == nil {
							panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
						}
						// el := reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface())
						// elementAttrs := _modelToAttributes(ctx, modelName+"."+fieldName, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))

						/*
							elementAttrs := structNameToTFType(elementtype)
							attrs[name] = schema.ObjectAttribute{
								AttributeTypes: elementAttrs,
								Required:       required,
								Optional:       optional,
								Computed:       computed,
								Sensitive:      sensitive,
							}*/

						elementAttrs2 := _modelToAttributes(ctx, modelName+"."+fieldName, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))
						attrs[name] = schema.SingleNestedAttribute{
							Attributes:  elementAttrs2,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(flags, "{}"),
						}

					case "basetypes.ListValue":
						if elementtype == "string" {
							var d defaults.List
							if defok && def == "" {
								f := types.ListNull(types.StringType)
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + modelName + "." + fieldName + ")")
							}
							attrs[name] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(flags, "[]"),
							}
						} else {
							elementModel := registeredTypes[elementtype]
							if elementModel == nil {
								panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
							}
							elementAttrs := _modelToAttributes(ctx, modelName+"."+fieldName, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))

							var d defaults.List
							if defok && def == "" {
								f := types.ListNull(types.ObjectType{})
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + modelName + "." + fieldName + ")")
							}
							attrs[name] = schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: elementAttrs,
								},
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(flags, "[]"),
							}
						}
					/*case "basetypes.ObjectValue":
					attrs[name] = schema.MapAttribute{
						ElementType: types.StringType,
						Required:    required,
						Optional:    optional,
						Computed:    computed,
						Sensitive:   sensitive,
					}*/
					case "basetypes.StringValue":
						var d defaults.String
						if defok {
							d = stringdefault.StaticString(def)
						}

						attrs[name] = schema.StringAttribute{
							Required:      required,
							Optional:      optional,
							Computed:      computed,
							Sensitive:     sensitive,
							Default:       d,
							Description:   flagsDescription(flags, "''"),
							PlanModifiers: s,
						}
					case "basetypes.BoolValue":
						var d defaults.Bool
						if defok && def != "" {
							if def == "true" {
								d = booldefault.StaticBool(true)
							} else if def == "false" {
								d = booldefault.StaticBool(false)
							} else {
								panic("unsupported default value: " + def + "(" + modelName + "." + fieldName + ")")
							}
						} else if defok && def == "" {
							d = booldefault.StaticBool(false)
						}
						attrs[name] = schema.BoolAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "false"),
						}
					case "basetypes.Int64Value":
						var d defaults.Int64
						if defok && def != "" {
							i, err := strconv.ParseInt(def, 10, 64)
							if err != nil {
								panic(err)
							}
							d = int64default.StaticInt64(i)
						} else if defok && def == "" {
							d = int64default.StaticInt64(0)
						}
						attrs[name] = schema.Int64Attribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "0"),
						}
					default:
						panic("unsupported type: " + typestr + "(" + modelName + "." + fieldName + ")")
					}
				} else {
					/*var d defaults.Object
					if defok && def != "" {
						panic("unsupported default value: " + def + "(" + modelName + "." + fieldName + ")")
					}
					if defok && def == "" {
						d = defaults.Object{}
					}*/

					attrs[name] = schema.SingleNestedAttribute{
						Attributes:  _modelToAttributes(ctx, modelName+"."+fieldName, reflectType.Field(i).Type),
						Required:    required,
						Optional:    optional,
						Computed:    computed,
						Sensitive:   sensitive,
						Description: flagsDescription(flags, "{}"),
						// Default:    d,
					}
				}
			case reflect.Ptr:
				attrs[name] = schema.SingleNestedAttribute{
					Attributes:  _modelToAttributes(ctx, modelName+"."+fieldName, reflectType.Field(i).Type.Elem()),
					Required:    required,
					Optional:    optional,
					Computed:    computed,
					Sensitive:   sensitive,
					Description: flagsDescription(flags, "{}"),
				}
			default:
				panic("unsupported: type" + kind.String() + " (" + modelName + "." + fieldName + ")")
			}
		}
	default:
		panic("unsupported: type" + value.Kind().String() + " (" + modelName + ")")
	}
	return attrs
}
