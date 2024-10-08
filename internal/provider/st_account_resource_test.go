package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestStAccountResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
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
		user ={
		  	name =  "` + name + `"
		  	password_credentials = {
				password = "zouzou"
		  	}
				#unknown_key = "unknownValue" #FIXME SHOULD NOT BE ALLOWED !
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
	contact = {
		email="zouzou@example.com"
	}
	bandwidth_limits = {
		inbound_limit = 1000
		outbound_limit = 1000
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
	contact = {
		email="zouzou@example.com"
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
