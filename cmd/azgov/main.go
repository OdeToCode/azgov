package main

import (
	"fmt"
	"log"
	"os"

	"github.com/odetocode/azgov/pkg/azure"
	"github.com/odetocode/azgov/pkg/configuration"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	path := "config.json"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	settings, err := configuration.Load(file)
	if err != nil {
		panic(err)
	}

	_, err = azure.InitializeAuthorizer(settings)
	if err != nil {
		panic(err)
	}

	_, err = azure.InitializeHub(settings)
	if err != nil {
		panic(err)
	}

	for _, subscription := range settings.Subscriptions {
		resources, err := azure.GetResourcesInSubscription(subscription.ID, settings)
		if err != nil {
			panic(err)
		}
		for _, r := range resources {
			log.Printf("Processing %s", r.Name)
			visit, err := r.GetVisitor()
			if visit != nil {
				visit(&r)
			}
			if err != nil {
				log.Println(err)
			}
		}
	}

	if settings.SendNotification {
		notifyComplete(settings)
	}
}

func notifyComplete(settings *configuration.AppSettings) {

	request := sendgrid.GetRequest(settings.SendGridKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	from := mail.NewEmail("Azure Audit", settings.FromAddress)
	to := mail.NewEmail(settings.ToAddress, settings.ToAddress)

	subject := "Audit Complete"
	content := mail.NewContent("text/plain", "Audit run complete. You can view the results at "+settings.Website)

	message := mail.NewV3MailInit(from, subject, to, content)
	request.Body = mail.GetRequestBody(message)

	_, err := sendgrid.API(request)

	if err != nil {
		fmt.Println(err)
	}
}
