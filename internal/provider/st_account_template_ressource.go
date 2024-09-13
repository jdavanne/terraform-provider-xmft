package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stAccountTemplateResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                  types.String `tfsdk:"type" helper:",default:template"`
	Uid                   types.String `tfsdk:"uid" helper:",default:10000"`
	Gid                   types.String `tfsdk:"gid" helper:",default:10000"`
	HomeFolder            types.String `tfsdk:"home_folder" helper:"homeFolder,required"`
	HomeFolderAccessLevel types.String `tfsdk:"home_folder_access_level" helper:"homeFolderAccessLevel,enum:/PRIVATE/PUBLIC/BUSINESS_UNIT,default:PRIVATE"`
	LastModified          types.String `tfsdk:"last_modified" helper:"lastModified,computed,nowrite"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Disabled types.Bool `tfsdk:"disabled" helper:"disabled,default:false"`

	FileArchivingPolicy     types.String `tfsdk:"file_archiving_policy" helper:"fileArchivingPolicy,enum:/DEFAULT/DISABLED/ENABLEDv,default:DEFAULT"`
	FileMaintenanceSettings struct {
		DeleteFilesDays               types.Int64  `tfsdk:"delete_files_days" helper:"deleteFilesDays,optional"`
		DeletionNotifications         types.Bool   `tfsdk:"deletion_notifications" helper:"deletionNotifications,default:false"`
		DeletionNotificationsTemplate types.String `tfsdk:"deletion_notifications_template" helper:"deletionNotificationsTemplate,optional"`
		DeletionNotifyAccount         types.Bool   `tfsdk:"deletion_notify_account" helper:"deletionNotifyAccount,optional"`
		ExpirationPeriod              types.Bool   `tfsdk:"expiration_period" helper:"expirationPeriod,optional"`
		NotifyDays                    types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
		Pattern                       types.String `tfsdk:"pattern" helper:"pattern,optional"`
		Policy                        types.String `tfsdk:"policy" helper:"policy,default:default"`
		RemoveFolders                 types.Bool   `tfsdk:"remove_folders" helper:"removeFolders,default:"`
		ReportNotified                types.String `tfsdk:"report_notified" helper:"reportNotified,optional"`
		SendSentinelAlert             types.Bool   `tfsdk:"send_sentinel_alert" helper:"sendSentinelAlert,default:"`
		WarnNotified                  types.String `tfsdk:"warn_notified" helper:"warnNotified,optional"`
		WarnNotifyAccount             types.Bool   `tfsdk:"warn_notify_account" helper:"warnNotifyAccount,default:"`
		WarningNotifications          types.Bool   `tfsdk:"warning_notifications" helper:"warningNotifications,optional"`
		WarningNotificationsTemplate  types.String `tfsdk:"warning_notifications_template" helper:"warningNotificationsTemplate,optional"`
	} `tfsdk:"file_maintenance_settings" helper:"fileMaintenanceSettings,default:"`

	AuthByEmail                    types.Bool   `tfsdk:"auth_by_email" helper:"authByEmail,default:false"`
	AccountCreationDate            types.Int64  `tfsdk:"account_creation_date" helper:"accountCreationDate,computed,nowrite"`
	AccountEncryptMode             types.String `tfsdk:"account_encrypt_mode" helper:"accountEncryptMode,enum:/UNSPECIFIED/ENABLED,default:UNSPECIFIED"`
	AccountSubmitForApprove        types.Bool   `tfsdk:"account_submit_for_approve" helper:"accountSubmitForApprove,computed,nowrite"`
	AccountVerificationStatus      types.String `tfsdk:"account_verification_status" helper:"accountVerificationStatus,computed,nowrite"`
	BusinessUnit                   types.String `tfsdk:"business_unit" helper:"businessUnit,emptyIsNull,default:"`
	EnrolledWithExternalPass       types.Bool   `tfsdk:"enrolled_with_external_pass" helper:"enrolledWithExternalPass,default:false"`
	IsUnlicensedUserAllowedToReply types.Bool   `tfsdk:"is_unlicensed_user_allowed_to_reply" helper:"isUnlicensedUserAllowedToReply,default:true"`
	LoginRestrictionPolicy         types.String `tfsdk:"login_restriction_policy" helper:"loginRestrictionPolicy,default:"`
	ManagedByCG                    types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,default:"`
	MappedUser                     types.String `tfsdk:"mapped_user" helper:"mappedUser,emptyIsNull,default:"`
	Notes                          types.String `tfsdk:"notes" helper:"notes,emptyIsNull,default:"`
	PesitId                        types.String `tfsdk:"pesit_id" helper:"pesitId,emptyIsNull,default:"`
	RejectReason                   types.String `tfsdk:"reject_reason" helper:"rejectReason,computed,nowrite"`
	Skin                           types.String `tfsdk:"skin" helper:"skin,default:Default HTML Template"`
	SubscriptionFolderDiscovery    types.String `tfsdk:"subscription_folder_discovery" helper:"subscriptionFolderDiscovery,enum:/ITERATIVE/RECURSIVE,default:ITERATIVE"`
	TemplateClass                  types.String `tfsdk:"template_class" helper:"templateClass,default:"`
	TransferType                   types.String `tfsdk:"transfer_type" helper:"transferType,enum:/false/I/E,default:N"`
	TransfersWebServiceAllowed     types.Bool   `tfsdk:"transfers_web_service_allowed" helper:"transfersWebServiceAllowed,default:false"`

	AddressBookSettings struct {
		Contacts []struct {
			FullName     types.String `tfsdk:"full_name" helper:"fullName,emptyIsNull,default:"`
			Id           types.String `tfsdk:"id" helper:",computed"`
			PrimaryEmail types.String `tfsdk:"primary_email" helper:"primaryEmail,emptyIsNull,default:"`
		} `tfsdk:"contacts" helper:"contacts,optional"`
		NonAddressBookCollaborationAllowed types.Bool   `tfsdk:"non_address_book_collaboration_allowed" helper:"nonAddressBookCollaborationAllowed,default:"`
		Policy                             types.String `tfsdk:"policy" helper:"policy,enum:/default/custom/disabled,default:default"`
		Sources                            []struct {
			// ignore for now CustomProperties types.Map    `tfsdk:"custom_properties" helper:"customProperties,elementtype:string,optional"`
			Id          types.String `tfsdk:"id" helper:",computed"`
			Name        types.String `tfsdk:"name" helper:",required"`
			Type        types.String `tfsdk:"type" helper:",enum:/LOCAL/LDAP/CUSTOM,default:LOCAL"`
			ParentGroup types.String `tfsdk:"parent_group" helper:"parentGroup,emptyIsNull,default:"`
			Enabled     types.Bool   `tfsdk:"enabled" helper:"enabled,default:false"`
		} `tfsdk:"sources" helper:"sources,elementtype:string,emptyIsNull,optional"`
	} `tfsdk:"address_book_settings" helper:"addressBookSettings,default:"`

	AdhocSettings struct {
		DeliveryMethod         types.String   `tfsdk:"delivery_method" helper:"deliveryMethod,enum:/DEFAULT/DISABLED/ANONYMOUS/ACCOUNT_WITHOUT_ENROLLMENT/ACCOUNT_WITH_ENROLLMENT/CUSTOM,emptyIsNull,default:DEFAULT"`
		EnrollmentTypes        []types.String `tfsdk:"enrollment_types" helper:"enrollmentTypes,elementtype:string,emptyIsNull,default:"`
		ImplicitEnrollmentType types.String   `tfsdk:"implicit_enrollment_type" helper:"implicitEnrollmentType,enum:/ANONYMOUS_LINK/CHALLENGED_LINK/EXISTING_ACCOUNT/ENROLL_UNLICENSED/ENROLL_LICENSED,emptyIsNull,default:"`
	} `tfsdk:"adhoc_settings" helper:"adhocSettings,default:"`

	BandwidthLimits struct {
		InboundLimit  types.Int64  `tfsdk:"inbound_limit" helper:"inboundLimit,default:"`
		OutboundLimit types.Int64  `tfsdk:"outbound_limit" helper:"outboundLimit,default:"`
		Policy        types.String `tfsdk:"policy" helper:"policy,enum:/default/custom/disabled,default:default"`
	} `tfsdk:"bandwidth_limits" helper:"bandwidthLimits,default:"`

	Contact struct {
		Email types.String `tfsdk:"email" helper:"email,emptyIsNull,default:"`
		Phone types.String `tfsdk:"phone" helper:"phone,emptyIsNull,default:"`
	} `tfsdk:"contact" helper:"contact,default:"`
}

func NewSTAccountTemplateResource() resource.Resource {
	return NewSTResource(&stAccountTemplateResourceModel{}, "st_account_template", "", "/api/v2.0/accounts", "/api/v2.0/accounts/{name}").AddDiscriminator("[type=template]")
}

func init() {
	registerResource(NewSTAccountTemplateResource)
	// tfhelper.RegisterType("stAccountContact", &stAccountContact{})
}
