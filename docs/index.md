# Algolia Provider

The Algolia provider is used to interact with the many resources supported by Algolia. The provider needs to be configured with the proper credentials before it can be used.
Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Set the variable value in *.tfvars file
# or using -var="application_id =..." -var="api_key=..." CLI option
variable "application_id" {}
variable "api_key" {}

# Configure the Algolia Provider
provider "algolia" {
    application_id = var.application_id
    api_key        = var.api_key
}

# Create a Algolia API key
resource "algolia_api_key" "example" {
  acl         = ["search"]
  description = "example"
  indexes     = ["example"]
}
```

## Argument Reference

The following arguments are supported:
* `application_id` - (Required) This is your unique application identifier. It's used to identify you when using Algolia's API (Defaults to the value of the `ALGOLIA_APPLICATION_ID` environment variable).
* `api_key` - (Required) This is the Algolia ADMIN API key, You can use it to manage your API keys or create, update and DELETE your indices (Defaults to the value of the `ALGOLIA_API_KEY` environment variable).
