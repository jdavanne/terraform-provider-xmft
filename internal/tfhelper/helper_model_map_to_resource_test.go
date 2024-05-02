package tfhelper

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestAttributeToResource(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	attrsRead := map[string]interface{}{
		"name":         "name1",
		"str":          "str1",
		"str_co":       "strCO1",
		"str_no_read":  "strNoRead1",
		"str_no_write": "strNoWrite1",
		"str_optional": "strOptional1",
		"bool_true":    true,
		"bool_false":   false,
		"int64":        567867,
		"str_list":     []string{"strList789"},
		"sub_list": []interface{}{
			map[string]interface{}{
				"id": "1", "sub_str": "SubStr3", "sub_str_co": "SubStrCO1", "sub_optional": "SubOptional1",
				"sub_sub_struct": map[string]interface{}{"subSubStr": "SubSubStr1", "sub_sub_bool": false},
			},
		},
		"sub_struct": map[string]interface{}{
			"id": "2", "sub_str": "SubStr42", "sub_str_co": "SubStrCO2", "sub_optional": "SubOptional46",
			"sub_sub_struct": map[string]interface{}{"subSubStr": "SubSubStr2", "sub_sub_bool": false},
		},
		"sub_struct_p": map[string]interface{}{
			"id": "2", "sub_str": "SubStr4", "sub_str_co": "SubStrCO2", "sub_optional": "SubOptional45",
			"sub_sub_struct": map[string]interface{}{"subSubStr": "SubSubStr3", "sub_sub_bool": true},
		},
		"sub_map_string": map[string]interface{}{
			"additionalProp2": "value5",
			"additionalProp1": "value4",
		},
		"sub_list_string": []string{"value1"},
		"sub_list_object": []interface{}{
			map[string]interface{}{"subSubStr": "SubSubStr2", "sub_sub_bool": false},
		},
		"poly": map[string]interface{}{"subSubStr": "SubSubStr1234", "sub_sub_bool": false},
	}

	var v TestResourceModel
	// v.Name = types.StringValue("name1")
	v.Str = types.StringValue("str1")
	v.StrCO = types.StringValue("strCO1")
	// v.StrNoRead = types.StringValue("strNoRead1")
	v.StrNoWrite = types.StringValue("strNoWrite1")
	v.StrOptional = types.StringValue("strOptional1")
	v.BoolTrue = types.BoolValue(true)
	v.BoolFalse = types.BoolValue(false)
	v.StrList = append(v.StrList, types.StringValue("strList789"))
	v.SubList = append(v.SubList, TestResourceModelSub{
		Id:          types.StringValue("1"),
		SubStr:      types.StringValue("SubStr3"),
		SubStrCO:    types.StringValue("SubStrCO1"),
		SubOptional: types.StringValue("SubOptional1"),
		SubSubStruct: TestResourceModelSubSub{
			SubSubStr:  types.StringValue("SubSubStr1"),
			SubSubBool: types.BoolValue(false),
		},
	})
	v.SubStruct = TestResourceModelSub{
		Id:          types.StringValue("2"),
		SubStr:      types.StringValue("SubStr42"),
		SubStrCO:    types.StringValue("SubStrCO2"),
		SubOptional: types.StringValue("SubOptional46"),
		SubSubStruct: TestResourceModelSubSub{
			SubSubStr:  types.StringValue("SubSubStr2"),
			SubSubBool: types.BoolValue(false),
		},
	}
	v.SubStructP = &TestResourceModelSub{
		Id:          types.StringValue("2"),
		SubStr:      types.StringValue("SubStr4"),
		SubStrCO:    types.StringValue("SubStrCO2"),
		SubOptional: types.StringValue("SubOptional45"),
		SubSubStruct: TestResourceModelSubSub{
			SubSubStr:  types.StringValue("SubSubStr3"),
			SubSubBool: types.BoolValue(true),
		},
	}
	v.Int64 = types.Int64Value(567867)
	v.SubMapString = PO(types.MapValue(types.StringType, map[string]attr.Value{
		"additionalProp1": types.StringValue("value4"),
		"additionalProp2": types.StringValue("value5"),
	}))
	v.SubListString = PO(types.ListValue(types.StringType, []attr.Value{types.StringValue("value1")}))
	v.SubListObject = PO(types.ListValue(types.ObjectType{}.WithAttributeTypes(structNameToTFType("TestResourceModelSubSub")),
		[]attr.Value{
			PO(types.ObjectValue(structNameToTFType("TestResourceModelSubSub"), map[string]attr.Value{
				"sub_sub_str":  types.StringValue("SubSubStr2"),
				"sub_sub_bool": types.BoolValue(false),
			})),
		},
	))
	v.Poly = PO(types.ObjectValue(structNameToTFType("TestResourceModelSubSub"), map[string]attr.Value{
		"sub_sub_str":  types.StringValue("SubSubStr1234"),
		"sub_sub_bool": types.BoolValue(false),
	}))

	var v2 TestResourceModel
	AttributesToResource(ctx, "testresource", attrsRead, &v2)
	// attrs3 := make(map[string]interface{})
	// ResourceToAttributes(ctx, "testresource", &v2, attrs3)

	assert.Equal(t, &v, &v2)
}
