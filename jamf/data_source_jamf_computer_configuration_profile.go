package jamf

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func dataSourceJamfComputerConfigurationProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJamfComputerConfigurationProfileRead,
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribution_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_removable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"level": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"redeploy_on_update": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"payload": {
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
						"all_users": {
							Type:     schema.TypeBool,
							Computed: true,
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
									"udid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"computer_group": {
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
						"building": {
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
						"department": {
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
						"jamf_user": {
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
						"jamf_user_group": {
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
						"limitation": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"user_group": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"network_segment": {
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
									"ibeacon": {
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
						"exclusion": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
												"udid": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"computer_group": {
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
									"building": {
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
									"department": {
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
									"user": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"user_group": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"jamf_user": {
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
									"jamf_user_group": {
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
									"network_segment": {
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
									"ibeacon": {
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
					},
				},
			},
			"self_service": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"self_service_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"install_button_text": {
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
						"removal_disallowed": {
							Type:     schema.TypeString,
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
		},
	}
}

func dataSourceJamfComputerConfigurationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	resp, err := c.GetComputerConfigurationProfileByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	deconstructJamfComputerConfigurationProfileStruct(d, resp)

	return diags
}

func deconstructJamfComputerConfigurationProfileStruct(d *schema.ResourceData, in *jamf.ComputerConfigurationProfile) {

	// General
	d.SetId(strconv.Itoa(in.General.ID))

	general := map[string]interface{}{}
	if generalInterface, ok := d.GetOk("general"); ok {
		generalSet := generalInterface.(*schema.Set)
		generalList := generalSet.List()
		general = generalList[0].(map[string]interface{})
	}

	general["id"] = in.General.ID
	general["name"] = in.General.Name
	general["description"] = in.General.Description
	general["distribution_method"] = in.General.DistributionMethod
	general["user_removable"] = in.General.UserRemovable
	general["level"] = in.General.Level
	general["uuid"] = in.General.UUID
	general["redeploy_on_update"] = in.General.RedeployOnUpdate
	if _, hasMobileConfigPayload := general["mobileconfig_path"]; !hasMobileConfigPayload {
		general["payload"] = in.General.Payload
	}
	general["category"] = []interface{}{
		map[string]interface{}{
			"id":   in.General.Category.ID,
			"name": in.General.Category.Name,
		},
	}
	general["site"] = []interface{}{
		map[string]interface{}{
			"id":   in.General.Site.ID,
			"name": in.General.Site.Name,
		},
	}

	d.Set("general", []interface{}{general})

	// Scope
	scope := map[string]interface{}{
		"all_computers": in.Scope.AllComputers,
		"all_users":     in.Scope.AllUsers,
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

	// Scope - Jamf Users
	jamfUsers := []interface{}{}
	for _, v := range in.Scope.JamfUsers {
		jamfUsers = append(jamfUsers, map[string]interface{}{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	scope["jamf_user"] = jamfUsers

	// Scope - Jamf User Groups
	jamfUserGroups := []interface{}{}
	for _, v := range in.Scope.JamfUserGroups {
		jamfUserGroups = append(jamfUserGroups, map[string]interface{}{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	scope["jamf_user_group"] = jamfUserGroups

	// Scope - Limitiations
	limitations := map[string]interface{}{}

	// Scope - Limitiations - Users
	users := []interface{}{}
	for _, v := range in.Scope.Limitiations.Users {
		users = append(users, map[string]interface{}{
			"name": v.Name,
		})
	}
	limitations["user"] = users

	// Scope - Limitiations - User Groups
	userGroups := []interface{}{}
	for _, v := range in.Scope.Limitiations.UserGroups {
		userGroups = append(userGroups, map[string]interface{}{
			"name": v.Name,
		})
	}
	limitations["user_group"] = userGroups

	// Scope - Limitiations - Network Segments
	networkSegments := []interface{}{}
	for _, v := range in.Scope.Limitiations.NetworkSegments {
		networkSegments = append(networkSegments, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	limitations["network_segment"] = networkSegments

	// Scope - Limitiations - iBeacons
	iBeacons := []interface{}{}
	for _, v := range in.Scope.Limitiations.IBeacons {
		iBeacons = append(iBeacons, map[string]interface{}{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	limitations["ibeacon"] = iBeacons
	scope["limitation"] = []interface{}{limitations}

	// Scope - Exclusions
	exclusions := map[string]interface{}{}

	// Scope - Exclusions - Users
	users = []interface{}{}
	for _, v := range in.Scope.Exclusions.Users {
		users = append(users, map[string]interface{}{
			"name": v.Name,
		})
	}
	exclusions["user"] = users

	// Scope - Exclusions - User Groups
	userGroups = []interface{}{}
	for _, v := range in.Scope.Exclusions.UserGroups {
		userGroups = append(userGroups, map[string]interface{}{
			"name": v.Name,
		})
	}
	exclusions["user_group"] = userGroups

	// Scope - Exclusions - Network Segments
	networkSegments = []interface{}{}
	for _, v := range in.Scope.Exclusions.NetworkSegments {
		networkSegments = append(networkSegments, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	exclusions["network_segment"] = networkSegments

	// Scope - Exclusions - iBeacons
	iBeacons = []interface{}{}
	for _, v := range in.Scope.Exclusions.IBeacons {
		iBeacons = append(iBeacons, map[string]interface{}{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	exclusions["ibeacon"] = iBeacons

	// Scope - Exclusions - Computers
	computers = []interface{}{}
	for _, v := range in.Scope.Computers {
		computers = append(computers, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
			"udid": v.UDID,
		})
	}
	exclusions["computer"] = computers

	// Scope - Exclusions - Computer Groups
	computerGroups = []interface{}{}
	for _, v := range in.Scope.Exclusions.ComputerGroups {
		computerGroups = append(computerGroups, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	exclusions["computer_group"] = computerGroups

	// Scope - Exclusions - Buildings
	buildings = []interface{}{}
	for _, v := range in.Scope.Exclusions.Buildings {
		buildings = append(buildings, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	exclusions["building"] = buildings

	// Scope - Exclusions - Departments
	departments = []interface{}{}
	for _, v := range in.Scope.Exclusions.Departments {
		departments = append(departments, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}
	exclusions["department"] = departments

	// Scope - Exclusions - Jamf Users
	jamfUsers = []interface{}{}
	for _, v := range in.Scope.Exclusions.JamfUsers {
		jamfUsers = append(jamfUsers, map[string]interface{}{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	exclusions["jamf_user"] = jamfUsers

	// Scope - Exclusions - Jamf User Groups
	jamfUserGroups = []interface{}{}
	for _, v := range in.Scope.Exclusions.JamfUserGroups {
		jamfUserGroups = append(jamfUserGroups, map[string]interface{}{
			"id":   v.Id,
			"name": v.Name,
		})
	}
	exclusions["jamf_user_group"] = jamfUserGroups
	scope["exclusion"] = []interface{}{exclusions}
	d.Set("scope", []interface{}{scope})

	// Self Service
	selfService := map[string]interface{}{
		"self_service_display_name":       in.SelfService.SelfServiceDisplayName,
		"install_button_text":             in.SelfService.InstallButtonText,
		"self_service_description":        in.SelfService.SelfServiceDescription,
		"force_users_to_view_description": in.SelfService.ForceUsersToViewDescription,
		"feature_on_main_page":            in.SelfService.FeatureOnMainPage,
		"removal_disallowed":              in.SelfService.RemovalDisallowed,
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

	return
}
