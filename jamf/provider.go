package jamf

import (
	"context"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
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
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This is the full url of jamf, xxxx.jamfcloud.com",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"JAMF_URL"}, nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"jamf_department":                     resourceJamfDepartment(),
			"jamf_category":                       resourceJamfCategory(),
			"jamf_building":                       resourceJamfBuilding(),
			"jamf_staticComputerGroup":            resourceJamfStaticComputerGroup(),
			"jamf_smartComputerGroup":             resourceJamfSmartComputerGroup(),
			"jamf_script":                         resourceJamfScript(),
			"jamf_policy":                         resourceJamfPolicy(),
			"jamf_computer_configuration_profile": resourceJamfComputerConfigurationProfile(),
			"jamf_computer_extension_attribute":   resourceJamfComputerExtensionAttribute(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"jamf_department":                     dataSourceJamfDepartment(),
			"jamf_category":                       dataSourceJamfCategory(),
			"jamf_building":                       dataSourceJamfBuilding(),
			"jamf_staticComputerGroup":            dataSourceJamfStaticComputerGroup(),
			"jamf_smartComputerGroup":             dataSourceJamfSmartComputerGroup(),
			"jamf_script":                         dataSourceJamfScript(),
			"jamf_package":                        dataSourceJamfPackage(),
			"jamf_policy":                         dataSourceJamfPolicy(),
			"jamf_computer_configuration_profile": dataSourceJamfComputerConfigurationProfile(),
		},
		ConfigureContextFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	c, err := jamf.NewClient(d.Get("username").(string), d.Get("password").(string), d.Get("url").(string))
	if err != nil {
		diag.FromErr(err)
	}
	c.ExtraHeader["User-Agent"] = AppName
	c.HttpClient = cleanhttp.DefaultClient()
	return c, diags
}
