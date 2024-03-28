package tools

import (
	"fmt"
	"reflect"
	"strings"
)

type JSONOptions struct {
	Indent          bool
	RemoveUppercase bool
	Cson            bool
}

func JSONMarchal(s interface{}) ([]byte, error) {
	var sb strings.Builder
	jsonValue(reflect.ValueOf(s), &sb, JSONOptions{}, 0)
	return []byte(sb.String()), nil
}

// JSONMarchal: naive JSON marchaller (remove uppercase)
func jsonValue(value reflect.Value, sb *strings.Builder, opt JSONOptions, depth int) (empty bool) {
	// valueValue := value.Interface()

	switch value.Kind() {
	case reflect.Struct:
		v := value.Interface()
		first := true
		reflectType := reflect.TypeOf(v)
		reflectValue := reflect.ValueOf(v)
		sb.WriteString("{")
		for i := 0; i < reflectType.NumField(); i++ {
			var sb2 strings.Builder
			typeName := reflectType.Field(i).Name
			if typeName[0] >= 'a' && typeName[0] <= 'z' {
				continue
			}
			empty2 := jsonValue(reflectValue.Field(i), &sb2, opt, depth+1)
			val := sb2.String()
			if !empty2 {
				if !first {
					sb.WriteString(",")
				}
				first = false
				typeName := reflectType.Field(i).Name
				if typeName[0] >= 'A' && typeName[0] <= 'Z' {
					typeName = strings.ToLower(typeName[0:1]) + typeName[1:]
				}
				if !opt.Cson {
					sb.WriteByte('"')
				}
				sb.WriteString(typeName)
				if !opt.Cson {
					sb.WriteByte('"')
				}
				sb.WriteString(":")
				sb.WriteString(val)
			}
		}
		sb.WriteString("}")
	case reflect.String:
		valueValue := value.Interface()
		val := valueValue.(string)
		if val == "" {
			empty = true
		}
		sb.WriteByte('"')
		sb.WriteString(val)
		sb.WriteByte('"')
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		valueValue := value.Interface()
		val := fmt.Sprintf("%d", valueValue)
		if val == "0" {
			empty = true
		}
		sb.WriteString(val)
	case reflect.Float64, reflect.Float32:
		valueValue := value.Interface()
		sb.WriteString(fmt.Sprintf("%f", valueValue))
	case reflect.Bool:
		valueValue := value.Interface()
		// fmt.Printf("%t", valueValue)
		if valueValue.(bool) == false {
			empty = true
		}
		sb.WriteString(fmt.Sprintf("%t", valueValue))
	case reflect.Slice, reflect.Array:
		// valueValue := value.Interface()
		sb.WriteString("[")
		// s := reflect.ValueOf(value)
		s := value
		for i := 0; i < s.Len(); i++ {
			if i != 0 {
				sb.WriteString(",")
			}
			jsonValue(s.Index(i), sb, opt, depth+1)
		}
		sb.WriteString("]")
	case reflect.Map:
		sb.WriteString("{")
		for i, key := range value.MapKeys() {
			if i != 0 {
				sb.WriteString(",")
			}
			if !opt.Cson {
				sb.WriteByte('"')
			}
			sb.WriteString(key.String())
			if !opt.Cson {
				sb.WriteByte('"')
			}
			sb.WriteString(":")
			val := value.MapIndex(key)
			jsonValue(val, sb, opt, depth+1)
		}
		sb.WriteString("}")
	case reflect.Interface:
		jsonValue(value.Elem(), sb, opt, depth)
	case reflect.Ptr:
		jsonValue(value.Elem(), sb, opt, depth)
	case reflect.Chan,
		reflect.Func,
		reflect.UnsafePointer,
		reflect.Complex64, reflect.Complex128,
		reflect.Uintptr,
		reflect.Invalid:
		// panic("unknown type :" + value.Kind().String())
		sb.WriteString("!ERROR: unsupported type: " + value.Kind().String())
	default:
		panic("unknown type :" + value.Kind().String())
	}
	return
}
