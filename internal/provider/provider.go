package provider

import (
	"context"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
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
				"algolia_index":   resourceIndex(),
			},

			DataSourcesMap: map[string]*schema.Resource{
				"algolia_index": dataSourceIndex(),
			},
		}
		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	algolia *search.Client
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		config := search.Configuration{
			AppID:          d.Get("application_id").(string),
			APIKey:         d.Get("api_key").(string),
			ExtraUserAgent: p.UserAgent("terraform-provider-algolia", version),
		}

		return &apiClient{algolia: search.NewClientWithConfig(config)}, nil
	}
}
