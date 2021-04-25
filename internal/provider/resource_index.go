package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIndex() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIndexCreate,
		ReadContext:   resourceIndexRead,
		UpdateContext: resourceIndexUpdate,
		DeleteContext: resourceIndexDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importIndexState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"searchable_attributes": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"attributes_for_faceting": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"unretrievable_attributes": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"attributes_to_retrieve": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"ranking": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"custom_ranking": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"replicas": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"max_values_per_facet": {
				Type:     schema.TypeInt,
				Set:      schema.HashString,
				Optional: true,
				Default:  100,
			},
			"sort_facet_values_by": {
				Type:     schema.TypeString,
				Set:      schema.HashString,
				Optional: true,
				Default:  "count",
			},
			"hits_per_page": {
				Type:     schema.TypeInt,
				Set:      schema.HashString,
				Optional: true,
				Default:  20,
			},
			"pagination_limited_to": {
				Type:     schema.TypeInt,
				Set:      schema.HashString,
				Optional: true,
				Default:  1000,
			},
			"attribute_for_distinct": {
				Type:     schema.TypeString,
				Set:      schema.HashString,
				Optional: true,
			},
			"distinct": {
				Type:     schema.TypeInt,
				Set:      schema.HashString,
				Optional: true,
			},
		},
	}
}

func resourceIndexCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	index := m.(*apiClient).algolia.InitIndex(d.Get("name").(string))
	exist, err := index.Exists()
	if err != nil {
		return diag.FromErr(err)
	} else if exist {
		return diag.Errorf("algolia api: index %s already exist", d.Get("name").(string))
	}

	res, err := index.SetSettings(getIndexSettings(d))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Wait(); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return resourceIndexRead(ctx, d, m)
}

func resourceIndexRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := refreshIndexState(d, m); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceIndexUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := m.(*apiClient).algolia.InitIndex(d.Get("name").(string)).SetSettings(getIndexSettings(d))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Wait(); err != nil {
		return diag.FromErr(err)
	}

	return resourceIndexRead(ctx, d, m)
}

func resourceIndexDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	res, err := m.(*apiClient).algolia.InitIndex(d.Get("name").(string)).Delete()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Wait(); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func importIndexState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("name", d.Id()); err != nil {
		return nil, err
	}

	if err := refreshIndexState(d, m); err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return []*schema.ResourceData{d}, nil
}

func refreshIndexState(d *schema.ResourceData, m interface{}) error {
	index := m.(*apiClient).algolia.InitIndex(d.Get("name").(string))
	exist, err := index.Exists()
	if err != nil {
		d.SetId("")
		return err
	} else if !exist {
		d.SetId("")
		return fmt.Errorf("algolia api: index %s doesn't exist", d.Get("name").(string))
	}

	settings, err := index.GetSettings()
	if err != nil {
		return err
	}

	if err := d.Set("searchable_attributes", settings.SearchableAttributes.Get()); err != nil {
		return err
	}

	if err := d.Set("attributes_for_faceting", settings.AttributesForFaceting.Get()); err != nil {
		return err
	}

	if err := d.Set("unretrievable_attributes", settings.UnretrievableAttributes.Get()); err != nil {
		return err
	}

	if err := d.Set("attributes_to_retrieve", settings.AttributesToRetrieve.Get()); err != nil {
		return err
	}

	if err := d.Set("ranking", settings.Ranking.Get()); err != nil {
		return err
	}

	if err := d.Set("custom_ranking", settings.CustomRanking.Get()); err != nil {
		return err
	}

	if err := d.Set("replicas", settings.Replicas.Get()); err != nil {
		return err
	}

	if err := d.Set("max_values_per_facet", settings.MaxValuesPerFacet.Get()); err != nil {
		return err
	}

	if err := d.Set("sort_facet_values_by", settings.SortFacetValuesBy.Get()); err != nil {
		return err
	}

	if err := d.Set("hits_per_page", settings.HitsPerPage.Get()); err != nil {
		return err
	}

	if err := d.Set("pagination_limited_to", settings.PaginationLimitedTo.Get()); err != nil {
		return err
	}

	if err := d.Set("attribute_for_distinct", settings.AttributeForDistinct.Get()); err != nil {
		return err
	}

	_, distinct := settings.Distinct.Get()
	if err := d.Set("distinct", distinct); err != nil {
		return err
	}

	return nil
}

func getIndexSettings(d *schema.ResourceData) search.Settings {
	var searchableAttributes []string
	if value, ok := d.GetOk("searchable_attributes"); ok {
		for _, v := range value.(*schema.Set).List() {
			searchableAttributes = append(searchableAttributes, v.(string))
		}
	}

	var attributesForFaceting []string
	if value, ok := d.GetOk("attributes_for_faceting"); ok {
		for _, v := range value.(*schema.Set).List() {
			attributesForFaceting = append(attributesForFaceting, v.(string))
		}
	}

	var unretrievableAttributes []string
	if value, ok := d.GetOk("unretrievable_attributes"); ok {
		for _, v := range value.(*schema.Set).List() {
			unretrievableAttributes = append(unretrievableAttributes, v.(string))
		}
	}

	var attributesToRetrieve []string
	if value, ok := d.GetOk("attributes_to_retrieve"); ok {
		for _, v := range value.(*schema.Set).List() {
			attributesToRetrieve = append(attributesToRetrieve, v.(string))
		}
	}

	var ranking []string
	if value, ok := d.GetOk("ranking"); ok {
		for _, v := range value.(*schema.Set).List() {
			ranking = append(ranking, v.(string))
		}
	}

	var customRanking []string
	if value, ok := d.GetOk("custom_ranking"); ok {
		for _, v := range value.(*schema.Set).List() {
			customRanking = append(customRanking, v.(string))
		}
	}

	var replicas []string
	if value, ok := d.GetOk("replicas"); ok {
		for _, v := range value.(*schema.Set).List() {
			replicas = append(replicas, v.(string))
		}
	}

	return search.Settings{
		SearchableAttributes:    opt.SearchableAttributes(searchableAttributes...),
		AttributesForFaceting:   opt.AttributesForFaceting(attributesForFaceting...),
		UnretrievableAttributes: opt.UnretrievableAttributes(unretrievableAttributes...),
		AttributesToRetrieve:    opt.AttributesToRetrieve(attributesToRetrieve...),
		Ranking:                 opt.Ranking(ranking...),
		CustomRanking:           opt.CustomRanking(customRanking...),
		Replicas:                opt.Replicas(replicas...),
		MaxValuesPerFacet:       opt.MaxValuesPerFacet(d.Get("max_values_per_facet").(int)),
		SortFacetValuesBy:       opt.SortFacetValuesBy(d.Get("sort_facet_values_by").(string)),
		HitsPerPage:             opt.HitsPerPage(d.Get("hits_per_page").(int)),
		PaginationLimitedTo:     opt.PaginationLimitedTo(d.Get("pagination_limited_to").(int)),
		AttributeForDistinct:    opt.AttributeForDistinct(d.Get("attribute_for_distinct").(string)),
		Distinct:                opt.DistinctOf(d.Get("distinct").(int)),
	}
}
