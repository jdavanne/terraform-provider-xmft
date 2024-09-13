package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
	{
	  "name": "string",
	  "type": "AccountFilePurge",
	  "notes": "string",
	  "managedByCG": false,
	  "additionalAttributes": {
	    "additionalProp1": "string",
	    "additionalProp2": "string",
	    "additionalProp3": "string"
	  },
	  "businessUnits": [
	    "string"
	  ],
	  "deleteFilesDays": 0,
	  "pattern": "string",
	  "expirationPeriod": false,
	  "removeFolders": false,
	  "warningNotifications": false,
	  "notifyDays": "string",
	  "sendSentinelAlert": false,
	  "warnNotifyAccount": false,
	  "warningNotificationsTemplate": "FileMaintenanceNotification.xhtml",
	  "warnNotifyEmails": "string",
	  "deletionNotifications": false,
	  "deletionNotificationsTemplate": "FileMaintenanceNotification.xhtml",
	  "deletionNotifyAccount": false,
	  "deletionNotifyEmails": "string",
	  "schedules": [
	    {
	      "tag": "string",
	      "type": "ONCE",
	      "executionTimes": [
	        "string"
	      ],
	      "startDate": {},
	      "skipHolidays": false
	    }
	  ]
	}
*/
type stFileMaintenanceApplication struct {
	Id          types.String `tfsdk:"id" helper:",computed,state"`
	Name        types.String `tfsdk:"name" helper:",required"`
	LastUpdated types.String `tfsdk:"last_updated" helper:",computed,noread,nowrite"`

	Type          types.String   `tfsdk:"type" helper:",default:AccountFilePurge"`
	Notes         types.String   `tfsdk:"notes"`
	BusinessUnits []types.String `tfsdk:"business_units" helper:"businessUnits,default:"`
	ManagedByCG   types.Bool     `tfsdk:"managed_by_cg" helper:"managedByCG,default:false"`

	AdditionalAttributes types.Map `tfsdk:"additional_attributes" helper:"additionalAttributes,elementtype:string,optional"`

	DeleteFilesDays               types.Int64  `tfsdk:"delete_files_days" helper:"deleteFilesDays,default:0"`
	Pattern                       types.String `tfsdk:"pattern" helper:"pattern,default:"`
	ExpirationPeriod              types.Bool   `tfsdk:"expiration_period" helper:"expirationPeriod,default:false"`
	RemoveFolders                 types.Bool   `tfsdk:"remove_folders" helper:"removeFolders,default:false"`
	WarningNotifications          types.Bool   `tfsdk:"warning_notifications" helper:"warningNotifications,default:false"`
	NotifyDays                    types.String `tfsdk:"notify_days" helper:"notifyDays,default:"`
	SendSentinelAlert             types.Bool   `tfsdk:"send_sentinel_alert" helper:"sendSentinelAlert,default:false"`
	WarnNotifyAccount             types.Bool   `tfsdk:"warn_notify_account" helper:"warnNotifyAccount,default:false"`
	WarningNotificationsTemplate  types.String `tfsdk:"warning_notifications_template" helper:"warningNotificationsTemplate,default:FileMaintenanceNotification.xhtml"`
	WarnNotifyEmails              types.String `tfsdk:"warn_notify_emails" helper:"warnNotifyEmails,default:"`
	DeletionNotifications         types.Bool   `tfsdk:"deletion_notifications" helper:"deletionNotifications,default:false"`
	DeletionNotificationsTemplate types.String `tfsdk:"deletion_notifications_template" helper:"deletionNotificationsTemplate,default:FileMaintenanceNotification.xhtml"`
	DeletionNotifyAccount         types.Bool   `tfsdk:"deletion_notify_account" helper:"deletionNotifyAccount,default:false"`
	DeletionNotifyEmails          types.String `tfsdk:"deletion_notify_emails" helper:"deletionNotifyEmails,default:"`

	Schedules []Schedule `tfsdk:"schedules" helper:"schedules,fold:type,optional"`
}

func NewSTFileMaintenanceApplicationResource() resource.Resource {
	return NewSTResource(&stFileMaintenanceApplication{}, "st_file_maintenance_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}").UseSwaggerUri("/api/v2.0/applications/{name}[type=AccountFilePurge]")
}

func init() {
	registerResource(NewSTFileMaintenanceApplicationResource)
}
