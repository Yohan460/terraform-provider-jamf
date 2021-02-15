package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfSmartComputerGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfSmartComputerGroupRead,
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
			"criteria": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"and_or": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"search_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"search_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"opening_paren": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"closing_paren": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"computer": {
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

func dataSourceJamfSmartComputerGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetComputerGroupByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if !resp.IsSmart {
		return diags
	}

	deconstructJamfComputerGroupStruct(d, resp)

	return diags
}

func deconstructJamfComputerGroupStruct(d *schema.ResourceData, in *jamf.ComputerGroup) {
	d.SetId(strconv.Itoa(in.ID))
	d.Set("name", in.Name)

	if in.Site.ID != -1 {
		d.Set("site", []interface{}{
			map[string]interface{}{
				"id":   in.Site.ID,
				"name": in.Site.Name,
			},
		})
	}

	criterias := make([]interface{}, len(in.Criteria), len(in.Criteria))
	for i, v := range in.Criteria {
		criterias[i] = map[string]interface{}{
			"priority":      v.Priority,
			"name":          v.Name,
			"and_or":        v.AndOr,
			"search_type":   v.SearchType,
			"search_value":  v.SearchValue,
			"opening_paren": v.OpeningParen,
			"closing_paren": v.ClosingParen,
		}
	}
	d.Set("criteria", criterias)

	comps := make([]interface{}, len(in.Computers), len(in.Computers))
	for i, v := range in.Computers {
		comps[i] = map[string]interface{}{
			"id":            v.ID,
			"name":          v.Name,
			"serial_number": v.SerialNumber,
		}
	}
	d.Set("computer", comps)

	return
}
