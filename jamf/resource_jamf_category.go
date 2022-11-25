package jamf

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfCategory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfCategoryCreate,
		ReadContext:   resourceJamfCategoryRead,
		UpdateContext: resourceJamfCategoryUpdate,
		DeleteContext: resourceJamfCategoryDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: importJamfCategoryState,
		},
		Schema: map[string]*schema.Schema{
			// Computed values.
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func buildJamfCategoryStruct(d *schema.ResourceData) *jamf.Category {
	var out jamf.Category
	out.SetId(d.Id())
	out.SetName(d.Get("name").(string))
	out.SetPriority(d.Get("priority").(int))

	return &out
}

func resourceJamfCategoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfCategoryStruct(d)

	resp, err := c.CreateCategory(b.Name, b.Priority)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.GetId())

	return resourceJamfCategoryRead(ctx, d, m)
}

func resourceJamfCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetCategoryByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", resp.GetName())
	d.Set("priority", resp.GetPriority())

	return diags
}

func resourceJamfCategoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfCategoryStruct(d)
	d.SetId(b.GetId())

	if _, err := c.UpdateCategory(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfCategoryRead(ctx, d, m)
}

func resourceJamfCategoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	b := buildJamfCategoryStruct(d)

	if err := c.DeleteCategory(*b.Name); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfCategoryState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	resp, err := c.GetCategory(d.Id())
	if err != nil {
		return nil, fmt.Errorf("cannot get Category data")
	}

	d.Set("name", resp.GetName())
	d.Set("priority", resp.GetPriority())

	return []*schema.ResourceData{d}, nil
}
