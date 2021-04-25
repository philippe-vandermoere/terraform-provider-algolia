package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIndex(t *testing.T) {
	resourceName := "algolia_index.test"
	resourceNameReplica := "algolia_index.test_replica"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIndex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckNoResourceAttr(resourceName, "searchable_attributes"),
					resource.TestCheckNoResourceAttr(resourceName, "attributes_for_faceting"),
					resource.TestCheckNoResourceAttr(resourceName, "unretrievable_attributes"),
					resource.TestCheckResourceAttr(resourceName, "attributes_to_retrieve.0", "*"),
					resource.TestCheckResourceAttr(resourceName, "ranking.0", "attribute"),
					resource.TestCheckResourceAttr(resourceName, "ranking.1", "custom"),
					resource.TestCheckResourceAttr(resourceName, "ranking.2", "exact"),
					resource.TestCheckResourceAttr(resourceName, "ranking.3", "filters"),
					resource.TestCheckResourceAttr(resourceName, "ranking.4", "geo"),
					resource.TestCheckResourceAttr(resourceName, "ranking.5", "proximity"),
					resource.TestCheckResourceAttr(resourceName, "ranking.6", "typo"),
					resource.TestCheckResourceAttr(resourceName, "ranking.7", "words"),
					resource.TestCheckNoResourceAttr(resourceName, "custom_ranking"),
					resource.TestCheckNoResourceAttr(resourceName, "replicas"),
					resource.TestCheckResourceAttr(resourceName, "max_values_per_facet", "100"),
					resource.TestCheckResourceAttr(resourceName, "sort_facet_values_by", "count"),
					resource.TestCheckResourceAttr(resourceName, "hits_per_page", "20"),
					resource.TestCheckResourceAttr(resourceName, "pagination_limited_to", "1000"),
					resource.TestCheckResourceAttr(resourceName, "attribute_for_distinct", ""),
					resource.TestCheckResourceAttr(resourceName, "distinct", "0"),
				),
			},
			{
				Config: testAccResourceIndexUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckResourceAttr(resourceName, "searchable_attributes.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "attributes_for_faceting.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "unretrievable_attributes.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "attributes_to_retrieve.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "ranking.0", "attribute"),
					resource.TestCheckResourceAttr(resourceName, "ranking.1", "custom"),
					resource.TestCheckResourceAttr(resourceName, "custom_ranking.0", "desc(test)"),
					resource.TestCheckResourceAttr(resourceName, "replicas.0", "test_replica"),
					resource.TestCheckResourceAttr(resourceName, "max_values_per_facet", "50"),
					resource.TestCheckResourceAttr(resourceName, "sort_facet_values_by", "alpha"),
					resource.TestCheckResourceAttr(resourceName, "hits_per_page", "10"),
					resource.TestCheckResourceAttr(resourceName, "pagination_limited_to", "100"),
					resource.TestCheckResourceAttr(resourceName, "attribute_for_distinct", "test"),
					resource.TestCheckResourceAttr(resourceName, "distinct", "2"),
					resource.TestCheckResourceAttr(resourceNameReplica, "name", "test_replica"),
					resource.TestCheckNoResourceAttr(resourceNameReplica, "searchable_attributes"),
					resource.TestCheckNoResourceAttr(resourceNameReplica, "attributes_for_faceting"),
					resource.TestCheckNoResourceAttr(resourceNameReplica, "unretrievable_attributes"),
					resource.TestCheckResourceAttr(resourceNameReplica, "attributes_to_retrieve.0", "*"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.0", "attribute"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.1", "custom"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.2", "exact"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.3", "filters"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.4", "geo"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.5", "proximity"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.6", "typo"),
					resource.TestCheckResourceAttr(resourceNameReplica, "ranking.7", "words"),
					resource.TestCheckNoResourceAttr(resourceNameReplica, "custom_ranking"),
					resource.TestCheckNoResourceAttr(resourceNameReplica, "replicas"),
					resource.TestCheckResourceAttr(resourceNameReplica, "max_values_per_facet", "100"),
					resource.TestCheckResourceAttr(resourceNameReplica, "sort_facet_values_by", "count"),
					resource.TestCheckResourceAttr(resourceNameReplica, "hits_per_page", "20"),
					resource.TestCheckResourceAttr(resourceNameReplica, "pagination_limited_to", "1000"),
					resource.TestCheckResourceAttr(resourceNameReplica, "attribute_for_distinct", ""),
					resource.TestCheckResourceAttr(resourceNameReplica, "distinct", "0"),
				),
			},
		},
	})
}

const testAccResourceIndex = `
resource "algolia_index" "test" {
    name = "test"
}
`

const testAccResourceIndexUpdate = `
resource "algolia_index" "test" {
    name                     = "test"
	searchable_attributes    = ["test"]
    unretrievable_attributes = ["test"]
    attributes_for_faceting  = ["test"]
    attributes_to_retrieve   = ["test"]
    ranking                  = ["attribute", "custom"]
    custom_ranking           = ["desc(test)"]
    hits_per_page            = 10
    max_values_per_facet     = 50
    pagination_limited_to    = 100
    sort_facet_values_by     = "alpha"
	replicas                 = [algolia_index.test_replica.name]
    distinct                 = 2
    attribute_for_distinct   = "test"
}

resource "algolia_index" "test_replica" {
    name  = "test_replica"
}
`
