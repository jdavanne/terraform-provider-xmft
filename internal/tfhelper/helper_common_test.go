package tfhelper

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TestResourceModelSubSub struct {
	SubSubStr  types.String `tfsdk:"sub_sub_str" helper:"subSubStr"`
	SubSubBool types.Bool   `tfsdk:"sub_sub_bool"`
}

type TestResourceModelSubSub2 struct {
	SubSubStr  types.String `tfsdk:"sub_sub_str" helper:"subSubStr"`
	SubSubInt  types.Int64  `tfsdk:"sub_sub_int"`
	SubSubBool types.Bool   `tfsdk:"sub_sub_bool"`
}

type TestResourceModelSub struct {
	Id           types.String            `tfsdk:"id" helper:",required"`
	SubStr       types.String            `tfsdk:"sub_str"`
	SubStrCO     types.String            `tfsdk:"sub_str_co" helper:",computed,optional"`
	SubOptional  types.String            `tfsdk:"sub_optional" helper:",optional"`
	SubSubStruct TestResourceModelSubSub `tfsdk:"sub_sub_struct"`
}

type TestResourceModelPolyObject struct {
	T1 types.Object `tfsdk:"t1" helper:",elementtype:TestResourceModelSubSub2"`
	T2 types.Object `tfsdk:"t2" helper:",elementtype:TestResourceModelSubSub"`
	T3 types.Object `tfsdk:"t3" helper:",elementtype:TestResourceModelSubSub"`
}

type TestResourceModelPolyPtr struct {
	T1 *TestResourceModelSubSub2 `tfsdk:"t1" helper:"t1"`
	T2 *TestResourceModelSubSub  `tfsdk:"t2" helper:"t2"`
	T3 *TestResourceModelSubSub  `tfsdk:"t3" helper:"t3"`
}

type TestResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required,noread,nowrite"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Str              types.String `tfsdk:"str"` // (array)
	StrDefault       types.String `tfsdk:"str_default"`
	StrCO            types.String `tfsdk:"str_co" helper:",computed,optional"`
	StrNoRead        types.String `tfsdk:"str_no_read" helper:",noread"`
	StrNoWrite       types.String `tfsdk:"str_no_write" helper:",nowrite"`
	StrOptional      types.String `tfsdk:"str_optional" helper:",optional"`
	BoolFalse        types.Bool   `tfsdk:"bool_false"`
	BoolTrue         types.Bool   `tfsdk:"bool_true"`
	Bool             types.Bool   `tfsdk:"bool"`
	BoolDefault      types.Bool   `tfsdk:"bool_default" helper:",default"`
	BoolDefaultTrue  types.Bool   `tfsdk:"bool_default_true" helper:",default:true"`
	BoolDefaultFalse types.Bool   `tfsdk:"bool_default_false" helper:",default:false"`
	Int64            types.Int64  `tfsdk:"int64"`
	Int64Default     types.Int64  `tfsdk:"int64_default" helper:",default"`
	Int64DefaultVal  types.Int64  `tfsdk:"int64_default_val" helper:",default:1234"`

	StrList    []types.String         `tfsdk:"str_list"`
	SubList    []TestResourceModelSub `tfsdk:"sub_list"`
	SubStruct  TestResourceModelSub   `tfsdk:"sub_struct"`
	SubStructP *TestResourceModelSub  `tfsdk:"sub_struct_p"`

	// SubObject     types.Object `tfsdk:"sub_object" helper:",optional,elementtype:TestResourceModelSub"`
	SubListString types.List   `tfsdk:"sub_list_string" helper:",optional,elementtype:string"`
	SubListObject types.List   `tfsdk:"sub_list_object" helper:",optional,elementtype:TestResourceModelSubSub"`
	SubMapString  types.Map    `tfsdk:"sub_map_string" helper:",elementtype:string,optional"`
	Object        types.Object `tfsdk:"object" helper:",elementtype:TestResourceModelSubSub"`
	// Poly          TestResourceModelPoly `tfsdk:"poly" helper:",fold:kind,elementtype:TestResourceModelPoly"`
	// PolyList      types.List            `tfsdk:"poly_list" helper:",fold:kind,elementtype:TestResourceModelPoly"`
	// SubMapObject  types.Map  `tfsdk:"additional_attributes" helper:"additional_attributes,elementtype:TestResourceModelSub,optional"`

	PolyPtr     TestResourceModelPolyPtr   `tfsdk:"poly_ptr" helper:",fold:kind"`
	ListPolyPtr []TestResourceModelPolyPtr `tfsdk:"list_poly_ptr" helper:",fold:kind"`
}

func init() {
	// Register the model
	RegisterType("TestResourceModelSub", &TestResourceModelSub{})
	RegisterType("TestResourceModelSubSub", &TestResourceModelSubSub{})
	RegisterType("TestResourceModelSubSub2", &TestResourceModelSubSub2{})
}

func PO[T any](a T, diags diag.Diagnostics) T {
	if diags.HasError() {
		s := ""
		s += fmt.Sprint(reflect.TypeOf(a)) + "\n"
		for _, d := range diags {
			s += d.Detail() + "\n"
		}
		panic(s)
	}
	return a
}
