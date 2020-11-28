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

## Testing

There are test and accTest.

test is the usual test　The accTest actually makes the resource and checks it.

So if you want to do accTest, you need to have the jamf environment to run it.

```shell
# test
$ make test

# acctest
$ JAMF_USERNAME=xxx JAMF_PASSWORD=xxx JAMF_URL=xxx make testacc

# Running a specific test in acctest
$ JAMF_USERNAME=xxx JAMF_PASSWORD=xxx JAMF_URL=xxx make testacc TESTARGS="-run TestAccJamfDepartments_basic"
```
