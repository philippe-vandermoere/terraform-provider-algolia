# Terraform Provider Algolia

This repository is a Algolia for a [Terraform](https://www.terraform.io) provider.
The provider needs to be configured with the proper credentials before it can be used.

## Usage Example

````hcl-terraform
# Configure the Algolia Provider
provider "algolia" {
    application_id = "{your application id}"
    api_key        = "{your api key}"
}

# Create a Algolia API key
resource "algolia_api_key" "example" {
  acl         = list("search")
  description = "example"
  indexes     = list("example*")
}
````

Full usage [documentation](/docs/README.md).

## Developer

### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.15

### Installation

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command:

````bash
git clone https://github.com/philippe-vandermoere/terraform-provider-algolia
cd terraform-provider-algolia
make install
````
