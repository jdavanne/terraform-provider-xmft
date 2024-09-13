package tfhelper

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var registeredTypes = make(map[string]interface{})

func RegisterType(name string, model interface{}) {
	registeredTypes[name] = model
}

func newRegisteredType(name string) interface{} {
	return reflect.New(reflect.ValueOf(registeredTypes[name]).Type()).Interface()
}

func structNameToTFType(modelName string) map[string]attr.Type {
	model := registeredTypes[modelName]
	return _structToTFType(modelName, model)
}

func _structToTFType(modelName string, model interface{}) map[string]attr.Type {
	attrs := make(map[string]attr.Type)
	value := reflect.ValueOf(model).Elem()
	switch value.Kind() {
	case reflect.Struct:
		reflectType := value.Type()
		for i := 0; i < reflectType.NumField(); i++ {
			fieldName := reflectType.Field(i).Name
			fieldTypeStr := reflectType.Field(i).Type.String()
			fieldTypeKind := reflectType.Field(i).Type.Kind()
			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			name := FlagsTfsdkGetName(tfsdk)
			elementtype, _ := FlagsGet(flags, "elementtype")
			// name := FlagsHelperName(tfsdk, flags)

			switch fieldTypeStr {
			case "basetypes.StringValue":
				attrs[name] = types.StringType
			case "basetypes.Int64Value":
				attrs[name] = types.Int64Type
			case "basetypes.Float64Value":
				attrs[name] = types.Float64Type
			case "basetypes.BoolValue":
				attrs[name] = types.BoolType
			case "basetypes.ObjectValue":
				t, ok := registeredTypes[elementtype]
				if !ok {
					panic("unsupported type for: " + modelName + "." + fieldName + ":" + elementtype)
				}

				attrs[name] = types.ObjectType{
					AttrTypes: _structToTFType(modelName+"."+fieldName, t),
				}

			default:
				if fieldTypeKind == reflect.Struct {
					if t, ok := registeredTypes[fieldTypeStr]; ok {
						attrs[name] = types.ObjectType{
							AttrTypes: _structToTFType(modelName+"."+fieldName, t),
						}
					} else {
						panic("unsupported type for: " + modelName + "." + fieldName + ":" + fieldTypeStr)
					}
				} else {
					panic("unsupported type for: " + modelName + "." + fieldName + ":" + fieldTypeStr)
				}
			}
		}
	default:
		panic("not a struct")
	}
	return attrs
}

