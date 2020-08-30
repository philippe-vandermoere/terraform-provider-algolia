package algolia

import (
	"context"
	"strconv"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIndex() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIndexRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"searchable_attributes": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"attributesfor_faceting": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"unretrievable_attributes": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"attributes_to_retrieve": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"ranking": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"custom_ranking": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"replicas": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"primary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_values_per_facet": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sort_facet_values_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes_to_highlight": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"attributes_to_snippet": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"snippet_ellipsis_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"restrict_highlight_and_snippet_arrays": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"hits_per_page": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pagination_limited_to": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_word_sizefor1_typo": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min_word_sizefor2_typos": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allow_typos_on_numeric_tokens": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"disable_typo_tolerance_on_attributes": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"disable_typo_tolerance_on_words": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"separators_to_index": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"camel_case_attributes": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"query_languages": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"index_languages": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"query_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_words_if_no_results": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"advanced_syntax": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"optional_words": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"disable_prefix_on_attributes": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"disable_exact_on_attributes": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"exact_on_single_word_query": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alternatives_as_exact": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"advanced_syntax_features": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"enable_rules": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"enable_personalization": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"numeric_attributes_for_filtering": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
							Computed: true,
						},
						"allow_compression_of_integer_array": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"attribute_for_distinct": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replace_synonyms_in_highlight": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"min_proximity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_facet_hits": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"attribute_criteria_computed_by_min_proximity": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIndexRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	index := m.(*search.Client).InitIndex(d.Get("name").(string))

	exist, err := index.Exists()
	if err != nil {
		return diag.FromErr(err)
	}

	if !exist {
		return diag.Errorf("Algolia API: Unable to find index \"%s\"", d.Get("name").(string))
	}

	settings, err := index.GetSettings()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("settings", flattenIndexSettings(settings)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenIndexSettings(settings search.Settings) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"searchable_attributes":                        settings.SearchableAttributes.Get(),
			"attributesfor_faceting":                       settings.AttributesForFaceting.Get(),
			"unretrievable_attributes":                     settings.UnretrievableAttributes.Get(),
			"attributes_to_retrieve":                       settings.AttributesToRetrieve.Get(),
			"ranking":                                      settings.Ranking.Get(),
			"custom_ranking":                               settings.CustomRanking.Get(),
			"replicas":                                     settings.Replicas.Get(),
			"primary":                                      settings.Primary.Get(),
			"max_values_per_facet":                         settings.MaxValuesPerFacet.Get(),
			"sort_facet_values_by":                         settings.SortFacetValuesBy.Get(),
			"attributes_to_highlight":                      settings.AttributesToHighlight.Get(),
			"attributes_to_snippet":                        settings.AttributesToSnippet.Get(),
			"snippet_ellipsis_text":                        settings.SnippetEllipsisText.Get(),
			"restrict_highlight_and_snippet_arrays":        settings.RestrictHighlightAndSnippetArrays.Get(),
			"hits_per_page":                                settings.HitsPerPage.Get(),
			"pagination_limited_to":                        settings.PaginationLimitedTo.Get(),
			"min_word_sizefor1_typo":                       settings.MinWordSizefor1Typo.Get(),
			"min_word_sizefor2_typos":                      settings.MinWordSizefor2Typos.Get(),
			"allow_typos_on_numeric_tokens":                settings.AllowTyposOnNumericTokens.Get(),
			"disable_typo_tolerance_on_attributes":         settings.DisableTypoToleranceOnAttributes.Get(),
			"disable_typo_tolerance_on_words":              settings.DisableTypoToleranceOnWords.Get(),
			"separators_to_index":                          settings.SeparatorsToIndex.Get(),
			"camel_case_attributes":                        settings.CamelCaseAttributes.Get(),
			"query_languages":                              settings.QueryLanguages.Get(),
			"index_languages":                              settings.IndexLanguages.Get(),
			"query_type":                                   settings.QueryType.Get(),
			"remove_words_if_no_results":                   settings.RemoveWordsIfNoResults.Get(),
			"advanced_syntax":                              settings.AdvancedSyntax.Get(),
			"optional_words":                               settings.OptionalWords.Get(),
			"disable_prefix_on_attributes":                 settings.DisablePrefixOnAttributes.Get(),
			"disable_exact_on_attributes":                  settings.DisableExactOnAttributes.Get(),
			"exact_on_single_word_query":                   settings.ExactOnSingleWordQuery.Get(),
			"alternatives_as_exact":                        settings.AlternativesAsExact.Get(),
			"advanced_syntax_features":                     settings.AdvancedSyntaxFeatures.Get(),
			"enable_rules":                                 settings.EnableRules.Get(),
			"enable_personalization":                       settings.EnablePersonalization.Get(),
			"numeric_attributes_for_filtering":             settings.NumericAttributesForFiltering.Get(),
			"allow_compression_of_integer_array":           settings.AllowCompressionOfIntegerArray.Get(),
			"attribute_for_distinct":                       settings.AttributeForDistinct.Get(),
			"replace_synonyms_in_highlight":                settings.ReplaceSynonymsInHighlight.Get(),
			"min_proximity":                                settings.MinProximity.Get(),
			"max_facet_hits":                               settings.MaxFacetHits.Get(),
			"attribute_criteria_computed_by_min_proximity": settings.AttributeCriteriaComputedByMinProximity.Get(),
		},
	}
}
