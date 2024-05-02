package tfhelper

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResourceToAttributes converts the fields in the Terraform resource to the corresponding attributes in the XCO MFT object.
// It takes the resource name, the resource object, and an empty attributes map as input.
// The function iterates over the fields in the resource object and sets the corresponding attributes in the attributes map.
// The function does not return any value.
func ResourceToAttributes(ctx context.Context, modelName string, model interface{}, attrs map[string]interface{}) {
	value := reflect.ValueOf(model).Elem()
	_resourceValueToAttributes(ctx, modelName, value, attrs)
}

func _resourceValueToAttributes(ctx context.Context, modelName string, value reflect.Value, attrs map[string]interface{}) {
	switch value.Kind() {
	case reflect.Struct:
		// v := value.Interface()
		// first := true
		reflectType := value.Type()
		reflectValue := value
		for i := 0; i < reflectType.NumField(); i++ {
			fieldName := reflectType.Field(i).Name
			// fieldType := reflectType.Field(i).Type
			fieldTypeStr := reflectType.Field(i).Type.String()
			fieldTypeKind := reflectType.Field(i).Type.Kind()

			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			name := FlagsHelperName(tfsdk, flags)
			nowrite := FlagsHas(flags, "nowrite")
			// required := Has(flags, "required")
			// noread := FlagsHas(flags, "noread")
			// readmap := FlagsHas(flags, "readmap")

			elementtype, _ := FlagsGet(flags, "elementtype")
			fieldValue := reflectValue.Field(i)
			fieldValKind := fieldValue.Kind()
			tflog.Info(ctx, ">>"+fieldName+" : "+fieldTypeStr+"/"+fieldTypeKind.String())

			switch fieldValKind {
			case reflect.Slice:
				eltype := reflectType.Field(i).Type.Elem()
				eltypestr := eltype.String()
				elkind := eltype.Kind()
				switch elkind {
				case reflect.Struct:
					if strings.HasPrefix(eltypestr, "basetypes.") {
						switch eltypestr {
						case "basetypes.StringValue":
							slice := make([]interface{}, 0)
							for i := 0; i < fieldValue.Len(); i++ {
								slice = append(slice, fieldValue.Index(i).Interface().(basetypes.StringValue).ValueString())
							}
							attrs[name] = slice
						default:
							panic("unsupported slice type: " + eltypestr + "(" + modelName + "." + fieldName + ")")
						}
					} else {
						slice := make([]interface{}, 0)
						for i := 0; i < fieldValue.Len(); i++ {
							attrs2 := make(map[string]interface{})
							_resourceValueToAttributes(ctx, modelName+"."+fieldName, fieldValue.Index(i), attrs2)
							slice = append(slice, attrs2)
						}
						attrs[name] = slice
					}
				default:
					panic("unsupported: Slice" + elkind.String() + " (" + modelName + "." + fieldName + ")")
				}
			case reflect.Struct:
				if strings.HasPrefix(fieldTypeStr, "basetypes.") {
					val := fieldValue.Interface()
					fieldValueTypeStr := fieldValue.Type().String()
					switch fieldValueTypeStr {
					case "basetypes.MapValue":
						bval := val.(basetypes.MapValue)
						if elementtype == "string" {
							elements := make(map[string]types.String, len(bval.Elements()))
							v2 := bval.ElementsAs(ctx, &elements, false)
							if v2.HasError() {
								panic("error: " + fmt.Sprint(v2))
							}
							attrs2 := make(map[string]interface{})
							for k, v := range elements {
								attrs2[k] = v.ValueString()
							}
							attrs[name] = attrs2
						} else {
							panic("unsupported type: " + elementtype + "(" + modelName + "." + fieldName + ")")
						}
					case "basetypes.ObjectValue":
						bval := val.(basetypes.ObjectValue)
						val := newRegisteredType(elementtype)
						diags := bval.As(ctx, val, basetypes.ObjectAsOptions{})
						if diags.HasError() {
							panic("error: " + fmt.Sprint(diags))
						}
						attrs2 := make(map[string]interface{})
						if !bval.IsNull() {
							_resourceValueToAttributes(ctx, modelName+"."+fieldName, reflect.ValueOf(val).Elem().Elem(), attrs2)
						}
						attrs[name] = attrs2

					case "basetypes.ListValue":
						bval := val.(basetypes.ListValue)
						if elementtype == "string" {
							elements := make([]types.String, len(bval.Elements()))
							v2 := bval.ElementsAs(ctx, &elements, false)
							if v2.HasError() {
								panic("error: " + fmt.Sprint(v2))
							}
							attrs2 := make([]interface{}, 0, len(elements))
							for _, v := range elements {
								attrs2 = append(attrs2, v.ValueString())
							}
							attrs[name] = attrs2
						} else {
							list := make([]interface{}, len(bval.Elements()))
							for i, v := range bval.Elements() {
								attrs2 := make(map[string]interface{})
								t := v.Type(ctx)

								switch v := v.(type) {
								case types.Object:
									val := newRegisteredType(elementtype)
									diags := v.As(ctx, val, basetypes.ObjectAsOptions{})
									if diags.HasError() {
										panic("error: " + fmt.Sprint(diags))
									}

									_resourceValueToAttributes(ctx, modelName+"."+fieldName, reflect.ValueOf(val).Elem().Elem(), attrs2)
								default:
									panic("unsupported type: " + t.String() + "(" + modelName + "." + fieldName + ")")
								}
								list[i] = attrs2
							}
							attrs[name] = list

							// panic("unsupported type: " + elementtype + "(" + modelName + "." + fieldName + ")")
						}
					case "basetypes.StringValue":
						bval := val.(basetypes.StringValue)
						val := bval.ValueString()
						if val != "" && !nowrite {
							attrs[name] = val
						}
					case "basetypes.BoolValue":
						bval := val.(basetypes.BoolValue)
						val := bval.ValueBool()
						if !nowrite {
							attrs[name] = val
						}
					case "basetypes.Int64Value":
						bval := val.(basetypes.Int64Value)
						val := bval.ValueInt64()
						if !nowrite {
							attrs[name] = val
						}
					default:
						panic("unsupported type: " + fieldTypeStr + "(" + modelName + "." + fieldName + ")")
					}
				} else {
					attrs2 := make(map[string]interface{})
					_resourceValueToAttributes(ctx, modelName+"."+fieldName, fieldValue, attrs2)
					attrs[name] = attrs2
				}
			case reflect.Ptr:
				if !fieldValue.IsNil() {
					attrs2 := make(map[string]interface{})
					_resourceValueToAttributes(ctx, modelName+"."+fieldName, fieldValue.Elem(), attrs2)
					attrs[name] = attrs2
				}
			default:
				panic("unsupported type: " + fieldValKind.String() + "(" + modelName + "." + fieldName + ")")
			}
		}

	default:
		panic("unsupported type: " + value.Kind().String() + "(" + modelName + ")")
	}
}
