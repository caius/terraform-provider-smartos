package smartos

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccSmartosImage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSmartosImageBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.smartos_image.base_64_lts", "id"),
					resource.TestCheckResourceAttrSet("data.smartos_image.base_64_lts", "name"),
				),
			},
		},
	})
}

var testAccSmartosImageBasic = `
data "smartos_image" "base_64_lts" {}
`
