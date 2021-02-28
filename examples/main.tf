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
    name = "test"
}

resource "algolia_api_key" "example" {
    acl         = ["search"]
    description = "example"
    indexes     = [data.algolia_index.example.name]
}

output "api_key" {
    value = algolia_api_key.example.key
}
