package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stAddressBook struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:"name,required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	// AdditionalAttributes types.Map  `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
	Enabled types.Bool `tfsdk:"enabled" helper:"enabled,default:false"`

	ParentGroup types.String `tfsdk:"parent_group" helper:"parentGroup,default:"`
	Type        types.String `tfsdk:"type" helper:"type,enum:/LOCAL/LDAP/CUSTOM,default:"`
}

func NewSTAddressBookModelResource() resource.Resource {
	return NewSTResource(&stAddressBook{}, "st_address_book", "", "/api/v2.0/addressBook/sources", "/api/v2.0/addressBook/sources/{id}")
}

func init() {
	registerResource(NewSTAddressBookModelResource)
}
