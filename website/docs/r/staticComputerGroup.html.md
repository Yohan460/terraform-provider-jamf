---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_staticComputerGroup"
description: |-
  Provides a static computer group.
---

# Resource: jamf_staticComputerGroup

Provides a static computer group.

## Example Usage

```hcl
resource "jamf_staticComputerGroup" "test_static_1" {
  computer {
    id = 1
  }
  computer {
    serial_number = "C06X4489JFD3"
  }
  name = "Test Static 1"
}
```

## Argument Reference

The following arguments are supported:

* `name`      - (Required) The name of the group.
* `computer`  - (Optional) List of computers to to be in the group.
   * `id`            - (Optional) Jamf ID of the computer.
   * `serial_number` - (Optional) Serial Number of the computer.
* `site`      - (Optional) Site the computer group should be a member of.
   * `name`          - (Required) Name of the site.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the static computer group.
