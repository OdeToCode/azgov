package azure

import (
	"testing"
)	

func TestReportRangeIsMidnight(t *testing.T) {
	reportStart, reportEnd := GetUsageReportRange()

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
