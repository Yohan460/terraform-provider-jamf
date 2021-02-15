<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Jamf


## Quick Start

- [Using the provider](https://registry.terraform.io/providers/yohan460/jamf/latest/docs)
- [Provider development](docs/DEVELOPMENT.md)

```hcl
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

## Documentation

Full, comprehensive documentation is available on the Terraform website:

https://registry.terraform.io/providers/yohan460/jamf/latest/docs

## Release

```shell
$ gGPG_FINGERPRINT=xxx GITHUB_TOKEN="xxx" goreleaser release --rm-dist
```
