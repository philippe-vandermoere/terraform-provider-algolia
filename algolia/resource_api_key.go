package algolia

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
	res, err := m.(*search.Client).AddAPIKey(getAlgoliaSearchKey(d))
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
	res, err := m.(*search.Client).UpdateAPIKey(getAlgoliaSearchKey(d))
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

	res, err := m.(*search.Client).DeleteAPIKey(d.Get("key").(string))
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
	key, err := m.(*search.Client).GetAPIKey(d.Get("key").(string))
	if err != nil {
		d.SetId("")
		return err
	}

	if err := d.Set("acl", key.ACL); err != nil {
		return err
	}

	if err := d.Set("description", key.Description); err != nil {
		return err
	}

	if err := d.Set("indexes", key.Indexes); err != nil {
		return err
	}

	if err := d.Set("max_queries_per_ip_peer_hour", key.MaxQueriesPerIPPerHour); err != nil {
		return err
	}

	if err := d.Set("max_hits_per_query", key.MaxHitsPerQuery); err != nil {
		return err
	}

	if err := d.Set("referers", key.Referers); err != nil {
		return err
	}

	return nil
}

func getAlgoliaSearchKey(d *schema.ResourceData) search.Key {
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
		Value:                  d.Get("key").(string),
		ACL:                    acl,
		Description:            d.Get("description").(string),
		Indexes:                indexes,
		MaxQueriesPerIPPerHour: d.Get("max_queries_per_ip_peer_hour").(int),
		MaxHitsPerQuery:        d.Get("max_hits_per_query").(int),
		Referers:               referers,
	}
}
