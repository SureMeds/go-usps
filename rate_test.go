package usps

import (
	"fmt"
	"os"
	"testing"
)

func TestRateDomestic(*testing.T) {
	var client Client
	client.Username = os.Getenv("USPSUsername")
	client.Production = true

	var rate RateRequest
	rate.Service = "PRIORITY"
	rate.ZipOrigination = "44106"
	rate.ZipDestination = "20770"
	rate.Pounds = "1"
	rate.Ounces = "8"
	rate.Container = "NONRECTANGULAR"
	rate.Size = "LARGE"
	rate.Width = "15"
	rate.Length = "30"
	rate.Height = "15"
	rate.Girth = "55"

	output := client.RateDomestic(rate)
	fmt.Println(output)

	client.Production = false
	//if output.Error != "API Authorization failure. User "+client.Username+" is not authorized to use API CarrierPickupAvailability." {
	//	t.Error("Pickup availability is incorrect.")
	//}
}
