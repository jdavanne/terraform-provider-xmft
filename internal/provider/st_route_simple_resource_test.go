package provider

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTRouteSimpleResource(t *testing.T) {
	// t.Setenv("TF_ACC", "1")
	r := time.Now().Format("-2006-01-02_15-04-05")
	resourceType := "xmft_st_route_simple"
	resourceName := "route_simple1"
	name := "route_simple1" + r
	account := "account1" + r
	// route_tmpl1 := "route_template1" + r
	arApp := "app_ar1" + r

	localConfig := `
resource "xmft_st_advanced_routing_application" "ar1" {
	provider       = xmft.st1
	name           = "` + arApp + `"
	type           = "AdvancedRouting"
	notes          = "mynotes"
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
						#type           = "SIMPLE"
						description    = "mydescription"
						#route_template = xmft_st_route_template.template1.id
						#subscriptions = [ xmft_st_subscription_ar.sub1.id ]
						#account        = xmft_st_account.account1.name

						#condition_type = "MATCH_ALL"
						condition      = ""
						steps = [{
							type                     = "SendToPartner"
							status                   = "ENABLED"
							transfer_site_expression = "ssh1#!#CVD#!#"
							}
						]
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						#type           = "SIMPLE"
						description    = "mydescription"

						condition_type = "ALWAYS"
						condition      = ""

						#steps = [{
						#	type                     = "SendToPartner"
						#	status                   = "ENABLED"
						#	transfer_site_expression = "ssh1"
						#	}
						#  ]
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
						description    = "mydescription"

						condition_type = "EL"
						condition      = "zouzou"
						#steps = [{
						#	type                     = "SendToPartner"
						#	status                   = "ENABLED"
						#	transfer_site_expression = "ssh1"
						#	}
						#  ]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						description    = "mydescription"

						condition_type = "EL"
						condition      = "zouzou"
						#steps = [{
						#	type                     = "SendToPartner"
						#	status                   = "ENABLED"
						#	transfer_site_expression = "ssh1"
						#	}
						#  ]
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
