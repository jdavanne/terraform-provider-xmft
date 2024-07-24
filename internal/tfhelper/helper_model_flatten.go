package tfhelper

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

func ModelFlatten(ctx context.Context, goPath string, tfPath string, apiPath string, parentFold bool, value reflect.Type) []map[string]string {
	fields := make([]map[string]string, 0)
	// attrs := make(map[string]schema.Attribute)
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
			tfname := FlagsTfsdkGetName(tfsdk)
			apiName := FlagsHelperName(tfsdk, flags)
			mustCheckSupportedAttributes(goPath+"."+goName, flags)
			// required := FlagsHas(flags, "required")
			computed := FlagsHas(flags, "computed")
			// sensitive := FlagsHas(flags, "sensitive")
			elementtype, _ := FlagsGet(flags, "elementtype")
			_, fold := FlagsGet(flags, "fold")
			if parentFold {
				apiName = "[type=" + apiName + "]"
			}
			// state := FlagsHas(flags, "state")
			def, defok := FlagsGet(flags, "default")

			// optional := (!required && !computed) || FlagsHas(flags, "optional") || defok
			computed = computed || defok

			// typ := reflectType.Field(i).Type.
			typestr := reflectType.Field(i).Type.String()
			kind := reflectType.Field(i).Type.Kind()

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
							fields = append(fields, map[string]string{
								"goPath":  goPath + "." + goName,
								"tfPath":  tfPath + "." + tfname,
								"apiPath": apiPath + "." + apiName + ".#",
								"apiName": apiName,
								"tfname":  tfname,
								"goName":  goName,

								"type":        "array[string]",
								"defaultOk":   fmt.Sprint(defok),
								"default":     def,
								"description": flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							})

							/*attrs[name] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(flags, "[]"),
							}*/
						case "basetypes.Int64Value":
							fields = append(fields, map[string]string{
								"goPath":  goPath + "." + goName,
								"tfPath":  tfPath + "." + tfname,
								"apiPath": apiPath + "." + apiName + ".#",
								"apiName": apiName,
								"tfname":  tfname,
								"goName":  goName,

								"type":        "array[int]",
								"defaultOk":   fmt.Sprint(defok),
								"default":     def,
								"description": flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							})

							/*attrs[name] = schema.ListAttribute{
								ElementType: types.Int64Type,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(flags, "[]"),
							}*/
						default:
							panic("unsupported slice type: " + typestr + "(" + goPath + "." + goName + ")")
						}
					} else {
						/*attrs[name] = schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: _modelFlatten(ctx, modelName+"."+fieldName, reflectType.Field(i).Type.Elem()),
							},
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "[]"),
						}*/
						fields = append(fields, ModelFlatten(ctx, goPath+"."+goName, tfPath+"."+tfname, apiPath+"."+apiName+".#", fold, reflectType.Field(i).Type.Elem())...)
					}
				default:
					panic("unsupported: Slice" + kind2.String() + " (" + goPath + "." + goName + ")")
				}

			case reflect.Struct:
				if strings.HasPrefix(typestr, "basetypes.") {
					switch typestr {
					case "basetypes.MapValue":
						if elementtype == "string" {
							/*attrs[name] = schema.MapAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Description: flagsDescription(flags, "{}"),
							}*/
							fields = append(fields, map[string]string{
								"goPath":  goPath + "." + goName,
								"tfPath":  tfPath + "." + tfname,
								"apiPath": apiPath + "." + apiName,
								"apiName": apiName,
								"tfname":  tfname,
								"goName":  goName,

								"type":        "map[string]string",
								"default":     def,
								"defaultOk":   fmt.Sprint(defok),
								"description": flagsDescription(ctx, flags, "{}", apiPath+"."+apiName),
							})
						} else {
							panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
						}
					case "basetypes.ObjectValue":
						elementModel := registeredTypes[elementtype]
						if elementModel == nil {
							panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
						}

						fields = append(fields, ModelFlatten(ctx, goPath+"."+goName, tfPath+"."+tfname, apiPath+"."+apiName, fold, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))...)
						/*attrs[name] = schema.SingleNestedAttribute{
							Attributes:  elementAttrs2,
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "{}"),
						}*/

					case "basetypes.ListValue":
						if elementtype == "string" {
							/*attrs[name] = schema.ListAttribute{
								ElementType: types.StringType,
								Required:    required,
								Optional:    optional,
								Computed:    computed,
								Sensitive:   sensitive,
								Default:     d,
								Description: flagsDescription(flags, "[]"),
							}*/
							fields = append(fields, map[string]string{
								"goPath":  goPath + "." + goName,
								"tfPath":  tfPath + "." + tfname,
								"apiPath": apiPath + "." + apiName + ".#",
								"apiName": apiName,
								"tfname":  tfname,
								"goName":  goName,

								"type":        "array[string]",
								"default":     def,
								"defaultOk":   fmt.Sprint(defok),
								"description": flagsDescription(ctx, flags, "[]", apiPath+"."+apiName+".#"),
							})
						} else {
							elementModel := registeredTypes[elementtype]
							if elementModel == nil {
								panic("unsupported element type: '" + elementtype + "' (" + goPath + "." + goName + ")")
							}
							fields = append(fields, ModelFlatten(ctx, goPath+"."+goName, tfPath+"."+tfname, apiPath+"."+apiName+".#", fold, reflect.TypeOf(reflect.ValueOf(elementModel).Elem().Interface()))...)
						}
					case "basetypes.StringValue":
						/*attrs[name] = schema.StringAttribute{
							Required:      required,
							Optional:      optional,
							Computed:      computed,
							Sensitive:     sensitive,
							Default:       d,
							Description:   flagsDescription(flags, "''"),
							PlanModifiers: s,
						}*/
						fields = append(fields, map[string]string{
							"goPath":  goPath + "." + goName,
							"tfPath":  tfPath + "." + tfname,
							"apiPath": apiPath + "." + apiName,
							"apiName": apiName,
							"tfname":  tfname,
							"goName":  goName,

							"type":        "string",
							"default":     def,
							"defaultOk":   fmt.Sprint(defok),
							"description": flagsDescription(ctx, flags, "''", apiPath+"."+apiName),
						})
					case "basetypes.BoolValue":
						/*attrs[name] = schema.BoolAttribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "false"),
						}*/
						fields = append(fields, map[string]string{
							"goPath":  goPath + "." + goName,
							"tfPath":  tfPath + "." + tfname,
							"apiPath": apiPath + "." + apiName,
							"apiName": apiName,
							"tfname":  tfname,
							"goName":  goName,

							"type":        "boolean",
							"default":     def,
							"defaultOk":   fmt.Sprint(defok),
							"description": flagsDescription(ctx, flags, "false", apiPath+"."+apiName),
						})
					case "basetypes.Int64Value":
						/*attrs[name] = schema.Int64Attribute{
							Required:    required,
							Optional:    optional,
							Computed:    computed,
							Sensitive:   sensitive,
							Default:     d,
							Description: flagsDescription(flags, "0"),
						}*/
						fields = append(fields, map[string]string{
							"goPath":  goPath + "." + goName,
							"tfPath":  tfPath + "." + tfname,
							"apiPath": apiPath + "." + apiName,
							"apiName": apiName,
							"tfname":  tfname,
							"goName":  goName,

							"type":        "int",
							"default":     def,
							"defaultOk":   fmt.Sprint(defok),
							"description": flagsDescription(ctx, flags, "0", apiPath+"."+apiName),
						})
					default:
						panic("unsupported type: " + typestr + "(" + goPath + "." + goName + ")")
					}
				} else {
					/*attrs[name] = schema.SingleNestedAttribute{
						Attributes:  _modelFlatten(ctx, modelName+"."+fieldName, reflectType.Field(i).Type),
						Required:    required,
						Optional:    optional,
						Computed:    computed,
						Sensitive:   sensitive,
						Description: flagsDescription(flags, "{}"),
						// Default:    d,
					}*/
					fields = append(fields, ModelFlatten(ctx, goPath+"."+goName, tfPath+"."+tfname, apiPath+"."+apiName, fold, reflectType.Field(i).Type)...)
				}
			case reflect.Ptr:

				/*attrs[name] = schema.SingleNestedAttribute{
					Attributes:  _modelFlatten(ctx, modelName+"."+fieldName, reflectType.Field(i).Type.Elem()),
					Required:    required,
					Optional:    optional,
					Computed:    computed,
					Sensitive:   sensitive,
					Description: flagsDescription(flags, "{}"),
				}*/
				fields = append(fields, ModelFlatten(ctx, goPath+"."+goName, tfPath+"."+tfname, apiPath+"."+apiName, fold, reflectType.Field(i).Type.Elem())...)
			default:
				panic("unsupported: type" + kind.String() + " (" + goPath + "." + goName + ")")
			}
		}
	default:
		panic("unsupported: type" + value.Kind().String() + " (" + goPath + ")")
	}
	return fields
}
