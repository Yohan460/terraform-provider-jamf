package jamf

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"jamf": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

// testAccPreCheck ...if non-nil, will be called before any test steps are executed.
// It is commonly used to verify that required values exist for testing,
// such as environment variables containing test keys that are used to
// configure the Provider or Resource under test.
func testAccPreCheck(t *testing.T) {
	if isClientIdSet() {
		if !isClientSecretSet() {
			t.Fatal("JAMF_CLIENT_SECRET environment variable must be set for acceptance tests (if JAMF_CLIENT_ID is set")
		}
	} else if isUsernameSet() {
		if !isPasswordSet() {
			t.Fatal("JAMF_PASSWORD environment variable must be set for acceptance tests (if JAMF_USERNAME is set)")
		}
	} else {
		t.Fatal("JAMF_USERNAME or JAMF_CLIENT_ID environment variable must be set for acceptance tests")
	}

	if !isURLSet() {
		t.Fatal("JAMF_URL environment variable must be set for acceptance tests")
	}
}

func isVarSet(variable string) bool {
	return os.Getenv(variable) != ""
}

func isUsernameSet() bool {
	return isVarSet("JAMF_USERNAME")
}

func isPasswordSet() bool {
	return isVarSet("JAMF_PASSWORD")
}

func isURLSet() bool {
	return isVarSet("JAMF_URL")
}

func isClientIdSet() bool {
	return isVarSet("JAMF_CLIENT_ID")
}

func isClientSecretSet() bool {
	return isVarSet("JAMF_CLIENT_SECRET")
}
