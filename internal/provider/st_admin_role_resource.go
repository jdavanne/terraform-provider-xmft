package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
 {
      "roleName": "Master Administrator",
      "isLimited": false,
      "isBounceAllowed": true,
      "menus": [
        "SSH Settings",
        "Login Restrictions",
        "Business Units",
        "PeSIT Settings",
        "FTP Commands",
        "Axway Sentinel/DI",
        "Server Log",
        "ICAP Settings",
        "Cluster Management",
        "Administrators",
        "Holiday Schedule",
        "HTTP Settings",
        "User Accounts",
        "Database Settings",
        "Command Logging",
        "Restrictions",
        "Site Templates",
        "Server Access Control",
        "FTP Settings",
        "File Tracking",
        "Mail Templates",
        "Application",
        "Miscellaneous",
        "Service Accounts",
        "Audit Log",
        "Address Books",
        "System",
        "Certificates",
        "Admin Access Control",
        "Secure Socket Layer",
        "Server Control",
        "Virtual Groups",
        "Unlicensed Users",
        "Login Settings",
        "Manage Roles",
        "Admin Settings",
        "LDAP Domains",
        "Account Templates",
        "Import/Export",
        "File Archiving",
        "Access Rules",
        "SiteMinder Settings",
        "AS2 Settings",
        "Transfer Logging",
        "User Classes",
        "Change Password",
        "TM Settings",
        "Active Users",
        "Server Configuration",
        "Network Zones",
        "Home Folders",
        "Server Usage Monitor",
        "Support Tool",
        "User Type Ranges",
        "AdHoc Settings",
        "Server License",
        "Route Packages"
      ],
}
*/

type stAdminRoleModel struct {
	// Id              types.String   `tfsdk:"id" helper:",computed,state"`
	Name            types.String   `tfsdk:"name" helper:"roleName,state,required"`
	LastUpdated     types.String   `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	IsLimited       types.Bool     `tfsdk:"is_limited" helper:"isLimited,optional"`
	IsBounceAllowed types.Bool     `tfsdk:"is_bounce_allowed" helper:"isBounceAllowed,optional"`
	Menus           []types.String `tfsdk:"menus" helper:",optional"`

	// AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTAdminRoleModelResource() resource.Resource {
	return NewSTResource(&stAdminRoleModel{}, "st_admin_role", "", "/api/v2.0/administrativeRoles", "/api/v2.0/administrativeRoles/{roleName}").UseSwaggerUri("/api/v2.0/administrativeRoles/{name}")
}

func init() {
	registerResource(NewSTAdminRoleModelResource)
}
