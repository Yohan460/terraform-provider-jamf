---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_policy"
description: |-
  Provides details about a policy.
---

# Data Source: jamf_policy

Provides a policy.

## Example Usage

```hcl
resource "jamf_policy" "test_policy_1" {
  general {
    name = "Test Policy 1"
    category {
      id = "-1"
      name = "No category assigned"
    }
    network_limitations {}
    override_default_settings {}
    site {}
  }
  self_service {
    self_service_icon {}
  }
  scope {}
  reboot {}
}
```

## Argument Reference

The following arguments are supported:

* `general`                     - (Required) General information of the Policy
  * `name`                          - (Required) The name of the policy
  * `enabled`                       - (Optional) Enabled state of the policy
  * `trigger`                       - (Optional) Trigger of the policy, Default `EVENT`
  * `trigger_checkin`               - (Optional) Check-in trigger state
  * `trigger_enrollment_complete`   - (Optional) Enrollment Complete trigger state
  * `trigger_login`                 - (Optional) Login trigger state
  * `trigger_logout`                - (Optional) Logout trigger state
  * `trigger_network_state_changed` - (Optional) Network state change trigger state
  * `trigger_startup`               - (Optional) Startup trigger state
  * `trigger_other`                 - (Optional) Custom trigger event
  * `frequency`                     - (Optional) Frequency of the policy, Default `Once per computer`
  * `retry_event`                   - (Optional) Retry event of the policy, Default `none`
  * `retry_attempts`                - (Optional) Number of retry events for the policy, Default `-1`
  * `notify_on_each_failed_retry`   - (Optional) Notify state of the retry event for the policy
  * `location_user_only`            - (Optional) ?
  * `target_drive`                  - (Optional) Target drive of the policy, Default `/`
  * `offline`                       - (Optional) State of the policy being available offline
  * `network_requirements`          - (Optional) Network requirements of the policy, Default `Any`
  * `category`                  - (Required) Category of the policy
    * `id`                          - (Required) ID of the category, for default set to `-1`
    * `name`                        - (Required) Name of the category, for default set to `No category assigned`
  * `date_time_limitations`     - (Optional) Date and Time limitations of the policy
    * `activation_date_epoch`       - (Optional) Activation date in epoch, Default `0`
    * `expiration_date_epoch`       - (Optional) Expiration date in epoch, Default `0`
    * `no_execute_on`               - (Optional) ?
    * `no_execute_start`            - (Optional) ?
    * `no_execute_end`              - (Optional) ?
  * `network_limitations`       - (Required) Network Limitations of the policy
    * `minimum_network_connection`  - (Optional) Minimum network connections required, Default `No Minimum`
    * `any_ip_address`              - (Optional) IP Address range required, Default `true`
    * `network_segments`            - (Optional) Network Segments required
  * `override_default_settings` - (Required) Default settings of the policy to override
    * `target_drive`                - (Optional) Default target installation drive, Default `/`
    * `distribution_point`          - (Optional) Default distribution point, Default `default`
    * `force_afp_smb`               - (Optional) ?
    * `sus`                         - (Optional) ?, Default `default`
  * `site`                      - (Required) Site of the policy
    * `id`                          - (Optional) ID of the site, Default `-1`
* `scope`                       - (Required) Scope of the policy
  * `all_computers`                 - (Optional) State of the policy being scoped to all computers
  * `computer`                      - (Optional) Details a computer the policy is scoped to
    * `id`                              - (Required) ID of the computer
  * `computer_group`                - (Optional) Details a computer group the policy is scoped to
    * `id`                              - (Required) ID of the computer group
  * `building`                      - (Optional) Details a building the policy is scoped to
    * `id`                              - (Required) ID of the building
  * `department`                    - (Optional) Details a department the policy is scoped to
    * `id`                              - (Required) ID of the department
* `self_service`                - (Required) Self Service configuration of the policy
  * `use_for_self_service`          - (Optional) Self Service enablement state
  * `self_service_display_name`     - (Optional) Display name
  * `install_button_text`           - (Optional) Install Button Text, Default `Install`
  * `reinstall_button_text`         - (Optional) Reinstall Button Text, Default `Reinstall`
  * `self_service_description`      - (Optional) Description
  * `force_users_to_view_description` - (Optional) State of description viewing enforcement
  * `feature_on_main_page`          - (Optional) State of main page featuring
  * `self_service_icon`             - (Required) Self Service icon configuration
    * `id`                              - (Optional) ID of the icon, Default `0`
  * `self_service_category`         - (Optional) Self Service category configuration and this block can have multiple settings, if the policy category is defined this will need to be set to the same category at least one
    * `id`                              - (Optional) ID of the category
    * `display_in`                      - (Optional) Display state in the category, Default `true`
    * `feature_in`                      - (Optional) Feature state in the category, Default `false`
