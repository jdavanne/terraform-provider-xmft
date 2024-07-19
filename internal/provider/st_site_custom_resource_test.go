package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTTransferSiteCustomResource(t *testing.T) {
	t.Parallel()

	providerConfig := `
	  provider "xmft" {
		product = "st"
		alias = "st1"
		username = "admin"
		password = "admin"
		host     = "https://10.128.57.198:8444"
		additional_attributes = {
		Env = "test"
		}
}
	`

	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_site_custom"
	resourceName := "ssh1"
	name := "ssh1" + r

	account := "account1" + r

	localConfig := `
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
						name             = "` + name + `"
						account 		 = xmft_st_account.account1.name
						#type		     = "ExternalPersistedCustomSite"
						protocol	     = "s3"
						custom_properties = {
       						s3Bucket = "storage1-s3bucket"
						}	

					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						name             = "`+name+`"
						account 		 = "`+account+`"
						type		     = "ExternalPersistedCustomSite"
						protocol	     = "s3"
						custom_properties = {
       						s3Bucket = "storage1-s3bucket"
						}	
						custom_properties_all = {
       						s3Bucket = "storage1-s3bucket"
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
				Config: providerConfig + localConfig + `
				resource "` + resourceType + `" "` + resourceName + `" {
					    provider         = xmft.st1
						name             = "` + name + `"
						account 		 = xmft_st_account.account1.name
						#type		     = "ExternalPersistedCustomSite"
						protocol	     = "s3"
						custom_properties = {
       						s3Bucket = "storage1-s3bucketzouzou"
						}	
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name             = "`+name+`"
						account 		 = "`+account+`"
						type		     = "ExternalPersistedCustomSite"
						protocol	     = "s3"
						custom_properties = {
       						s3Bucket = "storage1-s3bucketzouzou"
						}	
						custom_properties_all = {
       						s3Bucket = "storage1-s3bucketzouzou"
						}	
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
