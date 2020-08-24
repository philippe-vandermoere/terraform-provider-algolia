terraform {
    required_version = ">= 0.13"

    required_providers {
        algolia = {
            source  = "philippe-vandermoere/algolia"
            version = "0.0.1"
        }
    }
}

provider "algolia" {}

resource "algolia_api_key" "example" {
    acl         = ["search"]
    description = "example"
    indexes     = ["example"]
}

output "api_key" {
    value = algolia_api_key.example.id
}
