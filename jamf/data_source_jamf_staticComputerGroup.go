package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sioncojp/go-jamf-api"
)

func dataSourceJamfStaticComputerGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfStaticComputerGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"computers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceJamfStaticComputerGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetComputerGroupByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if resp.IsSmart {
		return diags
	}

	d.SetId(strconv.Itoa(resp.ID))
	d.Set("name", resp.Name)

	if resp.Site.ID != -1 {
		d.Set("site", []interface{}{
			map[string]interface{}{
				"id":   resp.Site.ID,
				"name": resp.Site.Name,
			},
		})
	}

	comps := make([]interface{}, len(resp.Computers), len(resp.Computers))
	for i, v := range resp.Computers {
		comps[i] = map[string]interface{}{
			"id":            v.ID,
			"name":          v.Name,
			"serial_number": v.SerialNumber,
		}
	}
	d.Set("computers", comps)

	return diags
}
