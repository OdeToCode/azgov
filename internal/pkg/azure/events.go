package azure

import (
	"github.com/Azure/azure-event-hubs-go"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

var hub *eventhub.Hub

func initHub(settings *configuration.AppSettings) (*eventhub.Hub, error) {

	_hub, err := eventhub.NewHubFromConnectionString(settings.EventHubConnection)
	if err != nil {
		return nil, err
	}

	hub = _hub
	return hub, nil
}
