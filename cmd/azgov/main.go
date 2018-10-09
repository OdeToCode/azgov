package main

import (
	"os"

	"github.com/odetocode/azuregovenor/internal/pkg/azure"

	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

func main() {

	if len(os.Args) < 2 {
		panic("Enter a path to the configuration file")
	}

	file, err := os.Open(os.Args[1])
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

	for _, subscription := range settings.Subscriptions {
		resources, err := azure.GetResourcesInSubscription(subscription.ID, settings)
		if err != nil {
			panic(err)
		}
		for _, r := range resources {
			visit := r.GetVisitor()
			visit(r)
		}
	}
}
