---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_computer_configuration_profile"
description: |-
  Provides details about a computer configuration profile.
---

# Data Source: jamf_computer_configuration_profile

Provides a computer configuration profile.

## Example Usage

```hcl
resource "jamf_computer_configuration_profile" "test_profile_1" {
  general {
    name = "test_profile_1"
    mobileconfig_path = "test profile.mobileconfig"
    category {}
    site {}
  }
  scope {
    limitation {}
    exclusion {}
  }
  self_service{
    self_service_icon{}
  }
}
```

### Signed Configuration Profiles

Currently signed configuration profile uploads are not supported by the Jamf API. Please file requests with them to have this be supported.

Once it's supported that functionality will be added into this terraform provider.

## Argument Reference

The following arguments are supported:

* `general`             - (Required) General information of the configuration profile
  * `name`                  - (Required) name of the configuration profile
  * `description`           - (Optional) description of the configuration profile
  * `distribution_method`   - (Optional) distribution method of the configuration profile, Default `Install Automatically`
  * `user_removable`        - (Optional) if the configuration profile is user removable, Default `false`
  * `level`                 - (Optional) installation level of the configuration profile, Default `computer`
  * `redeploy_on_update`    - (Optional) the configuration profiles redeploy on update behavior, Default `Newly Assigned`
  * `mobileconfig_path`     - (Optional) path to the mobileconfig payload to upload
  * `category`              - (Required) Category of the configuration profile
    * `id`                      - (Optional) ID of the category, Default `-1`
    * `name`                    - (Optional) Name of the category, Default `No category assigned`
  * `site`                  - (Required) Site of the configuration profile
    * `id`                      - (Optional) ID of the site, Default `-1`
* `scope`               - (Required) Scope of the configuration profile
  * `all_computers`         - (Optional) State of the configuration profile being scoped to all computers, Default `false`
  * `all_users`             - (Optional) State of the configuration profile being scoped to all users, Default `false`
  * `computer`              - (Optional) Details a computer the configuration profile is scoped to
    * `id`                      - (Required) ID of the computer
  * `computer_group`        - (Optional) Details a computer group the configuration profile is scoped to
    * `id`                      - (Required) ID of the computer group
  * `building`              - (Optional) Details a building the configuration profile is scoped to
    * `id`                      - (Required) ID of the building
  * `department`            - (Optional) Details a department the configuration profile is scoped to
    * `id`                      - (Required) ID of the department
  * `jamf_user`             - (Optional) Details a Jamf user the configuration profile is scoped to
    * `id`                      - (Required) ID of the Jamf user
  * `jamf_user_group`       - (Optional) Details a Jamf user group the configuration profile is scoped to
    * `id`                      - (Required) ID of the Jamf user group
  * `limitation`            - (Required) Details the scoping limitations of the configuration profile
    * `user`                    - (Optional) Details an LDAP user the configuration profile is scoped to
      * `name`                      - (Required) Name of the LDAP user
    * `user_group`              - (Optional) Details an LDAP user group the configuration profile is scoped to
      * `name`                      - (Required) ID of the LDAP user group
    * `network_segment`         - (Optional) Details a network segment the configuration profile is scoped to
      * `id`                        - (Required) ID of the network segment
    * `ibeacon`                 - (Optional) Details an iBeacon the configuration profile is scoped to
      * `id`                        - (Required) ID of the iBeacon
  * `exclusion`             - (Required) Details the scoping exclusions of the configuration profile
    * `computer`              - (Optional) Details a computer the configuration profile is scoped to
      * `id`                      - (Required) ID of the computer
    * `computer_group`        - (Optional) Details a computer group the configuration profile is scoped to
      * `id`                      - (Required) ID of the computer group
    * `building`              - (Optional) Details a building the configuration profile is scoped to
      * `id`                      - (Required) ID of the building
    * `department`            - (Optional) Details a department the configuration profile is scoped to
      * `id`                      - (Required) ID of the department
    * `jamf_user`             - (Optional) Details a Jamf user the configuration profile is scoped to
      * `id`                      - (Required) ID of the Jamf user
    * `jamf_user_group`       - (Optional) Details a Jamf user group the configuration profile is scoped to
      * `id`                      - (Required) ID of the Jamf user group
    * `user`                  - (Optional) Details an LDAP user the configuration profile is scoped to
      * `name`                    - (Required) Name of the LDAP user
    * `user_group`            - (Optional) Details an LDAP user group the configuration profile is scoped to
      * `name`                    - (Required) ID of the LDAP user group
    * `network_segment`       - (Optional) Details a network segment the configuration profile is scoped to
      * `id`                      - (Required) ID of the network segment
    * `ibeacon`               - (Optional) Details an iBeacon the configuration profile is scoped to
