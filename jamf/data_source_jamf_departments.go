package jamf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sioncojp/go-jamf-api"
)

func dataSourceJamfDepartment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfDepartmentRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Computed values.
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceJamfDepartmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetDepartment(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if len(*resp.Id) < 1 {
		return diag.Errorf(ErrMsgGetQueryNoData)
	}
	if len(*resp.Id) > 1 {
		return diag.Errorf(ErrMsgGetQueryMultipleData)
	}

	d.SetId(*resp.Id)
	d.Set("name", *resp.Name)

	return diags
}
