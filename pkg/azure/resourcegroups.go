package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
)

func GetAllResourceGroups(subscriptionID string) {
	client := resources.NewGroupsClient(subscriptionID)
	client.Authorizer = GetAuthorizer()
}
