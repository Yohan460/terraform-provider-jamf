---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_policy"
description: |-
  Provides details about a policy.
---

# Data Source: jamf_policy

Use this data source to get the policy information.

The policy data source allows access to details of a specific
policy within Jamf.

## Example Usage

```hcl
data "jamf_policy" "test_policy_1" {
  name = "Test Policy 1"
}
```

## Argument Reference

The following arguments are supported:

* `name`     - (Required) The name of the policy

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - ID of the policy
* `general` - General information of the Policy
  * `id` - ID of the policy
  * `name` - name of the policy
  * `enabled` - Enabled state of the policy
  * `trigger` - Trigger of the policy
  * `trigger_checkin` - Check-in trigger state
  * `trigger_enrollment_complete`- Enrollment Complete trigger state
  * `trigger_login` - Login trigger state
  * `trigger_logout` - Logout trigger state
  * `trigger_network_state_changed` - Network state change trigger state
  * `trigger_startup` - Startup trigger state
  * `trigger_other` - Custom trigger event
  * `frequency` - Frequency of the policy
  * `retry_event` - Retry event of the policy
  * `retry_attempts` - Number of retry events for the policy
  * `notify_on_each_failed_retry` - Notify state of the retry event for the policy
  * `location_user_only` - ?
  * `target_drive` - Target drive of the policy
  * `offline` - State of the policy being available offline
  * `network_requirements` - Network requirements of the policy
  * `category` - Category of the policy
    * `id` - ID of the category
    * `name` - Name of the category
  * `date_time_limitations` - Date and Time limitations of the policy
    * `activation_date` - Activation date written out
    * `activation_date_epoch` - Activation date in epoch
    * `activation_date_utc` - Activation date written out in UTC
    * `expiration_date` - Expiration date written out
    * `expiration_date_epoch` - Expiration date in epoch
    * `expiration_date_utc` - Expiration date written out in UTC
    * `no_execute_on` - ?
    * `no_execute_start` - ?
    * `no_execute_end` - ?
  * `network_limitations` - Network Limitations of the policy
    * `minimum_network_connection` - Minimum network connections required
    * `any_ip_address` - IP Address range required
    * `network_segments` - Network Segments required
  * `override_default_settings` - Default settings of the policy to override
    * `target_drive` - Default target installation drive
    * `distribution_point` - Default distribution point
    * `force_afp_smb` - ?
    * `sus` - ?
    * `netboot_server` - ?
  * `site` - Site of the policy
    * `id` - ID of the site
    * `name` - Name of the site
* `scope` - Scope of the policy
  * `all_computers` - State of the policy being scoped to all computers
  * `computer` - Details a computer the policy is scoped to
    * `id` - ID of the computer
    * `name` - Name of the computer
    * `udid` - UDID of the computer
  * `computer_group` - Details a computer group the policy is scoped to
    * `id` - ID of the computer group
    * `name`- Name of the computer group
  * `building` - Details a building the policy is scoped to
    * `id` - ID of the building
    * `name` - Name of the building
  * `department` - Details a department the policy is scoped to
    * `id` - ID of the department
    * `name` - Name of the department
* `self_service` - Self Service configuration of the policy
  * `use_for_self_service` - Self Service enablement state
  * `self_service_display_name` - Display name
  * `install_button_text` - Install Button Text
  * `reinstall_button_text`- Reinstall Button Text
  * `self_service_description` - Description
  * `force_users_to_view_description` - State of description viewing enforcement
  * `feature_on_main_page` - State of main page featuring
  * `self_service_icon` - Self Service icon configuration
    * `id` - ID of the icon
    * `filename` - Filename of the icon
    * `uri` - URI of the icon
  * `self_service_category` - Self Service category configuration
    * `id` - ID of the category
    * `name` - Name of the category
    * `display_in` - Display state in the category
    * `feature_in`  - Feature state in the category
* `package` - Package information assigned to the policy
  * `id` - ID of the package
  * `name` - Name of the package
  * `action` - Action of the package
  * `fut` - Fill user template state of the DMG
  * `feu` - Fill existing users state of the DMG
  * `update_autorun` - Update autorun state of the DMG
* `script `- Script information assigned to the policy
  * `id` - ID of the script
  * `name` - Name of the script
  * `priority` - Priority of the script
  * `parameter4` - Parameter 4 of the script
  * `parameter5` - Parameter 5 of the script
  * `parameter6` - Parameter 6 of the script
  * `parameter7` - Parameter 7 of the script
  * `parameter8` - Parameter 8 of the script
  * `parameter9` - Parameter 9 of the script
  * `parameter10` - Parameter 10 of the script
  * `parameter11` - Parameter 11 of the script
* `reboot` - Reboot information assigned to the policy
  * `message` - User message
  * `startup_disk` - Startup disk after reboot
  * `specify_startup` - ?
  * `no_user_logged_in` - No user logged in functionality of the reboot
  * `user_logged_in` - User logged in functionality of the reboot
  * `minutes_until_reboot` - Minutes until the reboot triggers after reboot
  * `start_reboot_timer_immediately` - Reboot timer immediate state
  * `file_vault_2_reboot` - ?
* `maintenance` - Maintenance information assigned to the policy
  * `recon` - Recon state after policy completes
  * `reset_name` - Reset name state
  * `install_all_cached_packages` - Cached package install state
  * `heal` - ?
  * `prebindings` - ?
  * `permissions` - Disk permissions repair state
  * `byhost` - ByHost repair state
  * `system_cache` - Flush system cache state
  * `user_cache` - Flush user cache state
  * `verify` - Verify startup disk state
* `files_and_processes` - File and Process information assigned to the policy
  * `search_by_path` - Path to search for file
  * `delete_file` - Deletion state of a file found by `search_by_path`
  * `locate_file` - Path to search for file
  * `update_locate_database` - Local database update state of a file found by `locate_file`
  * `spotlight_search` - Spotlight search for file
  * `search_for_process` - Process to search for
  * `kill_process` - Kill state for process found by `search_for_process`
  * `run_command` - Command to execute as `root`
* `user_interaction` - User Interaction information assigned to the policy
  * `message_start` - Message at start of policy
  * `message_finish` - Message at end of policy
  * `allow_users_to_defer` - Deferral state
  * `allow_deferral_until_utc` - Deferral until date in UTC
  * `allow_deferral_minutes` - Minutes to defer
