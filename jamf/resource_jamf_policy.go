package jamf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfPolicyCreate,
		ReadContext:   resourceJamfPolicyRead,
		UpdateContext: resourceJamfPolicyUpdate,
		DeleteContext: resourceJamfPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importJamfPolicyState,
		},
		Schema: map[string]*schema.Schema{
			"general": {
				Type:     schema.TypeSet,
				Required: true,
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
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "EVENT",
						},
						"trigger_checkin": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger_enrollment_complete": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger_login": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger_logout": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger_network_state_changed": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger_startup": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"trigger_other": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"frequency": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Once per computer",
						},
						"retry_event": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "none",
						},
						"retry_attempts": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},
						"notify_on_each_failed_retry": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"location_user_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"target_drive": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "/",
						},
						"offline": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"network_requirements": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Any",
						},
						"category": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"date_time_limitations": {
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"activation_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"activation_date_epoch": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  0,
									},
									"activation_date_utc": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"expiration_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"expiration_date_epoch": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  0,
									},
									"expiration_date_utc": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"no_execute_on": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"no_execute_start": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"no_execute_end": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"network_limitations": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"minimum_network_connection": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "No Minimum",
									},
									"any_ip_address": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"network_segments": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"override_default_settings": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_drive": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "/",
									},
									"distribution_point": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "default",
									},
									"force_afp_smb": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"sus": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "default",
									},
								},
							},
						},
						"site": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  -1,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"scope": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_computers": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"computer": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"udid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"computer_group": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"building": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"department": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"self_service": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_for_self_service": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"self_service_display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"install_button_text": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Install",
						},
						"reinstall_button_text": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Reinstall",
						},
						"self_service_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"force_users_to_view_description": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"feature_on_main_page": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"self_service_icon": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  0,
									},
									"filename": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"uri": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"self_service_category": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_in": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"feature_in": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},
					},
				},
			},
			"package": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "INSTALL",
						},
						"fut": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"feu": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"update_autorun": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"script": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "AFTER",
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
					},
				},
			},
			"reboot": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"startup_disk": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Current Startup Disk",
						},
						"specify_startup": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"no_user_logged_in": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Do not restart",
						},
						"user_logged_in": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Do not restart",
						},
						"minutes_until_reboot": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5,
						},
						"start_reboot_timer_immediately": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"file_vault_2_reboot": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"maintenance": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recon": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"reset_name": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"install_all_cached_packages": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"heal": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"prebindings": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"permissions": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"byhost": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"system_cache": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"user_cache": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"verify": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"files_and_processes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_by_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"delete_file": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"locate_file": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_locate_database": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"spotlight_search": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"search_for_process": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kill_process": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"run_command": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"user_interaction": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_start": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"allow_users_to_defer": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"allow_deferral_until_utc": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"allow_deferral_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"message_finish": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func buildJamfPolicyStruct(d *schema.ResourceData) *jamf.Policy {
	var out jamf.Policy

	// General
	id, _ := strconv.Atoi(d.Id())
	out.General.ID = id
	if g, ok := d.GetOk("general"); ok {
		v := g.(*schema.Set).List()
		general := v[0].(map[string]interface{})

		out.General.Name = general["name"].(string)

		if val, ok := general["enabled"]; ok {
			out.General.Enabled = val.(bool)
		}
		if val, ok := general["trigger"]; ok {
			out.General.Trigger = val.(string)
		}
		if val, ok := general["trigger_checkin"]; ok {
			out.General.TriggerCheckin = val.(bool)
		}
		if val, ok := general["trigger_enrollment_complete"]; ok {
			out.General.TriggerEnrollmentComplete = val.(bool)
		}
		if val, ok := general["trigger_logout"]; ok {
			out.General.TriggerLogout = val.(bool)
		}
		if val, ok := general["trigger_network_state_changed"]; ok {
			out.General.TriggerNetworkStateChanged = val.(bool)
		}
		if val, ok := general["trigger_startup"]; ok {
			out.General.TriggerStartup = val.(bool)
		}
		if val, ok := general["trigger_other"]; ok {
			out.General.TriggerOther = val.(string)
		}
		if val, ok := general["frequency"]; ok {
			out.General.Frequency = val.(string)
		}
		if val, ok := general["retry_event"]; ok {
			out.General.RetryEvent = val.(string)
		}
		if val, ok := general["retry_attempts"]; ok {
			out.General.RetryAttempts = val.(int)
		}
		if val, ok := general["notify_on_each_failed_retry"]; ok {
			out.General.NotifyOnEachFailedRetry = val.(bool)
		}
		if val, ok := general["location_user_only"]; ok {
			out.General.LocationUserOnly = val.(bool)
		}
		if val, ok := general["target_drive"]; ok {
			out.General.TargetDrive = val.(string)
		}
		if val, ok := general["offline"]; ok {
			out.General.Offline = val.(bool)
		}
		if val, ok := general["network_requirements"]; ok {
			out.General.NetworkRequirements = val.(string)
		}

		// General - Category
		if v, ok := general["category"]; ok {
			categoryList := v.(*schema.Set).List()
			category := categoryList[0].(map[string]interface{})
			if val, ok := category["name"].(string); ok {
				out.General.Category.Name = val
			}
			if val, ok := category["id"].(string); ok {
				out.General.Category.ID = val
			}
		}

		// General - Date and Time Limitations
		if v, ok := general["date_time_limitations"]; ok {
			dateAndTimeLimitationList := v.(*schema.Set).List()
			if len(dateAndTimeLimitationList) == 1 {
				dateAndTimeLimitation := dateAndTimeLimitationList[0].(map[string]interface{})
				if val, ok := dateAndTimeLimitation["activation_date"].(string); ok {
					out.General.DateTimeLimitations.ActivationDate = val
				}
				if val, ok := dateAndTimeLimitation["activation_date_epoch"].(int); ok {
					out.General.DateTimeLimitations.ActivationDateEpoch = val
				}
				if val, ok := dateAndTimeLimitation["activation_date_utc"].(string); ok {
					out.General.DateTimeLimitations.ActivationDateUtc = val
				}
				if val, ok := dateAndTimeLimitation["expiration_date"].(string); ok {
					out.General.DateTimeLimitations.ExpirationDate = val
				}
				if val, ok := dateAndTimeLimitation["expiration_date_epoch"].(int); ok {
					out.General.DateTimeLimitations.ExpirationDateEpoch = val
				}
				if val, ok := dateAndTimeLimitation["expiration_date_utc"].(string); ok {
					out.General.DateTimeLimitations.ExpirationDateUtc = val
				}
				if val, ok := dateAndTimeLimitation["no_execute_on"].(string); ok {
					out.General.DateTimeLimitations.NoExecuteOn = val
				}
				if val, ok := dateAndTimeLimitation["no_execute_start"].(string); ok {
					out.General.DateTimeLimitations.NoExecuteStart = val
				}
				if val, ok := dateAndTimeLimitation["no_execute_end"].(string); ok {
					out.General.DateTimeLimitations.NoExecuteEnd = val
				}
			}
		}

		// General - Network Limitations
		if v, ok := general["network_limitations"]; ok {
			networkLimitationList := v.(*schema.Set).List()
			networkLimitation := networkLimitationList[0].(map[string]interface{})
			if val, ok := networkLimitation["minimum_network_connection"].(string); ok {
				out.General.NetworkLimitations.MinimumNetworkConnection = val
			}
			if val, ok := networkLimitation["any_ip_address"].(bool); ok {
				out.General.NetworkLimitations.AnyIpAddress = val
			}
			if val, ok := networkLimitation["network_segments"].(string); ok {
				out.General.NetworkLimitations.NetworkSegments = val
			}
		}

		// General - Override Default Settings
		if v, ok := general["override_default_settings"]; ok {
			overrideDefaultSettingsList := v.(*schema.Set).List()
			overrideDefaultSettings := overrideDefaultSettingsList[0].(map[string]interface{})
			if val, ok := overrideDefaultSettings["target_drive"].(string); ok {
				out.General.OverrideDefaultSettings.TargetDrive = val
			}
			if val, ok := overrideDefaultSettings["distribution_point"].(string); ok {
				out.General.OverrideDefaultSettings.DistributionPoint = val
			}
			if val, ok := overrideDefaultSettings["force_afp_smb"].(bool); ok {
				out.General.OverrideDefaultSettings.ForceAfpSmb = val
			}
			if val, ok := overrideDefaultSettings["sus"].(string); ok {
				out.General.OverrideDefaultSettings.Sus = val
			}
		}

		// General - Site
		if v, ok := general["site"]; ok {
			siteList := v.(*schema.Set).List()
			site := siteList[0].(map[string]interface{})
			if val, ok := site["name"].(string); ok {
				out.General.Site.Name = val
			}
			if val, ok := site["id"].(int); ok {
				out.General.Site.ID = val
			}
		}
	}

	// Scope
	if s, ok := d.GetOk("scope"); ok {
		v := s.(*schema.Set).List()
		scope := v[0].(map[string]interface{})

		if val, ok := scope["all_computers"]; ok {
			out.Scope.AllComputers = val.(bool)
		}

		// Scope - Computers
		if v, ok := scope["computer"]; ok {
			computers := v.(*schema.Set).List()
			computerList := []jamf.ComputerScope{}
			for _, c := range computers {
				computerData := c.(map[string]interface{})
				computer := jamf.ComputerScope{}
				if val, ok := computerData["id"].(int); ok {
					computer.ID = val
				}
				computerList = append(computerList, computer)
			}
			out.Scope.Computers = computerList
		}

		// Scope - Computer Groups
		if v, ok := scope["computer_group"]; ok {
			computerGroups := v.(*schema.Set).List()
			computerGroupList := []jamf.ComputerGroupListResponse{}
			for _, c := range computerGroups {
				computerGroupData := c.(map[string]interface{})
				computerGroup := jamf.ComputerGroupListResponse{}
				if val, ok := computerGroupData["id"].(int); ok {
					computerGroup.ID = val
				}
				computerGroupList = append(computerGroupList, computerGroup)
			}
			out.Scope.ComputerGroups = computerGroupList
		}

		// Scope - Buildings
		if v, ok := scope["building"]; ok {
			buildings := v.(*schema.Set).List()
			buildingList := []jamf.BuildingScope{}
			for _, c := range buildings {
				buildingData := c.(map[string]interface{})
				building := jamf.BuildingScope{}
				if val, ok := buildingData["id"].(int); ok {
					building.ID = val
				}
				buildingList = append(buildingList, building)
			}
			out.Scope.Buildings = buildingList
		}

		// Scope - Departments
		if v, ok := scope["department"]; ok {
			departments := v.(*schema.Set).List()
			departmentList := []jamf.DepartmentScope{}
			for _, c := range departments {
				departmentData := c.(map[string]interface{})
				department := jamf.DepartmentScope{}
				if val, ok := departmentData["id"].(int); ok {
					department.ID = val
				}
				departmentList = append(departmentList, department)
			}
			out.Scope.Departments = departmentList
		}
	}

	// Self Service
	if s, ok := d.GetOk("self_service"); ok {
		v := s.(*schema.Set).List()
		selfService := v[0].(map[string]interface{})

		if val, ok := selfService["use_for_self_service"]; ok {
			out.SelfService.UseForSelfService = val.(bool)
		}
		if val, ok := selfService["self_service_display_name"]; ok {
			out.SelfService.SelfServiceDisplayName = val.(string)
		}
		if val, ok := selfService["install_button_text"]; ok {
			out.SelfService.InstallButtonText = val.(string)
		}
		if val, ok := selfService["reinstall_button_text"]; ok {
			out.SelfService.ReinstallButtonText = val.(string)
		}
		if val, ok := selfService["self_service_description"]; ok {
			out.SelfService.SelfServiceDescription = val.(string)
		}
		if val, ok := selfService["force_users_to_view_description"]; ok {
			out.SelfService.ForceUsersToViewDescription = val.(bool)
		}
		if val, ok := selfService["feature_on_main_page"]; ok {
			out.SelfService.FeatureOnMainPage = val.(bool)
		}
		if v, ok := selfService["self_service_icon"]; ok {
			selfServiceIconList := v.(*schema.Set).List()
			selfServiceIcon := selfServiceIconList[0].(map[string]interface{})
			if val, ok := selfServiceIcon["id"].(int); ok {
				out.SelfService.SelfServiceIcon.ID = val
			}
		}

		// Self Service - Category
		if v, ok := selfService["self_service_category"]; ok {
			selfServiceCategories := v.(*schema.Set).List()
			selfServiceCategoryList := []jamf.SelfServiceCategory{}
			for _, c := range selfServiceCategories {
				selfServiceCategoryData := c.(map[string]interface{})
				selfServiceCategory := jamf.SelfServiceCategory{}
				if val, ok := selfServiceCategoryData["id"].(int); ok {
					selfServiceCategory.ID = val
				}
				if val, ok := selfServiceCategoryData["display_in"].(bool); ok {
					selfServiceCategory.DisplayIn = val
				}
				if val, ok := selfServiceCategoryData["feature_in"].(bool); ok {
					selfServiceCategory.FeatureIn = val
				}
				selfServiceCategoryList = append(selfServiceCategoryList, selfServiceCategory)
			}
			out.SelfService.SelfServiceCategories = selfServiceCategoryList
		}
	}

	// Packages
	if v, ok := d.GetOk("package"); ok {
		packages := v.(*schema.Set).List()
		packageList := []jamf.PolicyPackageConfigurationPackage{}
		for _, c := range packages {
			packageData := c.(map[string]interface{})
			pkg := jamf.PolicyPackageConfigurationPackage{}
			if val, ok := packageData["id"].(int); ok {
				pkg.ID = val
			}
			if val, ok := packageData["action"].(string); ok {
				pkg.Action = val
			}
			if val, ok := packageData["fut"].(bool); ok {
				pkg.FillUserTemplate = val
			}
			if val, ok := packageData["feu"].(bool); ok {
				pkg.FillExistingUsers = val
			}
			if val, ok := packageData["update_autorun"].(bool); ok {
				pkg.UpdateAutorun = val
			}
			packageList = append(packageList, pkg)
		}
		out.PackageConfiguration.Packages = packageList
	}

	// Scripts
	if v, ok := d.GetOk("script"); ok {
		scripts := v.(*schema.Set).List()
		scriptList := []jamf.PolicyScript{}
		for _, c := range scripts {
			scriptData := c.(map[string]interface{})
			script := jamf.PolicyScript{}
			if val, ok := scriptData["id"].(string); ok {
				script.ID = val
			}
			if val, ok := scriptData["priority"].(string); ok {
				script.Priority = val
			}
			if val, ok := scriptData["parameter4"].(string); ok {
				script.Parameter4 = val
			}
			if val, ok := scriptData["parameter5"].(string); ok {
				script.Parameter5 = val
			}
			if val, ok := scriptData["parameter6"].(string); ok {
				script.Parameter6 = val
			}
			if val, ok := scriptData["parameter7"].(string); ok {
				script.Parameter7 = val
			}
			if val, ok := scriptData["parameter8"].(string); ok {
				script.Parameter8 = val
			}
			if val, ok := scriptData["parameter9"].(string); ok {
				script.Parameter9 = val
			}
			if val, ok := scriptData["parameter10"].(string); ok {
				script.Parameter10 = val
			}
			if val, ok := scriptData["parameter11"].(string); ok {
				script.Parameter11 = val
			}
			scriptList = append(scriptList, script)
		}
		out.ScriptsConfiguration.Scripts = scriptList
	}

	// Reboot
	if r, ok := d.GetOk("reboot"); ok {
		v := r.(*schema.Set).List()
		reboot := v[0].(map[string]interface{})

		if val, ok := reboot["message"]; ok {
			out.Reboot.Message = val.(string)
		}
		if val, ok := reboot["startup_disk"]; ok {
			out.Reboot.StartupDisk = val.(string)
		}
		if val, ok := reboot["specify_startup"]; ok {
			out.Reboot.SpecifyStartup = val.(string)
		}
		if val, ok := reboot["no_user_logged_in"]; ok {
			out.Reboot.NoUserLoggedIn = val.(string)
		}
		if val, ok := reboot["user_logged_in"]; ok {
			out.Reboot.UserLoggedIn = val.(string)
		}
		if val, ok := reboot["minutes_until_reboot"]; ok {
			out.Reboot.MinutesUntilReboot = val.(int)
		}
		if val, ok := reboot["start_reboot_timer_immediately"]; ok {
			out.Reboot.StartRebootTimerImmediately = val.(bool)
		}
		if val, ok := reboot["file_vault_2_reboot"]; ok {
			out.Reboot.FileVault2Reboot = val.(bool)
		}
	}

	// Maintenance
	if m, ok := d.GetOk("maintenance"); ok {
		v := m.(*schema.Set).List()
		maintenance := v[0].(map[string]interface{})

		if val, ok := maintenance["recon"]; ok {
			out.Maintenance.Recon = val.(bool)
		}
		if val, ok := maintenance["reset_name"]; ok {
			out.Maintenance.ResetName = val.(bool)
		}
		if val, ok := maintenance["install_all_cached_packages"]; ok {
			out.Maintenance.InstallAllCachedPackages = val.(bool)
		}
		if val, ok := maintenance["heal"]; ok {
			out.Maintenance.Heal = val.(bool)
		}
		if val, ok := maintenance["prebindings"]; ok {
			out.Maintenance.Prebindings = val.(bool)
		}
		if val, ok := maintenance["permissions"]; ok {
			out.Maintenance.Permissions = val.(bool)
		}
		if val, ok := maintenance["byhost"]; ok {
			out.Maintenance.Byhost = val.(bool)
		}
		if val, ok := maintenance["system_cache"]; ok {
			out.Maintenance.SystemCache = val.(bool)
		}
		if val, ok := maintenance["user_cache"]; ok {
			out.Maintenance.UserCache = val.(bool)
		}
		if val, ok := maintenance["verify"]; ok {
			out.Maintenance.Verify = val.(bool)
		}
	}

	// Files and Processses
	if f, ok := d.GetOk("files_and_processes"); ok {
		v := f.(*schema.Set).List()
		filesAndProcesses := v[0].(map[string]interface{})

		if val, ok := filesAndProcesses["search_by_path"]; ok {
			out.FilesAndProcesses.SearchByPath = val.(string)
		}
		if val, ok := filesAndProcesses["delete_file"]; ok {
			out.FilesAndProcesses.DeleteFile = val.(bool)
		}
		if val, ok := filesAndProcesses["locate_file"]; ok {
			out.FilesAndProcesses.LocateFile = val.(string)
		}
		if val, ok := filesAndProcesses["update_locate_database"]; ok {
			out.FilesAndProcesses.UpdateLocateDatabase = val.(bool)
		}
		if val, ok := filesAndProcesses["spotlight_search"]; ok {
			out.FilesAndProcesses.SpotlightSearch = val.(string)
		}
		if val, ok := filesAndProcesses["search_for_process"]; ok {
			out.FilesAndProcesses.SearchForProcess = val.(string)
		}
		if val, ok := filesAndProcesses["kill_process"]; ok {
			out.FilesAndProcesses.KillProcess = val.(bool)
		}
		if val, ok := filesAndProcesses["run_command"]; ok {
			out.FilesAndProcesses.RunCommand = val.(string)
		}
	}

	// User Interaction
	if u, ok := d.GetOk("user_interaction"); ok {
		v := u.(*schema.Set).List()
		userInteraction := v[0].(map[string]interface{})

		if val, ok := userInteraction["message_start"]; ok {
			out.UserInteraction.MessageStart = val.(string)
		}
		if val, ok := userInteraction["allow_users_to_defer"]; ok {
			out.UserInteraction.AllowUsersToDefer = val.(bool)
		}
		if val, ok := userInteraction["allow_deferral_until_utc"]; ok {
			out.UserInteraction.AllowDeferralUntilUtc = val.(string)
		}
		if val, ok := userInteraction["allow_deferral_minutes"]; ok {
			out.UserInteraction.AllowDeferralMinutes = val.(int)
		}
		if val, ok := userInteraction["message_finish"]; ok {
			out.UserInteraction.MessageFinish = val.(string)
		}
	}

	return &out
}

func resourceJamfPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfPolicyStruct(d)
	id, err := c.CreatePolicy(b)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceJamfPolicyRead(ctx, d, m)
}

func resourceJamfPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	resp, err := c.GetPolicy(id)
	if err != nil {
		if jamfErr, ok := err.(jamf.Error); ok && jamfErr.StatusCode() == 404 {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	} else {
		deconstructJamfPolicyStruct(d, resp)
	}

	return diags
}

func resourceJamfPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b := buildJamfPolicyStruct(d)

	if _, err := c.UpdatePolicy(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfPolicyRead(ctx, d, m)
}

func resourceJamfPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err := c.DeletePolicy(id); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfPolicyState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, err
	}
	resp, err := c.GetPolicy(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get Computer Group data")
	}

	deconstructJamfPolicyStruct(d, resp)

	return []*schema.ResourceData{d}, nil
}
