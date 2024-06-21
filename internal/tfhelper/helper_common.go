package tfhelper

import (
	"fmt"
	"reflect"
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
	attrs := make(map[string]attr.Value)
	typs := make(map[string]attr.Type)

	value := reflect.ValueOf(model).Elem()
	switch value.Kind() {
	case reflect.Struct:
		reflectType := value.Type()
		for i := 0; i < reflectType.NumField(); i++ {
			fieldName := reflectType.Field(i).Name
			fieldTypeStr := reflectType.Field(i).Type.String()
			// fieldTypeKind := reflectType.Field(i).Type.Kind()
			fieldValue := value.Field(i)
			tag := reflectType.Field(i).Tag
			tfsdk := tag.Get(tfsdkTagName)
			// flags := tag.Get(helperTagName)
			name := FlagsTfsdkGetName(tfsdk)
			// elementtype, _ := FlagsGet(flags, "elementtype")
			// name := FlagsHelperName(tfsdk, flags)

			switch v := fieldValue.Interface().(type) {
			case basetypes.StringValue:
				attrs[name] = v
				typs[name] = types.StringType
			case basetypes.Int64Value:
				attrs[name] = v
				typs[name] = types.Int64Type
			case basetypes.Float64Value:
				attrs[name] = v
				typs[name] = types.Float64Type
			case basetypes.BoolValue:
				attrs[name] = v
				typs[name] = types.BoolType
			case basetypes.ObjectValue:
				attrs[name] = v // structToObjectValue(modelName+"."+fieldName, fieldValue.Interface())
				typs[name] = types.ObjectType{}.WithAttributeTypes(v.AttributeTypes(nil))
			default:
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
				//}
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

func flagsDescription(flags string, def string) string {
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

	desc, _ := FlagsGet(flags, "desc")
	if desc == "" {
		return strings.Join(description, ", ")
	}
	return strings.Join([]string{strings.Join(description, ", "), desc}, "")
}
