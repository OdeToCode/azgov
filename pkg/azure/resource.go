package azure

import (
	"context"
	"errors"

	"github.com/Azure/azure-amqp-common-go/uuid"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/odetocode/azgov/pkg/configuration"
)

func noop(info *ResourceInfo) {

}

var resourceMap = map[string]func(*ResourceInfo){
	"Microsoft.Sql/servers":           visitSQLServer,
	"Microsoft.Sql/servers/databases": noop,
	"Microsoft.Cache/Redis":           visitRedisCache,
	"Microsoft.Web/sites":             visitWebSite,
	"Microsoft.Portal/dashboards":     noop,

	"Microsoft.Compute/disks":                             noop,
	"Microsoft.Network/publicIPAddresses":                 noop,
	"Microsoft.Web/serverFarms":                           noop,
	"Microsoft.Compute/virtualMachines":                   noop,
	"Microsoft.Compute/virtualMachines/extensions":        noop,
	"Microsoft.Network/networkInterfaces":                 noop,
	"Microsoft.Network/networkSecurityGroups":             noop,
	"Microsoft.Network/virtualNetworks":                   noop,
	"Microsoft.OperationalInsights/workspaces":            noop,
	"Microsoft.OperationsManagement/solutions":            noop,
	"Microsoft.Insights/alertrules":                       noop,
	"microsoft.insights/alertrules":                       noop,
	"Microsoft.Insights/components":                       noop,
	"microsoft.insights/components":                       noop,
	"Microsoft.Storage/storageAccounts":                   noop,
	"Microsoft.Web/certificates":                          noop,
	"Microsoft.AzureActiveDirectory/b2cDirectories":       noop,
	"Microsoft.Web/sites/slots":                           noop,
	"Microsoft.KeyVault/vaults":                           noop,
	"Microsoft.DataFactory/dataFactories":                 noop,
	"Microsoft.DomainRegistration/domains":                noop,
	"Microsoft.Network/localNetworkGateways":              noop,
	"Microsoft.Network/virtualNetworkGateways":            noop,
	"Microsoft.DataFactory/factories":                     noop,
	"Microsoft.Insights/autoscalesettings":                noop,
	"Microsoft.Network/dnszones":                          noop,
	"Microsoft.CertificateRegistration/certificateOrders": noop,
	"Microsoft.ServiceBus/namespaces":                     noop,
	"Microsoft.Web/connections":                           noop,
	"Sendgrid.Email/accounts":                             noop,
	"Microsoft.Network/connections":                       noop,
	"Microsoft.Logic/workflows":                           noop,
}

func newResourceInfo(r *resources.GenericResource, run string) *ResourceInfo {
	info := new(ResourceInfo)
	info.Type = *r.Type
	info.Name = *r.Name
	info.ID = *r.ID
	info.RunID = run
	info.GroupName = extractResourceGroupNameFromResourceID(*r.ID)
	info.SubscriptionID = extractSubscriptionIDFromResourceID(*r.ID)
	return info
}

func getClient(subscriptionID string) resources.Client {
	client := resources.NewClient(subscriptionID)
	client.Authorizer = GetAuthorizer()
	return client
}

// ResourceInfo carries attributes common to all resources in Azure
type ResourceInfo struct {
	ID             string
	SubscriptionID string
	GroupName      string
	Name           string
	Type           string
	RunID          string
}

// GetVisitor finds a function to invoke for a given Azure resource
func (info *ResourceInfo) GetVisitor() (func(*ResourceInfo), error) {
	visitor := resourceMap[info.Type]
	if visitor == nil {
		return nil, errors.New("no visitor for " + info.Type)
	}
	return visitor, nil
}

// GetResourcesInSubscription will retrieve all the Azure resources in the specified subscription
func GetResourcesInSubscription(subscriptionID string, settings *configuration.AppSettings) ([]ResourceInfo, error) {
	allResources := make([]ResourceInfo, 0)

	run, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	client := getClient(subscriptionID)
	context := context.Background()

	listResult, err := client.List(context, "", "", nil)
	if err != nil {
		return nil, err
	}

	for listResult.NotDone() {
		for _, r := range listResult.Values() {
			info := newResourceInfo(&r, run.String())
			allResources = append(allResources, *info)
		}
		err = listResult.Next()
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return allResources, err
	}

	return allResources, nil
}
