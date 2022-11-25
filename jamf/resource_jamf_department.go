package jamf

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfDepartment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfDepartmentCreate,
		ReadContext:   resourceJamfDepartmentRead,
		UpdateContext: resourceJamfDepartmentUpdate,
		DeleteContext: resourceJamfDepartmentDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: importJamfDepartmentState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func buildJamfDepartmentStruct(d *schema.ResourceData) *jamf.Department {
	var out jamf.Department
	out.SetId(d.Id())
	out.SetName(d.Get("name").(string))

	return &out
}

func resourceJamfDepartmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfDepartmentStruct(d)

	resp, err := c.CreateDepartment(b.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())

	return resourceJamfDepartmentRead(ctx, d, m)
}

func resourceJamfDepartmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetDepartmentByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", resp.GetName())

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
	resp, err := c.GetDepartment(d.Id())
	if err != nil {
		return nil, fmt.Errorf("cannot get department data")
	}

	d.Set("name", resp.GetName())

	return []*schema.ResourceData{d}, nil
}
