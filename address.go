package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
)

// Address ...
type Address struct {
	Address1 string `xml:"Address1"`
	Address2 string `xml:"Address2"`
	City     string `xml:"City"`
	State    string `xml:"State"`
	Zip5     string `xml:"Zip5"`
	Zip4     string `xml:"Zip4"`
}

// ZipCode ...
type ZipCode struct {
	Zip5 string `xml:"Zip5"`
}

// AddressValidateResponse ...
type AddressValidateResponse struct {
	Address struct {
		Address1 string `xml:"Address1"`
		Address2 string `xml:"Address2"`
		City     string `xml:"City"`
		State    string `xml:"State"`
		Zip5     string `xml:"Zip5"`
		Zip4     string `xml:"Zip4"`
	} `xml:"Address"`
}

// ZipCodeLookupResponse ...
type ZipCodeLookupResponse struct {
	Address struct {
		Address1 string `xml:"Address1"`
		Address2 string `xml:"Address2"`
		City     string `xml:"City"`
		State    string `xml:"State"`
		Zip5     string `xml:"Zip5"`
		Zip4     string `xml:"Zip4"`
	} `xml:"Address"`
}

// CityStateLookupResponse ...
type CityStateLookupResponse struct {
	ZipC struct {
		Zip5  string `xml:"Zip5"`
		City  string `xml:"City"`
		State string `xml:"State"`
	} `xml:"ZipCode"`
}

// AddressVerification ...
func (U *Client) AddressVerification(address Address) AddressValidateResponse {
	result := AddressValidateResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(address)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("Verify&XML=")
	urlToEncode := "<AddressValidateRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</AddressValidateRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequest(requestURL.String())
	if body == nil {
		return result
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}

// ZipCodeLookup ...
func (U *Client) ZipCodeLookup(address Address) ZipCodeLookupResponse {
	result := ZipCodeLookupResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(address)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("ZipCodeLookup&XML=")
	urlToEncode := "<ZipCodeLookupRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</ZipCodeLookupRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequest(requestURL.String())
	if body == nil {
		return result
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}

// CityStateLookup ...
func (U *Client) CityStateLookup(zipcode ZipCode) CityStateLookupResponse {
	result := CityStateLookupResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	xmlOut, err := xml.Marshal(zipcode)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CityStateLookup&XML=")
	urlToEncode := "<CityStateLookupRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CityStateLookupRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequest(requestURL.String())
	if body == nil {
		return result
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}
