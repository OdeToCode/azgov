package main

import (
	"fmt"
	"os"

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

	config, err := configuration.Load(file)
	if err != nil {
		panic(err)
	}

	for _, subscription := range config.Subscriptions {
		fmt.Println(subscription.Name)
	}
}
