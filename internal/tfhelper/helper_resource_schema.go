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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// tru to use NestedBlock instead of NestedAttribute : but fail on default :
const useBlock = false

func ModelToSchema(ctx context.Context, goPath string, apiPath string, model interface{}) schema.Schema {
	attrs, blocks := _modelToAttributes(ctx, goPath, apiPath, false, reflect.TypeOf(reflect.ValueOf(model).Elem().Interface()), useBlock)
	return schema.Schema{Attributes: attrs, Blocks: blocks}
}

func _modelToAttributesWrap(ctx context.Context, goPath string, apiPath string, parentFold bool, value reflect.Type) map[string]schema.Attribute {
	attrs, _ := _modelToAttributes(ctx, goPath, apiPath, parentFold, value, false)
	return attrs
}

func _modelToAttributes(ctx context.Context, goPath string, apiPath string, parentFold bool, value reflect.Type, useBlock bool) (map[string]schema.Attribute, map[string]schema.Block) {
	attrs := make(map[string]schema.Attribute)
	var blocks map[string]schema.Block
	if useBlock {
		blocks = make(map[string]schema.Block)
	}
	switch value.Kind() {
	case reflect.Struct:
		// v := value.Interface()
		// first := true
		// reflectType := reflect.TypeOf(v)
		reflectType := value
		// reflectValue := reflect.ValueOf(v)
		for i := 0; i < reflectType.NumField(); i++ {
			goName := reflectType.Field(i).Name
			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			tfName := FlagsTfsdkGetName(tfsdk)
			apiName := FlagsHelperName(tfsdk, flags)

			if tfName == "" {
				panic("tfname is empty" + goName + "/" + apiName + "/" + goPath)
			}

			mustCheckSupportedAttributes(goPath+"."+goName, flags)
			required := FlagsHas(flags, "required")
			computed := FlagsHas(flags, "computed")
			sensitive := FlagsHas(flags, "sensitive")
			elementtype, _ := FlagsGet(flags, "elementtype")
			_, fold := FlagsGet(flags, "fold")
			if parentFold {
				apiName = "[type=" + apiName + "]"
			}

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
								panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
							}
							attrs[tfName] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							}
						case "basetypes.Int64Value":
							var d defaults.List
							if defok && def == "" {
								f := types.ListNull(types.Int64Type)
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
							}
							attrs[tfName] = schema.ListAttribute{
								ElementType: types.Int64Type,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName),
							}
						default:
							panic("unsupported slice type: " + typestr + "(" + goPath + "." + goName + ")")
						}
					} else {
						var d defaults.List
						if defok && def == "" {
							f, _ := types.ListValue(types.ObjectType{}, make([]attr.Value, 0))
							d = listdefault.StaticValue(f)
						} else if defok {
							panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
						}
						attrs[tfName] = schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: _modelToAttributesWrap(ctx, goPath+"."+goName, apiPath+"."+apiName+".#", fold, reflectType.Field(i).Type.Elem()),
							},
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName),
						}
					}
				default:
					panic("unsupported: Slice" + kind2.String() + " (" + goPath + "." + goName + ")")
				}

			case reflect.Struct:
				if strings.HasPrefix(typestr, "basetypes.") {
					switch typestr {
					case "basetypes.MapValue":
						if elementtype == "string" {
							attrs[tfName] = schema.MapAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
							}
						} else {
							panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
						}
					case "basetypes.ObjectValue":
						/*if true {
							panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
						}*/
						elementModel := registeredTypes[elementtype]
						if elementModel == nil {
							panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
						}

						var d defaults.Object
						if defok && def == "" {
							// var v types.Object
							v, _ := structToObjectValue(elementtype, elementModel)
							d = objectdefault.StaticValue(v)
						}

						elementAttrs := _modelToAttributesWrap(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))
						attrs[tfName] = schema.SingleNestedAttribute{
							Attributes:  elementAttrs,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
						}

					case "basetypes.ListValue":
						if elementtype == "string" {
							var d defaults.List
							if defok && def == "" {
								f, _ := types.ListValueFrom(ctx, types.StringType, []string{})
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
							}
							attrs[tfName] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							}
						} else if elementtype == "int" {
							var d defaults.List
							if defok && def == "" {
								f, _ := types.ListValueFrom(ctx, types.Int64Type, []string{})
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
							}
							attrs[tfName] = schema.ListAttribute{
								ElementType: types.Int64Type,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							}
						} else {
							elementModel := registeredTypes[elementtype]
							if elementModel == nil {
								panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
							}
							elementAttrs := _modelToAttributesWrap(ctx, goPath+"."+goName, apiPath+"."+apiName+".#", fold, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))

							var d defaults.List
							if defok && def == "" {
								f := types.ListNull(types.ObjectType{})
								d = listdefault.StaticValue(f)
							} else if defok {
								panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
							}
							attrs[tfName] = schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: elementAttrs,
								},
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
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

						attrs[tfName] = schema.StringAttribute{
							Required:      required,
							Optional:      optional,
							Computed:      computed,
							Sensitive:     sensitive,
							Default:       d,
							Description:   flagsDescription(ctx, flags, "''", apiPath+"."+apiName),
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
								panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
							}
						} else if defok && def == "" {
							d = booldefault.StaticBool(false)
						}
						attrs[tfName] = schema.BoolAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(ctx, flags, "false", apiPath+"."+apiName),
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
						attrs[tfName] = schema.Int64Attribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(ctx, flags, "0", apiPath+"."+apiName),
						}
					default:
						panic("unsupported type: " + typestr + "(" + goPath + "." + goName + ")")
					}
				} else {

					if defok && def != "" {
						panic("unsupported default value: " + def + "(" + goPath + "." + goName + ")")
					}

					if !useBlock {
						var d defaults.Object
						var defaultVal basetypes.ObjectValue
						if defok && def == "" {
							defaultVal, _ = structToObjectValue(goPath+"."+goName, reflect.New(reflectType.Field(i).Type).Interface())
							d = objectdefault.StaticValue(defaultVal)
						}

						attrs[tfName] = schema.SingleNestedAttribute{
							Attributes:  _modelToAttributesWrap(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflectType.Field(i).Type),
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
							Default:     d,
							// Validators:  []validator.Object{onlyFields(0, 1024)},
							// PlanModifiers: []planmodifier.Object{useDefaultObjectModifier{defaultVal}},
						}
					} else {
						var d basetypes.ObjectValue
						if defok && def == "" {
							d, _ = structToObjectValue(goPath+"."+goName, reflect.New(reflectType.Field(i).Type).Interface())
							// d = objectdefault.StaticValue(defaultVal)
						}

						a2, b2 := _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflectType.Field(i).Type, false)
						blocks[tfName] = schema.SingleNestedBlock{
							Attributes:  a2,
							Blocks:      b2,
							Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
							// Validators:  []validator.Object{onlyFields(0, 1024)},
							PlanModifiers: []planmodifier.Object{useDefaultObjectModifier{d}},
						}
					}
				}
			case reflect.Ptr:
				if !useBlock {
					attrs[tfName] = schema.SingleNestedAttribute{
						Attributes:  _modelToAttributesWrap(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflectType.Field(i).Type.Elem()),
						Required:    required,
						Optional:    optional,
						Computed:    computed,
						Sensitive:   sensitive,
						Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
					}
				} else {
					a2, b2 := _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflectType.Field(i).Type.Elem(), false)
					blocks[tfName] = schema.SingleNestedBlock{
						Attributes:  a2,
						Blocks:      b2,
						Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
					}
				}
			default:
				panic("unsupported: type" + kind.String() + " (" + goPath + "." + goName + ")")
			}
		}
	default:
		panic("unsupported: type" + value.Kind().String() + " (" + goPath + ")")
	}
	return attrs, blocks
}
