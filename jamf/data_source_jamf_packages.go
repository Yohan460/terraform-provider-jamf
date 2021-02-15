package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfPackage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfPackageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"category_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fill_existing_users": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"boot_volume_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_uninstalled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"os_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"required_processor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hash_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hash_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJamfPackageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetPackageByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	deconstructPackageStruct(d, resp)

	return diags
}

func deconstructPackageStruct(d *schema.ResourceData, in *jamf.Package) {
	d.SetId(strconv.Itoa(in.ID))
	d.Set("name", in.Name)
	d.Set("category_name", in.CategoryName)
	d.Set("filename", in.Filename)
	d.Set("info", in.Info)
	d.Set("notes", in.Notes)
	d.Set("priority", in.Priority)
	d.Set("reboot_required", in.RebootRequired)
	d.Set("fill_existing_users", in.FillExistingUsers)
	d.Set("boot_volume_required", in.BootVolumeRequired)
	d.Set("allow_uninstalled", in.AllowUninstalled)
	d.Set("os_requirements", in.OsRequirements)
	d.Set("required_processor", in.RequiredProcessor)
	d.Set("hash_type", in.HashType)
	d.Set("hash_value", in.HashValue)
	return
}
