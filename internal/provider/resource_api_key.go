package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApiKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApiKeyCreate,
		ReadContext:   resourceApiKeyRead,
		UpdateContext: resourceApiKeyUpdate,
		DeleteContext: resourceApiKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importApiKeyState,
		},
		Schema: map[string]*schema.Schema{
			"key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"acl": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"indexes": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"max_queries_per_ip_per_hour": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15000,
			},
			"max_hits_per_query": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"referers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
				Computed: true,
			},
			"validity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceApiKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := m.(*apiClient).algolia.AddAPIKey(getAlgoliaSearchKey(d))
	if err != nil {
		return diag.FromErr(err)
	}

	err = res.Wait()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("key", res.Key); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return resourceApiKeyRead(ctx, d, m)
}

func resourceApiKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if err := refreshApiKeyState(d, m); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceApiKeyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := m.(*apiClient).algolia.UpdateAPIKey(getAlgoliaSearchKey(d))
	if err != nil {
		return diag.FromErr(err)
	}

	err = res.Wait()
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceApiKeyRead(ctx, d, m)
}

func resourceApiKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	res, err := m.(*apiClient).algolia.DeleteAPIKey(d.Get("key").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	err = res.Wait()
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func importApiKeyState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	if err := d.Set("key", d.Id()); err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	if err := refreshApiKeyState(d, m); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func refreshApiKeyState(d *schema.ResourceData, m interface{}) error {
	key, err := m.(*apiClient).algolia.GetAPIKey(d.Get("key").(string))
	if err != nil {
		d.SetId("")
		return err
	}

	return setValues(d, map[string]interface{}{
		"acl":                         key.ACL,
		"description":                 key.Description,
		"indexes":                     key.Indexes,
		"max_queries_per_ip_per_hour": key.MaxQueriesPerIPPerHour,
		"max_hits_per_query":          key.MaxHitsPerQuery,
		"referers":                    key.Referers,
		"validity":                    key.Validity.Seconds(),
	})
}

func getAlgoliaSearchKey(d *schema.ResourceData) search.Key {
	var acl []string
	if value, ok := d.GetOk("acl"); ok {
		for _, v := range value.(*schema.Set).List() {
			acl = append(acl, v.(string))
		}
	}

	var indexes []string
	if value, ok := d.GetOk("indexes"); ok {
		for _, v := range value.(*schema.Set).List() {
			indexes = append(indexes, v.(string))
		}
	}

	var referers []string
	if value, ok := d.GetOk("referers"); ok {
		for _, v := range value.(*schema.Set).List() {
			referers = append(referers, v.(string))
		}
	}

	return search.Key{
		Value:                  d.Get("key").(string),
		ACL:                    acl,
		Description:            d.Get("description").(string),
		Indexes:                indexes,
		MaxQueriesPerIPPerHour: d.Get("max_queries_per_ip_per_hour").(int),
		MaxHitsPerQuery:        d.Get("max_hits_per_query").(int),
		Referers:               referers,
		Validity:               time.Duration(d.Get("validity").(int)) * time.Second,
	}
}
