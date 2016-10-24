package usps

const (
	devbase       string = "http://production.shippingapis.com/ShippingAPITest.dll?API="
	prodbase      string = "http://production.shippingapis.com/ShippingAPI.dll?API="
	devhttpsbase  string = "https://secure.shippingapis.com/ShippingAPITest.dll?API="
	prodhttpsbase string = "https://secure.shippingapis.com/ShippingAPI.dll?API="
)

// Client ...
type Client struct {
	Username   string
	Password   string
	Production bool `default:"false"`
}
