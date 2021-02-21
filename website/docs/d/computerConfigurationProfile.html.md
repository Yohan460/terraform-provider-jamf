---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_computer_configuration_profile"
description: |-
  Provides details about a computer configuration profile.
---

# Data Source: jamf_computer_configuration_profile

Use this data source to get the computer configuration profile information.

The computer configuration profile data source allows access to details of a specific
computer configuration profile within Jamf.

## Example Usage

```hcl
data "jamf_computer_configuration_profile" "test_profile_1" {
  name = "Test Profile 1"
}

```

## Argument Reference

The following arguments are supported:

* `name`     - (Required) The name of the configuration profile

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - ID of the configuration profile
* `general` - General information of the configuration profile
  * `id` - ID of the configuration profile
  * `name` - name of the configuration profile
  * `description` - description of the configuration profile
  * `distribution_method` - distribution method of the configuration profile
  * `user_removable` - if the configuration profile is user removable
  * `level` - installation level of the configuration profile
  * `uuid` - uuid of the configuration profile
  * `redeploy_on_update` - the configuration profiles redeploy on update behavior
  * `payload` - payload of the configuration profile
  * `category` - Category of the configuration profile
    * `id` - ID of the category
    * `name` - Name of the category
  * `site` - Site of the configuration profile
    * `id` - ID of the site
    * `name` - Name of the site
* `scope` - Scope of the configuration profile
  * `all_computers` - State of the configuration profile being scoped to all computers
  * `all_users` - State of the configuration profile being scoped to all users
  * `computer` - Details a computer the configuration profile is scoped to
    * `id` - ID of the computer
    * `name` - Name of the computer
    * `udid` - UDID of the computer
  * `computer_group` - Details a computer group the configuration profile is scoped to
    * `id` - ID of the computer group
    * `name`- Name of the computer group
  * `building` - Details a building the configuration profile is scoped to
    * `id` - ID of the building
    * `name` - Name of the building
  * `department` - Details a department the configuration profile is scoped to
    * `id` - ID of the department
    * `name` - Name of the department
  * `jamf_user` - Details a Jamf user the configuration profile is scoped to
    * `id` - ID of the Jamf user
    * `name` - Name of the Jamf user
  * `jamf_user_group` - Details a Jamf user group the configuration profile is scoped to
    * `id` - ID of the Jamf user group
    * `name` - Name of the Jamf user group
  * `limitation` - Details the scoping limitations of the configuration profile
    * `user` - Details an LDAP user the configuration profile is scoped to
      * `name` - Name of the LDAP user
    * `user_group` - Details an LDAP user group the configuration profile is scoped to
      * `name` - Name of the LDAP user group
    * `network_segment` - Details a network segment the configuration profile is scoped to
      * `id` - ID of the network segment
      * `name` - Name of the network segment
    * `ibeacon` - Details an iBeacon the configuration profile is scoped to
      * `id` - ID of the iBeacon
      * `name` - Name of the iBeacon
  * `exclusion` - Details the scoping exclusions of the configuration profile
    * `computer` - Details a computer the configuration profile is scoped to
      * `id` - ID of the computer
      * `name` - Name of the computer
      * `udid` - UDID of the computer
    * `computer_group` - Details a computer group the configuration profile is scoped to
      * `id` - ID of the computer group
      * `name`- Name of the computer group
    * `building` - Details a building the configuration profile is scoped to
      * `id` - ID of the building
      * `name` - Name of the building
    * `department` - Details a department the configuration profile is scoped to
      * `id` - ID of the department
      * `name` - Name of the department
    * `jamf_user` - Details a Jamf user the configuration profile is scoped to
      * `id` - ID of the Jamf user
      * `name` - Name of the Jamf user
    * `jamf_user_group` - Details a Jamf user group the configuration profile is scoped to
      * `id` - ID of the Jamf user group
      * `name` - Name of the Jamf user group
    * `user` - Details an LDAP user the configuration profile is scoped to
      * `name` - Name of the LDAP user
    * `user_group` - Details an LDAP user group the configuration profile is scoped to
      * `id` - ID of the LDAP user group
    * `network_segment` - Details a network segment the configuration profile is scoped to
      * `id` - ID of the network segment
      * `name` - Name of the network segment
    * `ibeacon` - Details an iBeacon the configuration profile is scoped to
      * `id` - ID of the iBeacon
      * `name` - Name of the iBeacon
* `self_service` - Self Service configuration of the configuration profile
  * `self_service_display_name` - Display name
  * `install_button_text` - Install Button Text
  * `self_service_description` - Description
  * `force_users_to_view_description` - State of description viewing enforcement
  * `feature_on_main_page` - State of main page featuring
  * `removal_disallowed` - If removal of the profile is not allowed
  * `self_service_icon` - Self Service icon configuration
    * `id` - ID of the icon
    * `filename` - Filename of the icon
    * `uri` - URI of the icon
  * `self_service_category` - Self Service category configuration
    * `id` - ID of the category
    * `name` - Name of the category
    * `display_in` - Display state in the category
    * `feature_in`  - Feature state in the category
