package azure

import (
	"testing"

	"github.com/Azure/azure-amqp-common-go/uuid"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
)

func TestResourceMapCanProvideVisitor(t *testing.T) {

	id := "/subscriptions/23adad27-b425-4a88-85a5-e60614bfde1b/resourceGroups/hereisthename/foo/bar/foo/bar"
	resourceType := "Microsoft.Cache/Redis"
	name := "somename"

	generic := &resources.GenericResource{ID: &id, Type: &resourceType, Name: &name}
	run, _ := uuid.NewV4()
	info := newResourceInfo(generic, run.String())
	visitor, _ := info.GetVisitor()

	if visitor == nil {
		t.Error("could not find resource type visitor")
	}
}
