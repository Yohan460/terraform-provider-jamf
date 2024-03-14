package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfComputerExtensionAttribute() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfComputerExtensionAttributeRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "String",
				ValidateFunc: validation.StringInSlice([]string{"String", "Integer", "Date"}, false),
			},
			"inventory_display": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Extension Attributes",
				ValidateFunc: validation.StringInSlice([]string{"General", "Hardware", "Operating System", "User and Location", "Purchasing", "Extension Attributes"}, false),
			},
			"recon_display": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"General", "Hardware", "Operating System", "User and Location", "Purchasing", "Extension Attributes"}, false),
			},
			"script": {
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"script", "text_field", "popup_menu"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"platform": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "Mac",
							ValidateFunc: validation.StringInSlice([]string{"Mac", "Windows"}, false),
						},
						"script_contents": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"script.0.file_path"},
						},
						"file_path": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"script.0.script_contents"},
						},
					},
				},
			},
			"text_field": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// set as a placeholder to `text_field` is recognized,
						// this schema is not used anywhere
						"input_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"popup_menu": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"choices": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceJamfComputerExtensionAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetComputerExtensionAttributeByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	deconstructJamfComputerExtensionAttributeStruct(d, resp)

	return diags
}

func deconstructJamfComputerExtensionAttributeStruct(d *schema.ResourceData, in *jamf.ComputerExtensionAttribute) {
	d.SetId(strconv.Itoa(in.Id))
	d.Set("name", in.Name)
	d.Set("description", in.Description)
	d.Set("data_type", in.DataType)
	d.Set("inventory_display", in.InventoryDisplay)
	d.Set("recon_display", in.ReconDisplay)

	// Input Type
	switch inputType := in.InputType.Type; inputType {
	case "script":
		scriptInterface := map[string]interface{}{
			"enabled":  in.Enabled,
			"platform": in.InputType.Platform,
		}

		if s, ok := d.GetOk("script"); ok {
			for _, v := range s.([]interface{}) {
				script := v.(map[string]interface{})

				// since file_path is always set in TypeList
				if script["file_path"] == "" {
					scriptInterface["script_contents"] = in.InputType.Script
				}
			}
		}

		d.Set("script", scriptInterface)
	case "Text Field":
		d.Set("text_field", []interface{}{
			map[string]interface{}{
				"input_type": "text_field",
			},
		})
	case "Pop-up Menu":
		d.Set("popup_menu", []interface{}{
			map[string]interface{}{
				"choices": in.InputType.Choices,
			},
		})
	}

	return
}
