package azure

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-event-hubs-go"
	"github.com/odetocode/azgov/pkg/configuration"
)

var hub *eventhub.Hub
var sendEvents bool

// InitializeHub will establish a connection to the destination hub
func InitializeHub(settings *configuration.AppSettings) (*eventhub.Hub, error) {

	_hub, err := eventhub.NewHubFromConnectionString(settings.EventHubConnection)

	if err != nil {
		return nil, err
	}

	hub = _hub
	sendEvents = settings.SendEvents
	return hub, nil
}

// SendReport will deliver a message to event hub in Azure
func SendReport(report interface{}) error {

	if sendEvents {

		message, err := json.Marshal(report)

		if err != nil {
			log.Println(err)
			return err
		}

		event := eventhub.NewEvent(message)
		err = hub.Send(context.Background(), event)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil

}
