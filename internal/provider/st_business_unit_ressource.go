package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stBusinessUnitResourceModel struct {
	// Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	BaseFolder                 types.String   `tfsdk:"base_folder" helper:"baseFolder,required"`
	Parent                     types.String   `tfsdk:"parent" helper:",optional"`
	BusinessUnitHierarchy      types.String   `tfsdk:"business_unit_hierarchy" helper:"businessUnitHierarchy,computed"`
	BaseFolderModifyingAllowed types.Bool     `tfsdk:"base_folder_modifying_allowed" helper:"baseFolderModifyingAllowed,default:false"`
	HomeFolderModifyingAllowed types.Bool     `tfsdk:"home_folder_modifying_allowed" helper:"homeFolderModifyingAllowed,default:false"`
	Dmz                        types.String   `tfsdk:"dmz" helper:",optional"`
	ManagedByCG                types.Bool     `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	EnabledIcapServers         []types.String `tfsdk:"enabled_icap_servers" helper:"enabledIcapServers,default:"`

	BandwidthLimits struct {
		InboundLimit        types.Int64  `tfsdk:"inbound_limit" helper:"inboundLimit,optional"`
		OutboundLimit       types.Int64  `tfsdk:"outbound_limit" helper:"outboundLimit,optional"`
		ModifyLimitsAllowed types.Bool   `tfsdk:"modify_limits_allowed" helper:"modifyLimitsAllowed,default:false"`
		Policy              types.String `tfsdk:"policy" helper:"policy,emptyIsNull,default:default"`
	} `tfsdk:"bandwidth_limits" helper:"bandwidthLimits,default:"`

	AccountMaintenanceSettings struct {
		Criteria struct {
			DaysAfterCreation types.Int64  `tfsdk:"days_after_creation" helper:"daysAfterCreation,optional"`
			DaysOfInactivity  types.Int64  `tfsdk:"days_of_inactivity" helper:"daysOfInactivity,optional"`
			SpecificDate      types.String `tfsdk:"specific_date" helper:"specificDate,optional"`
		} `tfsdk:"criteria" helper:"criteria,default:"`
		Action struct {
			Action             types.String `tfsdk:"action" helper:"action,enum:/DELETE/DISABLE/PURGE,optional"`
			DaysDeleteDisabled types.Int64  `tfsdk:"days_delete_disabled" helper:"daysDeleteDisabled,optional"`
		} `tfsdk:"action" helper:"action,default:"`
		EmailNotificationBeforeAction struct {
			EmailTemplate types.String `tfsdk:"email_template" helper:"emailTemplate,default:"`
			NotifyAccount types.Bool   `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
			NotifyEmails  types.String `tfsdk:"notify_emails" helper:"notifyEmails,default:"`
			NotifyDays    types.String `tfsdk:"notify_days" helper:"notifyDays,default:"`
		} `tfsdk:"email_notification_before_action" helper:"emailNotificationBeforeAction,default:"`
		EmailNotificationForUserPassword struct {
			EmailTemplate types.String `tfsdk:"email_template" helper:"emailTemplate,default:"`
			NotifyAccount types.Bool   `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
			NotifyEmails  types.String `tfsdk:"notify_emails" helper:"notifyEmails,default:"`
			NotifyDays    types.String `tfsdk:"notify_days" helper:"notifyDays,default:"`
		} `tfsdk:"email_notification_for_user_password" helper:"emailNotificationForUserPassword,default:"`
		EmailNotificationForUserCertificate struct {
			EmailTemplate types.String `tfsdk:"email_template" helper:"emailTemplate,default:"`
			NotifyAccount types.Bool   `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
			NotifyEmails  types.String `tfsdk:"notify_emails" helper:"notifyEmails,default:"`
			NotifyDays    types.String `tfsdk:"notify_days" helper:"notifyDays,default:"`
		} `tfsdk:"email_notification_for_user_certificate" helper:"emailNotificationForUserCertificate,default:"`
		Policy                 types.String `tfsdk:"policy" helper:"policy,default:default"`
		PolicyModifyingAllowed types.Bool   `tfsdk:"policy_modifying_allowed" helper:"policyModifyingAllowed,default:false"`
	} `tfsdk:"account_maintenance_settings" helper:"accountMaintenanceSettings,default:"`

	HtmlTemplateSettings struct {
		HtmlTemplateFolderPath types.String `tfsdk:"html_template_folder_path" helper:"htmlTemplateFolderPath,default:Default HTML Template"`
		IsAllowedForModifying  types.Bool   `tfsdk:"is_allowed_for_modifying" helper:"isAllowedForModifying,default:false"`
	} `tfsdk:"html_template_settings" helper:"htmlTemplateSettings,default:"`

	TransfersApiSettings struct {
		IsWebServiceRightsModifyingAllowed types.Bool `tfsdk:"is_web_service_rights_modifying_allowed" helper:"isWebServiceRightsModifyingAllowed,default:false"`
		TransfersWebServiceAllowed         types.Bool `tfsdk:"transfers_web_service_allowed" helper:"transfersWebServiceAllowed,default:false"`
	} `tfsdk:"transfers_api_settings" helper:"transfersApiSettings,default:"`

	AdHocSettings struct {
		AuthByEmail types.Bool `tfsdk:"auth_by_email" helper:"authByEmail,default:false"`
		// IsAuthByEmailModifyingAllowed types.Bool `tfsdk:"is_auth_by_email_modifying_allowed" helper:"isAuthByEmailModifyingAllowed,default:false"`
		// IsDeliveryMethodModifyingAllowed types.Bool     `tfsdk:"is_delivery_method_modifying_allowed" helper:"isDeliveryMethodModifyingAllowed,default:false"`
		DeliveryMethod         types.String   `tfsdk:"delivery_method" helper:"deliveryMethod,enum:/DEFAULT/EMAIL,default:DEFAULT"`
		EnrollmentTypes        []types.String `tfsdk:"enrollment_types" helper:"enrollmentTypes,optional"`
		ImplicitEnrollmentType types.String   `tfsdk:"implicit_enrollment_type" helper:"implicitEnrollmentType,optional"`
		NotificationTemplate   types.String   `tfsdk:"notification_template" helper:"notificationTemplate,optional"`
		EnrollmentTemplate     types.String   `tfsdk:"enrollment_template" helper:"enrollmentTemplate,default:default"`
	} `tfsdk:"ad_hoc_settings" helper:"adHocSettings,default:"`

	FileArchivingSettings struct {
		Policy                      types.String `tfsdk:"policy" helper:"policy,default:default"`
		PolicyModifyingAllowed      types.Bool   `tfsdk:"policy_modifying_allowed" helper:"policyModifyingAllowed,default:false"`
		FolderPolicy                types.String `tfsdk:"folder_policy" helper:"folderPolicy,default:default"`
		CustomFolder                types.String `tfsdk:"custom_folder" helper:"customFolder,default:"`
		EncryptionCertificatePolicy types.String `tfsdk:"encryption_certificate_policy" helper:"encryptionCertificatePolicy,enum:/default/custom/disabled,default:default"`
		CustomEncryptionCertificate types.String `tfsdk:"custom_encryption_certificate" helper:"customEncryptionCertificate,optional"`
		CustomFileSizePolicy        types.String `tfsdk:"custom_file_size_policy" helper:"customFileSizePolicy,default:default"`
		CustomFileSize              types.Int64  `tfsdk:"custom_file_size" helper:"customFileSize,default:0"`
	} `tfsdk:"file_archiving_settings" helper:"fileArchivingSettings,default:"`

	LoginRestrictionSettings struct {
		IsPolicyModifyingAllowed types.Bool   `tfsdk:"is_policy_modifying_allowed" helper:"isPolicyModifyingAllowed,default:false"`
		Policy                   types.String `tfsdk:"policy" helper:"policy,optional"`
	} `tfsdk:"login_restriction_settings" helper:"loginRestrictionSettings,default:"`

	AddressBookSettings struct {
		Policy                             types.String `tfsdk:"policy" helper:"policy,optional"`
		NonAddressBookCollaborationAllowed types.Bool   `tfsdk:"non_address_book_collaboration_allowed" helper:"nonAddressBookCollaborationAllowed,default:false"`
		Sources                            []struct {
			Id               types.String `tfsdk:"id" helper:"id,required"`
			Name             types.String `tfsdk:"name" helper:"name,required"`
			Type             types.String `tfsdk:"type" helper:"type,required"`
			ParentGroup      types.String `tfsdk:"parent_group" helper:"parentGroup,required"`
			Enabled          types.Bool   `tfsdk:"enabled" helper:"enabled,default:false"`
			CustomProperties struct {
				Value1 types.String `tfsdk:"value1" helper:"value1,default:"`
				Value2 types.String `tfsdk:"value2" helper:"value2,default:"`
			} `tfsdk:"custom_properties" helper:"customProperties,default:"`
		} `tfsdk:"sources" helper:"sources,optional"`
		ModifySourcesAllowed       types.Bool `tfsdk:"modify_sources_allowed" helper:"modifySourcesAllowed,default:false"`
		ModifyCollaborationAllowed types.Bool `tfsdk:"modify_collaboration_allowed" helper:"modifyCollaborationAllowed,default:false"`
	} `tfsdk:"address_book_settings" helper:"addressBookSettings,default:"`

	FileMaintenanceSettings struct {
		AllowPolicyModifying          types.Bool   `tfsdk:"allow_policy_modifying" helper:"allowPolicyModifying,default:false"`
		Policy                        types.String `tfsdk:"policy" helper:"policy,enum:/default/custom/disabled,default:default"`
		DeleteFilesDays               types.Int64  `tfsdk:"delete_files_days" helper:"deleteFilesDays,optional"`
		Pattern                       types.String `tfsdk:"pattern" helper:"pattern,optional"`
		ExpirationPeriod              types.Bool   `tfsdk:"expiration_period" helper:"expirationPeriod,optional"`
		RemoveFolders                 types.Bool   `tfsdk:"remove_folders" helper:"removeFolders,optional"`
		WarningNotifications          types.Bool   `tfsdk:"warning_notifications" helper:"warningNotifications,optional"`
		NotifyDays                    types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
		SendSentinelAlert             types.Bool   `tfsdk:"send_sentinel_alert" helper:"sendSentinelAlert,optional"`
		WarnNotifyAccount             types.Bool   `tfsdk:"warn_notify_account" helper:"warnNotifyAccount,optional"`
		WarningNotificationsTemplate  types.String `tfsdk:"warning_notifications_template" helper:"warningNotificationsTemplate,optional"`
		WarnNotifyEmails              types.String `tfsdk:"warn_notify_emails" helper:"warnNotifyEmails,optional"`
		DeletionNotifications         types.Bool   `tfsdk:"deletion_notifications" helper:"deletionNotifications,optional"`
		DeletionNotificationsTemplate types.String `tfsdk:"deletion_notifications_template" helper:"deletionNotificationsTemplate,optional"`
		DeletionNotifyAccount         types.Bool   `tfsdk:"deletion_notify_account" helper:"deletionNotifyAccount,optional"`
		DeletionNotifyEmails          types.String `tfsdk:"deletion_notify_emails" helper:"deletionNotifyEmails,optional"`
	} `tfsdk:"file_maintenance_settings" helper:"fileMaintenanceSettings,default:"`

	SharedFoldersCollaborationAllowed types.Bool `tfsdk:"shared_folders_collaboration_allowed" helper:"sharedFoldersCollaborationAllowed,optional"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`
}

func NewSTBusinessUnitResource() resource.Resource {
	return NewSTResource(&stBusinessUnitResourceModel{}, "st_business_unit", "", "/api/v2.0/businessUnits", "/api/v2.0/businessUnits/{name}")
}

func init() {
	registerResource(NewSTBusinessUnitResource)
}
