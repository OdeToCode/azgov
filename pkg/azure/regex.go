package azure

import (
	"regexp"
)

var groupNameFinder = regexp.MustCompile(`\/subscriptions\/.*?\/resourceGroups\/(.*?)\/.*`)
var subscriptionFinder = regexp.MustCompile(`\/subscriptions\/(.*?)\/.*`)
var resourceURIFinder = regexp.MustCompile(`resourceUri":"(.*?)"`)

func extractResourceGroupNameFromResourceID(resourceID string) string {
	result := groupNameFinder.FindAllStringSubmatch(resourceID, len(resourceID))
	return result[0][1]
}

func extractSubscriptionIDFromResourceID(resourceID string) string {
	result := subscriptionFinder.FindAllStringSubmatch(resourceID, len(resourceID))
	return result[0][1]
}

func extractResourceURI(instanceData string) string {
	result := resourceURIFinder.FindAllStringSubmatch(instanceData, len(instanceData))
	return result[0][1]
}
