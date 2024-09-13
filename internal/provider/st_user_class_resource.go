package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "className": "string",
  "userType": "*",
  "order": 0,
  "enabled": false,
  "userName": "string",
  "group": "string",
  "address": "string",
  "expression": "string"
}
*/

type stUserClassModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:"className,required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	UserType   types.String `tfsdk:"user_type" helper:"userType,enum:/*/real/virtual,default:*"`
	Order      types.Int64  `tfsdk:"order" helper:"order,default:1"`
	Enabled    types.Bool   `tfsdk:"enabled" helper:"enabled,default:true"`
	UserName   types.String `tfsdk:"user_name" helper:"userName,default:*"`
	Group      types.String `tfsdk:"group" helper:"group,default:*"`
	Address    types.String `tfsdk:"address" helper:"address,default:*"`
	Expression types.String `tfsdk:"expression" helper:"expression,default:"`
}

func NewSTUserClassModelModelResource() resource.Resource {
	return NewSTResource(&stUserClassModel{}, "st_user_class", "", "/api/v2.0/userClasses", "/api/v2.0/userClasses/{id}")
}

func init() {
	registerResource(NewSTUserClassModelModelResource)
}
