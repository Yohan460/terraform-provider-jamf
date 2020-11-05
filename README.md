# Terraform Provider for Jamf

## Quick Start

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

## Release

```shell
$ goreleaser
```

## Development

```shell
### local install terraform & build provider to bin/
$ make terraform/install
$ make build

# create .tf file
$ vim jamf.tf
terraform {
  required_providers {
    jamf = {
    }
  }
}

provider "jamf" {
    username = "xxxx"
    password = "xxxx"

    # "This is the xxxx part of xxxx.jamfcloud.com"
    organization = "xxxx"
}

data "jamf_department" "example" {
    name = "hoge"
}

# Please prepare the following directory structure and files
$ tree
.
├── jamf.tf
├── plugins
│   └── registry.terraform.io
│       └── hashicorp
│           └── jamf
│               └── 1.0.0
│                   └── darwin_amd64
│                       └── terraform-provider-jamf_v1.0.0
└── terraform


# After moving to the top directory
# exec (-plugin-dir is full path)
$ ./terraform init -plugin-dir=/....../plugins/
$ ./terraform plan
```
