package azure

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Azure/azure-event-hubs-go"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

var hub *eventhub.Hub

// InitializeHub will establish a connection to the destination hub
func InitializeHub(settings *configuration.AppSettings) (*eventhub.Hub, error) {

	_hub, err := eventhub.NewHubFromConnectionString(settings.EventHubConnection)

	if err != nil {
		return nil, err
	}

	hub = _hub
	return hub, nil
}

// SendReport will deliver a message to event hub in Azure
func SendReport(report interface{}) error {

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message, err := json.Marshal(report)

	if err != nil {
		return err
	}

	event := eventhub.NewEvent(message)
	hub.Send(context, event)

	return nil
}
