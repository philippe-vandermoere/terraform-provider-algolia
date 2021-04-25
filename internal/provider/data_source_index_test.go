package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIndex(t *testing.T) {
	resourceName := "data.algolia_index.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIndexCreate,
			},
			{
				Config: testAccDataSourceIndex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "data_test"),
					// todo test settings
				),
			},
		},
	})
}

const testAccDataSourceIndexCreate = `
resource "algolia_index" "test" {
    name = "data_test"
}
`

const testAccDataSourceIndex = `
resource "algolia_index" "test" {
    name = "data_test"
}

data "algolia_index" "test" {
    name = "data_test"
}
`
