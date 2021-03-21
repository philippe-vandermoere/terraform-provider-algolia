terraform {
    required_version = ">= 0.13"

    required_providers {
        algolia = {
            source  = "philippe-vandermoere/algolia"
        }
    }
}

resource "algolia_index" "example" {
    name = "example"
}

resource "algolia_index" "example2" {
    name                     = "example2"
    searchable_attributes    = []
    unretrievable_attributes = []
    attributes_for_faceting  = []
    attributes_to_retrieve   = ["*"]
    ranking                  = ["attribute", "custom", "exact", "filters", "geo", "proximity", "typo", "words"]
    custom_ranking           = []
    hits_per_page            = 20
    max_values_per_facet     = 100
    pagination_limited_to    = 1000
    sort_facet_values_by     = "count"
}

resource "algolia_index" "example3" {
    name                     = "example3"
    searchable_attributes    = ["example"]
    unretrievable_attributes = ["unretrievable"]
    attributes_for_faceting  = ["example"]
    attributes_to_retrieve   = ["example"]
    ranking                  = ["attribute", "custom", "exact", "filters", "geo", "proximity", "typo", "words"]
    custom_ranking           = ["desc(example)"]
    hits_per_page            = 10
    max_values_per_facet     = 50
    pagination_limited_to    = 100
    sort_facet_values_by     = "alpha"
    replicas                 = [algolia_index.example3_replica.name]
}

resource "algolia_index" "example3_replica" {
    name           = "example3_replica"
    custom_ranking = ["asc(example)"]
}

resource "algolia_api_key" "example" {
    acl         = ["search"]
    description = "example"
    indexes = [
        algolia_index.example.name,
        algolia_index.example2.name,
        algolia_index.example3.name,
        algolia_index.example3_replica.name,
    ]
}

output "api_key" {
    value = algolia_api_key.example.key
}

output "algolia_index_example" {
    value = algolia_index.example
}

output "algolia_index_example2" {
    value = algolia_index.example2
}

output "algolia_index_example3" {
    value = algolia_index.example3
}

output "algolia_index_example3_replica" {
    value = algolia_index.example3_replica
}
