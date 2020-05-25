package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ts "github.com/ryanwersal/teamspeak"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TEAMSPEAK_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"teamspeak_channel": resourceChannel(),
		},
		ConfigureFunc: Configure,
	}
}

func Configure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	key := d.Get("api_token").(string)

	cfg := ts.NewConfiguration()
	cfg.BasePath = endpoint
	cfg.Debug = true
	cfg.AddDefaultHeader("x-api-key", key)

	client := ts.NewAPIClient(cfg)
	return client, nil
}
