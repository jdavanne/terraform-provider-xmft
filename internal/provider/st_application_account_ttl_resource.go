package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type stAccountTTLApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type                                     types.String   `tfsdk:"type" helper:"type,default:AccountTTL"`
	BusinessUnits                            []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	AccountAction                            types.String   `tfsdk:"account_action" helper:"accountAction,enum:/DELETE/DISABLE/PURGE,default:"`
	AccountCreationDaysCriteria              types.Int64    `tfsdk:"account_creation_days_criteria" helper:"accountCreationDaysCriteria,default:"`
	AccountInactivityDaysCriteria            types.Int64    `tfsdk:"account_inactivity_days_criteria" helper:"accountInactivityDaysCriteria,default:"`
	ActionNotificationEmails                 types.String   `tfsdk:"action_notification_emails" helper:"actionNotificationEmails,default:"`
	CertificateNotificationDays              types.String   `tfsdk:"certificate_notification_days" helper:"certificateNotificationDays,default:"`
	CertificateNotificationEmailTemplate     types.String   `tfsdk:"certificate_notification_email_template" helper:"certificateNotificationEmailTemplate,enum:/AccountMaintenanceNotification.xhtml/FileMaintenanceNotification.xhtml/LoginThresholdReport.xhtml/AdhocDefault.xhtml/LoginThresholdNotification.xhtml/RoutingFailedNotification.xhtml/RoutingSucceededNotification.xhtml/RoutingTriggeredNotification.xhtml,default:AccountMaintenanceNotification.xhtml"`
	DeleteDisabledAccountAfterXDaysCriteria  types.Int64    `tfsdk:"delete_disabled_account_after_x_days_criteria" helper:"deleteDisabledAccountAfterXDaysCriteria,default:"`
	EnableAccountNotifications               types.Bool     `tfsdk:"enable_account_notifications" helper:"enableAccountNotifications,default:false"`
	EnableCertificateExpirationNotifications types.Bool     `tfsdk:"enable_certificate_expiration_notifications" helper:"enableCertificateExpirationNotifications,default:false"`
	EnablePasswordExpirationNotifications    types.Bool     `tfsdk:"enable_password_expiration_notifications" helper:"enablePasswordExpirationNotifications,default:false"`
	FirstMaintenanceRun                      types.Int64    `tfsdk:"first_maintenance_run" helper:"firstMaintenanceRun,default:"`
	ManagedByCG                              types.Bool     `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`
	Notes                                    types.String   `tfsdk:"notes" helper:"notes,default:"`
	NotificationCertificateEmails            types.String   `tfsdk:"notification_certificate_emails" helper:"notificationCertificateEmails,default:"`
	NotificationDays                         types.String   `tfsdk:"notification_days" helper:"notificationDays,default:"`
	PasswordExpirationNotificationEmails     types.String   `tfsdk:"password_expiration_notification_emails" helper:"passwordExpirationNotificationEmails,default:"`
	PasswordNotificationDays                 types.String   `tfsdk:"password_notification_days" helper:"passwordNotificationDays,default:"`
	PasswordNotificationEmailTemplate        types.String   `tfsdk:"password_notification_email_template" helper:"passwordNotificationEmailTemplate,enum:/AccountMaintenanceNotification.xhtml/FileMaintenanceNotification.xhtml/LoginThresholdReport.xhtml/AdhocDefault.xhtml/LoginThresholdNotification.xhtml/RoutingFailedNotification.xhtml/RoutingSucceededNotification.xhtml/RoutingTriggeredNotification.xhtml,default:AccountMaintenanceNotification.xhtml"`
	ReportNotificationEmailTemplate          types.String   `tfsdk:"report_notification_email_template" helper:"reportNotificationEmailTemplate,enum:/AccountMaintenanceNotification.xhtml/FileMaintenanceNotification.xhtml/LoginThresholdReport.xhtml/AdhocDefault.xhtml/LoginThresholdNotification.xhtml/RoutingFailedNotification.xhtml/RoutingSucceededNotification.xhtml/RoutingTriggeredNotification.xhtml,default:AccountMaintenanceNotification.xhtml"`
	AdditionalAttributes                     types.Map      `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
}

func NewSTAccountTTLApplicationResource() resource.Resource {
	return NewSTResource(&stAccountTTLApplication{}, "st_account_ttl_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=AccountTTL]")
}

func init() {
	registerResource(NewSTAccountTTLApplicationResource)
}
