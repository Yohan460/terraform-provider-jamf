package jamf

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sioncojp/go-jamf-api"
)

func resourceJamfDepartment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfDepartmentCreate,
		ReadContext:   resourceJamfDepartmentRead,
		UpdateContext: resourceJamfDepartmentUpdate,
		DeleteContext: resourceJamfDepartmentDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		Schema: map[string]*schema.Schema{
			"department_id": {
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

	return &schema.Resource{}
}

func buildJamfDepartmentStruct(d *schema.ResourceData) *jamf.Department {
	var out jamf.Department
	out.SetName(d.Get("name").(string))
	out.SetId(d.Get("department_id").(string))

	return &out
}

func resourceJamfDepartmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	b := buildJamfDepartmentStruct(d)

	if _, err := c.CreateDepartment(b.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceJamfDepartmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	out, err := c.GetDepartment(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("department_id", out.GetId())
	d.Set("name", out.GetName())

	return diags
}

func resourceJamfDepartmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	b := buildJamfDepartmentStruct(d)

	if _, err := c.UpdateDepartment(b); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceJamfDepartmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	b := buildJamfDepartmentStruct(d)

	if err := c.DeleteDepartment(*b.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
