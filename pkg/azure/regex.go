package azure

import (
	"regexp"
)

var groupNameFinder = regexp.MustCompile(`\/subscriptions\/.*?\/resourceGroups\/(.*?)\/.*`)
var subscriptionFinder = regexp.MustCompile(`\/subscriptions\/(.*?)\/.*`)
var resourceUriFinder = regexp.MustCompile(`resourceUri\\":\\"(.*?)\\`)

func extractResourceGroupNameFromResourceID(resourceID string) string {
	result := groupNameFinder.FindAllStringSubmatch(resourceID, len(resourceID))
	return result[0][1]
}

func extractSubscriptionIDFromResourceID(resourceID string) string {
	result := subscriptionFinder.FindAllStringSubmatch(resourceID, len(resourceID))
	return result[0][1]
}

func extractResourceUri(instanceData string) string {
	result := resourceUriFinder.FindAllStringSubmatch(instanceData, len(instanceData))
	return result[0][1]
}
