package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
  "loginName": "string",
  "roleName": "string",
  "isLimited": false,
  "localAuthentication": true,
  "certificateDN": "string",
  "dualAuthentication": false,
  "locked": false,
  "parent": "string",
  "fullCreationPath": "string",
  "passwordCredentials": {
    "password": "string",
    "passwordExpired": false,
    "lastPasswordChangeTime": "string",
    "lastOwnPasswordChange": "string",
    "loginFailureCount": 0,
    "lastLoginTime": "string",
    "lastFailureLoginTime": "string"
  },
  "administratorRights": {
    "canReadOnly": false,
    "isMaker": false,
    "isChecker": false,
    "canCreateUsers": false,
    "canUpdateUsers": false,
    "canAccessHelpDesk": false,
    "canSeeFullAuditLog": false,
    "canManageAdministrators": false,
    "canManageApplications": false,
    "canManageSharedFolders": false,
    "canManageBusinessUnits": false,
    "canManageRouteTemplates": false,
    "canManageExternalScriptStep": false,
    "canManageLoginRestrictionPolicies": false,
    "canManageIcapSettings": false,
    "canManageExternalScriptRootExecution": false
  },
  "businessUnits": [
    "string"
  ]
}
*/

type stAdminRoleAdministratorRights struct {
	CanReadOnly                          types.Bool `tfsdk:"can_read_only" helper:"canReadOnly,default:false"`
	IsMaker                              types.Bool `tfsdk:"is_maker" helper:"isMaker,default:false"`
	IsChecker                            types.Bool `tfsdk:"is_checker" helper:"isChecker,default:false"`
	CanCreateUsers                       types.Bool `tfsdk:"can_create_users" helper:"canCreateUsers,default:false"`
	CanUpdateUsers                       types.Bool `tfsdk:"can_update_users" helper:"canUpdateUsers,default:false"`
	CanAccessHelpDesk                    types.Bool `tfsdk:"can_access_help_desk" helper:"canAccessHelpDesk,default:false"`
	CanSeeFullAuditLog                   types.Bool `tfsdk:"can_see_full_audit_log" helper:"canSeeFullAuditLog,default:false"`
	CanManageAdministrators              types.Bool `tfsdk:"can_manage_administrators" helper:"canManageAdministrators,default:false"`
	CanManageApplications                types.Bool `tfsdk:"can_manage_applications" helper:"canManageApplications,default:false"`
	CanManageSharedFolders               types.Bool `tfsdk:"can_manage_shared_folders" helper:"canManageSharedFolders,default:false"`
	CanManageBusinessUnits               types.Bool `tfsdk:"can_manage_business_units" helper:"canManageBusinessUnits,default:false"`
	CanManageRouteTemplates              types.Bool `tfsdk:"can_manage_route_templates" helper:"canManageRouteTemplates,default:false"`
	CanManageExternalScriptStep          types.Bool `tfsdk:"can_manage_external_script_step" helper:"canManageExternalScriptStep,default:false"`
	CanManageLoginRestrictionPolicies    types.Bool `tfsdk:"can_manage_login_restriction_policies" helper:"canManageLoginRestrictionPolicies,default:false"`
	CanManageIcapSettings                types.Bool `tfsdk:"can_manage_icap_settings" helper:"canManageIcapSettings,default:false"`
	CanManageExternalScriptRootExecution types.Bool `tfsdk:"can_manage_external_script_root_execution" helper:"canManageExternalScriptRootExecution,default:false"`
}

type stAdminModel struct {
	// Id                  types.String `tfsdk:"id" helper:",computed,state"`
	Name                types.String `tfsdk:"name" helper:"loginName,state,required"`
	LastUpdated         types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	RoleName            types.String `tfsdk:"role_name" helper:"roleName,required"`
	IsLimited           types.Bool   `tfsdk:"is_limited" helper:"isLimited,default:false"`
	LocalAuthentication types.Bool   `tfsdk:"local_authentication" helper:"localAuthentication,default:false"`
	CertificateDN       types.String `tfsdk:"certificate_dn" helper:"certificateDN,default:"`
	DualAuthentication  types.Bool   `tfsdk:"dual_authentication" helper:"dualAuthentication,default:false"`
	Locked              types.Bool   `tfsdk:"locked" helper:",computed"`
	Parent              types.String `tfsdk:"parent" helper:",default:"`
	FullCreationPath    types.String `tfsdk:"full_creation_path" helper:"fullCreationPath,computed"`
	PasswordCredentials struct {
		Password               types.String `tfsdk:"password" helper:",noread,sensitive"`
		PasswordExpired        types.Bool   `tfsdk:"password_expired" helper:"passwordExpired,nowrite,computed"`
		LastPasswordChangeTime types.String `tfsdk:"last_password_change_time" helper:"lastPasswordChangeTime,computed"`
		// LastOwnPasswordChange  types.String `tfsdk:"last_own_password_change" helper:"lastOwnPasswordChange,computed"`
		LoginFailureCount    types.Int64  `tfsdk:"login_failure_count" helper:"loginFailureCount,nowrite,computed"`
		LastLoginTime        types.String `tfsdk:"last_login_time" helper:"lastLoginTime,nowrite,computed"`
		LastFailureLoginTime types.String `tfsdk:"last_failure_login_time" helper:"lastFailureLoginTime,nowrite,computed"`
	} `tfsdk:"password_credentials" helper:"passwordCredentials,optional"`
	AdministratorRights types.Object `tfsdk:"administrator_rights" helper:"administratorRights,elementtype:stAdminRoleAdministratorRights,default:"`

	BusinessUnits types.List `tfsdk:"business_units" helper:",elementtype:string,optional"`
	// AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTAdminModelResource() resource.Resource {
	return NewSTResource(&stAdminModel{}, "st_admin", "", "/api/v2.0/administrators", "/api/v2.0/administrators/{loginName}")
}

func init() {
	registerResource(NewSTAdminModelResource)
	tfhelper.RegisterType("stAdminRoleAdministratorRights", &stAdminRoleAdministratorRights{})
}
