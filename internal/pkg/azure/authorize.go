package azure

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/odetocode/azuregovenor/internal/pkg/configuration"
)

var authorizer autorest.Authorizer

// InitializeAuthorizer creates an authorizer to manage tokens for Azure API calls
func InitializeAuthorizer(settings *configuration.AppSettings) (autorest.Authorizer, error) {
	config := auth.NewClientCredentialsConfig(settings.ClientID, settings.ClientSecret, settings.TenantID)
	config.AADEndpoint = settings.ActiveDirectoryEndpoint
	config.Resource = settings.Resource

	a, err := config.Authorizer()

	if err == nil {
		return nil, err
	}

	authorizer = a
	return authorizer, nil
}

// GetAuthorizer returns the authorizer created by InitializeAuthorizer
func GetAuthorizer() autorest.Authorizer {
	return authorizer
}
