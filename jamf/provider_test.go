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
	if !isUsernameSet() {
		t.Fatal("JAMF_USERNAME environment variable must be set for acceptance tests")
	}
	if !isPasswordSet() {
		t.Fatal("JAMF_PASSWORD environment variable must be set for acceptance tests")
	}
	if !isURLSet() {
		t.Fatal("JAMF_URL environment variable must be set for acceptance tests")
	}
}

func isUsernameSet() bool {
	if os.Getenv("JAMF_USERNAME") != "" {
		return true
	}
	return false
}

func isPasswordSet() bool {
	if os.Getenv("JAMF_PASSWORD") != "" {
		return true
	}
	return false
}

func isURLSet() bool {
	if os.Getenv("JAMF_URL") != "" {
		return true
	}
	return false
}
