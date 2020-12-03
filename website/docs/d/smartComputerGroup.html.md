---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_smartComputerGroup"
description: |-
  Provides details about a smart computer group.
---

# Data Source: jamf_smartComputerGroup

Use this data source to get the smart computer group information.

The smart computer group data source allows access to details of a specific
smart computer group within Jamf.

## Example Usage

```hcl
data "jamf_smartComputerGroup" "test_smart_1" {
  name = "Test Smart 1"
}
```

## Argument Reference

The following arguments are supported:

* `name`     - (Required) The name of the group.
* `criteria` - (Optional) List of computers to to be in the group.
   * `priority`      - (Optional) Priority of the criteria to be evaluated, must start at zero and increase by one.
   * `and_or`        - (Optional) Join your search with `and` or `or`. Default: `and`
   * `name`          - (Optional) Name of the criteria to compare with.
   * `search_type`   - (Optional) The search comparator to use.
   * `search_value`  - (Optional) The search value to compare against.
   * `opening_paren` - (Optional) If the opening parenthese is enabled. Default `false`
   * `closing_paren` - (Optional) If the closing parenthese is enabled. Default `false`
* `site`    - (Optional) Site the computer group should be a member of.
   * `name`          - (Optional) Name of the site.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the smart computer group.
