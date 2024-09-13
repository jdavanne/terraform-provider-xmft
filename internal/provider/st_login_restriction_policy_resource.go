package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stLoginRestrictionPolicyModel struct {
	Id                   types.String   `tfsdk:"id" helper:",computed,state"`
	Name                 types.String   `tfsdk:"name" helper:"name,required"`
	LastUpdated          types.String   `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	AdditionalAttributes types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
	BusinessUnits        []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`

	Description types.String `tfsdk:"description" helper:"description,default:"`

	IsDefault types.Bool `tfsdk:"is_default" helper:"isDefault,default:false"`
	Rules     []struct {
		ClientAddress types.String `tfsdk:"client_address" helper:"clientAddress,default:"`
		Description   types.String `tfsdk:"description" helper:"description,default:"`
		Expression    types.String `tfsdk:"expression" helper:"expression,default:"`
		Id            types.String `tfsdk:"id" helper:"id,default:"`
		IsEnabled     types.Bool   `tfsdk:"is_enabled" helper:"isEnabled,default:false"`
		Name          types.String `tfsdk:"name" helper:"name,default:string"`
		Type          types.String `tfsdk:"type" helper:"type,default:"`
	} `tfsdk:"rules" helper:"rules,optional"`

	Type types.String `tfsdk:"type" helper:"type,enum:/ALLOW_THEN_DENY/DENY_THEN_ALLOW,default:ALLOW_THEN_DENY"`
}

func NewSTLoginRestrictionPolicyModelModelResource() resource.Resource {
	return NewSTResource(&stLoginRestrictionPolicyModel{}, "st_login_restriction_policy", "", "/api/v2.0/loginRestrictionPolicies", "/api/v2.0/loginRestrictionPolicies/{name}")
}

func init() {
	registerResource(NewSTLoginRestrictionPolicyModelModelResource)
}
