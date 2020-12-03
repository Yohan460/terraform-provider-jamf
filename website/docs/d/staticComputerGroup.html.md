---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_staticComputerGroup"
description: |-
  Provides details about a static computer group.
---

# Data Source: jamf_staticComputerGroup

Use this data source to get the static computer group information.

The static computer group data source allows access to details of a specific
static computer group within Jamf.

## Example Usage

```hcl
data "jamf_staticComputerGroup" "test_static_1" {
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
   * `name`          - (Optional) Name of the site.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the static computer group.
