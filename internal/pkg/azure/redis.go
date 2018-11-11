package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2018-03-01/redis"
)

func visitRedisCache(info ResourceInfo) {

	client := redis.NewClient(info.SubscriptionID)
	client.Authorizer = GetAuthorizer()

	redis, err := client.Get(context.Background(), info.GroupName, info.Name)
	if err != nil {
		return
	}

	if *redis.EnableNonSslPort == true {
		// raise warning
	}

}
