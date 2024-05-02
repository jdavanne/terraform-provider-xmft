package tfhelper

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	return structToTFType(modelName, model)
}

func structToTFType(modelName string, model interface{}) map[string]attr.Type {
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
			// flags := tag.Get(helperTagName)
			name := FlagsTfsdkGetName(tfsdk)
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
			default:
				if fieldTypeKind == reflect.Struct {
					if _, ok := registeredTypes[fieldTypeStr]; ok {
						attrs[name] = types.ObjectType{}
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
