package jamf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccJamfComputerExtensionAttribute_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckJamfComputerExtensionAttributeScript,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJamfComputerExtensionAttributeExists("jamf_computer_extension_attribute.extensionattribute-script"),
				),
			},
			{
				Config: testAccCheckJamfComputerExtensionAttributeTextField,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJamfComputerExtensionAttributeExists("jamf_computer_extension_attribute.extensionattribute-textfield"),
				),
			},
			{
				Config: testAccCheckJamfComputerExtensionAttributePopup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJamfComputerExtensionAttributeExists("jamf_computer_extension_attribute.extensionattribute-popup"),
				),
			},
		},
	})
}

const (
	testAccCheckJamfComputerExtensionAttributeScript = `resource "jamf_computer_extension_attribute" "extensionattribute-script" {
		name = "Terraform test script"
		description = "testing jamf extension attribute resource"
		data_type = "String"
		inventory_display = "Extension Attributes" 

		script {
			enabled = false
			script_contents = "#!/bin/bash\nprint(\"hello world\")"
		}
	}`

	testAccCheckJamfComputerExtensionAttributeTextField = `resource "jamf_computer_extension_attribute" "extensionattribute-textfield" {
		name = "Terraform test textfield"
		text_field {}
	}`

	testAccCheckJamfComputerExtensionAttributePopup = `resource "jamf_computer_extension_attribute" "extensionattribute-popup" {
		name = "Terraform test popup"
		popup_menu {
			choices = ["choice1", "choice2"]
		}
	}`
)

func testAccCheckJamfComputerExtensionAttributeExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		extensionattribute, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if extensionattribute.Primary.ID == "" {
			return fmt.Errorf("No resource id set")
		}

		return nil
	}
}
