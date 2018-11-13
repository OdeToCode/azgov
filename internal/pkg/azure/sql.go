package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql"
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

func visitThreatDetectionPolicies(info *ResourceInfo, report *SQLReport) {

	client := sql.NewServerSecurityAlertPoliciesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	policy, err := client.Get(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
		return
	}

	if policy.SecurityAlertPolicyProperties.State != sql.SecurityAlertPolicyStateEnabled {
		report.Failed = true
		report.ThreatDetectionNotEnabled = true
	}

}

func visitAuditingPolicies(info *ResourceInfo, report *SQLReport) {

	client := sql.NewServerBlobAuditingPoliciesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	policy, err := client.Get(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
		return
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
		log.Println(err)
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
			}
		}
	}

}

func visitFirewallRules(info *ResourceInfo, report *SQLReport) {

	client := sql.NewFirewallRulesClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	listResult, err := client.ListByServer(context.Background(), info.GroupName, info.Name)
	if err != nil {
		log.Println(err)
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

	visitFirewallRules(info, report)
	visitDatabases(info, report)
	visitAuditingPolicies(info, report)
	visitThreatDetectionPolicies(info, report)

	SendReport(report)
}
