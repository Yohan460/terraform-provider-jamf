.PHONY: terraform/* build fmtcheck test testacc

##### Terraform
OS_TYPE              := $(shell echo $(shell uname) | tr A-Z a-z)
OS_ARCH              := amd64
OS_BUILD              := ${OS_TYPE}_amd64
BINDIR               := ~/bin

TERRAFORM            := ~/bin/terraform
TERRAFORM_VERSION    := 1.6.5

### install
terraform/install: file         = terraform_$(TERRAFORM_VERSION)_$(OS_TYPE)_$(OS_ARCH).zip
terraform/install: download_url = https://releases.hashicorp.com/terraform/$(TERRAFORM_VERSION)/$(file)
terraform/install:
	rm -f ${TERRAFORM}
	curl -L -fsS --retry 2 -o $(BINDIR)/$(file) $(download_url) && \
	unzip -qq $(BINDIR)/$(file) -d $(BINDIR) && rm -f $(BINDIR)/$(file)


##### Go
NAME    := $(notdir $(PWD))
VERSION := 1.0.0
TEST    ?= $$(go list ./... |grep -v 'vendor')

build: ## go build
	CGO_ENABLED=0 go build -o $(BINDIR)/$(NAME)_v$(VERSION) main.go

install: build
	mkdir -p ${BINDIR}/plugins/registry.terraform.io/hashicorp/jamf/${VERSION}/${OS_BUILD}
	mv $(BINDIR)/$(NAME)_v$(VERSION) ${BINDIR}/plugins/registry.terraform.io/hashicorp/jamf/${VERSION}/${OS_BUILD}

release:
	mkdir -p ${BINDIR}/release/${VERSION}
	GOOS=darwin GOARCH=amd64 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ${BINDIR}/release/${VERSION}/${NAME}_${VERSION}_windows_amd64

test:
	go test $(TEST) -v  -timeout=5m -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test ./jamf -v -count 1 -parallel 20 $(TESTARGS) -timeout 120m

fmtcheck:
	@sh -c "'$(CURDIR)/website/docs/scripts/gofmtcheck.sh'"
