package jamf

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfScript() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfScriptRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter5": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter6": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter7": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter8": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter9": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter10": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameter11": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"script_contents": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJamfScriptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetScriptByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	deconstructJamfScriptStruct(d, resp)

	return diags
}

func deconstructJamfScriptStruct(d *schema.ResourceData, in *jamf.Script) {
	d.SetId(in.ID)
	d.Set("name", in.Name)
	d.Set("category_id", in.CategoryID)
	d.Set("category_name", in.CategoryName)
	d.Set("info", in.Info)
	d.Set("notes", in.Notes)
	d.Set("priority", in.Priority)
	d.Set("parameter4", in.Parameter4)
	d.Set("parameter5", in.Parameter5)
	d.Set("parameter6", in.Parameter6)
	d.Set("parameter7", in.Parameter7)
	d.Set("parameter8", in.Parameter8)
	d.Set("parameter9", in.Parameter9)
	d.Set("parameter10", in.Parameter10)
	d.Set("parameter11", in.Parameter11)
	d.Set("os_requirements", in.OsRequirements)
	if _, hasFilePath := d.GetOk("file_path"); !hasFilePath {
		d.Set("script_contents", in.ScriptContents)
	}

	return
}
