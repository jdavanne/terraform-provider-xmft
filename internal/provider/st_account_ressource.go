package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
{
        "accountCreationDate": 1713774059845,
        "accountEncryptMode": "UNSPECIFIED",
        "accountMaintenanceSettings": {
            "accountCertificateNotified": null,
            "accountDisabledDate": null,
            "accountNotified": null,
            "accountPasswordNotified": null,
            "action": {
                "action": null,
                "daysDeleteDisabled": null
            },
            "criteria": {
                "daysAfterCreation": null,
                "daysOfInactivity": null
            },
            "emailNotificationBeforeAction": {
                "emailTemplate": null,
                "notifyAccount": false,
                "notifyDays": null
            },
            "emailNotificationForUserCertificate": {
                "emailTemplate": null,
                "notifyAccount": false,
                "notifyDays": null
            },
            "emailNotificationForUserPassword": {
                "emailTemplate": null,
                "notifyAccount": false,
                "notifyDays": null
            },
            "policy": "default"
        },
        "accountSubmitForApprove": null,
        "accountVerificationStatus": null,
        "additionalAttributes": {},
        "addressBookSettings": {
            "contacts": [],
            "nonAddressBookCollaborationAllowed": null,
            "policy": "default",
            "sources": []
        },
        "adhocSettings": {
            "deliveryMethod": "DEFAULT",
            "enrollmentTypes": [],
            "implicitEnrollmentType": null
        },
        "authByEmail": false,
        "bandwidthLimits": {
            "inboundLimit": null,
            "outboundLimit": null,
            "policy": "default"
        },
        "businessUnit": null,
        "contact": {
            "email": null,
            "phone": null
        },
        "disabled": false,
        "fileArchivingPolicy": "DEFAULT",
        "fileMaintenanceSettings": {
            "deleteFilesDays": null,
            "deletionNotifications": null,
            "deletionNotificationsTemplate": null,
            "deletionNotifyAccount": null,
            "expirationPeriod": null,
            "notifyDays": null,
            "pattern": null,
            "policy": "default",
            "removeFolders": null,
            "reportNotified": null,
            "sendSentinelAlert": null,
            "warnNotified": null,
            "warnNotifyAccount": null,
            "warningNotifications": null,
            "warningNotificationsTemplate": null
        },
        "gid": "1000",
        "homeFolder": "/files/account1",
        "homeFolderAccessLevel": "PRIVATE",
        "id": "2cd008858ee428f2018f04e4e14500fb",
        "isUnlicensedUserAllowedToReply": true,
        "lastModified": null,
        "loginRestrictionPolicy": null,
        "managedByCG": null,
        "mappedUser": null,
        "name": "account1-2024-04-22_10-20-59",
        "notes": null,
        "pesitId": null,
        "rejectReason": null,
        "routingMode": "reject",
        "skin": "Default HTML Template",
        "subscriptionFolderDiscovery": "ITERATIVE",
        "transferType": "N",
        "transfersWebServiceAllowed": false,
        "type": "user",
        "uid": "1000",
        "unlicensed": false,
        "user": {
            "authExternal": false,
            "failedAuthAttempts": 0,
            "failedAuthMaximum": 3,
            "failedSshKeyAuthAttempts": 0,
            "failedSshKeyAuthMaximum": 3,
            "lastFailedAuth": null,
            "lastFailedSshKeyAuth": null,
            "lastLogin": null,
            "locked": false,
            "name": "account1-2024-04-22_10-20-59",
            "passwordCredentials": {
                "forcePasswordChange": false,
                "lastOwnPasswordChange": null,
                "lastPasswordChange": null,
                "password": null,
                "passwordDigest": "A/Re91oxrssYKjtoYIvJda9gJP1DQwon2JUXKdLCnlTG28pd+yWDRW3h4EimQ58N+Pg3Yk+LgA4c\r\nz5N6YMfCN7E9jyguILd4t0xdWjWaUAmJugnZTo/bJVczZNBfdYH2gtaw5Q1A+MpetvlOujA5R86e\r\nYtv4NmMeEX5q5i5mJ3YSCrreL3rIsVivlmHTf7+GDAYVEK21w4R/Lou3W/yOP82WGSqX8tCl1OsB\r\n0ApUs9b4hXx6UBWnp9va/QAkhRTWIMmn0TfjZ6YRLL4bRtbyojl8J9BMYoe5KSYYPDfcG30Fqd1H\r\nmNogiFZk9A2UJ2lOo4Gec1XIpUYiUTEIOWl6Hg==",
                "passwordExpiryInterval": null
            },
            "secretQuestion": {
                "forceSecretQuestionChange": false,
                "secretAnswerGuessFailures": 0,
                "secretQuestion": null
            },
            "successfulAuthMaximum": null,
            "successfulLogins": 0
        }
    }
