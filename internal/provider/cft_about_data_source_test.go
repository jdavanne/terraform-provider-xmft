package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCFTAboutDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "xmft_cft_about" "test" { 
					provider = xmft.cft1 
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.xmft_cft_about.test", "instance_id", "cft_docker_01"),
					// resource.TestCheckResourceAttr("data.xmft_about.test", "name", ""),
					resource.TestCheckResourceAttr("data.xmft_cft_about.test", "version", "3.10"),
					resource.TestCheckResourceAttr("data.xmft_cft_about.test", "level", "2209"),
					resource.TestCheckResourceAttr("data.xmft_cft_about.test", "multinode_enabled", "false"),
					resource.TestCheckResourceAttr("data.xmft_cft_about.test", "system", "unix"),
				),
			},
		},
	})
}
