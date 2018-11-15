package configuration

import (
	"strings"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {

	reader := strings.NewReader(`{
		"subscriptions": [{}, {}]	
	}`)

	config, err := Load(reader)
	if err != nil {
		t.Error(err)
	}

	if len(config.Subscriptions) != 2 {
		t.Error("Should have 2 subscriptions")
	}
}
