package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
)

// SQLReport to report on SQL database instances
type SQLReport struct {
	ResourceInfo
	Failed                    bool
	HolesInFirewall           bool
	EncryptionNotEnabled      bool
	AuditingNotEnabled        bool
	ThreatDetectionNotEnabled bool
}

func visitThreatDetectionPolicies(info *ResourceInfo, report *SQLReport, server string, database string) {
	client := sql.NewDatabaseThreatDetectionPoliciesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	detection, err := client.Get(context.Background(), info.GroupName, server, database)
	if err != nil {
		fmt.Println(err)
		return
	}

	if detection.State != sql.SecurityAlertPolicyStateEnabled {
		report.Failed = true
		report.ThreatDetectionNotEnabled = true
	}
}

func visitAuditingPolicies(info *ResourceInfo, report *SQLReport, server string, database string) {
	client := sql.NewDatabaseBlobAuditingPoliciesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	policy, err := client.Get(context.Background(), info.GroupName, server, database)
	if err != nil {
		fmt.Println(err)
	}

	if policy.State != sql.BlobAuditingPolicyStateEnabled {
		report.AuditingNotEnabled = true
		report.Failed = true
	}
}

func visitDatabases(info *ResourceInfo, report *SQLReport) {
	client := sql.NewDatabasesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	listResult, err := client.ListByServer(context.Background(), info.GroupName, info.Name, "transparentDataEncryption", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, database := range *listResult.Value {
		if *database.Name != "master" {
			tde := *database.TransparentDataEncryption
			for _, setting := range tde {
				if setting.TransparentDataEncryptionProperties.Status != sql.TransparentDataEncryptionStatusEnabled {
					report.EncryptionNotEnabled = true
					report.Failed = true
				}
				visitAuditingPolicies(info, report, info.Name, *database.Name)
				visitThreatDetectionPolicies(info, report, info.Name, *database.Name)
			}
		}
	}
}

func visitFirewallRules(info *ResourceInfo, report *SQLReport) {
	client := sql.NewFirewallRulesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	listResult, err := client.ListByServer(context.Background(), info.GroupName, info.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(*listResult.Value) > 0 {
		report.Failed = true
		report.HolesInFirewall = true
	}
}

func visitSQLServer(info *ResourceInfo) {

	report := new(SQLReport)
	report.ResourceInfo = *info

	// s := sql.NewServersClient(info.SubscriptionID)
	// f, _ := s.Get(context.Background(), "", "")
	// f

	visitFirewallRules(info, report)
	visitDatabases(info, report)

	SendReport(report)
}
