package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTSubscriptionBasicResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_subscription_basic"
	resourceName := "subbasic1"
	name := "subbasic1" + r

	account := "account1" + r
	app := "basic1" + r

	localConfig := `
resource "xmft_st_basic_application" "basic1" {
	provider       = xmft.st1
	name           = "` + app + `"
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
	`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + localConfig + `
					resource "` + resourceType + `" "` + resourceName + `" {
						provider         = xmft.st1
						#name             = "` + name + `"
						account 		 = xmft_st_account.account1.name
						folder		     = "/folder"
						application 	 = xmft_st_basic_application.basic1.name	

						depends_on = [xmft_st_basic_application.basic1]
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						#name             = "`+name+`"
						account 		 = "`+account+`"
						folder		     = "/folder"
						application 	 = "`+app+`"
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
						provider         = xmft.st1
						#name             = "` + name + `"
						account 		 = xmft_st_account.account1.name
						folder		     = "/folder2"
						application 	 = xmft_st_basic_application.basic1.name	

						depends_on = [xmft_st_basic_application.basic1]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						#name             = "`+name+`"
						account 		 = "`+account+`"
						folder		     = "/folder2"
						application 	 = "`+app+`"
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
