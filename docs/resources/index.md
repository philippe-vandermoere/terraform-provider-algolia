# algolia_api_key

Provides an Algolia API Key.

## Example Usage

### Create Algolia index

```hcl
resource "algolia_index" "example" {
    name = "example"
}
```

### Create Algolia index with replica

```hcl
resource "algolia_index" "example" {
    name     = "example"
    replicas = [algolia_index.example_replica.name]
}

resource "algolia_index" "example_replica" {
    name = "example_replica"
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Required) The name of Algolia index.
* `searchable_attributes` - (Optional) The complete list of attributes used for searching.
* `attributes_for_faceting` - (Optional) The complete list of attributes that will be used for faceting.
* `unretrievable_attributes` - (Optional) List of attributes that cannot be retrieved at query time.
* `attributes_to_retrieve` - (Optional) Gives control over which attributes to retrieve and which not to retrieve.
* `ranking` - (Optional) Controls how Algolia should [sort](https://www.algolia.com/doc/guides/managing-results/relevance-overview) your results.
* `custom_ranking` - (Optional) Specifies the custom ranking criterion.
* `replicas` - (Optional) Creates replicas, exact copies of an index.
* `max_values_per_facet` - (Optional) Maximum number of facet values to return for each facet during a regular search.
* `sort_facet_values_by` - (Optional) Controls how facet values are sorted.
* `hits_per_page` - (Optional) Set the number of hits per page.
* `pagination_limited_to` - (Optional) Set the maximum number of hits accessible via pagination.

## Import

Algolia index can be imported using the index's name

```shell
terraform import algolia_index.example {index name}
```
