---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_building"
description: |-
  Provides details about a building.
---

# Data Source: jamf_building

Use this data source to get the building information.

The building data source allows access to details of a specific
building within Jamf.

## Example Usage

```hcl
data "jamf_building" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the building.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the building.
* `street_address1` - (Optional) The street address1 of the building.
* `street_address2` - (Optional) The street address2 of the building.
* `city` - (Optional) The vity of the building.
* `state_province` - (Optional) The state province of the building. 
* `zip_postal_code` - (Optional) The zip postal code of the building.
* `country` - (Optional) The country of the building.
