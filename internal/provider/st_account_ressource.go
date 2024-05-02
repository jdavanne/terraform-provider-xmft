package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
name       = "account1"
type       = "user"
uid        = "1000"
gid        = "1000"
homeFolder = "/files/$name"
user = {
  name = "login1"
  passwordCredentials = {
	password = "mypassword1"
  }
}
*/

type stAccountResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type       types.String `tfsdk:"type" helper:",default:user"`
	Uid        types.String `tfsdk:"uid" helper:",default:1000"`
	Gid        types.String `tfsdk:"gid" helper:",default:1000"`
	HomeFolder types.String `tfsdk:"home_folder" helper:"homeFolder,required"`
	User       struct {
		Name                types.String `tfsdk:"name" helper:",required"`
		FailedAuthAttempts  types.Int64  `tfsdk:"failed_auth_attempts" helper:"failedAuthAttempts,computed"`
		LastFailedAuth      types.String `tfsdk:"last_failed_auth" helper:"lastFailedAuth,computed"`
		PasswordCredentials struct {
			ForcePasswordChange types.Bool   `tfsdk:"force_password_change" helper:"forcePasswordChange,computed,optional"`
			Password            types.String `tfsdk:"password" helper:",noread"`
			PasswordDigest      types.String `tfsdk:"password_digest" helper:"passwordDigest,computed,optional,nowrite"`
		} `tfsdk:"password_credentials" helper:"passwordCredentials,state,required"`
	} `tfsdk:"user" helper:",state,required"`
}

func NewSTAccountResource() resource.Resource {
	return NewSTResource(&stAccountResourceModel{}, "st_account", "", "/api/v2.0/accounts", "/api/v2.0/accounts/{name}")
}
