---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_script"
description: |-
  Provides a script.
---

# Resource: jamf_script

Provides a script.

## Example Usage

```hcl
resource "jamf_script" "test_script_1" {
  name = "Test Script 1"
  file_path = "example_script-v1.0.0.sh"
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
* `file_path`       - (Optional) The path of a file to be read in as the script contents.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id`            - The ID of the smart computer group.
* `category_name` - The name of the category the script is under.

## Notes

When using the `file_path` parameter the file name must be changed to trigger a terraform update.
