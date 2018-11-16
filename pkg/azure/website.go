package azure

import (
	"context"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2018-02-01/web"
)

// WebSiteReport contains data for an AppService
type WebSiteReport struct {
	ResourceInfo
	Failed                    bool
	ClientAffinityNotDisabled bool
	HTTPSOnlyNotSet           bool
	HTTP2NotEnabled           bool
	TLS12NotRequired          bool
	ContainsConnectionStrings bool
	ContainsSensitiveSettings bool
	FTPNotDisabled            bool
}

func visitWebSite(info *ResourceInfo) {
	report := new(WebSiteReport)
	report.ResourceInfo = *info

	client := web.NewAppsClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	app, err := client.Get(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
		return
	}

	if *app.SiteProperties.ClientAffinityEnabled {
		report.Failed = true
		report.ClientAffinityNotDisabled = true
	}

	if !*app.SiteProperties.HTTPSOnly {
		report.Failed = true
		report.HTTPSOnlyNotSet = true
	}

	configuration, err := client.GetConfiguration(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
		return
	}

	if !*configuration.HTTP20Enabled {
		report.Failed = true
		report.HTTP2NotEnabled = true
	}

	if configuration.FtpsState != web.Disabled {
		report.Failed = true
		report.FTPNotDisabled = true
	}

	if configuration.MinTLSVersion != web.OneFullStopTwo {
		report.Failed = true
		report.TLS12NotRequired = true
	}

	connectionStrings, err := client.ListConnectionStrings(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
		return
	}

	if len(connectionStrings.Properties) > 0 {
		report.Failed = true
		report.ContainsConnectionStrings = true
	}

	appSettings, err := client.ListApplicationSettings(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
		return
	}

	bad := []string{"secret", "key", "password", "pwd", "sas"}
	for key, name := range appSettings.Properties {
		key := strings.ToLower(key)
		name := strings.ToLower(*name)
		for _, value := range bad {
			if strings.Contains(name, value) || strings.Contains(key, value) {
				report.Failed = true
				report.ContainsSensitiveSettings = true
				break
			}
		}
	}

	SendReport(report)
}
