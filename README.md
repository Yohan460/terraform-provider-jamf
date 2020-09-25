# Terraform Provider for Jamf

# Quick Start

```hcl
provider "jamf" {
    username = "xxxx"
    password = "xxxx"

    # "This is the xxxx part of xxxx.jamfcloud.com"
    organization = "xxxx"
}

data "jamf_department" "example" {
    name = "hoge"
}
```

