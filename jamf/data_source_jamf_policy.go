package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sioncojp/go-jamf-api"
)

func dataSourceJamfPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfPolicyRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"general": {
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
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trigger_checkin": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger_enrollment_complete": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger_login": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger_logout": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger_network_state_changed": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger_startup": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"trigger_other": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"frequency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"retry_event": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"retry_attempts": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"notify_on_each_failed_retry": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"location_user_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"target_drive": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"offline": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"date_time_limitations": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"network_limitations": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"override_default_settings": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"network_requirements": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"site": {
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
								},
							},
						},
						// TODO All advanced General structures
					},
					// TODO All other structures
				},
			},
		},
	}
}

func dataSourceJamfPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetPolicyByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	deconstructJamfPolicyStruct(d, resp)

	return diags
}

func deconstructJamfPolicyStruct(d *schema.ResourceData, in *jamf.Policy) {

	// General
	d.SetId(strconv.Itoa(in.General.ID))
	d.Set("general", []interface{}{
		map[string]interface{}{
			"id":                            in.General.ID,
			"name":                          in.General.Name,
			"enabled":                       in.General.Enabled,
			"trigger":                       in.General.Trigger,
			"trigger_checkin":               in.General.TriggerCheckin,
			"trigger_enrollment_complete":   in.General.TriggerEnrollmentComplete,
			"trigger_login":                 in.General.TriggerLogin,
			"trigger_network_state_changed": in.General.TriggerNetworkStateChanged,
			"trigger_startup":               in.General.TriggerStartup,
			"trigger_other":                 in.General.TriggerOther,
			"frequency":                     in.General.Frequency,
			"retry_event":                   in.General.RetryEvent,
			"retry_attempts":                in.General.RetryAttempts,
			"notify_on_each_failed_retry":   in.General.NotifyOnEachFailedRetry,
			"location_user_only":            in.General.LocationUserOnly,
			"target_drive":                  in.General.TargetDrive,
			"offline":                       in.General.Offline,
			"network_requirements":          in.General.NetworkRequirements,

			// TODO All advanced General structures
		},
	})

	// TODO All other structures

	return
}
