terraform {
    required_version = ">= 0.13"

    required_providers {
        algolia = {
            source  = "philippe-vandermoere/algolia"
        }
    }
}

provider "algolia" {}

data "algolia_index" "example" {
    name = "ctf_products_fr"
}

resource "algolia_api_key" "example" {
    acl         = ["search"]
    description = "example"
    indexes     = [data.algolia_index.example.name]
}

output "api_key" {
    value = algolia_api_key.example.key
}

output "hits_per_page" {
    value = data.algolia_index.example.settings[0]
}
