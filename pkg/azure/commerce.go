package azure

import (
	"context"
	"log"
	"time"

	"github.com/Azure/go-autorest/autorest/date"

	"github.com/satori/go.uuid"

	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
)

var meters map[uuid.UUID]*commerce.MeterInfo
var usage map[string]*commerce.UsageAggregation

func GetSubscriptionRateCards(subscriptionID string) error {

	rateClient := commerce.NewRateCardClient(subscriptionID)
	rateClient.Authorizer = GetAuthorizer()

	filter := "OfferDurableId eq 'MS-AZR-0003P' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'"
	cardInfo, err := rateClient.Get(context.Background(), filter)
	if err != nil {
		return err
	}

	meters = make(map[uuid.UUID]*commerce.MeterInfo, len(*cardInfo.Meters))
	for k, v := range *cardInfo.Meters {
		meters[*v.MeterID] = &(*cardInfo.Meters)[k]
	}

	return nil
}

func GetSubscriptionUsage(subscriptionID string) {

	details := true
	now := date.Time{Time: time.Now()}
	monthAgo := date.Time{Time: time.Now().AddDate(0, -1, 0)}
	usage = make(map[string]*commerce.UsageAggregation)

	client := commerce.NewUsageAggregatesClient(subscriptionID)
	client.Authorizer = GetAuthorizer()

	result, err := client.List(context.Background(), monthAgo, now, &details, commerce.Daily, "")
	if err != nil {
		log.Println(err)
		return
	}

	for result.NotDone() {
		values := result.Values()
		for k, v := range values {
			id := extractResourceUri(*v.InstanceData)
			usage[id] = &values[k]
		}

		err = result.NextWithContext(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
	}

}
