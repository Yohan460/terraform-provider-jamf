package jamf

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfCategory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfCategoryRead,
		Schema: map[string]*schema.Schema{
			// Computed values.
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed values.
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceJamfCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetCategoryByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())
	d.Set("name", resp.GetName())
	d.Set("priority", resp.GetPriority())

	return diags
}
