package jamf

import (
	"context"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfScript() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfScriptCreate,
		ReadContext:   resourceJamfScriptRead,
		UpdateContext: resourceJamfScriptUpdate,
		DeleteContext: resourceJamfScriptDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importJamfScriptState,
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
			"category_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "-1",
			},
			"category_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AFTER",
			},
			"info": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter4": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter5": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter7": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter8": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter9": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter10": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter11": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_requirements": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"script_contents": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"file_path"},
			},
			"file_path": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"script_contents"},
			},
		},
	}
}

func buildJamfScriptStruct(d *schema.ResourceData) (*jamf.Script, error) {
	var out jamf.Script
	out.ID = d.Id()
	out.Name = d.Get("name").(string)
	if v, ok := d.GetOk("category_id"); ok {
		out.CategoryID = v.(string)
	}
	if v, ok := d.GetOk("category_name"); ok {
		out.CategoryName = v.(string)
	}
	if v, ok := d.GetOk("priority"); ok {
		out.Priority = v.(string)
	}
	if v, ok := d.GetOk("info"); ok {
		out.Info = v.(string)
	}
	if v, ok := d.GetOk("notes"); ok {
		out.Notes = v.(string)
	}
	if v, ok := d.GetOk("parameter4"); ok {
		out.Parameter4 = v.(string)
	}
	if v, ok := d.GetOk("parameter5"); ok {
		out.Parameter5 = v.(string)
	}
	if v, ok := d.GetOk("parameter6"); ok {
		out.Parameter6 = v.(string)
	}
	if v, ok := d.GetOk("parameter7"); ok {
		out.Parameter7 = v.(string)
	}
	if v, ok := d.GetOk("parameter8"); ok {
		out.Parameter8 = v.(string)
	}
	if v, ok := d.GetOk("parameter9"); ok {
		out.Parameter9 = v.(string)
	}
	if v, ok := d.GetOk("parameter10"); ok {
		out.Parameter10 = v.(string)
	}
	if v, ok := d.GetOk("parameter11"); ok {
		out.Parameter11 = v.(string)
	}
	if v, ok := d.GetOk("os_requirements"); ok {
		out.OsRequirements = v.(string)
	}

	filePath, hasFilePath := d.GetOk("file_path")
	scriptContents, hasScriptContents := d.GetOk("script_contents")

	if hasFilePath && !hasScriptContents {
		content, err := loadFileContent(filePath.(string))
		if err != nil {
			return &out, err
		}
		out.ScriptContents = content
	} else if !hasFilePath && hasScriptContents {
		out.ScriptContents = scriptContents.(string)
	}

	return &out, nil
}

func resourceJamfScriptCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b, err := buildJamfScriptStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := c.CreateScript(b)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	return resourceJamfScriptRead(ctx, d, m)
}

func resourceJamfScriptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetScript(d.Id())

	if err != nil {
		if jamfErr, ok := err.(jamf.Error); ok && jamfErr.StatusCode() == 404 {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	} else {
		deconstructJamfScriptStruct(d, resp)
	}

	return diags
}

func resourceJamfScriptUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b, err := buildJamfScriptStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}
	if _, err := c.UpdateScript(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfScriptRead(ctx, d, m)
}

func resourceJamfScriptDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	b, err := buildJamfScriptStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err := c.DeleteScript(b.ID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfScriptState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	resp, err := c.GetScript(d.Id())
	if err != nil {
		return nil, err
	}

	deconstructJamfScriptStruct(d, resp)

	return []*schema.ResourceData{d}, nil
}

// loadFileContent returns contents of a file in a given path
func loadFileContent(v string) (string, error) {
	filename, err := homedir.Expand(v)
	if err != nil {
		return "", err
	}
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(fileContent), err
}
