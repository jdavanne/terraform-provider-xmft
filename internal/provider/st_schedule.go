package provider

import (
	"terraform-provider-xmft/internal/tfhelper"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	DailyType      types.String   `tfsdk:"daily_type" helper:"dailyType,enum:/EVERYDAY/EVERY_WEEKDAY,default:EVERYDAY"`
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
	DaysOfWeek     []types.String `tfsdk:"days_of_week" helper:"daysOfWeek,enum:/SUNDAY/MONDAY/TUESDAY/WEDNESDAY/THURSDAY/FRIDAY/SATURDAY,uniqueItems,nullable"`
	DaysOfMonth    []types.Int64  `tfsdk:"days_of_month" helper:"daysOfMonth,uniqueItems,nullable"`
	MonthlyType    types.String   `tfsdk:"monthly_type" helper:"monthlyType,enum:/EXACT/SEQUENCE_WEEKLY/SEQUENCE_MONTHLY,default:EXACT"`
	WeekOfMonth    types.String   `tfsdk:"week_of_month" helper:"weekOfMonth,enum:/FIRST/SECOND/THIRD/FOURTH/LAST,default:FIRST"`
}

type ScheduleYearly struct {
	Tag            types.String   `tfsdk:"tag" helper:"tag"`
	Type           types.String   `tfsdk:"type" helper:"type,default:YEARLY"`
	ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String   `tfsdk:"start_date" helper:"startDate"`
	EndDate        types.String   `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool     `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
	Months         []types.String `tfsdk:"months" helper:"months,enum:/JANUARY/FEBRUARY/MARCH/APRIL/MAY/JUNE/JULY/AUGUST/SEPTEMBER/OCTOBER/NOVEMBER/DECEMBER,uniqueItems,nullable"`
	DaysOfMonth    []types.Int64  `tfsdk:"days_of_month" helper:"daysOfMonth,uniqueItems,nullable"`
}

type ScheduleExpression struct {
	Tag  types.String `tfsdk:"tag" helper:"tag"`
	Type types.String `tfsdk:"type" helper:"type,default:EXPRESSION"`
	// ExecutionTimes []types.String `tfsdk:"execution_times" helper:"executionTimes,uniqueItems,nullable"`
	StartDate      types.String `tfsdk:"start_date" helper:"startDate"`
	EndDate        types.String `tfsdk:"end_date" helper:"endDate,format:date-time"`
	SkipHolidays   types.Bool   `tfsdk:"skip_holidays" helper:"skipHolidays,default:false"`
	CronExpression types.String `tfsdk:"cron_expression" helper:"cronExpression,default:"`
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

func init() {
	tfhelper.RegisterType("Schedule", &Schedule{})
}
