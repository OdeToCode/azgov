package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/2017-08-01-preview/security"
)

var locations = make(map[string]string)

// GetSecurityCenterLocation will ask azure for the location of the ASC given a subscription ID
func GetSecurityCenterLocation(subscriptionID string) (string, error) {
	entry := locations[subscriptionID]
	if entry == "" {

		// TODO: configure location?
		client := security.NewLocationsClient(subscriptionID, "centralus")
		client.Authorizer = GetAuthorizer()

		location, err := client.Get(context.Background())
		if err != nil {
			return "", err
		}

		locations[subscriptionID] = *location.Name
	}
	return locations[subscriptionID], nil
}
