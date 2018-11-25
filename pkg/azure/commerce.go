package azure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/go-autorest/autorest/date"

	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
)

type ResourceUsage struct {
	ID           string
	Cost         float64
	DocumentType string
}

type MeterMap map[string]*commerce.MeterInfo
type UsageMap map[string]*ResourceUsage

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

func GetSubscriptionUsage(subscriptionID string, rates MeterMap) (UsageMap, error) {

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
			recordUsage(usage, usages, rates)
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

func recordUsage(usage commerce.UsageAggregation, usages UsageMap, meters MeterMap) {

	// TODO invetigate InstanceData nil cases, currently see this first with MeterName "Dynamic Public IP"

	if usage.InstanceData != nil {
		id := extractResourceUri(*usage.InstanceData)
		entry, ok := usages[id]
		if !ok {
			entry = new(ResourceUsage)
			entry.DocumentType = "cost"
			entry.ID = id
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
