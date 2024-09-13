package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTFileMaintenanceApplicationResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_file_maintenance_application"
	resourceName := "file_maintenance1"
	name := "FileMaintenance1" + r
	startDate := fmt.Sprint(time.Now().Unix()*1000 + 3610*2*1000)
	// startDate := time.Now().Format("2006-01-02T15:04:05Z")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "` + resourceType + `" "` + resourceName + `" {
						provider       = xmft.st1
						name           = "` + name + `"
						notes          = "mynotes"
						#business_units = []
						additional_attributes = {
							"userVars.additionalProp1" = "val1"
						}
						
						delete_files_days = 2
						send_sentinel_alert = true
						notify_days = 1
						schedules = [
							{
						       daily = {
									type = "DAILY"
									start_date = "` + startDate + `"
									skip_holidays = false
									tag = "AccountFilePurge"
									daily_type: "EVERYDAY"
									end_date = null
									execution_times = [
										"01:00"
									]
								}
			                }
						]
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
						resource "`+resourceType+`" "`+resourceName+`" {
							name        = "`+name+`"
							type        = "AccountFilePurge"
							notes       = "mynotes"
							#business_units = []
							additional_attributes = {
								"userVars.additionalProp1" = "val1"
							}
							
							delete_files_days = 2
							send_sentinel_alert = true
							notify_days = 1
							schedules = [
								{
								daily = {
										type = "DAILY"
										start_date = "`+startDate+`"
										skip_holidays = false
										tag = "AccountFilePurge"
										daily_type= "EVERYDAY"
										#end_date = null
										execution_times = [
											"01:00"
										]
									}
								}
							]
						}
						`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			/*// ImportState testing
			{
				ResourceName:      "xmft_staccount.account1",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},*/
			// Update and Read testing
			{
				Config: providerConfig + `
					resource "` + resourceType + `" "` + resourceName + `" {
						provider       = xmft.st1
						name           = "` + name + `"
						notes          = "mynotes"
						#business_units = []
						additional_attributes = {
							"userVars.additionalProp1" = "val2"
							"userVars.additionalProp2" = "val3"
						}
						
						delete_files_days = 10
						send_sentinel_alert = true
						notify_days = 1
						schedules = [
							{
						       daily = {
									type = "DAILY"
									start_date = "` + startDate + `"
									skip_holidays = false
									tag = "AccountFilePurge"
									daily_type: "EVERYDAY"
									#end_date = null
									execution_times = [
										"02:00"
									]
								}
			                }
						]
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
						resource "`+resourceType+`" "`+resourceName+`" {
							name        = "`+name+`"
							#type        = "AccountFilePurge"
							notes       = "mynotes"
							#business_units = []
							additional_attributes = {
								"userVars.additionalProp1" = "val2"
								"userVars.additionalProp2" = "val3"
							}
							
							delete_files_days = 10
							send_sentinel_alert = true
							notify_days = 1
							schedules = [
								{
								daily = {
										type = "DAILY"
										start_date = "`+startDate+`"
										skip_holidays = false
										tag = "AccountFilePurge"
										daily_type= "EVERYDAY"
										#end_date = null
										execution_times = [
											"02:00"
										]
									}
								}
							]
						}
						`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
