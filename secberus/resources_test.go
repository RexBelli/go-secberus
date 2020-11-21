package secberus

import "testing"

func TestGetResources(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Fatalf("could not create client: %s\n", err)
	}

	_, err = c.GetResources()
	if err != nil {
		t.Fatalf("could not get organizations: %s\n", err)
	}
}
