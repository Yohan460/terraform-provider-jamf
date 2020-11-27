---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_department"
description: |-
  Provides details about a department.
---

# Data Source: jamf_department

Use this data source to get the department information.

The Department data source allows access to details of a specific
department within Jamf.

## Example Usage

```hcl
data "jamf_department" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the department.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the department.
