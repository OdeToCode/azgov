package azure

import (
	"testing"
)

func TestExtractResourceGroupNameFromResourceId(t *testing.T) {
	id := "/subscriptions/23adad27-b425-4a88-85a5-e60614bfde1b/resourceGroups/hereisthename/foo/bar/foo/bar"
	result := extractResourceGroupNameFromResourceID(id)

	if result != "hereisthename" {
		t.Errorf("did not expect a group name of %s", result)
	}
}

func TestExtractSubscriptionFromResourceId(t *testing.T) {
	id := "/subscriptions/23adad27-b425-4a88-85a5-e60614bfde1b/resourceGroups/hereisthename/foo/bar/foo/bar"
	result := extractSubscriptionIDFromResourceID(id)

	if result != "23adad27-b425-4a88-85a5-e60614bfde1b" {
		t.Errorf("did not expect an id of %s", result)
	}
}