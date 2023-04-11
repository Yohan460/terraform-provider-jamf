---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_computer_extension_attribute"
description: |-
  Provides details about a computer extension attribute.
---

# Resource: jamf_computer_extension_attribute

Provides a computer extension attribute.

## Example Usage

```hcl
resource "jamf_computer_extension_attribute" "test-extension-attribute-script" {
  name = "test-extension-attribute-script"
  script {
    script_contents = file("${path.module}/extension-attributes/script-1.sh")
  }
}
```

```hcl
resource "jamf_computer_extension_attribute" "test-extension-attribute-text-field" {
  name = "test-extension-attribute-text-field"
  text_field { }
}
```

```hcl
resource "jamf_computer_extension_attribute" "test-extension-attribute-popup-menu" {
  name = "test-extension-attribute-popup-menu"
  popup_menu {
    choices = ["choice1", "choice2"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `name`              - (Required) Name of the extension attribute.
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

Only one of script, text_field, or popup_menu is required in a resource block.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id`              - ID of the extension attribute.
* `text_field`      - Display a text field to collect inventory data.
  * `input_type`      - Placeholder for an otherwise empty block, not used.

## Notes

When using the `file_path` parameter the file name must be changed to trigger a Terraform update.