package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceApiKey(t *testing.T) {
	resourceName := "algolia_api_key.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApiKey,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "key", regexp.MustCompile("^.{1,}$")),
					resource.TestCheckResourceAttr(resourceName, "acl.0", "search"),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckNoResourceAttr(resourceName, "indexes"),
					resource.TestCheckResourceAttr(resourceName, "max_queries_per_ip_per_hour", "15000"),
					resource.TestCheckResourceAttr(resourceName, "max_hits_per_query", "0"),
					resource.TestCheckNoResourceAttr(resourceName, "referers"),
					resource.TestCheckResourceAttr(resourceName, "validity", "0"),
				),
			},
			{
				Config: testAccResourceApiKeyUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "update"),
					resource.TestCheckResourceAttr(resourceName, "acl.0", "addObject"),
					resource.TestCheckResourceAttr(resourceName, "acl.1", "deleteIndex"),
					resource.TestCheckResourceAttr(resourceName, "acl.2", "deleteObject"),
					resource.TestCheckResourceAttr(resourceName, "indexes.0", "test1"),
					resource.TestCheckResourceAttr(resourceName, "indexes.1", "test1_price_asc"),
					resource.TestCheckResourceAttr(resourceName, "indexes.2", "test1_price_desc"),
					resource.TestCheckResourceAttr(resourceName, "max_queries_per_ip_per_hour", "50"),
					resource.TestCheckResourceAttr(resourceName, "max_hits_per_query", "100"),
					resource.TestCheckResourceAttr(resourceName, "referers.0", "https://test1.com/"),
					resource.TestCheckResourceAttr(resourceName, "referers.1", "https://test2.com/"),
					// resource.TestCheckResourceAttr(resourceName2, "validity", "3600"), to be refactor
				),
			},
		},
	})
}

const testAccResourceApiKey = `
resource "algolia_api_key" "test" {
    acl         = ["search"]
    description = "test"
}
`

const testAccResourceApiKeyUpdate = `
resource "algolia_api_key" "test" {
    acl = [
		"addObject",
		"deleteIndex",
		"deleteObject",
	]

	indexes = [
		"test1",
		"test1_price_asc",
		"test1_price_desc",
	]

	referers = [
		"https://test1.com/",
		"https://test2.com/",
	]

    description                 = "update"
	max_queries_per_ip_per_hour = 50
	max_hits_per_query          = 100
	//validity                    = 3600
}
`
