package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

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

type stAccountContact struct {
	Email types.String `tfsdk:"email" helper:"email,optional"`
	Phone types.String `tfsdk:"phone" helper:"phone,optional"`
}

type stAccountResourceModel struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type       types.String `tfsdk:"type" helper:",default:user"`
	Uid        types.String `tfsdk:"uid" helper:",default:1000"`
	Gid        types.String `tfsdk:"gid" helper:",default:1000"`
	HomeFolder types.String `tfsdk:"home_folder" helper:"homeFolder,required"`

	User struct {
		Name                types.String `tfsdk:"name" helper:",required"`
		FailedAuthAttempts  types.Int64  `tfsdk:"failed_auth_attempts" helper:"failedAuthAttempts,computed"`
		LastFailedAuth      types.String `tfsdk:"last_failed_auth" helper:"lastFailedAuth,computed"`
		PasswordCredentials struct {
			ForcePasswordChange types.Bool   `tfsdk:"force_password_change" helper:"forcePasswordChange,computed,optional"`
			Password            types.String `tfsdk:"password" helper:",noread"`
			PasswordDigest      types.String `tfsdk:"password_digest" helper:"passwordDigest,computed,optional,nowrite"`
		} `tfsdk:"password_credentials" helper:"passwordCredentials,state,required"`
		/*SecretQuestion struct {
			ForceSecretQuestionChange types.Bool   `tfsdk:"force_secret_question_change" helper:"forceSecretQuestionChange,computed,optional"`
			SecretAnswerGuessFailures types.Int64  `tfsdk:"secret_answer" helper:"secretAnswerGuessFailures,computed"`
			SecretQuestion            types.String `tfsdk:"secret_question" helper:"secretQuestion,default"`
		} `tfsdk:"secret_question" helper:"secretQuestion,state,required"`*/
	} `tfsdk:"user" helper:",state,required"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Disabled types.Bool `tfsdk:"disabled" helper:"disabled,default:false"`

	/*FileMaintenanceSettings *struct {
		DeleteFilesDays               types.Int64  `tfsdk:"delete_files_days" helper:"deleteFilesDays,optional"`
		DeletionNotifications         types.String `tfsdk:"deletion_notifications" helper:"deletionNotifications,optional"`
		DeletionNotificationsTemplate types.String `tfsdk:"deletion_notifications_template" helper:"deletionNotificationsTemplate,optional"`
		DeletionNotifyAccount         types.String `tfsdk:"deletion_notify_account" helper:"deletionNotifyAccount,optional"`
		ExpirationPeriod              types.Int64  `tfsdk:"expiration_period" helper:"expirationPeriod,optional"`
		NotifyDays                    types.Int64  `tfsdk:"notify_days" helper:"notifyDays,optional"`
		Pattern                       types.String `tfsdk:"pattern" helper:"pattern,optional"`
		Policy                        types.String `tfsdk:"policy" helper:"policy,default:default"`
		RemoveFolders                 types.String `tfsdk:"remove_folders" helper:"removeFolders,optional"`
		ReportNotified                types.String `tfsdk:"report_notified" helper:"reportNotified,optional"`
		SendSentinelAlert             types.String `tfsdk:"send_sentinel_alert" helper:"sendSentinelAlert,optional"`
		WarnNotified                  types.String `tfsdk:"warn_notified" helper:"warnNotified,optional"`
		WarnNotifyAccount             types.String `tfsdk:"warn_notify_account" helper:"warnNotifyAccount,optional"`
		WarningNotifications          types.String `tfsdk:"warning_notifications" helper:"warningNotifications,optional"`
		WarningNotificationsTemplate  types.String `tfsdk:"warning_notifications_template" helper:"warningNotificationsTemplate,optional"`
	} `tfsdk:"file_maintenance_settings" helper:"fileMaintenanceSettings,default:"`*/

	AuthByEmail         types.Bool   `tfsdk:"auth_by_email" helper:"authByEmail,default:false"`
	AccountCreationDate types.Int64  `tfsdk:"account_creation_date" helper:"accountCreationDate,computed"`
	AccountEncryptMode  types.String `tfsdk:"account_encrypt_mode" helper:"accountEncryptMode,enum:/UNSPECIFIED/ENABLED,default:UNSPECIFIED"`
	// AccountSubmitForApprove        types.Bool   `tfsdk:"account_submit_for_approve" helper:"accountSubmitForApprove,optional"`
	// AccountVerificationStatus      types.String `tfsdk:"account_verification_status" helper:"accountVerificationStatus,optional"`
	BusinessUnit                   types.String `tfsdk:"business_unit" helper:"businessUnit,emptyIsNull,default:"`
	IsUnlicensedUserAllowedToReply types.Bool   `tfsdk:"is_unlicensed_user_allowed_to_reply" helper:"isUnlicensedUserAllowedToReply,default:true"`
	LoginRestrictionPolicy         types.String `tfsdk:"login_restriction_policy" helper:"loginRestrictionPolicy,default:"`
	ManagedByCG                    types.String `tfsdk:"managed_by_cg" helper:"managedByCG,default:"`
	MappedUser                     types.String `tfsdk:"mapped_user" helper:"mappedUser,emptyIsNull,default:"`
	Notes                          types.String `tfsdk:"notes" helper:"notes,emptyIsNull,default:"`
	PesitId                        types.String `tfsdk:"pesit_id" helper:"pesitId,emptyIsNull,default:"`
	// RejectReason                   types.String `tfsdk:"reject_reason" helper:"rejectReason,optional"`
	RoutingMode                 types.String `tfsdk:"routing_mode" helper:"routingMode,enum:/accept/reject/ignore,default:reject"`
	Skin                        types.String `tfsdk:"skin" helper:"skin,default:Default HTML Template"`
	SubscriptionFolderDiscovery types.String `tfsdk:"subscription_folder_discovery" helper:"subscriptionFolderDiscovery,enum:/ITERATIVE/RECURSIVE,default:ITERATIVE"`
	TransferType                types.String `tfsdk:"transfer_type" helper:"transferType,enum:/N/Y,default:N"`
	TransfersWebServiceAllowed  types.Bool   `tfsdk:"transfers_web_service_allowed" helper:"transfersWebServiceAllowed,default:false"`

	/*AccountMaintenanceSettings *struct {
	    AccountCertificateNotified *types.String `tfsdk:"account_certificate_notified" helper:"accountCertificateNotified,optional"`
	    AccountDisabledDate *types.String `tfsdk:"account_disabled_date" helper:"accountDisabledDate,optional"`
	    AccountNotified *types.String `tfsdk:"account_notified" helper:"accountNotified,optional"`
	    AccountPasswordNotified *types.String `tfsdk:"account_password_notified" helper:"accountPasswordNotified,optional"`
	    Action *struct {
	        Action *types.String `tfsdk:"action" helper:"action,optional"`
	        DaysDeleteDisabled *types.String `tfsdk:"days_delete_disabled" helper:"daysDeleteDisabled,optional"`
	    } `tfsdk:"action" helper:"action,optional"`
	    Criteria *struct {
	        DaysAfterCreation *types.String `tfsdk:"days_after_creation" helper:"daysAfterCreation,optional"`
	        DaysOfInactivity *types.String `tfsdk:"days_of_inactivity" helper:"daysOfInactivity,optional"`
	    } `tfsdk:"criteria" helper:"criteria,optional"`
	    EmailNotificationBeforeAction *struct {
	        EmailTemplate *types.String `tfsdk:"email_template" helper:"emailTemplate,optional"`
	        NotifyAccount *types.Bool `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
	        NotifyDays *types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
	    } `tfsdk:"email_notification_before_action" helper:"emailNotificationBeforeAction,optional"`
	    EmailNotificationForUserCertificate *struct {
	        EmailTemplate *types.String `tfsdk:"email_template" helper:"emailTemplate,optional"`
	        NotifyAccount *types.Bool `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
	        NotifyDays *types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
	    } `tfsdk:"email_notification_for_user_certificate" helper:"emailNotificationForUserCertificate,optional"`
	    EmailNotificationForUserPassword *struct {
	        EmailTemplate *types.String `tfsdk:"email_template" helper:"emailTemplate,optional"`
	        NotifyAccount *types.Bool `tfsdk:"notify_account" helper:"notifyAccount,default:false"`
	        NotifyDays *types.String `tfsdk:"notify_days" helper:"notifyDays,optional"`
	    } `tfsdk:"email_notification_for_user_password" helper:"emailNotificationForUserPassword,optional"`
	    Policy *types.String `tfsdk:"policy" helper:"policy,default:default"`
	} `tfsdk:"account_maintenance_settings" helper:"accountMaintenanceSettings,optional"`*/

	/*AddressBookSettings *struct {
	    Contacts []*stAccountContact `tfsdk:"contacts" helper:"contacts,elementtype:stAccountContact,optional"`
	    NonAddressBookCollaborationAllowed *types.String `tfsdk:"non_address_book_collaboration_allowed" helper:"nonAddressBookCollaborationAllowed,optional"`
	    Policy *types.String `tfsdk:"policy" helper:"policy,default:default"`
	    Sources []*types.String `tfsdk:"sources" helper:"sources,elementtype:string,optional"`
	} `tfsdk:"address_book_settings" helper:"addressBookSettings,optional"`*/

	/*AdhocSettings *struct {
	    DeliveryMethod *types.String `tfsdk:"delivery_method" helper:"deliveryMethod,enum:/DEFAULT/EMAIL,default:DEFAULT"`
	    EnrollmentTypes []*types.String `tfsdk:"enrollment_types" helper:"enrollmentTypes,elementtype:string,optional"`
	    ImplicitEnrollmentType *types.String `tfsdk:"implicit_enrollment_type" helper:"implicitEnrollmentType,optional"`
	} `tfsdk:"adhoc_settings" helper:"adhocSettings,optional"`*/

	/*BandwidthLimits *struct {
	    InboundLimit *types.String `tfsdk:"inbound_limit" helper:"inboundLimit,optional"`
	    OutboundLimit *types.String `tfsdk:"outbound_limit" helper:"outboundLimit,optional"`
	    Policy *types.String `tfsdk:"policy" helper:"policy,default:default"`
	} `tfsdk:"bandwidth_limits" helper:"bandwidthLimits,optional"`*/

	// Contact types.Object `tfsdk:"contact" helper:"contact,elementtype:stAccountContact,optional,computed"`
}

func NewSTAccountResource() resource.Resource {
	return NewSTResource(&stAccountResourceModel{}, "st_account", "", "/api/v2.0/accounts", "/api/v2.0/accounts/{name}").AddDiscriminator("[type=user]")
}

func init() {
	registerResource(NewSTAccountResource)
	tfhelper.RegisterType("stAccountContact", &stAccountContact{})
}
