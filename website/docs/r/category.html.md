---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_category"
description: |-
  Provides an category.
---

# Resource: jamf_category

Provides an Category.

## Example Usage

```hcl
resource "jamf_category" "example" {
  name     = "example"
  priority = 9
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the category.
* `priority` - (Required) The name of the priority.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the category.
