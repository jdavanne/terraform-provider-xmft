package provider

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTRouteCompositeResource(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	r := time.Now().Format("-2006-01-02_15-04-05")
	resourceType := "xmft_st_route_composite"
	resourceName := "route_composite1"
	name := "route_composite1" + r
	account := "account1" + r
	route_tmpl1 := "route_template1" + r
	arApp := "app_ar1" + r
	routeSimple := "route_simple1" + r

	localConfig := `
resource "xmft_st_advanced_routing_application" "ar1" {
	provider       = xmft.st1
	name           = "` + arApp + `"
	type           = "AdvancedRouting"
	notes          = "mynotes"
	business_units = []
}

resource "xmft_st_route_template" "template1" {
	provider       = xmft.st1
	name           = "` + route_tmpl1 + `"
	type           = "TEMPLATE"
	description    = "mydescription"
	business_units = []
}

resource "xmft_st_account" "account1" {
	provider    = xmft.st1
	name        = "` + account + `"
	#type        = "user"
	uid         = "1000"	
	gid         = "1000"
	home_folder = "/files/account1"
	user = {
		name =  "` + account + `"
		password_credentials = {
			password = "zouzou"
		}
	}
}

resource "xmft_st_subscription_ar" "sub1" {
	provider         = xmft.st1
	account 		 = xmft_st_account.account1.name
	folder		     = "/folder+` + r + `"
	application 	 = xmft_st_advanced_routing_application.ar1.name	
}

resource "xmft_st_route_simple" "simple1" {
	provider       = xmft.st1
	name           = "` + routeSimple + `"
	#type           = "SIMPLE"
	description    = "mydescription"

	steps = [{
		send_to_partner = {
			transfer_site_expression = "ssh1#!#CVD#!#"
		}
	}]
}
`
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + localConfig + `
					resource "` + resourceType + `" "` + resourceName + `" {
						provider       = xmft.st1
						name           = "` + name + `"
						#type           = "COMPOSITE"
						description    = "mydescription"
						route_template = xmft_st_route_template.template1.id
						subscriptions = [ xmft_st_subscription_ar.sub1.id ]
						account        = xmft_st_account.account1.name

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
						#steps = [ { execute_route = xmft_st_route_simple.simple1.id } ]
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						type           = "COMPOSITE"
						description    = "mydescription"
						managed_by_cg  = false
						#route_template = "`+route_tmpl1+`"
						#account        = "myaccount"

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
				Config: providerConfig + localConfig + `
				resource "` + resourceType + `" "` + resourceName + `" {
					provider       = xmft.st1
					name           = "` + name + `"
					#type           = "COMPOSITE"
					description    = "mydescription3"
					route_template = xmft_st_route_template.template1.id
					account        = xmft_st_account.account1.name
					subscriptions = [ xmft_st_subscription_ar.sub1.id ]

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
					#steps = [ { execute_route = xmft_st_route_simple.simple1.id } ]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						type           = "COMPOSITE"
						description    = "mydescription3"
						managed_by_cg  = false
						#route_template = "`+route_tmpl1+`"
						#account        = "myaccount"

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
