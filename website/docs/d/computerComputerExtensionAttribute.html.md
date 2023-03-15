---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_computer_extension_attribute"
description: |-
  Provides details about a computer extension attribute.
---

# Data Source: jamf_computer_extension_attribute

Use this data source to get the computer extension attribute information.

The computer extension attribute data source allows access to details of a specific computer extension attribute within Jamf.

## Example Usage

```hcl
resource "jamf_computer_extension_attribute" "test-extension-attribute-script" {
  name = "test-extension-attribute-script"
}
```

## Argument Reference

The following arguments are supported:

* `name`              - (Required) Name of the extension attribute.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id`              - ID of the extension attribute.
* `description`       - (Optional) Description of the extension attribute.
* `data_type`         - (Optional) The type of data collected, defaults to `String`
* `inventory_display` - (Optional) Where the extension attribute is displayed in Jamf, defaults to `Extension Attributes`.
* `script`            - (Optional) Run a script to collect inventory data.
  * `enabled`           - (Optional) Enables collecting inventory data, defaults to `true`.
  * `platform`          - (Optional) Not documented, defaults to `Mac`. 
  * `script_contents`   - (Optional) The contents of the script.
  * `file_path`         - (Optional) The path of a file containing the script contents.
* `text_field`        - (Optional) Display a text field to collect inventory data.
* `popup_menu`        - (Optional) Display a pop-up menu to collect inventory data.
  * `choices`           - (Optional) List of values in the popup menu.
