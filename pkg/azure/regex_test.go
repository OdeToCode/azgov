package azure

import (
	"testing"
)

func TestExtractResourceGroupNameFromResourceId(t *testing.T) {
	id := "/subscriptions/23asp310-b425-4a88-85a5-e60614bfde1b/resourceGroups/hereisthename/foo/bar/foo/bar"
	result := extractResourceGroupNameFromResourceID(id)

	if result != "hereisthename" {
		t.Errorf("did not expect a group name of %s", result)
	}
}

func TestExtractSubscriptionFromResourceId(t *testing.T) {
	id := "/subscriptions/23asp310-b425-4a88-85a5-e60614bfde1b/resourceGroups/hereisthename/foo/bar/foo/bar"
	result := extractSubscriptionIDFromResourceID(id)

	if result != "23asp310-b425-4a88-85a5-e60614bfde1b" {
		t.Errorf("did not expect an id of %s", result)
	}
}

func TestExtractResourceUri(t *testing.T) {
	instanceData := `{"Microsoft.Resources":{"resourceUri":"/subscriptions/23asp310-b425-4a88-85a5-e60614bfde1b/resourceGroups/msc-snapshots-prod/providers/Microsoft.Storage/storageAccounts/mscsnapshotbackups","location":"useast"}}`
	result := extractResourceUri(instanceData)

	if result != "/subscriptions/23asp310-b425-4a88-85a5-e60614bfde1b/resourceGroups/msc-snapshots-prod/providers/Microsoft.Storage/storageAccounts/mscsnapshotbackups" {
		t.Errorf("did not expect an id of %s", result)
	}

}
