package configuration

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Subscription encapsulates the Azure subscription details
type Subscription struct {
	Name string
	ID   string
}

// Configuration holds the app.config data
type Configuration struct {
	Subscriptions []Subscription
}

// LoadConfig will read a JSON encoded file of Configuration
func Load(reader io.Reader) (Configuration, error) {
	var configuration Configuration

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return configuration, err
	}

	json.Unmarshal(bytes, &configuration)
	return configuration, nil
}
