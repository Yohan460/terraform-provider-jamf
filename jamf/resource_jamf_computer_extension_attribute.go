package jamf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfComputerExtensionAttribute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfComputerExtensionAttributeCreate,
		ReadContext:   resourceJamfComputerExtensionAttributeRead,
		UpdateContext: resourceJamfComputerExtensionAttributeUpdate,
		DeleteContext: resourceJamfComputerExtensionAttributeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importJamfComputerExtensionAttributeState,
		},
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

func buildJamfComputerExtensionAttributeStruct(d *schema.ResourceData) (*jamf.ComputerExtensionAttribute, error) {
	var out jamf.ComputerExtensionAttribute

	id, _ := strconv.Atoi(d.Id())
	out.Id = id

	out.Name = d.Get("name").(string)

	if v, ok := d.GetOk("description"); ok {
		out.Description = v.(string)
	}

	if v, ok := d.GetOk("data_type"); ok {
		out.DataType = v.(string)
	}

	if s, ok := d.GetOk("script"); ok {
		for _, v := range s.([]interface{}) {
			script := v.(map[string]interface{})

			out.InputType.Type = "script"

			if v, ok := script["enabled"]; ok {
				out.Enabled = v.(bool)
			}

			if v, ok := script["platform"]; ok {
				out.InputType.Platform = v.(string)
			}

			filePath, hasFilePath := script["file_path"]
			if hasFilePath {
				if filePath == "" {
					hasFilePath = false // since file_path is always set in TypeList
				}
			}
			scriptContents, hasScriptContents := script["script_contents"]
			if hasScriptContents {
				if scriptContents == "" {
					hasScriptContents = false // since script_contents is always set in TypeList
				}
			}

			if hasFilePath && !hasScriptContents {
				content, err := loadFileContent(filePath.(string))
				if err != nil {
					return &out, err
				}
				out.InputType.Script = content
			} else if !hasFilePath && hasScriptContents && scriptContents != "" {
				out.InputType.Script = scriptContents.(string)
			} else {
				return &out, fmt.Errorf("only one of file_path and script_contents must be set")
			}
		}
	}

	if _, ok := d.GetOk("text_field"); ok {
		out.InputType.Type = "Text Field"
	}

	if s, ok := d.GetOk("popup_menu"); ok {
		val := s.(*schema.Set).List()
		popup := val[0].(map[string]interface{})

		out.InputType.Type = "Pop-up Menu"

		if v, ok := popup["choices"]; ok {
			choices := v.([]interface{})
			choicesList := make([]string, len(choices))
			for i, c := range choices {
				choicesList[i] = c.(string)
			}
			out.InputType.Choices = choicesList
		}
	}

	if v, ok := d.GetOk("inventory_display"); ok {
		out.InventoryDisplay = v.(string)
	}

	if v, ok := d.GetOk("recon_display"); ok {
		out.ReconDisplay = v.(string)
	}

	return &out, nil
}

func resourceJamfComputerExtensionAttributeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b, err := buildJamfComputerExtensionAttributeStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := c.CreateComputerExtensionAttribute(b)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceJamfComputerExtensionAttributeRead(ctx, d, m)
}

func resourceJamfComputerExtensionAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	id, _ := strconv.Atoi(d.Id())
	resp, err := c.GetComputerExtensionAttribute(id)

	if err != nil {
		if jamfErr, ok := err.(jamf.Error); ok && jamfErr.StatusCode() == 404 {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	} else {
		deconstructJamfComputerExtensionAttributeStruct(d, resp)
	}

	return diags
}

func resourceJamfComputerExtensionAttributeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b, err := buildJamfComputerExtensionAttributeStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err := c.UpdateComputerExtensionAttribute(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfComputerExtensionAttributeRead(ctx, d, m)
}

func resourceJamfComputerExtensionAttributeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err := c.DeleteComputerExtensionAttribute(id); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfComputerExtensionAttributeState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, err
	}
	resp, err := c.GetComputerExtensionAttribute(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get computer extension attribute data")
	}

	deconstructJamfComputerExtensionAttributeStruct(d, resp)

	return []*schema.ResourceData{d}, nil
}
