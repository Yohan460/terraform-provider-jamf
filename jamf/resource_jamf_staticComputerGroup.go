package jamf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfStaticComputerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfStaticComputerGroupCreate,
		ReadContext:   resourceJamfStaticComputerGroupRead,
		UpdateContext: resourceJamfStaticComputerGroupUpdate,
		DeleteContext: resourceJamfStaticComputerGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importJamfStaticComputerGroupState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"computer": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
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

func buildJamfStaticComputerGroupStruct(d *schema.ResourceData) *jamf.ComputerGroup {
	var out jamf.ComputerGroup
	id, _ := strconv.Atoi(d.Id())
	out.ID = id
	out.Name = d.Get("name").(string)
	out.IsSmart = false
	if v, ok := d.GetOk("site"); ok {
		siteList := v.(*schema.Set).List()
		site := siteList[0].(map[string]interface{})
		if val, ok := site["name"].(string); ok {
			out.Site.Name = val
		}
		if val, ok := site["id"].(int); ok {
			out.Site.ID = val
		}
	}

	if v, ok := d.GetOk("computer"); ok {
		comps := v.(*schema.Set).List()
		compList := []jamf.ComputerGroupComputerEntry{}
		for _, c := range comps {
			compData := c.(map[string]interface{})
			comp := jamf.ComputerGroupComputerEntry{}
			if val, ok := compData["id"].(int); ok {
				comp.ID = val
			}
			if val, ok := compData["serial_number"].(string); ok {
				comp.SerialNumber = val
			}
			if val, ok := compData["name"].(string); ok {
				comp.Name = val
			}
			compList = append(compList, comp)
		}
		out.Computers = compList
	}

	return &out
}

func resourceJamfStaticComputerGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfStaticComputerGroupStruct(d)

	group := &jamf.ComputerGroupRequest{}
	group.Name = b.Name
	if b.Site.Name != "" {
		group.Site = b.Site
	}
	group.IsSmart = b.IsSmart
	group.Computers = b.Computers

	id, err := c.CreateComputerGroup(group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceJamfStaticComputerGroupRead(ctx, d, m)
}

func resourceJamfStaticComputerGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	resp, err := c.GetComputerGroup(id)
	if err != nil {
		if jamfErr, ok := err.(jamf.Error); ok && jamfErr.StatusCode() == 404 {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	} else {
		deconstructJamfComputerGroupStruct(d, resp)
	}

	return diags
}

func resourceJamfStaticComputerGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfStaticComputerGroupStruct(d)

	if _, err := c.UpdateComputerGroup(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfStaticComputerGroupRead(ctx, d, m)
}

func resourceJamfStaticComputerGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	b := buildJamfStaticComputerGroupStruct(d)

	if _, err := c.DeleteComputerGroup(b.ID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfStaticComputerGroupState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, err
	}
	resp, err := c.GetComputerGroup(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get Computer Group data")
	}

	deconstructJamfComputerGroupStruct(d, resp)

	return []*schema.ResourceData{d}, nil
}
