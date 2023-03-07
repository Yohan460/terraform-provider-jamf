package jamf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceJamfComputerExtensionAttribute() *schema.Resource {
	return &schema.Resource{
		// CreateContext: resourceJamfComputerExtensionAttributeCreate,
		// ReadContext:   resourceJamfComputerExtensionAttributeRead,
		// UpdateContext: resourceJamfComputerExtensionAttributeUpdate,
		// DeleteContext: resourceJamfComputerExtensionAttributeDelete,
		// Importer: &schema.ResourceImporter{
		// 	StateContext: importJamfScriptState,
		// },
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
				Default:      "string",
				ValidateFunc: validation.StringInSlice([]string{"string", "integer", "date"}, true),
			},
			"inventory_display": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Extension Attributes",
				ValidateFunc: validation.StringInSlice([]string{"General", "Hardware", "Operating System", "User and Location", "Purchasing", "Extension Attributes"}, true),
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
				Elem:     &schema.Resource{},
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
