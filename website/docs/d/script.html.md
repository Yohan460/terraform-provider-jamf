---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_script"
description: |-
  Provides details about a script.
---

# Data Sources: jamf_script

Use this data source to get the script information.

The script data source allows access to details of a specific script within Jamf.

## Example Usage

```hcl
data "jamf_script" "test_script_1" {
  name = "Test Script 1"
}
```

## Argument Reference

The following arguments are supported:

* `name`            - (Required) The name of the department.
* `category_id`     - (Optional) The id of the category the script is under, Default: `-1`.
* `priority`        - (Optional) The script execution priority, Default: `AFTER`.
* `info`            - (Optional) The info of the script.
* `notes`           - (Optional) The notes of the script.
* `parameter4`      - (Optional) The name of parameter 4.
* `parameter5`      - (Optional) The name of parameter 5.
* `parameter6`      - (Optional) The name of parameter 6.
* `parameter7`      - (Optional) The name of parameter 7.
* `parameter8`      - (Optional) The name of parameter 8.
* `parameter9`      - (Optional) The name of parameter 9.
* `parameter10`     - (Optional) The name of parameter 10.
* `os_requirements` - (Optional) The OS requirements of the script.
* `script_contents` - (Optional) The contents of the script if defined in the resources declaration.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id`            - The ID of the smart computer group.
* `category_name` - The name of the category the script is under.
