---
layout: "jamf"
page_title: "Provider: Jamf"
description: |-
  The Jamf provider is used to interact with the many resources supported by Jamf. The provider needs to be configured with the proper credentials before it can be used.
---

# Jamf Provider

The Jamf provider is used to interact with the
many resources supported by Jamf. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

Terraform 0.13 and later:

```hcl
terraform {
  required_providers {
    yohan460/jamf = {}
  }
}

provider "jamf" {
    username = "xxxx"
    password = "xxxx"

    # "This is the full url of jamf, xxxx.jamfcloud.com"
    url = "xxxx"
}

data "jamf_department" "example" {
    name = "hoge"
}
```

## Authentication

The Jamf provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static Credentials

!> **Warning:** Hard-coded credentials are not recommended in any Terraform
configuration and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `username`, `password` and `url`
in-line in the Jamf provider block:

Usage:

```hcl
provider "jamf" {
    username = "xxxx"
    password = "xxxx"

    # "This is the xxxx part of xxxx.jamfcloud.com"
    url = "xxxx"
}
```

Starting with Jamf Pro 10.49, you can alternatively provide a client ID and 
client secret from the 
["API Roles and Clients"](https://learn.jamf.com/bundle/jamf-pro-documentation-10.49.0/page/API_Roles_and_Clients.html) 
section of Jamf Pro preferences.

Usage:

```hcl
provider "jamf" {
    client_id = "xxxx"
    client_secret = "xxxx"

    # "This is the xxxx part of xxxx.jamfcloud.com"
    url = "xxxx"
}
```

### Environment Variables

You can provide your credentials via the `JAMF_USERNAME`, `JAMF_PASSWORD`, `JAMF_CLIENT_ID`, `JAMF_CLIENT_SECRET`, and
`JAMF_URL` environment variables.

```hcl
provider "jamf" {}
```

Usage:

```sh
$ export JAMF_USERNAME="xxxx"
$ export JAMF_PASSWORD="xxxx"
$ export JAMF_URL="xxxx"
$ terraform plan
```

Or, with API Clients (Jamf Pro 10.49+):

Usage:

```sh
$ export JAMF_CLIENT_ID="xxxx"
$ export JAMF_CLIENT_SECRET="xxxx"
$ export JAMF_URL="xxxx"
$ terraform plan
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the Jamf
 `provider` block:

* `username` - (Optional) This is the Jamf username.

* `password` - (Optional) This is the Jamf user password.

* `client_id` - (Optional) This is the Jamf API client ID.

* `client_secret` - (Optional) This is the Jamf API client secret.

* `url` - (Optional) This is the Jamf server url.