*/

/*type stAccountContact struct {
	Email types.String `tfsdk:"email" helper:"email,default:"`
	Phone types.String `tfsdk:"phone" helper:"phone,default:"`
}*/

type stAccountResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                  types.String `tfsdk:"type" helper:",default:user"`
	Uid                   types.String `tfsdk:"uid" helper:",default:10000"`
	Gid                   types.String `tfsdk:"gid" helper:",default:10000"`
	HomeFolder            types.String `tfsdk:"home_folder" helper:"homeFolder,required"`
	HomeFolderAccessLevel types.String `tfsdk:"home_folder_access_level" helper:"homeFolderAccessLevel,enum:/PRIVATE/PUBLIC/BUSINESS_UNIT,default:PRIVATE"`
	LastModified          types.String `tfsdk:"last_modified" helper:"lastModified,computed,nowrite"`
	Unlicensed            types.Bool   `tfsdk:"unlicensed" helper:"unlicensed,computed,nowrite"`

	User struct {
		Name types.String `tfsdk:"name" helper:",required"`

		AuthExternal             types.Bool   `tfsdk:"auth_external" helper:"authExternal,computed,nowrite"`
		FailedAuthAttempts       types.Int64  `tfsdk:"failed_auth_attempts" helper:"failedAuthAttempts,computed,nowrite"`
		FailedAuthMaximum        types.Int64  `tfsdk:"failed_auth_maximum" helper:"failedAuthMaximum,default:3"`
		FailedSshKeyAuthAttempts types.Int64  `tfsdk:"failed_ssh_key_auth_attempts" helper:"failedSshKeyAuthAttempts,computed,nowrite"`
		FailedSshKeyAuthMaximum  types.Int64  `tfsdk:"failed_ssh_key_auth_maximum" helper:"failedSshKeyAuthMaximum,default:3"`
		LastFailedAuth           types.String `tfsdk:"last_failed_auth" helper:"lastFailedAuth,computed,nowrite"`
		LastFailedSshKeyAuth     types.String `tfsdk:"last_failed_ssh_key_auth" helper:"lastFailedSshKeyAuth,computed,nowrite"`
		LastLogin                types.String `tfsdk:"last_login" helper:"lastLogin,computed,nowrite"`
		Locked                   types.Bool   `tfsdk:"locked" helper:"locked,computed,nowrite"`
		SuccessfulAuthMaximum    types.Int64  `tfsdk:"successful_auth_maximum" helper:"successfulAuthMaximum,optional"`
		SuccessfulLogins         types.Int64  `tfsdk:"successful_logins" helper:"successfulLogins,computed,nowrite"`
		/*SecretQuestion           struct {
			ForceSecretQuestionChange types.Bool   `tfsdk:"force_secret_question_change" helper:"forceSecretQuestionChange,computed,nowrite"`
			SecretAnswerGuessFailures types.Int64  `tfsdk:"secret_answer_guess_failures" helper:"secretAnswerGuessFailures,computed,nowrite"`
			SecretQuestion            types.String `tfsdk:"secret_question" helper:"secretQuestion,computed,nowrite"`
		} `tfsdk:"secret_question" helper:"secretQuestion,optional"`*/
		PasswordCredentials struct {
			// ForcePasswordChange    types.Bool   `tfsdk:"force_password_change" helper:"forcePasswordChange,computed,optional"`
			Password               types.String `tfsdk:"password" helper:",noread"`
			PasswordDigest         types.String `tfsdk:"password_digest" helper:"passwordDigest,computed,nowrite,optional"`
			PasswordExpiryInterval types.Int64  `tfsdk:"password_expiry_interval" helper:"passwordExpiryInterval,optional"`
			LastOwnPasswordChange  types.String `tfsdk:"last_own_password_change" helper:"lastOwnPasswordChange,computed,nowrite"`
			LastPasswordChange     types.String `tfsdk:"last_password_change" helper:"lastPasswordChange,computed,nowrite"`
		} `tfsdk:"password_credentials" helper:"passwordCredentials,state,required"`
	} `tfsdk:"user" helper:",required"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Disabled types.Bool `tfsdk:"disabled" helper:"disabled,default:false"`

	FileArchivingPolicy     types.String `tfsdk:"file_archiving_policy" helper:"fileArchivingPolicy,enum:/DEFAULT/DISABLED/ENABLED,default:DEFAULT"`
	FileMaintenanceSettings struct {
		DeleteFilesDays               types.Int64  `tfsdk:"delete_files_days" helper:"deleteFilesDays,optional"`
		DeletionNotifications         types.Bool   `tfsdk:"deletion_notifications" helper:"deletionNotifications,optional"`
		DeletionNotificationsTemplate types.String `tfsdk:"deletion_notifications_template" helper:"deletionNotificationsTemplate,optional"`
		DeletionNotifyAccount         types.Bool   `tfsdk:"deletion_notify_account" helper:"deletionNotifyAccount,optional"`
		ExpirationPeriod              types.Bool   `tfsdk:"expiration_period" helper:"expirationPeriod,optional"`
		NotifyDays                    types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
		Pattern                       types.String `tfsdk:"pattern" helper:"pattern,optional"`
		Policy                        types.String `tfsdk:"policy" helper:"policy,enum:/default/custom/disabled,default:default"`
		RemoveFolders                 types.Bool   `tfsdk:"remove_folders" helper:"removeFolders,optional"`
		ReportNotified                types.String `tfsdk:"report_notified" helper:"reportNotified,optional"`
		SendSentinelAlert             types.Bool   `tfsdk:"send_sentinel_alert" helper:"sendSentinelAlert,optional"`
		WarnNotified                  types.String `tfsdk:"warn_notified" helper:"warnNotified,optional"`
		WarnNotifyAccount             types.Bool   `tfsdk:"warn_notify_account" helper:"warnNotifyAccount,optional"`
		WarningNotifications          types.Bool   `tfsdk:"warning_notifications" helper:"warningNotifications,optional"`
		WarningNotificationsTemplate  types.String `tfsdk:"warning_notifications_template" helper:"warningNotificationsTemplate,optional"`
	} `tfsdk:"file_maintenance_settings" helper:"fileMaintenanceSettings,default:"`

	AuthByEmail                    types.Bool   `tfsdk:"auth_by_email" helper:"authByEmail,default:false"`
	AccountCreationDate            types.Int64  `tfsdk:"account_creation_date" helper:"accountCreationDate,computed,nowrite"`
	AccountEncryptMode             types.String `tfsdk:"account_encrypt_mode" helper:"accountEncryptMode,enum:/UNSPECIFIED/ENABLED,default:UNSPECIFIED"`
	AccountSubmitForApprove        types.Bool   `tfsdk:"account_submit_for_approve" helper:"accountSubmitForApprove,computed,nowrite"`
	AccountVerificationStatus      types.String `tfsdk:"account_verification_status" helper:"accountVerificationStatus,computed,nowrite"`
	BusinessUnit                   types.String `tfsdk:"business_unit" helper:"businessUnit,optional"`
	IsUnlicensedUserAllowedToReply types.Bool   `tfsdk:"is_unlicensed_user_allowed_to_reply" helper:"isUnlicensedUserAllowedToReply,default:true"`
	LoginRestrictionPolicy         types.String `tfsdk:"login_restriction_policy" helper:"loginRestrictionPolicy,optional"`
	ManagedByCG                    types.Bool   `tfsdk:"managed_by_cg" helper:"managedByCG,optional"`
	MappedUser                     types.String `tfsdk:"mapped_user" helper:"mappedUser,optional"`
	Notes                          types.String `tfsdk:"notes" helper:"notes,optional"`
	PesitId                        types.String `tfsdk:"pesit_id" helper:"pesitId,optional"`
	RejectReason                   types.String `tfsdk:"reject_reason" helper:"rejectReason,computed,nowrite"`
	RoutingMode                    types.String `tfsdk:"routing_mode" helper:"routingMode,enum:/reject/accept/ignore,default:reject"`
	Skin                           types.String `tfsdk:"skin" helper:"skin,default:Default HTML Template"`
	SubscriptionFolderDiscovery    types.String `tfsdk:"subscription_folder_discovery" helper:"subscriptionFolderDiscovery,enum:/ITERATIVE/RECURSIVE,default:ITERATIVE"`
	TransferType                   types.String `tfsdk:"transfer_type" helper:"transferType,enum:/false/I/E,default:N"`
	TransfersWebServiceAllowed     types.Bool   `tfsdk:"transfers_web_service_allowed" helper:"transfersWebServiceAllowed,default:false"`

	AccountMaintenanceSettings struct {
		AccountCertificateNotified types.Bool  `tfsdk:"account_certificate_notified" helper:"accountCertificateNotified,default:false"`
		AccountDisabledDate        types.Int64 `tfsdk:"account_disabled_date" helper:"accountDisabledDate,optional"`
		AccountNotified            types.Bool  `tfsdk:"account_notified" helper:"accountNotified,optional"`
		AccountPasswordNotified    types.Bool  `tfsdk:"account_password_notified" helper:"accountPasswordNotified,optional"`
		Action                     struct {
			Action             types.String `tfsdk:"action" helper:"action,enum:/DELETE/DISABLE/PURGE,optional"`
			DaysDeleteDisabled types.Int64  `tfsdk:"days_delete_disabled" helper:"daysDeleteDisabled,optional"`
		} `tfsdk:"action" helper:"action,default:"`
		Criteria struct {
			DaysAfterCreation types.Int64 `tfsdk:"days_after_creation" helper:"daysAfterCreation,optional"`
			DaysOfInactivity  types.Int64 `tfsdk:"days_of_inactivity" helper:"daysOfInactivity,optional"`
		} `tfsdk:"criteria" helper:"criteria,default:"`
		EmailNotificationBeforeAction struct {
			EmailTemplate types.String `tfsdk:"email_template" helper:"emailTemplate,optional"`
			NotifyAccount types.Bool   `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
			NotifyDays    types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
		} `tfsdk:"email_notification_before_action" helper:"emailNotificationBeforeAction,,default:"`
		EmailNotificationForUserCertificate struct {
			EmailTemplate types.String `tfsdk:"email_template" helper:"emailTemplate,optional"`
			NotifyAccount types.Bool   `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
			NotifyDays    types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
		} `tfsdk:"email_notification_for_user_certificate" helper:"emailNotificationForUserCertificate,default:"`
		EmailNotificationForUserPassword struct {
			EmailTemplate types.String `tfsdk:"email_template" helper:"emailTemplate,optional"`
			NotifyAccount types.Bool   `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
			NotifyDays    types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
		} `tfsdk:"email_notification_for_user_password" helper:"emailNotificationForUserPassword,optional"`
		Policy types.String `tfsdk:"policy" helper:"policy,enum:/default/custom/disabled,default:default"`
	} `tfsdk:"account_maintenance_settings" helper:"accountMaintenanceSettings,default:"`

	AddressBookSettings struct {
		Contacts []struct {
			FullName     types.String `tfsdk:"full_name" helper:"fullName,emptyIsNull,default:"`
			Id           types.String `tfsdk:"id" helper:",computed"`
			PrimaryEmail types.String `tfsdk:"primary_email" helper:"primaryEmail,emptyIsNull,default:"`
		} `tfsdk:"contacts" helper:"contacts,optional"`
		NonAddressBookCollaborationAllowed types.Bool   `tfsdk:"non_address_book_collaboration_allowed" helper:"nonAddressBookCollaborationAllowed,optional"`
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
		ImplicitEnrollmentType types.String   `tfsdk:"implicit_enrollment_type" helper:"implicitEnrollmentType,enum:/ANONYMOUS_LINK/CHALLENGED_LINK/EXISTING_ACCOUNT/ENROLL_UNLICENSED/ENROLL_LICENSED,optional"`
	} `tfsdk:"adhoc_settings" helper:"adhocSettings,default:"`

	BandwidthLimits struct {
		InboundLimit  types.Int64  `tfsdk:"inbound_limit" helper:"inboundLimit,optional"`
		OutboundLimit types.Int64  `tfsdk:"outbound_limit" helper:"outboundLimit,optional"`
		Policy        types.String `tfsdk:"policy" helper:"policy,enum:/default/custom/disabled,default:default"`
	} `tfsdk:"bandwidth_limits" helper:"bandwidthLimits,default:"`

	Contact struct {
		Email types.String `tfsdk:"email" helper:"email,optional"`
		Phone types.String `tfsdk:"phone" helper:"phone,optional"`
	} `tfsdk:"contact" helper:"contact,default:"`
}

func NewSTAccountResource() resource.Resource {
	return NewSTResource(&stAccountResourceModel{}, "st_account", "", "/api/v2.0/accounts", "/api/v2.0/accounts/{name}").AddDiscriminator("[type=user]")
}

func init() {
	registerResource(NewSTAccountResource)
	// tfhelper.RegisterType("stAccountContact", &stAccountContact{})
}