* `self_service`                  - (Required) Self Service configuration of the configuration profile
  * `self_service_display_name`       - (Optional) Display name, Default `Placeholder Name`
  * `install_button_text`             - (Optional) Install Button Text, Default `Install`
  * `self_service_description`        - (Optional) Description
  * `force_users_to_view_description` - (Optional) State of description viewing enforcement
  * `feature_on_main_page`            - (Optional) State of main page featuring
  * `removal_disallowed`              - (Optional) Removal disallowed, Default `Never`
  * `self_service_icon`               - (Required) Self Service icon configuration
    * `id`                                - (Optional) ID of the icon, Default `0`
  * `self_service_category`           - (Optional) Self Service category configuration
    * `id`                                - (Required) ID of the category
    * `display_in`                        - (Optional) Display state in the category
    * `feature_in`                        - (Optional) Feature state in the category

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id`              - ID of the configuration profile
* `general`             - General information of the configuration profile
  * `id`                    - ID of the configuration profile
  * `uuid`                  - uuid of the configuration profile
  * `site`                  - Site of the configuration profile
    * `name`                - Name of the site
* `scope`               - Scope of the configuration profile
  * `computer`              - Details a computer the configuration profile is scoped to
    * `name`                    - Name of the computer
    * `udid`                    - UDID of the computer
  * `computer_group`        - Details a computer group the configuration profile is scoped to
    * `name`                    - Name of the computer group
  * `building`              - Details a building the configuration profile is scoped to
    * `name`                    - Name of the building
  * `department`            - Details a department the configuration profile is scoped to
    * `name`                    - Name of the department
  * `jamf_user`             - Details a Jamf user the configuration profile is scoped to
    * `name`                    - Name of the Jamf user
  * `jamf_user_group`       - Details a Jamf user group the configuration profile is scoped to
    * `name`                    - Name of the Jamf user group
  * `limitation`            - Details the scoping limitations of the configuration profile
    * `network_segment`         - Details a network segment the configuration profile is scoped to
      * `name`                      - Name of the network segment
    * `ibeacon`                 - Details an iBeacon the configuration profile is scoped to
      * `name`                      - Name of the iBeacon
  * `exclusion`             - Details the scoping exclusions of the configuration profile
    * `computer`              - Details a computer the configuration profile is scoped to
      * `name`                    - Name of the computer
      * `udid`                    - UDID of the computer
    * `computer_group`        - Details a computer group the configuration profile is scoped to
      * `name`                    - Name of the computer group
    * `building`              - Details a building the configuration profile is scoped to
      * `name`                    - Name of the building
    * `department`            - Details a department the configuration profile is scoped to
      * `name`                    - Name of the department
    * `jamf_user`             - Details a Jamf user the configuration profile is scoped to
      * `id`                      - ID of the Jamf user
      * `name`                    - Name of the Jamf user
    * `jamf_user_group`       - Details a Jamf user group the configuration profile is scoped to
      * `id`                      - ID of the Jamf user group
      * `name`                    - Name of the Jamf user group
    * `network_segment`       - Details a network segment the configuration profile is scoped to
      * `name`                    - Name of the network segment
    * `ibeacon`               - Details an iBeacon the configuration profile is scoped to
      * `name`                    - Name of the iBeacon
* `self_service`                  - Self Service configuration of the configuration profile
  * `self_service_icon`               - Self Service icon configuration
    * `filename`                          - Filename of the icon
    * `uri`                               - URI of the icon
  * `self_service_category`           - Self Service category configuration
    * `name`                              - Name of the category
