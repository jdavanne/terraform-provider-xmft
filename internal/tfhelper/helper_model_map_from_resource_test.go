package tfhelper

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestResourceToAttributes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Define the input data for the GenerateToSchema function
	var v TestResourceModel
	v.Name = types.StringValue("name1")
	v.Str = types.StringValue("str1")
	v.StrCO = types.StringValue("strCO1")
	v.StrNoRead = types.StringValue("strNoRead1")
	v.StrNoWrite = types.StringValue("strNoWrite1")
	v.StrOptional = types.StringValue("strOptional1")
	v.BoolTrue = types.BoolValue(true)
	v.BoolFalse = types.BoolValue(false)
	v.Bool = types.BoolValue(false)
	v.BoolDefault = types.BoolValue(false)
	v.BoolDefaultTrue = types.BoolValue(true)
	v.BoolDefaultFalse = types.BoolValue(false)
	v.Int64 = types.Int64Value(5678)
	v.Int64Default = types.Int64Value(0)
	v.Int64DefaultVal = types.Int64Value(1234)
	v.StrList = append(v.StrList, types.StringValue("strList1"))
	v.SubList = append(v.SubList, TestResourceModelSub{
		Id:          types.StringValue("1"),
		SubStr:      types.StringValue("SubStr3"),
		SubStrCO:    types.StringValue("SubStrCO1"),
		SubOptional: types.StringValue("SubOptional1"),
		SubSubStruct: TestResourceModelSubSub{
			SubSubStr:  types.StringValue("SubSubStr1"),
			SubSubBool: types.BoolValue(true),
		},
	})
	v.SubStruct = TestResourceModelSub{
		Id:          types.StringValue("2"),
		SubStr:      types.StringValue("SubStr4"),
		SubStrCO:    types.StringValue("SubStrCO2"),
		SubOptional: types.StringValue("SubOptional45"),
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
	v.SubListString = PO(types.ListValue(types.StringType, []attr.Value{types.StringValue("value1")}))
	/*v.SubListObject = PO(types.ListValue(types.ObjectType{}.WithAttributeTypes(structNameToTFType("TestResourceModelSubSub")),
		[]attr.Value{
			PO(types.ObjectValue(structNameToTFType("TestResourceModelSubSub"), map[string]attr.Value{
				"sub_sub_str":  types.StringValue("SubSubStr65"),
				"sub_sub_bool": types.BoolValue(true),
			})),
		},
	))*/
	v.SubMapString = PO(types.MapValue(types.StringType, map[string]attr.Value{
		"additionalProp1": types.StringValue("value1"),
		"additionalProp2": types.StringValue("value2"),
	}))
	/*v.Object = PO(types.ObjectValue(structNameToTFType("TestResourceModelSubSub"), map[string]attr.Value{
		"sub_sub_str":  types.StringValue("SubSubStr567"),
		"sub_sub_bool": types.BoolValue(true),
	}))*/
	v.PolyPtr.T2 = &TestResourceModelSubSub{
		SubSubStr:  types.StringValue("SubSubStr56734"),
		SubSubBool: types.BoolValue(true),
	}
	v.ListPolyPtr = append(v.ListPolyPtr, TestResourceModelPolyPtr{
		T2: &TestResourceModelSubSub{
			SubSubStr:  types.StringValue("SubSubStr56734"),
			SubSubBool: types.BoolValue(true),
		},
	}, TestResourceModelPolyPtr{
		T1: &TestResourceModelSubSub2{
			SubSubStr:  types.StringValue("SubSubStr534"),
			SubSubInt:  types.Int64Value(452),
			SubSubBool: types.BoolValue(false),
		},
	})

	/*v.Poly.T1 = PO(types.ObjectValue(structNameToTFType("TestResourceModelSubSub2"), map[string]attr.Value{
		"sub_sub_str":  types.StringValue("SubSubStr23"),
		"sub_sub_int":  types.Int64Value(4523),
		"sub_sub_bool": types.BoolValue(true),
	}))*/

	attrs := make(map[string]interface{})
	// Call the GenerateToSchema function
	ResourceToAttributes(ctx, "testresource", &v, attrs)

	attrsWrite := map[string]interface{}{
		//"name":         "name1",
		"str":         "str1",
		"str_co":      "strCO1",
		"str_no_read": "strNoRead1",
		//"str_no_write": "strNoWrite1",
		"str_optional":       "strOptional1",
		"bool_true":          true,
		"bool_false":         false,
		"bool":               false,
		"bool_default":       false,
		"bool_default_true":  true,
		"bool_default_false": false,
		"int64":              int64(5678),
		"int64_default":      int64(0),
		"int64_default_val":  int64(1234),
		"str_list":           []interface{}{"strList1"},
		"sub_list": []interface{}{
			map[string]interface{}{
				"id": "1", "sub_str": "SubStr3", "sub_str_co": "SubStrCO1", "sub_optional": "SubOptional1",
				"sub_sub_struct": map[string]interface{}{"subSubStr": "SubSubStr1", "sub_sub_bool": true},
			},
		},
		"sub_struct": map[string]interface{}{
			"id": "2", "sub_str": "SubStr4", "sub_str_co": "SubStrCO2", "sub_optional": "SubOptional45",
			"sub_sub_struct": map[string]interface{}{"subSubStr": "SubSubStr2", "sub_sub_bool": false},
		},
		"sub_struct_p": map[string]interface{}{
			"id": "2", "sub_str": "SubStr4", "sub_str_co": "SubStrCO2", "sub_optional": "SubOptional45",
			"sub_sub_struct": map[string]interface{}{"subSubStr": "SubSubStr3", "sub_sub_bool": true},
		},
		"sub_list_string": []interface{}{"value1"},
		/*"sub_list_object": []interface{}{
			map[string]interface{}{"subSubStr": "SubSubStr65", "sub_sub_bool": true},
		},*/
		"sub_map_string": map[string]interface{}{
			"additionalProp1": "value1",
			"additionalProp2": "value2",
		},
		/*"object": map[string]interface{}{
			"subSubStr":    "SubSubStr567",
			"sub_sub_bool": true,
		},
		*/
		/*"poly": map[string]interface{}{
			"subSubStr": "SubSubStr56734", "sub_sub_bool": true,
		},*/
		"poly_ptr": map[string]interface{}{"subSubStr": "SubSubStr56734", "sub_sub_bool": true},
		"list_poly_ptr": []interface{}{
			map[string]interface{}{"subSubStr": "SubSubStr56734", "sub_sub_bool": true},
			map[string]interface{}{"subSubStr": "SubSubStr534", "sub_sub_int": int64(452), "sub_sub_bool": false},
		},
	}

	// Check proper map from ResourceModel to Attributes map
	assert.Equal(t, attrsWrite, attrs)
}
