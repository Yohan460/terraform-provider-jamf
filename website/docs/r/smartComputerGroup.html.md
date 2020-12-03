---
layout: "jamf"
subcategory: "Resources"
page_title: "Jamf: jamf_smartComputerGroup"
description: |-
  Provides a smart computer group.
---

# Resource: jamf_smartComputerGroup

Provides a smart computer group.

## Example Usage

```hcl
resource "jamf_smartComputerGroup" "test_smart_1" {
  name = "Test Smart 1"
  criteria {
    priority = 0
    name = "UDID"
    search_type = "is not"
    search_value = "FAKE-UDID-THAT-ALSO-DOES-NOT-EXIST"
  }
  criteria {
    priority = 1
    name = "UDID"
    search_type = "is not"
    search_value = "FAKE-UDID-THAT-DOES-NOT-EXIST-LIKE-REALLY"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name`     - (Required) The name of the group.
* `criteria` - (Optional) List of computers to to be in the group.
   * `priority`      - (Required) Priority of the criteria to be evaluated, must start at zero and increase by one.
   * `and_or`        - (Optional) Join your search with `and` or `or`. Default: `and`
   * `name`          - (Required) Name of the criteria to compare with.
   * `search_type`   - (Required) The search comparator to use.
   * `search_value`  - (Required) The search value to compare against.
   * `opening_paren` - (Optional) If the opening parenthese is enabled. Default `false`
   * `closing_paren` - (Optional) If the closing parenthese is enabled. Default `false`
* `site`    - (Optional) Site the computer group should be a member of.
   * `name`          - (Required) Name of the site.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the smart computer group.
