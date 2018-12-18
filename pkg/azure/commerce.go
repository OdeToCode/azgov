package azure

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/date"

	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
)

// ResourceUsage holds the computed cost for a resource
type ResourceUsage struct {
	ResourceInfo
	Cost float64
}

// MeterMap holds a maaping from MeterID to MeterInfo
type MeterMap map[string]*commerce.MeterInfo

// UsageMap maps from resource ID to usage reports
type UsageMap map[string]*ResourceUsage

// GetSubscriptionRateCards returns all rate cards for a subscription
func GetSubscriptionRateCards(subscriptionID string) (MeterMap, error) {

	rateClient := commerce.NewRateCardClient(subscriptionID)
	rateClient.Authorizer = GetAuthorizer()

	filter := "OfferDurableId eq 'MS-AZR-0003P' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'"
	cardInfo, err := rateClient.Get(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	meters := makeMeterMap(&cardInfo)
	return meters, nil
}

func makeMeterMap(cardInfo *commerce.ResourceRateCardInfo) MeterMap {
	meters := make(MeterMap, len(*cardInfo.Meters))
	for k, v := range *cardInfo.Meters {
		meters[v.MeterID.String()] = &(*cardInfo.Meters)[k]
	}
	return meters
}

// GetSubscriptionUsage fetches all usage records for a subscription
func GetSubscriptionUsage(subscriptionID string, rates MeterMap, runID string) (UsageMap, error) {

	details := true
	usages := make(UsageMap)
	reportStart, reportEnd := getUsageReportRange()

	client := commerce.NewUsageAggregatesClient(subscriptionID)
	client.Authorizer = GetAuthorizer()

	result, err := client.List(context.Background(), reportStart, reportEnd, &details, commerce.Daily, "")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.NotDone() {
		values := result.Values()
		for _, usage := range values {
			recordUsage(usage, usages, rates, runID)
		}

		err = result.NextWithContext(context.Background())
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return usages, nil

}

func getUsageReportRange() (date.Time, date.Time) {
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	reportEnd := date.Time{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
	year, month, day = time.Now().AddDate(0, 0, -8).Date()
	reportStart := date.Time{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}

	return reportStart, reportEnd
}

func recordUsage(usage commerce.UsageAggregation, usages UsageMap, meters MeterMap, runID string) {

	// TODO invetigate InstanceData nil cases, currently see this first with MeterName "Dynamic Public IP"

	if usage.InstanceData != nil {
		id := strings.ToLower(extractResourceURI(*usage.InstanceData))
		entry, ok := usages[id]
		if !ok {
			entry = new(ResourceUsage)
			entry.DocumentType = "cost"
			entry.RunID = runID
			entry.ResourceID = (id)
			entry.Cost = 0
			usages[id] = entry
		}
		rate, ok := meters[*usage.MeterID]
		if !ok {
			fmt.Printf("MeterID %s not found in rate card\n", *usage.MeterID)
		} else {
			entry.Cost += *usage.Quantity * *rate.MeterRates["0"]
		}
	}
}

// ProcessResourceUsage can compute the cost for a resource
func ProcessResourceUsage(usages UsageMap, resource ResourceInfo) {
	var usage *ResourceUsage

	for k, v := range usages {
		if k == resource.ResourceID {
			usage = v
			break
		}
	}

	if usage == nil {
		fmt.Printf("No usages found for resource %s\n", resource.ResourceID)
		return
	}

	usage.GroupName = resource.GroupName
	usage.Name = resource.Name
	usage.Type = resource.Type
	usage.SubscriptionID = resource.SubscriptionID

	SendReport(usage)
}
