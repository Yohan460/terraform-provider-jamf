---
layout: "jamf"
subcategory: "Data Sources"
page_title: "Jamf: jamf_package"
description: |-
  Provides details about a package.
---

# Data Source: jamf_package

Use this data source to get the package information.

The package data source allows access to details of a specific
package within Jamf.

## Example Usage

```hcl
data "jamf_package" "google_chrome" {
  name = "GoogleChrome-v87.0.4280.88.pkg"
}
```

## Argument Reference

The following arguments are supported:

* `name`      - (Required) The name of the package.

## Attributes Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The ID of the package.
* `category_name` - The name of the category of which the package pertains.
* `filename` - The name of the package file, Default: `ThisPackageDoesNotExist.pkg`.
* `notes` - The notes of the package.
* `priority` - The priority of the package, Default: `10`
* `fill_existing_users` - If the DMG should fill existing users.
* `boot_volume_required` - If the package should be installed on the boot volume after imaging.
* `allow_uninstalled` - If the package should be allowed to be uninstalled using Jamf remote.
* `os_requirements` - The Os requirements of the package.
* `required_processor` - The required processor of the package, Default: `None`.
* `hash_type` - The hash type of the package.
* `hash_value` - The hash value of the package.
