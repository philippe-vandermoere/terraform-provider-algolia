package algolia

import (
	"context"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Algolia application ID",
				DefaultFunc: schema.EnvDefaultFunc("ALGOLIA_APPLICATION_ID", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API key",
				DefaultFunc: schema.EnvDefaultFunc("ALGOLIA_API_KEY", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"algolia_api_key": resourceApiKey(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"algolia_index": dataSourceIndex(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	applicationId := d.Get("application_id").(string)
	apiKey := d.Get("api_key").(string)

	return search.NewClient(applicationId, apiKey), diags
}
