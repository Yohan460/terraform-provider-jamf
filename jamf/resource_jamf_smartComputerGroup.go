package jamf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfSmartComputerGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfSmartComputerGroupCreate,
		ReadContext:   resourceJamfSmartComputerGroupRead,
		UpdateContext: resourceJamfSmartComputerGroupUpdate,
		DeleteContext: resourceJamfSmartComputerGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importJamfSmartComputerGroupState,
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
			"criteria": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"and_or": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "and",
						},
						"search_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"search_value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"opening_paren": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"closing_paren": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
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

func buildJamfSmartComputerGroupStruct(d *schema.ResourceData) *jamf.ComputerGroup {
	var out jamf.ComputerGroup
	id, _ := strconv.Atoi(d.Id())
	out.ID = id
	out.Name = d.Get("name").(string)
	out.IsSmart = true
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

	if v, ok := d.GetOk("criteria"); ok {
		criteria := v.(*schema.Set).List()
		criteriaList := []jamf.ComputerGroupCriterion{}
		for _, c := range criteria {
			criteriaData := c.(map[string]interface{})
			criterion := jamf.ComputerGroupCriterion{}
			criterion.Priority = criteriaData["priority"].(int)
			criterion.Name = criteriaData["name"].(string)
			andOr := criteriaData["and_or"].(string)
			if andOr == "and" {
				criterion.AndOr = jamf.And
			} else {
				criterion.AndOr = jamf.Or
			}
			criterion.SearchType = criteriaData["search_type"].(string)
			criterion.SearchValue = criteriaData["search_value"].(string)

			if val, ok := criteriaData["opening_paren"].(bool); ok {
				criterion.OpeningParen = val
			}
			if val, ok := criteriaData["closing_paren"].(bool); ok {
				criterion.ClosingParen = val
			}
			criteriaList = append(criteriaList, criterion)
		}
		out.Criteria = criteriaList
	}

	return &out
}

func resourceJamfSmartComputerGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfSmartComputerGroupStruct(d)

	group := &jamf.ComputerGroupRequest{}
	group.Name = b.Name
	if b.Site.Name != "" {
		group.Site = b.Site
	}
	group.IsSmart = b.IsSmart
	group.Criteria = b.Criteria

	id, err := c.CreateComputerGroup(group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceJamfSmartComputerGroupRead(ctx, d, m)
}

func resourceJamfSmartComputerGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceJamfSmartComputerGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfSmartComputerGroupStruct(d)

	if _, err := c.UpdateComputerGroup(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfSmartComputerGroupRead(ctx, d, m)
}

func resourceJamfSmartComputerGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	b := buildJamfSmartComputerGroupStruct(d)

	if _, err := c.DeleteComputerGroup(b.ID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfSmartComputerGroupState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
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
