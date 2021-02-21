package jamf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yohan460/go-jamf-api"
)

func resourceJamfComputerConfigurationProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJamfComputerConfigurationProfileCreate,
		ReadContext:   resourceJamfComputerConfigurationProfileRead,
		UpdateContext: resourceJamfComputerConfigurationProfileUpdate,
		DeleteContext: resourceJamfComputerConfigurationProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importJamfComputerConfigurationProfileState,
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"distribution_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Install Automatically",
						},
						"user_removable": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"level": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "computer",
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"redeploy_on_update": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Newly Assigned",
						},
						"mobileconfig_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"category": {
							Type:     schema.TypeSet,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  -1,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "No category assigned",
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
						"all_users": {
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
						"jamf_user": {
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
						"jamf_user_group": {
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
						"limitation": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"user_group": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"network_segment": {
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
									"ibeacon": {
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
						"exclusion": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"user": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"user_group": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"jamf_user": {
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
									"jamf_user_group": {
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
									"network_segment": {
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
									"ibeacon": {
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
					},
				},
			},
			"self_service": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"self_service_display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Placeholder Name",
						},
						"install_button_text": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Install",
						},
						"self_service_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"force_users_to_view_description": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"feature_on_main_page": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"removal_disallowed": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Never",
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
									},
									"feature_in": {
										Type:     schema.TypeBool,
										Optional: true,
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

func buildJamfComputerConfigurationProfileStruct(d *schema.ResourceData) (*jamf.ComputerConfigurationProfile, error) {
	var out jamf.ComputerConfigurationProfile

	// General
	id, _ := strconv.Atoi(d.Id())
	out.General.ID = id
	if g, ok := d.GetOk("general"); ok {
		v := g.(*schema.Set).List()
		general := v[0].(map[string]interface{})

		out.General.Name = general["name"].(string)

		if val, ok := general["description"]; ok {
			out.General.Description = val.(string)
		}
		if val, ok := general["distribution_method"]; ok {
			out.General.DistributionMethod = val.(string)
		}
		if val, ok := general["user_removable"]; ok {
			out.General.UserRemovable = val.(bool)
		}
		if val, ok := general["level"]; ok {
			out.General.Level = val.(string)
		}
		if val, ok := general["redeploy_on_update"]; ok {
			out.General.RedeployOnUpdate = val.(string)
		}

		// MobileConfig Payload
		if val, ok := general["mobileconfig_path"]; ok {
			if str := val.(string); str != "" {
				content, err := loadFileContent(str)
				if err != nil {
					return &out, err
				}
				out.General.Payload = content
			}
		}

		// General - Category
		if v, ok := general["category"]; ok {
			categoryList := v.(*schema.Set).List()
			if len(categoryList) > 0 {
				category := categoryList[0].(map[string]interface{})
				if val, ok := category["name"].(string); ok {
					out.General.Category.Name = val
				}
				if val, ok := category["id"].(string); ok {
					out.General.Category.ID = val
				}
			}
		}

		// General - Site
		if v, ok := general["site"]; ok {
			siteList := v.(*schema.Set).List()
			if len(siteList) > 0 {
				site := siteList[0].(map[string]interface{})
				if val, ok := site["name"].(string); ok {
					out.General.Site.Name = val
				}
				if val, ok := site["id"].(int); ok {
					out.General.Site.ID = val
				}
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
		if val, ok := scope["all_users"]; ok {
			out.Scope.AllUsers = val.(bool)
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

		// Scope - Jamf User
		if v, ok := scope["jamf_user"]; ok {
			jamfUsers := v.(*schema.Set).List()
			jamfUserList := []jamf.JamfUserScope{}
			for _, c := range jamfUsers {
				jamfUserData := c.(map[string]interface{})
				jamfUser := jamf.JamfUserScope{}
				if val, ok := jamfUserData["id"].(int); ok {
					jamfUser.Id = val
				}
				jamfUserList = append(jamfUserList, jamfUser)
			}
			out.Scope.JamfUsers = jamfUserList
		}

		// Scope - Jamf User Group
		if v, ok := scope["jamf_user_group"]; ok {
			JamfUserGroups := v.(*schema.Set).List()
			JamfUserGroupList := []jamf.UserGroupScope{}
			for _, c := range JamfUserGroups {
				JamfUserGroupData := c.(map[string]interface{})
				JamfUserGroup := jamf.UserGroupScope{}
				if val, ok := JamfUserGroupData["id"].(int); ok {
					JamfUserGroup.Id = val
				}
				JamfUserGroupList = append(JamfUserGroupList, JamfUserGroup)
			}
			out.Scope.JamfUserGroups = JamfUserGroupList
		}

		// Scope - Limitiations
		if l, ok := scope["limitation"]; ok {
			v := l.(*schema.Set).List()
			if len(v) > 0 {
				limitations := v[0].(map[string]interface{})

				// Scope - Limitiations - User
				if v, ok := limitations["user"]; ok {
					Users := v.(*schema.Set).List()
					UserList := []jamf.UserScope{}
					for _, c := range Users {
						UserData := c.(map[string]interface{})
						User := jamf.UserScope{}
						if val, ok := UserData["name"].(string); ok {
							User.Name = val
						}
						UserList = append(UserList, User)
					}
					out.Scope.Limitiations.Users = UserList
				}

				// Scope - Limitiations - User Group
				if v, ok := limitations["user_group"]; ok {
					UserGroups := v.(*schema.Set).List()
					UserGroupList := []jamf.UserGroupScope{}
					for _, c := range UserGroups {
						UserGroupData := c.(map[string]interface{})
						UserGroup := jamf.UserGroupScope{}
						if val, ok := UserGroupData["name"].(string); ok {
							UserGroup.Name = val
						}
						UserGroupList = append(UserGroupList, UserGroup)
					}
					out.Scope.Limitiations.UserGroups = UserGroupList
				}

				// Scope - Limitiations - Network Segment
				if v, ok := limitations["network_segment"]; ok {
					NetworkSegments := v.(*schema.Set).List()
					NetworkSegmentList := []jamf.NetworkSegmentScope{}
					for _, c := range NetworkSegments {
						NetworkSegmentData := c.(map[string]interface{})
						NetworkSegment := jamf.NetworkSegmentScope{}
						if val, ok := NetworkSegmentData["id"].(int); ok {
							NetworkSegment.ID = val
						}
						NetworkSegmentList = append(NetworkSegmentList, NetworkSegment)
					}
					out.Scope.Limitiations.NetworkSegments = NetworkSegmentList
				}

				// Scope - Limitiations - iBeacon
				if v, ok := limitations["ibeacon"]; ok {
					iBeacons := v.(*schema.Set).List()
					iBeaconList := []jamf.IBeaconScope{}
					for _, c := range iBeacons {
						iBeaconData := c.(map[string]interface{})
						iBeacon := jamf.IBeaconScope{}
						if val, ok := iBeaconData["id"].(int); ok {
							iBeacon.Id = val
						}
						iBeaconList = append(iBeaconList, iBeacon)
					}
					out.Scope.Limitiations.IBeacons = iBeaconList
				}
			}
		}

		// Scope - Exclusions
		if e, ok := scope["exclusion"]; ok {
			v := e.(*schema.Set).List()
			if len(v) > 0 {
				exclusions := v[0].(map[string]interface{})

				// Scope - Exclusions - User
				if v, ok := exclusions["user"]; ok {
					Users := v.(*schema.Set).List()
					UserList := []jamf.UserScope{}
					for _, c := range Users {
						UserData := c.(map[string]interface{})
						User := jamf.UserScope{}
						if val, ok := UserData["name"].(string); ok {
							User.Name = val
						}
						UserList = append(UserList, User)
					}
					out.Scope.Exclusions.Users = UserList
				}

				// Scope - Exclusions - User Group
				if v, ok := exclusions["user_group"]; ok {
					UserGroups := v.(*schema.Set).List()
					UserGroupList := []jamf.UserGroupScope{}
					for _, c := range UserGroups {
						UserGroupData := c.(map[string]interface{})
						UserGroup := jamf.UserGroupScope{}
						if val, ok := UserGroupData["name"].(string); ok {
							UserGroup.Name = val
						}
						UserGroupList = append(UserGroupList, UserGroup)
					}
					out.Scope.Exclusions.UserGroups = UserGroupList
				}

				// Scope - Exclusions - Network Segment
				if v, ok := exclusions["network_segment"]; ok {
					NetworkSegments := v.(*schema.Set).List()
					NetworkSegmentList := []jamf.NetworkSegmentScope{}
					for _, c := range NetworkSegments {
						NetworkSegmentData := c.(map[string]interface{})
						NetworkSegment := jamf.NetworkSegmentScope{}
						if val, ok := NetworkSegmentData["id"].(int); ok {
							NetworkSegment.ID = val
						}
						NetworkSegmentList = append(NetworkSegmentList, NetworkSegment)
					}
					out.Scope.Exclusions.NetworkSegments = NetworkSegmentList
				}

				// Scope - Exclusions - iBeacon
				if v, ok := exclusions["ibeacon"]; ok {
					iBeacons := v.(*schema.Set).List()
					iBeaconList := []jamf.IBeaconScope{}
					for _, c := range iBeacons {
						iBeaconData := c.(map[string]interface{})
						iBeacon := jamf.IBeaconScope{}
						if val, ok := iBeaconData["id"].(int); ok {
							iBeacon.Id = val
						}
						iBeaconList = append(iBeaconList, iBeacon)
					}
					out.Scope.Exclusions.IBeacons = iBeaconList
				}

				// Scope - Exclusions - Computers
				if v, ok := exclusions["computer"]; ok {
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
					out.Scope.Exclusions.Computers = computerList
				}

				// Scope - Exclusions - Computer Groups
				if v, ok := exclusions["computer_group"]; ok {
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
					out.Scope.Exclusions.ComputerGroups = computerGroupList
				}

				// Scope - Exclusions - Buildings
				if v, ok := exclusions["building"]; ok {
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
					out.Scope.Exclusions.Buildings = buildingList
				}

				// Scope - Exclusions - Departments
				if v, ok := exclusions["department"]; ok {
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
					out.Scope.Exclusions.Departments = departmentList
				}

				// Scope - Exclusions - Jamf User
				if v, ok := exclusions["jamf_user"]; ok {
					jamfUsers := v.(*schema.Set).List()
					jamfUserList := []jamf.JamfUserScope{}
					for _, c := range jamfUsers {
						jamfUserData := c.(map[string]interface{})
						jamfUser := jamf.JamfUserScope{}
						if val, ok := jamfUserData["id"].(int); ok {
							jamfUser.Id = val
						}
						jamfUserList = append(jamfUserList, jamfUser)
					}
					out.Scope.Exclusions.JamfUsers = jamfUserList
				}

				// Scope - Exclusions - Jamf User Group
				if v, ok := exclusions["jamf_user_group"]; ok {
					JamfUserGroups := v.(*schema.Set).List()
					JamfUserGroupList := []jamf.UserGroupScope{}
					for _, c := range JamfUserGroups {
						JamfUserGroupData := c.(map[string]interface{})
						JamfUserGroup := jamf.UserGroupScope{}
						if val, ok := JamfUserGroupData["id"].(int); ok {
							JamfUserGroup.Id = val
						}
						JamfUserGroupList = append(JamfUserGroupList, JamfUserGroup)
					}
					out.Scope.Exclusions.JamfUserGroups = JamfUserGroupList
				}
			}
		}
	}

	// Self Service
	if s, ok := d.GetOk("self_service"); ok {
		v := s.(*schema.Set).List()
		selfService := v[0].(map[string]interface{})

		if val, ok := selfService["self_service_display_name"]; ok {
			out.SelfService.SelfServiceDisplayName = val.(string)
		}
		if val, ok := selfService["install_button_text"]; ok {
			out.SelfService.InstallButtonText = val.(string)
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
		if val, ok := selfService["removal_disallowed"]; ok {
			out.SelfService.RemovalDisallowed = val.(string)
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

	return &out, nil
}

func resourceJamfComputerConfigurationProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b, err := buildJamfComputerConfigurationProfileStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}
	id, err := c.CreateComputerConfigurationProfile(b)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceJamfComputerConfigurationProfileRead(ctx, d, m)
}

func resourceJamfComputerConfigurationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	resp, err := c.GetComputerConfigurationProfile(id)
	if err != nil {
		if jamfErr, ok := err.(jamf.Error); ok && jamfErr.StatusCode() == 404 {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	} else {
		deconstructJamfComputerConfigurationProfileStruct(d, resp)
	}

	return diags
}

func resourceJamfComputerConfigurationProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*jamf.Client)

	b, err := buildJamfComputerConfigurationProfileStruct(d)
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err := c.UpdateComputerConfigurationProfile(b); err != nil {
		return diag.FromErr(err)
	}

	return resourceJamfComputerConfigurationProfileRead(ctx, d, m)
}

func resourceJamfComputerConfigurationProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*jamf.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if _, err := c.DeleteComputerConfigurationProfile(id); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func importJamfComputerConfigurationProfileState(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	c := m.(*jamf.Client)
	d.SetId(d.Id())
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return nil, err
	}
	resp, err := c.GetComputerConfigurationProfile(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get Computer Group data")
	}

	deconstructJamfComputerConfigurationProfileStruct(d, resp)

	return []*schema.ResourceData{d}, nil
}
