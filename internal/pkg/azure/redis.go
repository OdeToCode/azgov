package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
)

// RedisReport to report on Redis cache instances
type RedisReport struct {
	ResourceInfo
	Failed     bool
	NonSslPort bool
}

func visitRedisCache(info *ResourceInfo) {

	client := redis.NewClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	redis, err := client.Get(context.Background(), info.GroupName, info.Name)
	report := new(RedisReport)
	report.ResourceInfo = *info

	if err != nil {
		return
	}

	if *redis.EnableNonSslPort == true {
		report.Failed = true
		report.NonSslPort = true
	}

	SendReport(report)

}
