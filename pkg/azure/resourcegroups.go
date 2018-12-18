package azure

import (
	"context"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
)

// CheckLocation to see if resource location is the same as the group location
func CheckLocation(groups map[string]resources.Group, info ResourceInfo) {
	name := strings.ToLower(info.GroupName)
	group, ok := groups[name]

	if !ok {
		log.Printf("Could not find group name %s\n", name)
	}
	if strings.ToLower(info.Location) != strings.ToLower(*group.Location) {
		info.LocationMisAligned = true
	}

}

// GetAllResourceGroups fetch all resource groups in a subscription
func GetAllResourceGroups(subscriptionID string) (map[string]resources.Group, error) {
	client := resources.NewGroupsClient(subscriptionID)
	client.Authorizer = GetAuthorizer()
	result, err := client.List(context.Background(), "", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	groups := make(map[string]resources.Group)

	for result.NotDone() {
		for _, g := range result.Values() {
			name := strings.ToLower(*g.Name)
			groups[name] = g
		}

		err = result.Next()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return groups, nil
}
