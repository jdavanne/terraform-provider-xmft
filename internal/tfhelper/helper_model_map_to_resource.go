package tfhelper

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AttributesToResource converts the attributes from the XCO MFT object to the corresponding fields in the Terraform resource.
// It takes the resource name, the attributes map, and the resource object as input.
// The function iterates over the attributes map and sets the corresponding field values in the resource object.
// If a field is not present in the attributes map, it is left unchanged in the resource object.
// The function does not return any value.
func AttributesToResource(ctx context.Context, modelName string, attrs map[string]interface{}, model interface{}) {
	tflog.Info(ctx, "AttributesToResource : "+fmt.Sprint(modelName, attrs))
	value := reflect.ValueOf(model).Elem()
	switch value.Kind() {
	case reflect.Struct:
		// v := value.Interface()
		// first := true
		// reflectType := reflect.TypeOf(v)
		reflectType := value.Type()
		// reflectValue := reflect.ValueOf(v)
		for i := 0; i < reflectType.NumField(); i++ {
			fieldName := reflectType.Field(i).Name
			// fieldType := reflectType.Field(i).Type
			fieldTypeStr := reflectType.Field(i).Type.String()
			fieldTypeKind := reflectType.Field(i).Type.Kind()

			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			name := FlagsHelperName(tfsdk, flags)
			// nowrite := FlagsHas(flags, "nowrite")
			// required := Has(flags, "required")
			noread := FlagsHas(flags, "noread")
			readmap := FlagsHas(flags, "readmap")

			elementtype, _ := FlagsGet(flags, "elementtype")
			// sensitive := Has(flags, "sensitive")
			if readmap {
				name = name[0 : len(name)-1]
			}
			attrValue, ok := attrs[name]
			tflog.Info(ctx, "AttributesToResource : "+fmt.Sprint(modelName, ".", name, "=", attrValue, " attrs=", attrs, "hasValue=", ok))
			if !noread && ok {
				if attrValue == nil {
					attrValue = ""
				}
				attrType := reflect.TypeOf(attrValue)
				// v := reflect.ValueOf(val)
				switch fieldTypeKind {
				case reflect.Slice:
					eltype := reflectType.Field(i).Type.Elem()
					eltypestr := eltype.String()
					elkind := eltype.Kind()
					switch elkind {
					case reflect.Struct:
						if strings.HasPrefix(eltypestr, "basetypes.") {
							switch eltypestr {
							case "basetypes.StringValue":
								switch v := attrValue.(type) {
								/*case string:
									attrs[name] = strings.Split(v, ",")
								case []string:
									attrs[name] = v
								*/
								case []interface{}:
									slice := make([]basetypes.StringValue, 0, len(v))
									for _, v := range v {
										v2 := v.(string)
										slice = append(slice, types.StringValue(v2))
									}
									value.Field(i).Set(reflect.ValueOf(slice))
								case []string:
									slice := make([]basetypes.StringValue, 0, len(v))
									for _, v := range v {
										v2 := v //.(string)
										slice = append(slice, types.StringValue(v2))
									}
									value.Field(i).Set(reflect.ValueOf(slice))
								default:
									panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
								}
							default:
								panic("unsupported slice type: " + eltypestr + "(" + modelName + "." + fieldName + ")")
							}
						} else {
							tflog.Info(ctx, "set resource field resource : "+fieldName+"=<slice>")
							val2 := attrValue.([]interface{})
							slice := reflect.MakeSlice(reflect.SliceOf(eltype), 0, len(val2))
							for j := 0; j < len(val2); j++ {
								if val3, ok := val2[j].(map[string]interface{}); ok {
									resource := reflect.New(eltype)
									AttributesToResource(ctx, modelName+"."+fieldName, val3, resource.Interface())
									slice = reflect.Append(slice, resource.Elem())

									value.Field(i).Set(slice)
								} else {
									panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
								}
							}
						}
					default:
						panic("unsupported: Slice" + elkind.String() + " (" + modelName + "." + fieldName + ")")
					}
				/*case reflect.Bool:
				val1 := attrValue.(bool)
				val2 := types.BoolValue(val1)
				tflog.Info(ctx, "set resource field resource : "+fieldName+"="+fmt.Sprint(val1))
				value.Field(i).Set(reflect.ValueOf(val2))*/
				case reflect.Struct:
					if strings.HasPrefix(fieldTypeStr, "basetypes.") {
						switch fieldTypeStr {
						case "basetypes.MapValue":
							if elementtype == "string" {
								switch v := attrValue.(type) {
								case map[string]interface{}:
									if len(v) != 0 {
										elements := make(map[string]attr.Value)
										for k, v := range v {
											elements[k] = types.StringValue(v.(string))
										}
										val1, diags := types.MapValue(types.StringType, elements)
										if diags.HasError() {
											panic("error: " + fmt.Sprint(diags))
										}
										value.Field(i).Set(reflect.ValueOf(val1))
									}
								default:
									panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
								}
							} else {
								panic("unsupported type: " + fmt.Sprint(elementtype) + "(" + modelName + "." + fieldName + ")")
							}
						case "basetypes.ObjectValue":

							elementModel := registeredTypes[elementtype]
							if elementModel == nil {
								panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
							}
							objTypes := structToTFType(modelName+"."+fieldName, elementModel)

							switch v := attrValue.(type) {
							case map[string]interface{}:
								elValueI := reflect.New(reflect.ValueOf(elementModel).Type().Elem())
								AttributesToResource(ctx, modelName+"."+fieldName, v, elValueI.Interface())

								obj, diags := types.ObjectValueFrom(ctx, objTypes, elValueI.Interface())
								if diags.HasError() {
									panic("error: " + fmt.Sprint(diags))
								}
								value.Field(i).Set(reflect.ValueOf(obj))
							default:
								panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
							}

						case "basetypes.ListValue":
							if elementtype == "string" {
								switch v := attrValue.(type) {
								case []string:
									elements := make([]attr.Value, len(v))
									for i, v := range v {
										elements[i] = types.StringValue(v)
									}
									val1, diags := types.ListValue(types.StringType, elements)
									if diags.HasError() {
										panic("error: " + fmt.Sprint(diags))
									}
									value.Field(i).Set(reflect.ValueOf(val1))
								default:
									// FIXME: should be error?
									panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
								}
							} else {
								elementModel := registeredTypes[elementtype]
								if elementModel == nil {
									panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
								}
								objTypes := structToTFType(modelName+"."+fieldName, elementModel)

								elements := make([]attr.Value, 0)
								switch v := attrValue.(type) {
								case []interface{}:
									for _, v := range v {
										switch v := v.(type) {
										case map[string]interface{}:
											elValueI := reflect.New(reflect.ValueOf(elementModel).Type().Elem())
											AttributesToResource(ctx, modelName+"."+fieldName, v, elValueI.Interface())

											obj, diags := types.ObjectValueFrom(ctx, objTypes, elValueI.Interface())
											if diags.HasError() {
												panic("error: " + fmt.Sprint(diags))
											}
											elements = append(elements, obj)
										default:
											// FIXME: should be error?
											panic("unsupported type for value: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
										}
									}
								default:
									// FIXME: should be error?
									panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
								}
								if len(elements) != 0 {
									val1, diags := types.ListValue(types.ObjectType{}.WithAttributeTypes(objTypes), elements)
									if diags.HasError() {
										panic("error: " + fmt.Sprint(diags))
									}

									value.Field(i).Set(reflect.ValueOf(val1))
								}
							}

						case "basetypes.StringValue":
							switch v := attrValue.(type) {
							case string:
								val1 := v
								val2 := types.StringValue(val1)
								tflog.Info(ctx, "set resource field resource : "+fieldName+"="+val1)
								value.Field(i).Set(reflect.ValueOf(val2))
							case []string:
								val1 := strings.Join(v, ",")
								val2 := types.StringValue(val1)
								value.Field(i).Set(reflect.ValueOf(val2))
							case []interface{}:
								arr := make([]string, len(v))
								for i, v2 := range v {
									arr[i] = v2.(string)
								}
								val1 := strings.Join(arr, ",")
								val2 := types.StringValue(val1)
								value.Field(i).Set(reflect.ValueOf(val2))
							default:
								panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ")")
							}
						case "basetypes.BoolValue":
							switch v := attrValue.(type) {
							case bool:
								val1 := v
								val2 := types.BoolValue(val1)
								tflog.Info(ctx, "set resource field resource : "+fieldName+"="+fmt.Sprint(val1))
								value.Field(i).Set(reflect.ValueOf(val2))
							case string:
								val1 := v
								if v == "" {
									val2 := types.BoolValue(false)
									tflog.Info(ctx, "set resource field resource : "+fieldName+"="+fmt.Sprint(val1))
									value.Field(i).Set(reflect.ValueOf(val2))
								} else {
									panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ") [" + fmt.Sprint(v) + "]")
								}
							default:
								panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ") [" + fmt.Sprint(v) + "]")
							}
						case "basetypes.Int64Value":
							switch v := attrValue.(type) {
							case float64:
								val1 := int64(v)
								val2 := types.Int64Value(val1)
								tflog.Info(ctx, "set resource field resource : "+fieldName+"="+fmt.Sprint(val1))
								value.Field(i).Set(reflect.ValueOf(val2))
							case int:
								val1 := v
								val2 := types.Int64Value(int64(val1))
								tflog.Info(ctx, "set resource field resource : "+fieldName+"="+fmt.Sprint(val1))
								value.Field(i).Set(reflect.ValueOf(val2))
							default:
								panic("unsupported type: " + fmt.Sprint(attrType) + "(" + modelName + "." + fieldName + ") [" + fmt.Sprint(v) + "]")
							}

						default:
							panic("unsupported type: " + fieldTypeStr + "(" + modelName + "." + fieldName + ")")
						}
					} else {
						// resource := reflect.New(fieldType)
						resource := value.Field(i).Addr()
						AttributesToResource(ctx, modelName+"."+fieldName, attrValue.(map[string]interface{}), resource.Interface())
						value.Field(i).Set(resource.Elem())
					}
				case reflect.Ptr:
					tflog.Info(ctx, "set resource field resource : "+fieldName+"=<ptr>")
					elValue := reflect.New(reflectType.Field(i).Type.Elem())
					if !value.Field(i).IsNil() {
						elValue = value.Field(i)
					}
					AttributesToResource(ctx, modelName+"."+fieldName, attrValue.(map[string]interface{}), elValue.Interface())
					value.Field(i).Set(elValue)

				default:
					panic("unsupported type: " + fieldTypeKind.String() + "(" + modelName + "." + fieldName + ")")
				}
			} else {
				tflog.Info(ctx, "skip resource field : "+fieldName+fmt.Sprint(noread, attrValue))
			}
		}
	default:
		panic("unsupported type: " + value.Kind().String() + "(" + modelName + ")")
	}
}
