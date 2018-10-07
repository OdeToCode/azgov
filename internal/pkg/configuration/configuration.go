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

// AppSettings holds the app.config data
type AppSettings struct {
	Subscriptions           []Subscription
	ClientID                string
	ClientSecret            string
	TenantID                string
	ActiveDirectoryEndpoint string
	Resource                string
}

// Load will read a JSON encoded file of Configuration
func Load(reader io.Reader) (AppSettings, error) {
	var settings AppSettings

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return settings, err
	}

	json.Unmarshal(bytes, &settings)
	return settings, nil
}
