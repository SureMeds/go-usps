package usps

import (
	"os"
	"testing"
)

func TestTrackPackage(t *testing.T) {
	var client Client
	client.Username = os.Getenv("USPSUsername")

	output := client.TrackPackage("9341989949036022338924")
	if output.TrackInfo.TrackSummary != "The Postal Service could not locate the tracking information for your request. Please verify your tracking number and try again later." {
		t.Error("Tracker is incorrect")
	}
}
