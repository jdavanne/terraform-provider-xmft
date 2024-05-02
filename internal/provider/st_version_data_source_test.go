package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSTVersionDataSource(t *testing.T) {
	// t.Setenv("TF_ACC", "1")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "xmft_st_version" "test" { 
					provider = xmft.st1 
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.xmft_st_version.test", "server_type", "ST-Core-Server"),
					resource.TestCheckResourceAttr("data.xmft_st_version.test", "version", "5.5-20240328"),
					resource.TestCheckResourceAttr("data.xmft_st_version.test", "build", "${env.BUILD_NUMBER}"),
					resource.TestCheckResourceAttr("data.xmft_st_version.test", "os", "Linux"),
					resource.TestCheckResourceAttr("data.xmft_st_version.test", "docker_mode", "true"),
					// resource.TestCheckResourceAttr("data.xmft_about.test", "about.multinode_enabled", "true"),
					// resource.TestCheckResourceAttr("data.xmft_about.test", "about.system", "unix"),
				),
			},
		},
	})
}
