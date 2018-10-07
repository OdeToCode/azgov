package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

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
