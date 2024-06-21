package tfhelper

import (
	"context"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceModelToSchema(ctx context.Context, modelName string, model interface{}) schema.Schema {
	return schema.Schema{Attributes: _model2ToAttributes(ctx, modelName, reflect.TypeOf(reflect.ValueOf(model).Elem().Interface()))}
}

func _model2ToAttributes(ctx context.Context, modelName string, value reflect.Type) map[string]schema.Attribute {
	_modelToAttributes := _model2ToAttributes
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
			optional := (!required && !computed) || FlagsHas(flags, "optional")
			elementtype, _ := FlagsGet(flags, "elementtype")

			typestr := reflectType.Field(i).Type.String()
			kind := reflectType.Field(i).Type.Kind()
			// tflog.Debug(ctx, "dataSourceModelToSchema >> "+modelName+"."+fieldName+" : "+typestr+"/"+kind.String())

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
							attrs[name] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(flags, "[]"),
							}
						default:
							panic("unsupported slice type: " + typestr + "(" + modelName + "." + fieldName + ")")
						}
					} else {
						attrs[name] = schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: _modelToAttributes(ctx, modelName+"."+fieldName, reflectType.Field(i).Type.Elem()),
							},
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
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
						attrs[name] = schema.MapAttribute{
							ElementType: types.StringType,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(flags, "{}"),
						}
					case "basetypes.ObjectValue":
						elementModel := registeredTypes[elementtype]
						if elementModel == nil {
							panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
						}
						elementAttrs := _modelToAttributes(ctx, modelName+"."+fieldName, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))
						attrs[name] = schema.SingleNestedAttribute{
							Attributes:  elementAttrs,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(flags, "{}"),
						}
					case "basetypes.ListValue":
						if elementtype == "string" {
							attrs[name] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(flags, "[]"),
							}
						} else {
							elementModel := registeredTypes[elementtype]
							if elementModel == nil {
								panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
							}
							elementAttrs := _modelToAttributes(ctx, modelName+"."+fieldName, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))

							attrs[name] = schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: elementAttrs,
								},
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(flags, "[]"),
							}
						}
					case "basetypes.StringValue":
						attrs[name] = schema.StringAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(flags, "''"),
						}
					case "basetypes.BoolValue":
						attrs[name] = schema.BoolAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(flags, "false"),
						}
					case "basetypes.Int64Value":
						attrs[name] = schema.Int64Attribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(flags, "0"),
						}
					default:
						panic("unsupported type: " + typestr + "(" + modelName + "." + fieldName + ")")
					}
				} else {
					attrs[name] = schema.SingleNestedAttribute{
						Attributes:  _modelToAttributes(ctx, modelName+"."+fieldName, reflectType.Field(i).Type),
						Required:    required,
						Optional:    optional,
						Computed:    computed,
						Sensitive:   sensitive,
						Description: flagsDescription(flags, "{}"),
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
	}
	return attrs
}
