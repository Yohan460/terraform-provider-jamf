---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_category"
description: |-
  Provides details about a category.
---

# Data Source: jamf_category

Use this data source to get the category information.

The Category data source allows access to details of a specific
category within Jamf.

## Example Usage

```hcl
data "jamf_category" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the category.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the category.
* `priority` - The Priority of the category.