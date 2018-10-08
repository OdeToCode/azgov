package azure

import (
	"context"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

var resourceMap = map[string]func(ResourceInfo){
	"Microsoft.Cache/Redis": visitRedisCache,
}

// ResourceInfo carries attributes common to all resources in Azure
type ResourceInfo struct {
	SubscriptionID string
	GroupName      string
	Name           string
	Type           string
}

// NewResourceInfo will create and initialize a ResourceInfo
func NewResourceInfo(r *resources.GenericResource) *ResourceInfo {
	info := new(ResourceInfo)
	info.Type = *r.Type
	info.Name = *r.Name
	info.GroupName = extractResourceGroupNameFromResourceID(*r.ID)
	info.SubscriptionID = extractSubscriptionIDFromResourceID(*r.ID)
	return info
}

// GetVisitor finds a function to invoke for a given Azure resource
func (info *ResourceInfo) GetVisitor() func(ResourceInfo) {
	return resourceMap[info.Type]
}

// GetResourcesInSubscription will retrieve all the Azure resources in the specified subscription
func GetResourcesInSubscription(subscriptionID string, settings *configuration.AppSettings) ([]resources.GenericResource, error) {
	allResources := make([]resources.GenericResource, 0)

	client, err := getClient(subscriptionID, settings)
	if err != nil {
		return nil, err
	}

	context := context.Background()
	listResult, err := client.List(context, "", "", nil)
	if err != nil {
		return nil, err
	}

	for listResult.NotDone() {
		for _, r := range listResult.Values() {
			allResources = append(allResources, r)
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

func getClient(subscriptionID string, settings *configuration.AppSettings) (resources.Client, error) {
	client := resources.NewClient(subscriptionID)
	authorizer, err := GetAuthorizer(settings)
	if err != nil {
		return client, err
	}
	client.Authorizer = authorizer
	return client, nil
}

var groupNameFinder = regexp.MustCompile(`\/subscriptions\/.*?\/resourceGroups\/(.*?)\/.*`)
var subscriptionFinder = regexp.MustCompile(`\/subscriptions\/(.*?)\/.*`)

func extractResourceGroupNameFromResourceID(resourceID string) string {
	result := groupNameFinder.FindAllStringSubmatch(resourceID, len(resourceID)+1)
	return result[0][1]
}

func extractSubscriptionIDFromResourceID(resourceID string) string {
	result := subscriptionFinder.FindAllStringSubmatch(resourceID, len(resourceID)+1)
	return result[0][1]
}
