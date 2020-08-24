package algolia

import (
	"context"

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
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"acl": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"indexes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"max_queries_per_ip_peer_hour": {
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
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceApiKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := m.(*search.Client).AddAPIKey(getAlgoliaKey(d))
	if err != nil {
		return diag.FromErr(err)
	}

	err = res.Wait()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.Key)

	return resourceApiKeyRead(ctx, d, m)
}

func resourceApiKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	key, err := m.(*search.Client).GetAPIKey(d.Id())
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	d.SetId(key.Value)
	_ = d.Set("acl", key.ACL)
	_ = d.Set("description", key.Description)
	_ = d.Set("indexes", key.Indexes)
	_ = d.Set("max_queries_per_ip_peer_hour", key.MaxQueriesPerIPPerHour)
	_ = d.Set("max_hits_per_query", key.MaxHitsPerQuery)
	_ = d.Set("referers", key.Referers)

	return diags
}

func resourceApiKeyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	res, err := m.(*search.Client).UpdateAPIKey(getAlgoliaKey(d))
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

	res, err := m.(*search.Client).DeleteAPIKey(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = res.Wait()
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func getAlgoliaKey(d *schema.ResourceData) search.Key {
	var acl []string
	if value := d.Get("acl"); value != nil {
		for _, v := range value.(*schema.Set).List() {
			acl = append(acl, v.(string))
		}
	}

	var indexes []string
	if value := d.Get("indexes"); value != nil {
		for _, v := range value.(*schema.Set).List() {
			indexes = append(indexes, v.(string))
		}
	}

	var referers []string
	if value := d.Get("referers"); value != nil {
		for _, v := range value.(*schema.Set).List() {
			referers = append(referers, v.(string))
		}
	}

	return search.Key{
		Value:                  d.Id(),
		ACL:                    acl,
		Description:            d.Get("description").(string),
		Indexes:                indexes,
		MaxQueriesPerIPPerHour: d.Get("max_queries_per_ip_peer_hour").(int),
		MaxHitsPerQuery:        d.Get("max_hits_per_query").(int),
		Referers:               referers,
	}
}
