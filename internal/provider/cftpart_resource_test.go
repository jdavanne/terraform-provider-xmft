package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCFTPARTResource(t *testing.T) {
	t.Parallel()
	// r := time.Now().Format("-2006-01-02_15-04-05") + "-" + fmt.Sprint(time.Now().UnixNano())
	name := "cft2"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
		resource "xmft_cftpart" "test" {
				provider = xmft.cft1
				name     = "cft2"
				prot     = "PESIT"
				sap      = "1788"
			  
				nrpart  = "CFT2"
				nrpassw = "cft2*"
				nspart  = "CFT1"
				nspassw = "cft1*"
			  
				tcp = [{
				  id     = "1"
				  cnxout = "100"
				  host   = "cft2"
				}]
		}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("xmft_cftpart.test", "name", name),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "prot", "PESIT"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "sap", "1788"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nrpart", "CFT2"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nrpassw", "cft2*"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nspart", "CFT1"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nspassw", "cft1*"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.#", "1"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.0.id", "1"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.0.cnxout", "100"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.0.host", "cft2"),
					resource.TestCheckResourceAttrSet("xmft_cftpart.test", "last_updated"),
				),
			},
			/*// ImportState testing
			{
				ResourceName:      "xmft_cftpart.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},*/
			// Update and Read testing
			{
				Config: providerConfig + `
				resource "xmft_cftpart" "test" {
					provider = xmft.cft1
					name     = "cft2"
					prot     = "PESIT"
					sap      = "1788"
				  
					nrpart  = "CFT2"
					nrpassw = "cft3*"
					nspart  = "CFT1"
					nspassw = "cft1*"
				  
					tcp = [{
					  id     = "1"
					  cnxout = "10"
					  host   = "cft3"
					}]
			}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("xmft_cftpart.test", "name", name),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "prot", "PESIT"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "sap", "1788"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nrpart", "CFT2"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nrpassw", "cft3*"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nspart", "CFT1"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "nspassw", "cft1*"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.#", "1"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.0.id", "1"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.0.cnxout", "10"),
					resource.TestCheckResourceAttr("xmft_cftpart.test", "tcp.0.host", "cft3"),
					resource.TestCheckResourceAttrSet("xmft_cftpart.test", "last_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
