package azure

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

var resourceMap = map[string]func(ResourceInfo){
	"Microsoft.Cache/Redis": visitRedisCache,
}

func newResourceInfo(r *resources.GenericResource) *ResourceInfo {
	info := new(ResourceInfo)
	info.Type = *r.Type
	info.Name = *r.Name
	info.GroupName = extractResourceGroupNameFromResourceID(*r.ID)
	info.SubscriptionID = extractSubscriptionIDFromResourceID(*r.ID)
	return info
}

func getClient(subscriptionID string) resources.Client {
	client := resources.NewClient(subscriptionID)
	client.Authorizer = GetAuthorizer()
	return client
}

// ResourceInfo carries attributes common to all resources in Azure
type ResourceInfo struct {
	SubscriptionID string
	GroupName      string
	Name           string
	Type           string
}

// GetVisitor finds a function to invoke for a given Azure resource
func (info *ResourceInfo) GetVisitor() (func(ResourceInfo), error) {
	visitor := resourceMap[info.Type]
	if visitor == nil {
		return nil, errors.New("no visitor for " + info.Type)
	}
	return visitor, nil
}

// GetResourcesInSubscription will retrieve all the Azure resources in the specified subscription
func GetResourcesInSubscription(subscriptionID string, settings *configuration.AppSettings) ([]ResourceInfo, error) {
	allResources := make([]ResourceInfo, 0)

	client := getClient(subscriptionID)

	context := context.Background()
	listResult, err := client.List(context, "", "", nil)
	if err != nil {
		return nil, err
	}

	for listResult.NotDone() {
		for _, r := range listResult.Values() {
			info := newResourceInfo(&r)
			allResources = append(allResources, *info)
		}
		err = listResult.Next()
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return allResources, err
	}

	return allResources, nil
}
