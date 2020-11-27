---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_building"
description: |-
  Provides an building.
---

# Resource: jamf_building

Provides an Building.

## Example Usage

```hcl
resource "jamf_building" "example" {
  name     = "example"

  street_address1 = "1-1-1"
  street_address2 = "example Building"
  city            = "Shibuya-ku"
  state_province  = "Tokyo"
  zip_postal_code = "111-1111"
  country         = "Japan"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the building.
* `street_address1` - (Optional) The street address1 of the building.
* `street_address2` - (Optional) The street address2 of the building.
* `city` - (Optional) The vity of the building.
* `state_province` - (Optional) The state province of the building. 
* `zip_postal_code` - (Optional) The zip postal code of the building.
* `country` - (Optional) The country of the building.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the building.
