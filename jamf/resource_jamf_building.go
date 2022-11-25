package jamf

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfBuilding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfBuildingCreate,
		ReadContext:   resourceJamfBuildingRead,
		UpdateContext: resourceJamfBuildingUpdate,
		DeleteContext: resourceJamfBuildingDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Read:   schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: importJamfBuildingState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"street_address1": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"street_address2": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state_province": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zip_postal_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func buildJamfBuildingStruct(d *schema.ResourceData) *jamf.Building {
	var out jamf.Building
	out.SetId(d.Id())
	out.SetName(d.Get("name").(string))
	out.SetStreetAddress1(d.Get("street_address1").(string))
	out.SetStreetAddress2(d.Get("street_address2").(string))
	out.SetCity(d.Get("city").(string))
	out.SetStateProvince(d.Get("state_province").(string))
	out.SetZipPostalCode(d.Get("zip_postal_code").(string))
	out.SetCountry(d.Get("country").(string))

	return &out
}

func resourceJamfBuildingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfBuildingStruct(d)

	out, err := c.CreateBuilding(b.Name, b.StreetAddress1, b.StreetAddress2, b.City, b.StateProvince, b.ZipPostalCode, b.Country)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(out.GetId())

	return resourceJamfBuildingRead(ctx, d, m)
}

func resourceJamfBuildingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetBuildingByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", resp.GetName())
	d.Set("street_address1", resp.GetStreetAddress1())
	d.Set("street_address2", resp.GetStreetAddress2())
	d.Set("city", resp.GetCity())
	d.Set("state_province", resp.GetStateProvince())
	d.Set("zip_postal_code", resp.GetZipPostalCode())
	d.Set("country", resp.GetCountry())

	return diags
}

func resourceJamfBuildingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfBuildingStruct(d)
	d.SetId(b.GetId())

	if _, err := c.UpdateBuilding(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfBuildingRead(ctx, d, m)
}

func resourceJamfBuildingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	b := buildJamfBuildingStruct(d)

	if err := c.DeleteBuilding(*b.Name); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfBuildingState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	resp, err := c.GetBuilding(d.Id())
	if err != nil {
		return nil, fmt.Errorf("cannot get Building data")
	}

	d.Set("name", resp.GetName())
	d.Set("street_address1", resp.GetStreetAddress1())
	d.Set("street_address2", resp.GetStreetAddress2())
	d.Set("city", resp.GetCity())
	d.Set("state_province", resp.GetStateProvince())
	d.Set("zip_postal_code", resp.GetZipPostalCode())
	d.Set("country", resp.GetCountry())

	return []*schema.ResourceData{d}, nil
}
