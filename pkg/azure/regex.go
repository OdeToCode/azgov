package azure

import (
	"regexp"
)

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
