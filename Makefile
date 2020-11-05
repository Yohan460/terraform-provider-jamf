.PHONY: terraform/* build

##### Terraform
TERRAFORM            := bin/terraform
TERRAFORM_VERSION    := 0.13.4

### install
terraform/install: file         = terraform_$(TERRAFORM_VERSION)_$(OS_TYPE)_$(OS_ARCH).zip
terraform/install: download_url = https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/$(file)
terraform/install:
	rm -f bin/terraform
	curl -L -fsS --retry 2 -o bin/$(file) $(download_url) && \
	unzip -qq bin/$(file) -d $(BINDIR) && rm -f bin/$(file); \


##### Go
NAME := $(notdir $(PWD))
VERSION := 1.0.0

build: ## go build
	CGO_ENABLED=0 go build -o bin/$(NAME)_v$(VERSION) main.go