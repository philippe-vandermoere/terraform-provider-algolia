# algolia_api_key

Provides an Algolia API Key.

## Example Usage

```hcl
# Create a Algolia API key
resource "algolia_api_key" "example" {
  acl         = ["search"]
  description = "example"
  indexes     = ["example"]
}

# Output the API Key
output "api_key" {
  value = algolia_api_key.example.key
}
```

## Argument Reference

The following arguments are supported:
* `acl` - (Required) Specify the list of permissions associated to the key, the possible acls:
  * `search` Allows search.
  * `browse` Allows retrieval of all index contents via the browse API.
  * `addObject` Allows adding/updating an object in the index. (Copying/moving indices are also allowed with this permission).
  * `deleteObject` Allows deleting an existing object.
  * `deleteIndex` Allows deleting index content.
  * `settings` allows getting index settings.
  * `editSettings` Allows changing index settings.
  * `analytics` Allows retrieval of analytics through the analytics API.
  * `recommendation` Allows usage of the Personalization dashboard and the Recommendation API.
  * `listIndexes` Allows listing all accessible indices.
  * `logs` Allows getting the logs.
  * `seeUnretrievableAttributes` Disables the unretrievable Attributes feature for all operations returning records.
* `description` - (Required) Specify a description to describe where the key is used.
* `indexes` - (Optional) Specify the list of targeted indices. You can target all indices starting with a prefix or ending with a suffix using the ‘*’ character. For example, “dev_*” matches all indices starting with “dev_” and “*_dev” matches all indices ending with “_dev”.
* `max_queries_per_ip_peer_hour` - (Deprecated) Use `max_queries_per_ip_per_hour` instead.
* `max_queries_per_ip_per_hour` - (Optional) Specify the maximum number of API calls allowed from an IP address per hour. Each time an API call is performed with this key, a check is performed. If the IP at the source of the call did more than this number of calls in the last hour, a 429 code is returned. This parameter can be used to protect you from attempts at retrieving your entire index contents by massively querying the index.
* `max_hits_per_query` - (Optional) Specify the maximum number of hits this API key can retrieve in one call. This parameter can be used to protect you from attempts at retrieving your entire index contents by massively querying the index.
* `referers` - (Optional) Specify the list of query parameters. You can force the query parameters for a query using the url string format. Example: “typoTolerance=strict&ignorePlurals=false”.
* `validity` - (Optional) How long this API key is valid, in seconds. A value of 0 means the API key does not expire.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `key` - The API Key

## Import

Algolia API Key can be imported using the `key`

```shell
terraform import algolia_api_key.example {my algolia api key}
```
