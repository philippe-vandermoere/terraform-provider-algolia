package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//nolint:unparam
var providerFactories = map[string]func() (*schema.Provider, error){
	"algolia": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("ALGOLIA_APPLICATION_ID") == "" {
		t.Fatal("env variable 'ALGOLIA_APPLICATION_ID' is not set")
	}
	if os.Getenv("ALGOLIA_API_KEY") == "" {
		t.Fatal("env variable 'ALGOLIA_API_KEY' is not set")
	}
}
