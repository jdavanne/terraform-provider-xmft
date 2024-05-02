package tfhelper

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TestResourceModelSubSub struct {
	SubSubStr  types.String `tfsdk:"sub_sub_str" helper:"subSubStr"`
	SubSubBool types.Bool   `tfsdk:"sub_sub_bool"`
}

type TestResourceModelSub struct {
	Id           types.String            `tfsdk:"id" helper:",required"`
	SubStr       types.String            `tfsdk:"sub_str"`
	SubStrCO     types.String            `tfsdk:"sub_str_co" helper:",computed,optional"`
	SubOptional  types.String            `tfsdk:"sub_optional" helper:",optional"`
	SubSubStruct TestResourceModelSubSub `tfsdk:"sub_sub_struct"`
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
	Poly          types.Object `tfsdk:"poly" helper:",fold,elementtype:TestResourceModelSubSub"`
	// SubMapObject  types.Map  `tfsdk:"additional_attributes" helper:"additional_attributes,elementtype:TestResourceModelSub,optional"`
}

func init() {
	// Register the model
	RegisterType("TestResourceModelSub", &TestResourceModelSub{})
	RegisterType("TestResourceModelSubSub", &TestResourceModelSubSub{})
}

func PO[T any](a T, diags diag.Diagnostics) T {
	if diags.HasError() {
		s := ""
		for _, d := range diags {
			s += d.Detail() + "\n"
		}
		panic(s)
	}
	return a
}