func structToObjectValue(modelName string, model interface{}) (types.Object, map[string]attr.Type) {
	buildDef := true
	attrs := make(map[string]attr.Value)
	typs := make(map[string]attr.Type)

	value := reflect.ValueOf(model).Elem()
	switch value.Kind() {
	case reflect.Struct:
		reflectType := value.Type()
		for i := 0; i < reflectType.NumField(); i++ {
			fieldName := reflectType.Field(i).Name
			fieldTypeStr := reflectType.Field(i).Type.String()
			fieldTypeKind := reflectType.Field(i).Type.Kind()
			fieldValue := value.Field(i)
			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			flags := tag.Get(helperTagName)
			name := FlagsTfsdkGetName(tfsdk)
			elementtype, _ := FlagsGet(flags, "elementtype")
			optional := FlagsHas(flags, "optional")
			// name := FlagsHelperName(tfsdk, flags)
			def, _ := FlagsGet(flags, "default")

			// fmt.Println(modelName, name, fieldTypeStr, fieldTypeKind)
			if strings.ToUpper(fieldName[0:1]) != fieldName[0:1] {
				panic("unsupported field name: " + modelName + "." + fieldName)
			}

			switch v := fieldValue.Interface().(type) {
			case basetypes.StringValue:
				attrs[name] = v
				if buildDef {
					if optional {
						attrs[name] = types.StringNull()
					} else {
						attrs[name] = types.StringValue(def)
					}
				}
				typs[name] = types.StringType
			case basetypes.Int64Value:
				attrs[name] = v
				if buildDef {
					if optional {
						attrs[name] = types.Int64Null()
					} else {
						v2, _ := strconv.ParseInt(def, 10, 64)
						attrs[name] = types.Int64Value(v2)
					}
				}
				typs[name] = types.Int64Type
			case basetypes.Float64Value:
				attrs[name] = v
				if buildDef {
					if optional {
						attrs[name] = types.Float64Null()
					} else {
						v2, _ := strconv.ParseFloat(def, 64)
						attrs[name] = types.Float64Value(v2)
					}
				}
				typs[name] = types.Float64Type
			case basetypes.BoolValue:
				attrs[name] = v
				if buildDef {
					if optional {
						attrs[name] = types.BoolNull()
					} else {
						v2, _ := strconv.ParseBool(def)
						attrs[name] = types.BoolValue(v2)
					}
				}
				typs[name] = types.BoolType
			/*case basetypes.ListValue:
			attrs[name] = v
			typs[name] = types.ListType{}.WithElementType(types.ObjectType{
				AttrTypes: _structToTFType(modelName+"."+fieldName, fieldValue.Interface().([]interface{})[0]),
			})*/
			case basetypes.ObjectValue:
				attrs[name] = v // structToObjectValue(modelName+"."+fieldName, fieldValue.Interface())
				typs[name] = types.ObjectType{}.WithAttributeTypes(v.AttributeTypes(nil))
			case basetypes.ListValue:
				panic("unsupported ListValue:" + modelName + "." + fieldName)
			case basetypes.MapValue:
				attrs[name] = v
				if elementtype == "string" {
					typs[name] = types.MapType{}.WithElementType(types.StringType)
				} else {
					panic("unsupported element type: '" + elementtype + "' (" + modelName + "." + fieldName + ")")
				}
			case basetypes.SetValue:
				panic("unsupported SetValue:" + modelName + "." + fieldName)
			default:
				if fieldTypeKind == reflect.Struct || fieldTypeKind == reflect.Ptr {
					element := fieldValue.Addr().Interface()
					if fieldTypeKind == reflect.Ptr {
						element = reflect.New(fieldValue.Type().Elem()).Interface()
					}
					obj, objTypes := structToObjectValue(modelName+"."+fieldName, element)
					attrs[name] = obj
					typs[name] = types.ObjectType{}.WithAttributeTypes(objTypes)
				} else if fieldTypeKind == reflect.Slice {
					element := reflect.New(fieldValue.Type().Elem()).Interface()
					// fmt.Println(reflect.TypeOf(element))
					switch element.(type) {
					case *basetypes.StringValue:
						attrs[name], _ = types.ListValue(types.StringType, []attr.Value{})
						typs[name] = types.ListType{}.WithElementType(types.StringType)
					case *basetypes.Int64Value:
						attrs[name], _ = types.ListValue(types.Int64Type, []attr.Value{})
						typs[name] = types.ListType{}.WithElementType(types.Int64Type)
					case *basetypes.Float64Value:
						attrs[name], _ = types.ListValue(types.Float64Type, []attr.Value{})
						typs[name] = types.ListType{}.WithElementType(types.Float64Type)
					case *basetypes.BoolValue:
						attrs[name], _ = types.ListValue(types.BoolType, []attr.Value{})
						typs[name] = types.ListType{}.WithElementType(types.BoolType)
					case *basetypes.ObjectValue:
						panic("unsupported ObjectValue:" + modelName + "." + fieldName)
					case *basetypes.ListValue:
						panic("unsupported ListValue:" + modelName + "." + fieldName)
					case *basetypes.MapValue:
						panic("unsupported MapValue:" + modelName + "." + fieldName)
					case *basetypes.SetValue:
						panic("unsupported SetValue:" + modelName + "." + fieldName)
					default:
						switch fieldValue.Type().Elem().Kind() {
						case reflect.Struct:
							_, objTypes := structToObjectValue(modelName+"."+fieldName, element)
							typ := types.ObjectType{
								AttrTypes: objTypes,
							}
							attrs[name], _ = types.ListValue(typ, []attr.Value{})
							typs[name] = types.ListType{}.WithElementType(types.ObjectType{
								AttrTypes: objTypes,
							})
						default:
							panic("unsupported type for: " + modelName + "." + fieldName + ":" + fieldTypeStr)
						}
					}
				} else {
					/*if fieldTypeKind == reflect.Struct {
						if t, ok := registeredTypes[fieldTypeStr]; ok {
							attrs[name] = types.ObjectType{
								AttrTypes: structToTFType(modelName+"."+fieldName, t),
							}
						} else {
							panic("unsupported type for: " + modelName + "." + fieldName + ":" + fieldTypeStr)
						}
					} else {*/
					panic("unsupported type for: " + modelName + "." + fieldName + ":" + fieldTypeStr)
				}
			}
		}
	default:
		panic("not a struct")
	}

	// types := structToTFType(modelName, model)
	obj, diags := basetypes.NewObjectValue(typs, attrs)

	if diags.HasError() {
		panic("error: " + modelName + ":" + fmt.Sprint(diags))
	}

	return obj, typs
}

func flagsDescription(ctx context.Context, flags string, def string, apiPath string) string {
	k := []string{"enum", "min", "max", "default", "regex", "length"}
	var description []string
	for _, v := range k {
		s, ok := FlagsGet(flags, v)
		if ok {
			if strings.Contains(s, " ") {
				s = "\"" + s + "\""
			}
			if s == "" {
				s = def
			}
			description = append(description, v+":"+s)
		}
	}

	if ctx != nil {
		m := ctx.Value("fieldDescription")
		if m != nil {
			if v, ok := m.(map[string]string); ok {
				if v2, ok := v[apiPath]; ok {
					if v2 != "" {
						description = append(description, v2)
					}
				} else {
					// description = append(description, apiPath)
				}
			}
		}
	}
	desc, _ := FlagsGet(flags, "desc")
	if desc == "" {
		return strings.Join(description, ", ")
	}
	return strings.Join([]string{strings.Join(description, ", "), desc}, "")
}
