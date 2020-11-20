package jamf

import (
	"context"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sioncojp/go-jamf-api"
)

type ProviderConfiguration struct {
	Client *jamf.Client
}

// Provider ... Define variables for the provider to use in the .tf file
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"JAMF_USERNAME"}, nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"JAMF_PASSWORD"}, nil),
			},
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This is the xxxx part of xxxx.jamfcloud.com",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"JAMF_ORGANIZATION"}, nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"jamf_department": resourceJamfDepartment(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"jamf_department": dataSourceJamfDepartment(),
		},
		ConfigureContextFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	c, err := jamf.NewClient(d.Get("username").(string), d.Get("password").(string), d.Get("organization").(string))
	if err != nil {
		diag.FromErr(err)
	}
	c.ExtraHeader["User-Agent"] = AppName
	c.HttpClient = cleanhttp.DefaultClient()
	return c, diags
}
