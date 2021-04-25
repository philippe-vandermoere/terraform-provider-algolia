# algolia_api_key

Use this data source to access information about an existing Algolia index.

## Example Usage

```hcl
data "algolia_index" "example" {
    name = "example"
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Required) The name of the Algolia index.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:
* `name` - The name of the Algolia index.
* `settings` - A settings block as documented below.

---

A `settings` block exports the following:
* `searchable_attributes` - The complete list of attributes used for searching.
* `attributesfor_faceting` - The complete list of attributes that will be used for faceting.
* `unretrievable_attributes` - List of attributes that cannot be retrieved at query time.
* `attributes_to_retrieve` - Gives control over which attributes to retrieve and which not to retrieve.
* `ranking` - Controls the way results are sorted.
* `custom_ranking` - Specifies the custom ranking criterion.
* `replicas` - Creates replicas, exact copies of an index.
* `primary` - If replica, the reference of primary index.
* `max_values_per_facet` - Maximum number of facet values to return for each facet during a regular search.
* `sort_facet_values_by` - Controls how facet values are sorted.
* `attributes_to_highlight` - List of attributes to highlight.
* `attributes_to_snippet` - List of attributes to snippet, with an optional maximum number of words to snippet.
* `snippet_ellipsis_text` - Restrict highlighting and snippeting to items that matched the query.
* `restrict_highlight_and_snippet_arrays` - Restrict highlighting and snippeting to items that matched the query.
* `hits_per_page` - Set the number of hits per page.
* `pagination_limited_to` - Set the maximum number of hits accessible via pagination.
* `min_word_sizefor1_typo` - Minimum number of characters a word in the query string must contain to accept matches with 1 typo.
* `min_word_sizefor2_typos` - Minimum number of characters a word in the query string must contain to accept matches with 2 typos.
* `allow_typos_on_numeric_tokens` - Whether to allow typos on numbers (“numeric tokens”) in the query string.
* `disable_typo_tolerance_on_attributes` - List of attributes on which you want to disable typo tolerance.
* `disable_typo_tolerance_on_words` - List of words on which you want to disable typo tolerance.
* `separators_to_index` - Control which separators are indexed.
* `camel_case_attributes` - List of attributes on which to do a decomposition of camel case words.
* `query_languages` - Sets the languages to be used by language-specific settings and functionalities such as ignorePlurals, removeStopWords, and CJK word-detection.
* `index_languages` - Sets the languages at the index level for language-specific processing such as tokenization and normalization.
* `query_type` - Controls if and how query words are interpreted as prefixes.
* `remove_words_if_no_results` - Selects a strategy to remove words from the query when it doesn’t match any hits.
* `advanced_syntax` - Enables the advanced query syntax.
* `optional_words` - A list of words that should be considered as optional when found in the query.
* `disable_prefix_on_attributes` - List of attributes on which you want to disable prefix matching.
* `disable_exact_on_attributes` - List of attributes on which you want to disable the exact ranking criterion.
* `exact_on_single_word_query` - Controls how the exact ranking criterion is computed when the query contains only one word.
* `alternatives_as_exact` - List of alternatives that should be considered an exact match by the exact ranking criterion.
* `advanced_syntax_features` - Allows you to specify which advanced syntax features are active when ‘advancedSyntax’ is enabled.
* `enable_rules` - Whether Rules should be globally enabled.
* `enable_personalization` - Allows you to specify which advanced syntax features are active when ‘advancedSyntax’ is enabled.
* `numeric_attributes_for_filtering` - List of numeric attributes that can be used as numerical filters.
* `allow_compression_of_integer_array` - Enables compression of large integer arrays.
* `attribute_for_distinct` - Name of the de-duplication attribute to be used with the distinct feature.
* `distinct` - Enables de-duplication or grouping of results. Set 0 to disable, 1 to enable. When N > 1 it enable grouping.
* `replace_synonyms_in_highlight` - Whether to highlight and snippet the original word that matches the synonym or the synonym itself.
* `min_proximity` - Precision of the proximity ranking criterion.
* `max_facet_hits` - Maximum number of facet hits to return during a search for facet values.
* `attribute_criteria_computed_by_min_proximity` - When attribute is ranked above proximity in your ranking formula, proximity is used to select which searchable attribute is matched in the attribute ranking stage.
