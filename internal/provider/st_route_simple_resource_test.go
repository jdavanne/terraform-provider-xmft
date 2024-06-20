package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTRouteSimpleResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_route_simple"
	resourceName := "route_simple1"
	name := "route_simple1" + r
	// account := "account1" + r
	// route_tmpl1 := "route_template1" + r
	// arApp := "app_ar1" + r

	localConfig := ``

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
						#zouzou= ""
						#condition_type = "MATCH_ALL"
						condition      = ""
						steps = [
							{
								send_to_partner = { 
									transfer_site_expression = "ssh1#!#CVD#!#"
								}
							},
							{
								compress = {
									
								}
							},
							{
								decompress = {
									
								}
							},
							{
								publish_to_account = {
									target_account_expression = "zou"
								}
							},
							{
								pull_from_partner = {
									target_account_expression = "zou"
									transfer_site_expression = "ssh1#!#CVD#!#"
								}
							},
							{
								rename = {
									output_file_name = "zou"
								}
							},

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
						condition      = "zouzou2"
						steps = [
							{
								send_to_partner = { 
									type                     = "SendToPartner"
									status                   = "ENABLED"
									transfer_site_expression = "ssh2#!#CVD#!#"
								}
							}
						]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name           = "`+name+`"
						description    = "mydescription"

						condition_type = "EL"
						condition      = "zouzou2"
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
