---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_department"
description: |-
  Provides an department.
---

# Resource: jamf_department

Provides an Department.

## Example Usage

```hcl
resource "jamf_department" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the department.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the department.
