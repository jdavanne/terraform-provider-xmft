package tfhelper

import (
	"context"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DataSourceModelToSchema(ctx context.Context, goPath string, apiPath string, model interface{}) schema.Schema {
	return schema.Schema{Attributes: _model2ToAttributes(ctx, goPath, apiPath, false, reflect.TypeOf(reflect.ValueOf(model).Elem().Interface()))}
}

func _model2ToAttributes(ctx context.Context, goPath string, apiPath string, parentFold bool, value reflect.Type) map[string]schema.Attribute {
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
			goName := reflectType.Field(i).Name
			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			tfName := FlagsTfsdkGetName(tfsdk)
			apiName := FlagsHelperName(tfsdk, flags)

			mustCheckSupportedAttributes(goPath+"."+goName, flags)
			required := FlagsHas(flags, "required")
			computed := FlagsHas(flags, "computed")
			sensitive := FlagsHas(flags, "sensitive")
			optional := (!required && !computed) || FlagsHas(flags, "optional")
			elementtype, _ := FlagsGet(flags, "elementtype")
			_, fold := FlagsGet(flags, "fold")
			if parentFold {
				apiName = "[type=" + apiName + "]"
			}

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
							attrs[tfName] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							}
						default:
							panic("unsupported slice type: " + typestr + "(" + goPath + "." + goName + ")")
						}
					} else {
						attrs[tfName] = schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName+".#", fold, reflectType.Field(i).Type.Elem()),
							},
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
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
						attrs[tfName] = schema.MapAttribute{
							ElementType: types.StringType,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
						}
					case "basetypes.ObjectValue":
						elementModel := registeredTypes[elementtype]
						if elementModel == nil {
							panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
						}
						elementAttrs := _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))
						attrs[tfName] = schema.SingleNestedAttribute{
							Attributes:  elementAttrs,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
						}
					case "basetypes.ListValue":
						if elementtype == "string" {
							attrs[tfName] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							}
						} else {
							elementModel := registeredTypes[elementtype]
							if elementModel == nil {
								panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
							}
							elementAttrs := _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName+".#", fold, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))

							attrs[tfName] = schema.ListNestedAttribute{
								NestedObject: schema.NestedAttributeObject{
									Attributes: elementAttrs,
								},
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							}
						}
					case "basetypes.StringValue":
						attrs[tfName] = schema.StringAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(ctx, flags, "''", apiPath+"."+apiName),
						}
					case "basetypes.BoolValue":
						attrs[tfName] = schema.BoolAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(ctx, flags, "false", apiPath+"."+apiName),
						}
					case "basetypes.Int64Value":
						attrs[tfName] = schema.Int64Attribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Description: flagsDescription(ctx, flags, "0", apiPath+"."+apiName),
						}
					default:
						panic("unsupported type: " + typestr + "(" + goPath + "." + goName + ")")
					}
				} else {
					attrs[tfName] = schema.SingleNestedAttribute{
						Attributes:  _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflectType.Field(i).Type),
						Required:    required,
						Optional:    optional,
						Computed:    computed,
						Sensitive:   sensitive,
						Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
					}
				}
			case reflect.Ptr:
				attrs[tfName] = schema.SingleNestedAttribute{
					Attributes:  _modelToAttributes(ctx, goPath+"."+goName, apiPath+"."+apiName, fold, reflectType.Field(i).Type.Elem()),
					Required:    required,
					Optional:    optional,
					Computed:    computed,
					Sensitive:   sensitive,
					Description: flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
				}
			default:
				panic("unsupported: type" + kind.String() + " (" + goPath + "." + goName + ")")
			}
		}
	}
	return attrs
}