* `package`                     - (Optional) Package information assigned to the policy
  * `id`                            - (Required) ID of the package
  * `action`                        - (Optional) Action of the package, Default `INSTALL`
  * `fut`                           - (Optional) Fill user template state of the DMG
  * `feu`                           - (Optional) Fill existing users state of the DMG
  * `update_autorun`                - (Optional) Update autorun state of the DMG
* `script `                     - (Optional) Script information assigned to the policy
  * `id`                            - (Required) ID of the script
  * `name`                          - (Optional) Name of the script
  * `priority`                      - (Optional) Priority of the script, Default `AFTER`
  * `parameter4`                    - (Optional) Parameter 4 of the script
  * `parameter5`                    - (Optional) Parameter 5 of the script
  * `parameter6`                    - (Optional) Parameter 6 of the script
  * `parameter7`                    - (Optional) Parameter 7 of the script
  * `parameter8`                    - (Optional) Parameter 8 of the script
  * `parameter9`                    - (Optional) Parameter 9 of the script
  * `parameter10`                   - (Optional) Parameter 10 of the script
  * `parameter11`                   - (Optional) Parameter 11 of the script
* `reboot`                      - (Required) Reboot information assigned to the policy
  * `message`                       - (Optional) User message
  * `startup_disk`                  - (Optional) Startup disk after reboot, Default `Current Startup Disk`
  * `specify_startup`               - (Optional) Local startup disk to boot computers when startup disk is `Specify Local Startup Disk`
  * `no_user_logged_in`             - (Optional) No user logged in functionality of the reboot, Default `Do not restart`
  * `user_logged_in`                - (Optional) User logged in functionality of the reboot, Default `Do not restart`
  * `minutes_until_reboot`          - (Optional) Minutes until the reboot triggers after reboot, Default `5`
  * `start_reboot_timer_immediately`- (Optional) Reboot timer immediate state, Default `false`
  * `file_vault_2_reboot`           - (Optional) Restart FileVault 2-encrypted computers without requiring an unlock during the next startup, Default `false`
* `maintenance`                 - (Optional) Maintenance information assigned to the policy
  * `recon`                         - (Optional) Recon state after policy completes
  * `reset_name`                    - (Optional) Reset name state
  * `install_all_cached_packages`   - (Optional) Cached package install state
  * `heal`                          - (Optional) ?
  * `prebindings`                   - (Optional) ?
  * `permissions`                   - (Optional) Disk permissions repair state
  * `byhost`                        - (Optional) ByHost repair state
  * `system_cache`                  - (Optional) Flush system cache state
  * `user_cache`                    - (Optional) Flush user cache state
  * `verify`                        - (Optional) Verify startup disk state
* `files_and_processes`         - (Optional) File and Process information assigned to the policy
  * `search_by_path`                - (Optional) Path to search for file
  * `delete_file`                   - (Optional) Deletion state of a file found by `search_by_path`
  * `locate_file`                   - (Optional) Path to search for file
  * `update_locate_database`        - (Optional) Local database update state of a file found by `locate_file`
  * `spotlight_search`              - (Optional) Spotlight search for file
  * `search_for_process`            - (Optional) Process to search for
  * `kill_process`                  - (Optional) Kill state for process found by `search_for_process`
  * `run_command`                   - (Optional) Command to execute as `root`
* `user_interaction`            - (Optional) User Interaction information assigned to the policy
  * `message_start`                 - (Optional) Message at start of policy
  * `message_finish`                - (Optional) Message at end of policy
  * `allow_users_to_defer`          - (Optional) Deferral state
  * `allow_deferral_until_utc`      - (Optional) Deferral until date in UTC
  * `allow_deferral_minutes`        - (Optional) Minutes to defer

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - ID of the policy
* `general` - General information of the Policy
  * `id` - ID of the policy
  * `date_time_limitations` - Date and Time limitations of the policy
    * `activation_date` - Activation date written out
    * `activation_date_utc` - Activation date written out in UTC
    * `expiration_date` - Expiration date written out
    * `expiration_date_utc` - Expiration date written out in UTC
  * `site` - Site of the policy
    * `name` - Name of the site
* `scope` - Scope of the policy
  * `computer` - Details a computer the policy is scoped to
    * `name` - Name of the computer
    * `udid` - UDID of the computer
  * `computer_group` - Details a computer group the policy is scoped to
    * `name`- Name of the computer group
  * `building` - Details a building the policy is scoped to
    * `name` - Name of the building
  * `department` - Details a department the policy is scoped to
    * `name` - Name of the department
* `self_service` - Self Service configuration of the policy
  * `self_service_icon` - Self Service icon configuration
    * `filename` - Filename of the icon
    * `uri` - URI of the icon
  * `self_service_category` - Self Service category configuration
    * `name` - Name of the category
* `package` - Package information assigned to the policy
  * `name` - Name of the package
* `script `- Script information assigned to the policy
  * `name` - Name of the script
