package ovmhelper

import (
	"testing"
)

func TestAccRepos(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	t.Run("GetIdFromName", testAccGetIDFromName(c))
}

func testAccGetIDFromName(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		i, err := c.Repos.GetIDFromName("ovm-corp-repository")

		if err != nil {
			t.Fatalf("[error] Could not get ID: %s", err)
		}

		if i.Name != "ovm-corp-repository" {
			t.Fatalf("expected: ovm-corp-repository; got %s", i.Name)
		}
	}
}
