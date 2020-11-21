package secberus

import (
	"fmt"
	"testing"
)

func TestGetOrganizations(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Fatalf("could not create client: %s\n", err)
	}

	orgs, err := c.GetOrganizations()
	if err != nil {
		t.Fatalf("could not get organizations: %s\n", err)
	}

	for _, org := range *orgs {
		fmt.Printf("%+v\n", org)
	}
}
