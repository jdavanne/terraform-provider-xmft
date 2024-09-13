package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stLdapDomainModel struct {
	Id               types.String `tfsdk:"id" helper:",computed,state"`
	LastUpdated      types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`
	IsDefault        types.Bool   `tfsdk:"is_default" helper:"isDefault,optional"`
	Precedence       types.Int64  `tfsdk:"precedence" helper:"precedence,optional"`
	ProtocolVersion  types.Int64  `tfsdk:"protocol_version" helper:"protocolVersion,optional"`
	SslEnabled       types.Bool   `tfsdk:"ssl_enabled" helper:"sslEnabled,optional"`
	TlsEnabled       types.Bool   `tfsdk:"tls_enabled" helper:"tlsEnabled,optional"`
	ReferralsAllowed types.Bool   `tfsdk:"referrals_allowed" helper:"referralsAllowed,optional"`

	Name                  types.String `tfsdk:"name" helper:"name,state,required"`
	AnonymousBindsAllowed types.Bool   `tfsdk:"anonymous_binds_allowed" helper:"anonymousBindsAllowed,optional"`
	BindDnPassword        types.String `tfsdk:"bind_dn_password" helper:"bindDnPassword,optional"`
	BindDn                types.String `tfsdk:"bind_dn" helper:"bindDn,optional"`
	CommonCase            types.String `tfsdk:"common_case" helper:"commonCase,enum:/NONE/UPPER/LOWER,optional"`
	ClientCertificate     types.String `tfsdk:"client_certificate" helper:"clientCertificate,optional"`
	CertVerify            types.Bool   `tfsdk:"cert_verify" helper:"certVerify,optional"`
	Description           types.String `tfsdk:"description" helper:"description,optional"`

	UserSettings *struct {
		DefaultUid         types.Int64  `tfsdk:"default_uid" helper:"defaultUid,optional"`
		DefaultGid         types.Int64  `tfsdk:"default_gid" helper:"defaultGid,optional"`
		DefaultUserType    types.String `tfsdk:"default_user_type" helper:"defaultUserType,enum:/virtual/real,optional"`
		DefaultUserShell   types.String `tfsdk:"default_user_shell" helper:"defaultUserShell,optional"`
		SysUser            types.String `tfsdk:"sys_user" helper:"sysUser,optional"`
		DefaultAuthByEmail types.Bool   `tfsdk:"default_auth_by_email" helper:"defaultAuthByEmail,default:false"`
	} `tfsdk:"user_settings" helper:"userSettings,optional"`
	LdapSearches struct {
		BaseDn                 types.String `tfsdk:"base_dn" helper:"baseDn,required"`
		AliasQuery             types.String `tfsdk:"alias_query" helper:"aliasQuery,required"`
		UseGenericSearchFilter types.Bool   `tfsdk:"use_generic_search_filter" helper:"useGenericSearchFilter,optional"`
		GenericSearchAttribute types.String `tfsdk:"generic_search_attribute" helper:"genericSearchAttribute,optional"`
		GenericSearchFilter    types.String `tfsdk:"generic_search_filter" helper:"genericSearchFilter,optional"`
		SearchAttribute        types.String `tfsdk:"search_attribute" helper:"searchAttribute,optional"`
	} `tfsdk:"ldap_searches" helper:"ldapSearches,optional"`
	LdapServers []struct {
		Id    types.String `tfsdk:"id" helper:"id,default:1"`
		Host  types.String `tfsdk:"host" helper:"host,required"`
		Port  types.Int64  `tfsdk:"port" helper:"port,default:389"`
		Order types.Int64  `tfsdk:"order" helper:"order,required"`
	} `tfsdk:"ldap_servers" helper:"ldapServers,optional"`
	Attributes []struct {
		Attribute       types.String `tfsdk:"attribute" helper:"attribute,required"`
		Description     types.String `tfsdk:"description" helper:"description,optional"`
		Enabled         types.Bool   `tfsdk:"enabled" helper:"enabled,optional"`
		MappedAttribute types.String `tfsdk:"mapped_attribute" helper:"mappedAttribute,optional"`
	} `tfsdk:"attributes" helper:"attributes,optional"`
	AddressBookAttributes []struct {
		Attribute       types.String `tfsdk:"attribute" helper:"attribute,required"`
		Description     types.String `tfsdk:"description" helper:"description,optional"`
		Enabled         types.Bool   `tfsdk:"enabled" helper:"enabled,optional"`
		MappedAttribute types.String `tfsdk:"mapped_attribute" helper:"mappedAttribute,optional"`
	} `tfsdk:"address_book_attributes" helper:"addressBookAttributes,optional"`
	DnFilters []struct {
		Id        types.String `tfsdk:"id" helper:"id,required"`
		Enabled   types.String `tfsdk:"enabled" helper:"enabled,optional"`
		Filter    types.String `tfsdk:"filter" helper:"filter,required"`
		UserClass types.String `tfsdk:"user_class" helper:"userClass,required"`
	} `tfsdk:"dn_filters" helper:"dnFilters,optional"`
	AddressBookSettings *struct {
		AddressBookaseDN             types.String `tfsdk:"address_book_base_dn" helper:"addressBookBaseDN,optional"`
		AddressBookQuery             types.String `tfsdk:"address_book_query" helper:"addressBookQuery,optional"`
		ShouldUseOnlyAdditionalQuery types.Bool   `tfsdk:"should_use_only_additional_query" helper:"shouldUseOnlyAdditionalQuery,optional"`
	} `tfsdk:"address_book_settings" helper:"addressBookSettings,optional"`
}

func NewSTLdapDomainModelResource() resource.Resource {
	return NewSTResource(&stLdapDomainModel{}, "st_ldap_domain", "", "/api/v2.0/ldapDomains", "/api/v2.0/ldapDomains/{name}")
}

func init() {
	registerResource(NewSTLdapDomainModelResource)
}
