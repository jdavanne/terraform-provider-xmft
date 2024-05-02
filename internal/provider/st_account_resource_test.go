package provider

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestStAccountResource(t *testing.T) {
	t.Setenv("TF_ACC", "1")
	r := time.Now().Format("-2006-01-02_15-04-05")
	resourceType := "xmft_st_account"
	resourceName := "account1"
	name := "account1" + r

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
				resource "` + resourceType + `" "` + resourceName + `" {
		provider    = xmft.st1
		name        = "` + name + `"
		#type        = "user"
		uid         = "1000"
		gid         = "1000"
		home_folder = "/files/account1"
		user = {
		  	name =  "` + name + `"
		  	password_credentials = {
				password = "zouzou"
		  	}
		}
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
	name        = "`+name+`"
	type        = "user"
	uid         = "1000"
	gid         = "1000"
	home_folder = "/files/account1"
	user = {
		name =  "`+name+`"
		password_credentials = {
			password = "zouzou"
		}
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
	provider    = xmft.st1
	name        = "` + name + `"
	type        = "user"
	uid         = "1000"
	gid         = "1000"
	home_folder = "/files/account1_modified"
	user = {
		name = "` + name + `"
		password_credentials = {
		password = "zouzou"
		}
	}
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
	name        = "`+name+`"
	type        = "user"
	uid         = "1000"
	gid         = "1000"
	home_folder = "/files/account1_modified"
	user = {
		name =  "`+name+`"
		password_credentials = {
			password = "zouzou"
		}
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
