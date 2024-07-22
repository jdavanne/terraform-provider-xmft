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

type ScheduleOnce struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:ONCE"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
}

type ScheduleHourly struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:HOURLY"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate,format:date-time,default:"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`

	HourlyStep types.Int64  `tfsdk:"hourly_step" helper:"hourlyStep,default:1"`
	HourlyType types.String `tfsdk:"hourly_type" helper:"hourlyType,enum:/PERHOURS/PERMINUTES,default:PERHOURS"`
}

type ScheduleDaily struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:DAILY"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate,format:date-time"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time,default:"`
	DailyType      types.String   `tfsdk:"daily_type" helper:"dailyType,enum:/EVERYDAY/EVERY_WEEK_DAY,default:EVERYDAY"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
}

type ScheduleWeekly struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:WEEKLY"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
	DaysOfWeek     []types.String `tfsdk:"days_of_week" helper:"daysOfWeek,enum:/SUNDAY/MONDAY/TUESDAY/WEDNESDAY/THURSDAY/FRIDAY/SATURDAY,uniqueItems,nullable"`
}

type ScheduleMonthly struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:MONTHLY"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate,format:date-time"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
	DaysOfWeek     []types.String `tfsdk:"days_of_week" helper:"daysOfWeek,enum:enum:/SUNDAY/MONDAY/TUESDAY/WEDNESDAY/THURSDAY/FRIDAY/SATURDAY,uniqueItems,nullable"`
	DaysOfMonth    []types.Int64  `tfsdk:"days_of_month" helper:"daysOfMonth,uniqueItems,nullable"`
	MonthlyType    types.String   `tfsdk:"monthly_type" helper:"monthlyType,enum:/EXACT/SEQUENCE_WEEKLY/SEQUENCE_MONTHLY,default:EXACT"`
}

type ScheduleYearly struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:YEARLY"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
	Months         []types.String `tfsdk:"months" helper:"months,enum:/JANUARY/FEBRUARY/MARCH/APRIL/MAY/JUNE/JULY/AUGUST/SEPTEMBER/OCTOBER/NOVEMBER/DECEMBER,uniqueItems,nullable"`
}

type ScheduleExpression struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:EXPRESSION"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
	CronExpression types.String   `tfsdk:"cron_expression" helper:"cronExpression,default:"`
}

type Schedule struct {
	Once       *ScheduleOnce       `tfsdk:"once" helper:"ONCE"`
	Hourly     *ScheduleHourly     `tfsdk:"hourly" helper:"HOURLY"`
	Daily      *ScheduleDaily      `tfsdk:"daily" helper:"DAILY"`
	Monthly    *ScheduleMonthly    `tfsdk:"monthly" helper:"MONTHLY"`
	Weekly     *ScheduleWeekly     `tfsdk:"weekly" helper:"WEEKLY"`
	Yearly     *ScheduleYearly     `tfsdk:"yearly" helper:"YEARLY"`
	Expression *ScheduleExpression `tfsdk:"expression" helper:"EXPRESSION"`
}

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
	return NewSTResource(&stFileMaintenanceApplication{}, "st_file_maintenance_application", "", "/api/v2.0/applications", "/api/v2.0/applications/{name}")
}

func init() {
	registerResource(NewSTFileMaintenanceApplicationResource)
}
