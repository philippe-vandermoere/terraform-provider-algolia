# Algolia Provider

The Algolia provider is used to interact with the many resources supported by Algolia. The provider needs to be configured with the proper credentials before it can be used.

## Authentication

### Static credentials

Warning: Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.

Static credentials can be provided by adding an application_id and api_key in-line in the Algolia provider block:

usage:

````hcl-terraform
provider "algolia" {
    application_id = "{your application id}"
    api_key        = "{your api key}"
}
````

### Environment variables

You can provide your credentials via the ALGOLIA_APPLICATION_ID and ALGOLIA_API_KEY, environment variables, representing your Algolia Application Id and Algolia Api Key, respectively.

usage:

````bash
export ALGOLIA_APPLICATION_ID="{your application id}"
export ALGOLIA_API_KEY="{your api key}"
````

````hcl-terraform
provider "algolia" {}
````

## Resources

- [api_key](/docs/RESOURCES_API_KEY.md)

## Data Sources
