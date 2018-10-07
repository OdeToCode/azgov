package azure

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

// GetAuthorizer creates an authorizer to manage tokens for Azure API calls
func GetAuthorizer(settings *configuration.AppSettings) (autorest.Authorizer, error) {
	config := auth.NewClientCredentialsConfig(settings.ClientID, settings.ClientSecret, settings.TenantID)
	config.AADEndpoint = settings.ActiveDirectoryEndpoint
	config.Resource = settings.Resource
	return config.Authorizer()
}
