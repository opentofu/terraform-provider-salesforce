package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUser_basic(t *testing.T) {
	t.Parallel()
	t.Skip("Users cannot be deleted and there are limited licenses, skipping, comment out this line to run locally")

	email := os.Getenv("SALESFORCE_USERNAME")
	parts := strings.Split(email, "@")
	var username string
	if !strings.Contains(parts[0], "+") {
		username = fmt.Sprintf("%s+%s@%s", parts[0], acctest.RandString(10), parts[1])
	} else {
		username = fmt.Sprintf("%s-%s@%s", parts[0], acctest.RandString(10), parts[1])
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUser_basic(email, username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("salesforce_user.test", "alias", "test"),
				),
			},
		},
	})
}

func testAccResourceUser_basic(email, username string) string {
	return fmt.Sprintf(`
data "salesforce_profile" "chatter_free" {
  name = "Chatter Free User"
}

resource "salesforce_user" "test" {
  alias = "test"
  email = "%s"
  last_name = "test"
  username = "%s"
  profile_id = data.salesforce_profile.chatter_free.id
}
`, email, username)
}
