package jamf

import (
	"context"
	"fmt"

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
		Importer: &schema.ResourceImporter{
			StateContext: importJamfDepartmentState,
		},
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
	c := m.(*jamf.Client)

	b := buildJamfDepartmentStruct(d)

	out, err := c.CreateDepartment(b.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(out.GetId())

	return resourceJamfDepartmentRead(ctx, d, m)
}

func resourceJamfDepartmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	out, err := c.GetDepartmentByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("department_id", out.GetId())
	d.Set("name", out.GetName())

	return diags
}

func resourceJamfDepartmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfDepartmentStruct(d)
	d.SetId(b.GetId())

	if _, err := c.UpdateDepartment(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfDepartmentRead(ctx, d, m)
}

func resourceJamfDepartmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	b := buildJamfDepartmentStruct(d)

	if err := c.DeleteDepartment(*b.Name); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfDepartmentState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	out, err := c.GetDepartment(d.Id())
	if err != nil {
		return nil, fmt.Errorf("cannot get department data")
	}

	d.Set("department_id", out.GetId())
	d.Set("name", out.GetName())

	return []*schema.ResourceData{d}, nil
}
