package azure

import (
	"context"

	"github.com/satori/go.uuid"

	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
	"github.com/odetocode/azgov/pkg/configuration"
)

var meters map[uuid.UUID]commerce.MeterInfo

func InitializeRateCards(settings *configuration.AppSettings) error {

	rateClient := commerce.NewRateCardClient(settings.Subscriptions[0].ID)
	rateClient.Authorizer = GetAuthorizer()

	filter := "OfferDurableId eq 'MS-AZR-0003P' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'"
	cardInfo, err := rateClient.Get(context.Background(), filter)
	if err != nil {
		return err
	}

	meters = make(map[uuid.UUID]commerce.MeterInfo, 0)
	for _, v := range *cardInfo.Meters {
		meters[*v.MeterID] = v
	}

	return nil
}
