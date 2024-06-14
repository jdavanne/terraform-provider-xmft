package provider

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTTransferSiteSSHResource(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	r := time.Now().Format("-2006-01-02_15-04-05")
	resourceType := "xmft_st_site_ssh"
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
						#type		     = "ssh"
						host		     = "myhost"
						#port		     = "22"
						user_name	     = "myusername"
						password	     = "mypassword"
						#download_folder  = "/download"
						upload_folder    = "/upload"
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						name             = "`+name+`"
						account 		 = "`+account+`"
						type		     = "ssh"
						#protocol	     = "SFTP"
						host		     = "myhost"
						port             = "22"
						user_name	     = "myusername"
						password	     = "mypassword"
						#download_folder  = "/download"
						upload_folder    = "/upload"
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
						host		     = "myhost"
						port             = "8022"
						user_name	     = "myusername"
						password	     = "mypassword2"
						download_folder  = "/download"
						upload_folder    = "/upload"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider       = xmft.st1
						name             = "`+name+`"
						account 		 = "`+account+`"
						type		     = "ssh"
						protocol	     = "ssh"
						host		     = "myhost"
						port             = "8022"
						user_name	     = "myusername"
						password	     = "mypassword2"
						download_folder  = "/download"
						upload_folder    = "/upload"
					}
					`)),
					resource.TestCheckResourceAttrSet(resourceType+"."+resourceName, "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
