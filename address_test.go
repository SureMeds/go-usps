package usps

import (
	"os"
	"testing"
)

func TestAddressVerification(t *testing.T) {
	var client Client
	client.Username = os.Getenv("USPSUsername")

	var address Address
	address.Address2 = "6406 Ivy Lane"
	address.City = "Greenbelt"
	address.State = "MD"

	output := client.AddressVerification(address)
	if output.Address.Address2 != "6406 IVY LN" {
		t.Error("Address Lookup is incorrect")
	}
}

func TestZipCodeLookup(t *testing.T) {
	var client Client
	client.Username = os.Getenv("USPSUsername")

	var address Address
	address.Address2 = "6406 Ivy Lane"
	address.City = "Greenbelt"
	address.State = "MD"

	output := client.ZipCodeLookup(address)
	if output.Address.Address2 != "6406 IVY LN" {
		t.Error("Zipcode Lookup is incorrect")
	}
}

func TestCityStateLookup(t *testing.T) {
	var client Client
	client.Username = os.Getenv("USPSUsername")

	var address ZipCode
	address.Zip5 = "90210"

	output := client.CityStateLookup(address)
	if output.ZipC.Zip5 != "90210" {
		t.Error("City/State Lookup is incorrect")
	}
}
