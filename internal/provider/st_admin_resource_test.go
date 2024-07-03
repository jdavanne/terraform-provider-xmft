package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewSTAdminResource(t *testing.T) {
	t.Parallel()
	r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	resourceType := "xmft_st_admin"
	resourceName := "admin1"
	name := "admin1" + r
	role := "role1" + r

	localConfig := `
	resource xmft_st_admin_role role1 {
		provider         = xmft.st1
		name             = "` + role + `"
		is_limited       = true
		is_bounce_allowed = true	
		menus             = [ "Server Log" ]
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
						role_name		 = "` + role + `"
						is_limited       = true
						parent           = "admin"
						password_credentials = {
							password = "mypassword"
						}
							
						administrator_rights = {
							can_read_only = false
							is_maker = true
							can_create_users = true
							can_update_users = true
						}
						depends_on = [xmft_st_admin_role.role1]
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						name             = "`+name+`"
						role_name		 = "`+role+`"
						is_limited       = true
						password_credentials = {
							password = "mypassword"
						}
						administrator_rights = {
							can_read_only = false
							is_maker = true
							can_create_users = true
							can_update_users = true
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
				// API, therefore there is no value for it during import
				ImportStateVerifyIgnore: []string{"last_updated"},
			},*/
			// Update and Read testing
			{
				Config: providerConfig + localConfig + `
				resource "` + resourceType + `" "` + resourceName + `" {
					    provider         = xmft.st1
						name             = "` + name + `"
						role_name		 = "` + role + `"
						is_limited       = true
						parent           = "admin"
						password_credentials = {
							password = "mypassword2"
						}
						administrator_rights = {
							can_read_only = true
						}
						depends_on = [xmft_st_admin_role.role1]
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					checkResources([]byte(`
					resource "`+resourceType+`" "`+resourceName+`" {
						#provider        = xmft.st1
						name             = "`+name+`"
						role_name		 = "`+role+`"
						is_limited       = true
						parent           = "admin"
						password_credentials = {
							password = "mypassword2"
						}
						administrator_rights = {
							can_read_only = true
							is_maker = false
							can_create_users = false
							can_update_users = false
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
