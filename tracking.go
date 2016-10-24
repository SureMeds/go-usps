package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// TrackResponse ...
type TrackResponse struct {
	TrackInfo TrackInfo `xml:"TrackInfo"`
}

// TrackInfo ...
type TrackInfo struct {
	AdditionalInfo           string `xml:"AdditionalInfo"`
	ARCHDATA                 string `xml:"ARCHDATA"`
	ArchiveRestoreInfo       string `xml:"ArchiveRestoreInfo"`
	AssociatedLabel          string `xml:"AssociatedLabel"`
	CarrierRelease           string `xml:"CarrierRelease"`
	Class                    string `xml:"Class"`
	ClassOfMailCode          string `xml:"ClassOfMailCode"`
	DeliveryNotificationDate string `xml:"DeliveryNotificationDate"`
	DestinationCity          string `xml:"DestinationCity"`
	DestinationCountryCode   string `xml:"DestinationCountryCode"`
	DestinationState         string `xml:"DestinationState"`
	DestinationZip           string `xml:"DestinationZip"`
	EditedLabelID            string `xml:"EditedLabelID"`
	EmailEnabled             string `xml:"EmailEnabled"`
	ExpectedDeliveryDate     string `xml:"ExpectedDeliveryDate"`
	ExpectedDeliveryTime     string `xml:"ExpectedDeliveryTime"`
	GuaranteedDeliveryDate   string `xml:"GuaranteedDeliveryDate"`
	GuaranteedDeliveryTime   string `xml:"GuaranteedDeliveryTime"`
	GuaranteedDetails        string `xml:"GuaranteedDetails"`
	KahalaIndicator          string `xml:"KahalaIndicator"`
	MailTypeCode             string `xml:"MailTypeCode"`
	MPDATE                   string `xml:" MPDATE"`
	MPSUFFIX                 string `xml:"MPSUFFIX"`
	OriginCity               string `xml:"OriginCity"`
	OriginCountryCode        string `xml:"OriginCountryCode"`
	OriginState              string `xml:"OriginState"`
	OriginZip                string `xml:"OriginZip"`
	PodEnabled               string `xml:"PodEnabled"`
	PredictedDeliveryDate    string `xml:"PredictedDeliveryDate"`
	PredictedDeliveryTime    string `xml:"PredictedDeliveryTime"`
	PDWStart                 string `xml:"PDWStart"`
	PDWEnd                   string `xml:"PDWEnd"`
	RelatedRRID              string `xml:"RelatedRRID"`
	RestoreEnabled           string `xml:"RestoreEnabled"`
	RRAMenabled              string `xml:"RRAMenabled"`
	RreEnabled               string `xml:"RreEnabled"`
	Service                  string `xml:"Service"`
	ServiceTypeCode          string `xml:"ServiceTypeCode"`
	Status                   string `xml:"Status"`
	StatusCategory           string `xml:"StatusCategory"`
	StatusSummary            string `xml:"StatusSummary"`
	TABLECODE                string `xml:"TABLECODE"`
	ValueofArticle           string `xml:"ValueofArticle"`
	TrackSummary             string `xml:"TrackSummary"`
	TrackDetail              string `xml:"TrackDetail"`
}

// GetEstimatedDeliveryTime ...
func (tr TrackResponse) GetEstimatedDeliveryTime() time.Time, error {
	if tr.TrackInfo.ExpectedDeliveryDate == "" {
		return time.Time{}, errors.New("Missing expected delivery date")
	}
	if tr.TrackInfo.ExpectedDeliveryTime == "" {
		return time.parse("January 2, 2006", tr.TrackInfo.ExpectedDeliveryDate), nil
	}
	return time.parse("January 2, 2006 03:04 PM", fmt.Sprintf("%s %s", tr.TrackInfo.ExpectedDeliveryDate,
		tr.TrackInfo.ExpectedDeliveryTime)), nil
}

// TrackPackage ...
func (U *Client) TrackPackage(trackingID string) TrackResponse {
	result := TrackResponse{}
	if U.Username == "" {
		fmt.Println("Username is missing")
		return result
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("TrackV2&xml:")
	urlToEncode := "<TrackRequest USERID=\"" + U.Username + "\">"
	urlToEncode += "<TrackID ID=\"" + trackingID + "\"></TrackID>"
	urlToEncode += "</TrackRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body := U.GetRequest(requestURL.String())
	if body == nil {
		return result
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err := xml.Unmarshal([]byte(bodyHeaderless), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return result
	}

	return result
}
