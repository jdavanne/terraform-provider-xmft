package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTBasicApplicationResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_basic_application"
	resourceName := "basic1"
	name := "Basic1" + r

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
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
						resource "`+resourceType+`" "`+resourceName+`" {
							name        = "`+name+`"
							type        = "Basic"
							notes       = "mynotes"
							#business_units = []
							additional_attributes = {
								"userVars.additionalProp1" = "val1"
							}
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
						notes          = "mynotes2"
						business_units = []
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
						resource "`+resourceType+`" "`+resourceName+`" {
							name        = "`+name+`"
							type        = "Basic"
							notes       = "mynotes2"
							#business_units = []
						}
						`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
