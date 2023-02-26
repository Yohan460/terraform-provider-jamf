package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
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
						"network_requirements": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"date_time_limitations": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"activation_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"activation_date_epoch": {
										Type:     schema.TypeInt,
										Computed: true,
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
										Computed: true,
									},
									"expiration_date_utc": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"no_execute_on": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"no_execute_start": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"no_execute_end": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"network_limitations": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"minimum_network_connection": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"any_ip_address": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"network_segments": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"override_default_settings": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_drive": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"distribution_point": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"force_afp_smb": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sus": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"netboot_server": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
					},
				},
			},
			"scope": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_computers": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"computers": {
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
									"udid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"computer_groups": {
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
						"buildings": {
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
						"departments": {
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
					},
				},
			},
			"self_service": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_for_self_service": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"self_service_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"install_button_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reinstall_button_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"self_service_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"force_users_to_view_description": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"feature_on_main_page": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"self_service_icon": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
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
									"display_in": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"feature_in": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"package": {
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
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fut": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"feu": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"update_autorun": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"script": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter4": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter5": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter6": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter7": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter8": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter9": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter10": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter11": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"reboot": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"startup_disk": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"specify_startup": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"no_user_logged_in": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_logged_in": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"minutes_until_reboot": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_reboot_timer_immediately": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"file_vault_2_reboot": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"maintenance": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recon": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"reset_name": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"install_all_cached_packages": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"heal": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"prebindings": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"permissions": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"byhost": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"system_cache": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"user_cache": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"verify": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"files_and_processes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"search_by_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"delete_file": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"locate_file": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_locate_database": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"spotlight_search": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"search_for_process": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kill_process": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"run_command": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"user_interaction": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_users_to_defer": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"allow_deferral_until_utc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_deferral_minutes": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"message_finish": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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

	general := map[string]interface{}{
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
		"category": []interface{}{
			map[string]interface{}{
				"id":   in.General.Category.ID,
				"name": in.General.Category.Name,
			},
		},
		"site": []interface{}{
			map[string]interface{}{
				"id":   in.General.Site.ID,
				"name": in.General.Site.Name,
			},
		},
	}

	if in.General.DateTimeLimitations != (jamf.PolicyGeneralDateTimeLimitations{}) {
		general["date_time_limitations"] = []interface{}{
			map[string]interface{}{
				"activation_date":       in.General.DateTimeLimitations.ActivationDate,
				"activation_date_epoch": in.General.DateTimeLimitations.ActivationDateEpoch,
				"activation_date_utc":   in.General.DateTimeLimitations.ActivationDateUtc,
				"expiration_date":       in.General.DateTimeLimitations.ExpirationDate,
				"expiration_date_epoch": in.General.DateTimeLimitations.ExpirationDateEpoch,
				"expiration_date_utc":   in.General.DateTimeLimitations.ExpirationDateUtc,
				"no_execute_on":         in.General.DateTimeLimitations.NoExecuteOn,
				"no_execute_start":      in.General.DateTimeLimitations.NoExecuteStart,
				"no_execute_end":        in.General.DateTimeLimitations.NoExecuteEnd,
			},
		}
	}

	if in.General.NetworkLimitations != (jamf.PolicyGeneralNetworkLimitations{}) {
		general["network_limitations"] = []interface{}{
			map[string]interface{}{
				"minimum_network_connection": in.General.NetworkLimitations.MinimumNetworkConnection,
				"any_ip_address":             in.General.NetworkLimitations.AnyIpAddress,
				"network_segments":           in.General.NetworkLimitations.NetworkSegments,
			},
		}
	}

	if in.General.OverrideDefaultSettings != (jamf.PolicyGeneralOverrideDefaultSettings{}) {
		general["override_default_settings"] = []interface{}{
			map[string]interface{}{
				"target_drive":       in.General.OverrideDefaultSettings.TargetDrive,
				"distribution_point": in.General.OverrideDefaultSettings.DistributionPoint,
				"force_afp_smb":      in.General.OverrideDefaultSettings.ForceAfpSmb,
				"sus":                in.General.OverrideDefaultSettings.Sus,
			},
		}
	}
	d.Set("general", []interface{}{general})

	// Scope
	scope := map[string]interface{}{
		"all_computers": in.Scope.AllComputers,
	}

	// Scope - Computers
	computers := []interface{}{}
	for _, v := range in.Scope.Computers {
		computers = append(computers, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
			"udid": v.UDID,
		})
	}
	scope["computer"] = computers

	// Scope - Computer Groups
	computerGroups := []interface{}{}
	for _, v := range in.Scope.ComputerGroups {
		computerGroups = append(computerGroups, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	scope["computer_group"] = computerGroups

	// Scope - Buildings
	buildings := []interface{}{}
	for _, v := range in.Scope.Buildings {
		buildings = append(buildings, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	scope["building"] = buildings

	// Scope - Departments
	departments := []interface{}{}
	for _, v := range in.Scope.Departments {
		departments = append(departments, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	scope["department"] = departments
	d.Set("scope", []interface{}{scope})

	// Self Service
	selfService := map[string]interface{}{
		"use_for_self_service":            in.SelfService.UseForSelfService,
		"self_service_display_name":       in.SelfService.SelfServiceDisplayName,
		"install_button_text":             in.SelfService.InstallButtonText,
		"reinstall_button_text":           in.SelfService.ReinstallButtonText,
		"self_service_description":        in.SelfService.SelfServiceDescription,
		"force_users_to_view_description": in.SelfService.ForceUsersToViewDescription,
		"feature_on_main_page":            in.SelfService.FeatureOnMainPage,
		"self_service_icon": []interface{}{
			map[string]interface{}{
				"id":       in.SelfService.SelfServiceIcon.ID,
				"filename": in.SelfService.SelfServiceIcon.Filename,
				"uri":      in.SelfService.SelfServiceIcon.URI,
			},
		},
	}

	// Self Service - Category
	if len(in.SelfService.SelfServiceCategories) != 0 {
		selfServiceCategory := []interface{}{}
		for _, v := range in.SelfService.SelfServiceCategories {
			selfServiceCategory = append(selfServiceCategory, map[string]interface{}{
				"id":         v.ID,
				"name":       v.Name,
				"display_in": v.DisplayIn,
				"feature_in": v.FeatureIn,
			})
		}
		selfService["self_service_category"] = selfServiceCategory
	}
	d.Set("self_service", []interface{}{selfService})

	// Packages
	if len(in.PackageConfiguration.Packages) != 0 {
		packages := []interface{}{}
		for _, v := range in.PackageConfiguration.Packages {
			packages = append(packages, map[string]interface{}{
				"id":             v.ID,
				"name":           v.Name,
				"action":         v.Action,
				"fut":            v.FillUserTemplate,
				"feu":            v.FillExistingUsers,
				"update_autorun": v.UpdateAutorun,
			})
		}
		d.Set("package", packages)
	}

	// Scripts
	if len(in.ScriptsConfiguration.Scripts) != 0 {
		scripts := []interface{}{}
		for _, v := range in.ScriptsConfiguration.Scripts {
			scripts = append(scripts, map[string]interface{}{
				"id":          v.ID,
				"name":        v.Name,
				"priority":    v.Priority,
				"parameter4":  v.Parameter4,
				"parameter5":  v.Parameter5,
				"parameter6":  v.Parameter6,
				"parameter7":  v.Parameter7,
				"parameter8":  v.Parameter8,
				"parameter9":  v.Parameter9,
				"parameter10": v.Parameter10,
				"parameter11": v.Parameter11,
			})
		}
		d.Set("script", scripts)
	}

	// Reboot
	if in.Reboot != (jamf.PolicyReboot{}) {
		d.Set("reboot", []interface{}{
			map[string]interface{}{
				"message":                        in.Reboot.Message,
				"startup_disk":                   in.Reboot.StartupDisk,
				"specify_startup":                in.Reboot.SpecifyStartup,
				"no_user_logged_in":              in.Reboot.NoUserLoggedIn,
				"user_logged_in":                 in.Reboot.UserLoggedIn,
				"minutes_until_reboot":           in.Reboot.MinutesUntilReboot,
				"start_reboot_timer_immediately": in.Reboot.StartRebootTimerImmediately,
				"file_vault_2_reboot":            in.Reboot.FileVault2Reboot,
			},
		})
	}

	// Maintenance
	if in.Maintenance != (jamf.PolicyMaintenance{}) {
		d.Set("maintenance", []interface{}{
			map[string]interface{}{
				"recon":                       in.Maintenance.Recon,
				"reset_name":                  in.Maintenance.ResetName,
				"install_all_cached_packages": in.Maintenance.InstallAllCachedPackages,
				"heal":                        in.Maintenance.Heal,
				"prebindings":                 in.Maintenance.Prebindings,
				"permissions":                 in.Maintenance.Permissions,
				"byhost":                      in.Maintenance.Byhost,
				"system_cache":                in.Maintenance.SystemCache,
				"user_cache":                  in.Maintenance.UserCache,
				"verify":                      in.Maintenance.Verify,
			},
		})
	}

	// Files and Processses
	if in.FilesAndProcesses != (jamf.PolicyFilesAndProcesses{}) {
		d.Set("files_and_processes", []interface{}{
			map[string]interface{}{
				"search_by_path":         in.FilesAndProcesses.SearchByPath,
				"delete_file":            in.FilesAndProcesses.DeleteFile,
				"locate_file":            in.FilesAndProcesses.LocateFile,
				"update_locate_database": in.FilesAndProcesses.UpdateLocateDatabase,
				"spotlight_search":       in.FilesAndProcesses.SpotlightSearch,
				"search_for_process":     in.FilesAndProcesses.SearchForProcess,
				"kill_process":           in.FilesAndProcesses.KillProcess,
				"run_command":            in.FilesAndProcesses.RunCommand,
			},
		})
	}

	// User Interaction
	if in.UserInteraction != (jamf.PolicyUserInteraction{}) {
		d.Set("user_interaction", []interface{}{
			map[string]interface{}{
				"message_start":            in.UserInteraction.MessageStart,
				"allow_users_to_defer":     in.UserInteraction.AllowUsersToDefer,
				"allow_deferral_until_utc": in.UserInteraction.AllowDeferralUntilUtc,
				"allow_deferral_minutes":   in.UserInteraction.AllowDeferralMinutes,
				"message_finish":           in.UserInteraction.MessageFinish,
			},
		})
	}

	return
}
