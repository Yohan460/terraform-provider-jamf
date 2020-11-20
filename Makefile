.PHONY: terraform/* build fmtcheck test testacc

##### Terraform
OS_TYPE              := $(shell echo $(shell uname) | tr A-Z a-z)
OS_ARCH              := amd64
BINDIR               := $(realpath bin/)

TERRAFORM            := bin/terraform
TERRAFORM_VERSION    := 0.13.4

### install
terraform/install: file         = terraform_$(TERRAFORM_VERSION)_$(OS_TYPE)_$(OS_ARCH).zip
terraform/install: download_url = https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/$(file)
terraform/install:
	rm -f bin/terraform
	curl -L -fsS --retry 2 -o $(BINDIR)/$(file) $(download_url) && \
	unzip -qq $(BINDIR)/$(file) -d $(BINDIR) && rm -f $(BINDIR)/$(file)


##### Go
NAME    := $(notdir $(PWD))
VERSION := 1.0.0
TEST    ?= $$(go list ./... |grep -v 'vendor')

build: ## go build
	CGO_ENABLED=0 go build -o $(BINDIR)/$(NAME)_v$(VERSION) main.go

test:
	go test $(TEST) -v  -timeout=5m -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test ./jamf -v -count 1 -parallel 20 $(TESTARGS) -timeout 120m

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"