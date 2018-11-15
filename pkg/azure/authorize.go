package azure

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/odetocode/azgov/pkg/configuration"
)

var authorizer autorest.Authorizer

// InitializeAuthorizer creates an authorizer to manage tokens for Azure API calls
func InitializeAuthorizer(settings *configuration.AppSettings) (autorest.Authorizer, error) {

	config := auth.NewClientCredentialsConfig(settings.ClientID, settings.ClientSecret, settings.TenantID)
	config.AADEndpoint = settings.ActiveDirectoryEndpoint
	config.Resource = settings.Resource

	_authorizer, err := config.Authorizer()

	if err != nil {
		return nil, err
	}

	authorizer = _authorizer
	return authorizer, nil
}

// GetAuthorizer returns the authorizer created by InitializeAuthorizer
func GetAuthorizer() autorest.Authorizer {
	if authorizer == nil {
		panic("Failed to initialize authorizer")
	}
	return authorizer
}
