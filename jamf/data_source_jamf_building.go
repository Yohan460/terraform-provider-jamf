package jamf

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfBuilding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfBuildingRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"street_address1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"street_address2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state_province": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zip_postal_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJamfBuildingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetBuildingByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())
	d.Set("name", resp.GetName())
	d.Set("street_address1", resp.GetStreetAddress1())
	d.Set("street_address2", resp.GetStreetAddress2())
	d.Set("city", resp.GetCity())
	d.Set("state_province", resp.GetStateProvince())
	d.Set("zip_postal_code", resp.GetZipPostalCode())
	d.Set("country", resp.GetCountry())

	return diags
}
