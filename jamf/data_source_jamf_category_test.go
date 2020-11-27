package jamf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJamfCategoryDataSource_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDatasourceJamfCategoryConfigMissingFields,
				ExpectError: regexp.MustCompile("The argument \"name\" is required, but no definition was found."),
			},
		},
	})
}

const (
	testAccDatasourceJamfCategoryConfigMissingFields = `
data "jamf_category" "test" {
}`
)
