package azure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/go-autorest/autorest/date"

	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
)

type MeterMap map[string]*commerce.MeterInfo
type UsageMap map[string]*commerce.UsageAggregation



func SummarizeCharges(meters MeterMap, usage UsageMap) {
	for k, v := range usage {
		rate := meters[*v.MeterID]
		cost := *v.Quantity * *rate.MeterRates["0"]
		fmt.Printf("%s cost %f", k, cost)
	}
}

func GetSubscriptionRateCards(subscriptionID string) (MeterMap, error) {

	rateClient := commerce.NewRateCardClient(subscriptionID)
	rateClient.Authorizer = GetAuthorizer()

	filter := "OfferDurableId eq 'MS-AZR-0003P' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'"
	cardInfo, err := rateClient.Get(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	meters := make(MeterMap, len(*cardInfo.Meters))
	for k, v := range *cardInfo.Meters {
		meters[v.MeterID.String()] = &(*cardInfo.Meters)[k]
	}

	return meters, nil
}

func GetSubscriptionUsage(subscriptionID string) (UsageMap, error) {

	details := true
	usage := make(UsageMap)
	reportStart, reportEnd := GetUsageReportRange()

	client := commerce.NewUsageAggregatesClient(subscriptionID)
	client.Authorizer = GetAuthorizer()

	result, err := client.List(context.Background(), reportStart, reportEnd, &details, commerce.Daily, "")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.NotDone() {
		values := result.Values()
		for k, v := range values {

			// TODO handle nil cases
			// Currently see this first with MeterName "Dynamic Public IP"
			if v.InstanceData != nil {
				id := extractResourceUri(*v.InstanceData)
				usage[id] = &values[k]
			}
		}

		err = result.NextWithContext(context.Background())
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return usage, nil

}

func GetUsageReportRange() (date.Time, date.Time) {
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	reportEnd := date.Time{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
	year, month, day = time.Now().AddDate(0, 0, -8).Date()
	reportStart := date.Time{Time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}

	return reportStart, reportEnd
}
