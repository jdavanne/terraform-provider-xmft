package provider

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTRouteTemplateResource(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	r := time.Now().Format("-2006-01-02_15-04-05")
	resourceType := "xmft_st_route_template"
	resourceName := "route_template1"
	name := "route_template1" + r

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "` + resourceType + `" "` + resourceName + `" {
						provider       = xmft.st1
						name           = "` + name + `"
						type           = "TEMPLATE"
						description    = "mydescription"
						managed_by_cg  = false
						#route_template = ""
						#account        = "myaccount"
						business_units = []
						condition_type = "MATCH_ALL"
						condition      = "mycondition"
						failure_email_notification = false
						failure_email_template     = ""
						failure_email_name         = "myname"
						success_email_notification = false
						success_email_template     = ""
						success_email_name         = "myname"
						triggering_email_notification = false
						triggering_email_template     = ""
						triggering_email_name         = "myname"
						#subscriptions = []
						#additional_attributes = {
						#	additionalProp1 = "myprop1"
						#	additionalProp2 = "myprop2"
						#	additionalProp3 = "myprop3"
						#}
						#step_statuses = [
						#	{
						#		id     = "mystatus"
						#		step_id = "mystep"
						#		status  = "ENABLED"
						#	}
						#]
						#steps = []
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						type           = "TEMPLATE"
						description    = "mydescription"
						managed_by_cg  = false
						#route_template = "myroute"
						#account        = "myaccount"
						#business_units = []
						condition_type = "MATCH_ALL"
						condition      = "mycondition"
						failure_email_notification = false
						failure_email_template     = ""
						failure_email_name         = "myname"
						success_email_notification = false
						success_email_template     = ""
						success_email_name         = "myname"
						triggering_email_notification = false
						triggering_email_template     = ""
						triggering_email_name         = "myname"
						#subscriptions = []
						#additional_attributes = {
						#	additionalProp1 = "myprop1"
						#	additionalProp2 = "myprop2"
						#	additionalProp3 = "myprop3"
						#}
						#step_statuses = [
						#	{
						#		id     = "mystatus"
						#		step_id = "mystep"
						#		status  = "ENABLED"
						#	}
						#]
						#steps = []
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
					type           = "TEMPLATE"
					description    = "mydescription3"
					managed_by_cg  = false
					#route_template = ""
					#account        = "myaccount"
					business_units = []
					condition_type = "MATCH_ALL"
					condition      = "mycondition"
					failure_email_notification = false
					failure_email_template     = ""
					failure_email_name         = "myname"
					success_email_notification = false
					success_email_template     = ""
					success_email_name         = "myname"
					triggering_email_notification = false
					triggering_email_template     = ""
					triggering_email_name         = "myname"
					#subscriptions = []
					#additional_attributes = {
					#	additionalProp1 = "myprop1"
					#	additionalProp2 = "myprop2"
					#	additionalProp3 = "myprop3"
					#}
					#step_statuses = [
					#	{
					#		id     = "mystatus"
					#		step_id = "mystep"
					#		status  = "ENABLED"
					#	}
					#]
					#steps = []
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						type           = "TEMPLATE"
						description    = "mydescription3"
						managed_by_cg  = false
						#route_template = "myroute"
						#account        = "myaccount"
						#business_units = []
						condition_type = "MATCH_ALL"
						condition      = "mycondition"
						failure_email_notification = false
						failure_email_template     = ""
						failure_email_name         = "myname"
						success_email_notification = false
						success_email_template     = ""
						success_email_name         = "myname"
						triggering_email_notification = false
						triggering_email_template     = ""
						triggering_email_name         = "myname"
						#subscriptions = []
						#additional_attributes = {
						#	additionalProp1 = "myprop1"
						#	additionalProp2 = "myprop2"
						#	additionalProp3 = "myprop3"
						#}
						#step_statuses = [
						#	{
						#		id     = "mystatus"
						#		step_id = "mystep"
						#		status  = "ENABLED"
						#	}
						#]
						#steps = []
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
