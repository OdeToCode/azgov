package azure

import (
	"encoding/json"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
)

func TestUsageRecorder(t *testing.T) {
	meterJson := `
	{
		"Meters": [{
			"EffectiveDate": "2014-11-01T00:00:00Z",
			"IncludedQuantity": 0.0,
			"MeterCategory": "Redis Cache",
			"MeterId": "fea0de99-4dcb-4387-82b9-fc3348238b27",
			"MeterName": "C1 Cache",
			"MeterRates": {
				"0": 2.00
			},
			"MeterRegion": "",
			"MeterStatus": "Active",
			"MeterSubCategory": "Basic",
			"MeterTags": [],
			"Unit": "1 Hour"
		}]
	}`
	usageJson1 :=
		`{
		"id": "/subscriptions/23aff310-b425-4a88-85a5-e60614bfde1b/providers/Microsoft.Commerce/UsageAggregates/Daily_BRSDT_20180901_0000",
		"name": "Daily_BRSDT_20180901_0000",
		"type": "Microsoft.Commerce/UsageAggregate",
		"properties": {
		  "subscriptionId": "23aff310-b425-4a88-85a5-e60614bfde1b",
		  "usageStartTime": "2018-08-31T00:00:00+00:00",
		  "usageEndTime": "2018-09-01T00:00:00+00:00",
		  "meterName": "C1 Cache",
		  "meterCategory": "Redis Cache",
		  "meterSubCategory": "Basic",
		  "unit": "1 Hour",
		  "instanceData": "{\"Microsoft.Resources\":{\"resourceUri\":\"/subscriptions/23aff310-b425-4a88-85a5-e60614bfde1b/resourceGroups/msc-alder-prod/providers/Microsoft.Cache/Redis/msc-alder-redis\",\"location\":\"eastus\"}}",
		  "meterId": "fea0de99-4dcb-4387-82b9-fc3348238b27",
		  "infoFields": {},
		  "quantity": 2.0
		}
	  }`

	usageJson2 :=
		`{
		"id": "/subscriptions/23aff310-b425-4a88-85a5-e60614bfde1b/providers/Microsoft.Commerce/UsageAggregates/Daily_BRSDT_20180901_0000",
		"name": "Daily_BRSDT_20180901_0000",
		"type": "Microsoft.Commerce/UsageAggregate",
		"properties": {
		  "subscriptionId": "23aff310-b425-4a88-85a5-e60614bfde1b",
		  "usageStartTime": "2018-08-31T00:00:00+00:00",
		  "usageEndTime": "2018-09-01T00:00:00+00:00",
		  "meterName": "C1 Cache",
		  "meterCategory": "Redis Cache",
		  "meterSubCategory": "Basic",
		  "unit": "1 Hour",
		  "instanceData": "{\"Microsoft.Resources\":{\"resourceUri\":\"/subscriptions/23aff310-b425-4a88-85a5-e60614bfde1b/resourceGroups/msc-alder-prod/providers/Microsoft.Cache/Redis/msc-alder-redis\",\"location\":\"eastus\"}}",
		  "meterId": "fea0de99-4dcb-4387-82b9-fc3348238b27",
		  "infoFields": {},
		  "quantity": 1.0
		}
	  }`

	rates := &commerce.ResourceRateCardInfo{}
	e := json.Unmarshal([]byte(meterJson), rates)
	if e != nil {
		t.Error(e)
	}

	usage1 := commerce.UsageAggregation{}
	e = usage1.UnmarshalJSON([]byte(usageJson1))
	if e != nil {
		t.Error(e)
	}

	usage2 := commerce.UsageAggregation{}
	e = usage2.UnmarshalJSON([]byte(usageJson2))
	if e != nil {
		t.Error(e)
	}

	meters := makeMeterMap(rates)

	if len(meters) != 1 {
		t.Errorf("should have loaded a single rate card")
	}

	_, ok := (meters)["fea0de99-4dcb-4387-82b9-fc3348238b27"]
	if !ok {
		t.Errorf("could not find the test data rate card")
	}

	usages := make(UsageMap)
	recordUsage(usage1, usages, meters)

	if len(usages) != 1 {
		t.Errorf("should have a single usage entered")
	}

	recordUsage(usage2, usages, meters)

	if len(usages) != 1 {
		t.Errorf("should still have a single usage entered")
	}

	entry := usages["/subscriptions/23aff310-b425-4a88-85a5-e60614bfde1b/resourceGroups/msc-alder-prod/providers/Microsoft.Cache/Redis/msc-alder-redis"]
	if entry.Cost < 5.99 || entry.Cost > 6.01 {
		t.Errorf("should have two usage entries")
	}
}

func TestReportRangeIsMidnight(t *testing.T) {
	reportStart, reportEnd := getUsageReportRange()

	if reportStart.Hour() != 0 {
		t.Errorf("reportStart has an Hour of %d", reportStart.Hour())
	}

	if reportStart.Minute() != 0 {
		t.Errorf("reportStart has an Minute of %d", reportStart.Minute())
	}

	if reportStart.Second() != 0 {
		t.Errorf("reportStart has an Second of %d", reportStart.Second())
	}

	if reportEnd.Hour() != 0 {
		t.Errorf("reportEnd has an Hour of %d", reportStart.Hour())
	}

	if reportEnd.Minute() != 0 {
		t.Errorf("reportEnd has an Minute of %d", reportStart.Minute())
	}

	if reportEnd.Second() != 0 {
		t.Errorf("reportEnd has an Second of %d", reportStart.Second())
	}

}
