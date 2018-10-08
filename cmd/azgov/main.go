package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/odetocode/azuregovenor/internal/pkg/azure"

	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

func main() {

	if len(os.Args) < 2 {
		panic("Enter a path to the configuration file")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	settings, err := configuration.Load(file)
	if err != nil {
		panic(err)
	}

	for _, subscription := range settings.Subscriptions {
		resources, err := azure.GetResourcesInSubscription(subscription.ID, &settings)
		if err != nil {
			panic(err)
		}
		for _, r := range resources {
			dump, _ := json.Marshal(r)
			fmt.Printf(string(dump))
			fmt.Println("***") // ;
			// fmt.Printf("%s %s\n", *r.Name, *r.Type)
		}
	}
}
