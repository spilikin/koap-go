package ProductInformation_v1_1

import (
	"encoding/xml"
	"time"
)

type Local struct {
	XMLName   xml.Name `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 Local"`
	HWVersion string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 HWVersion"`
	FWVersion string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 FWVersion"`
}

type ProductIdentification struct {
	XMLName         xml.Name       `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductIdentification"`
	ProductVendorID string         `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductVendorID"`
	ProductCode     string         `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductCode"`
	ProductVersion  ProductVersion `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductVersion"`
}

type ProductInformation struct {
	XMLName                xml.Name               `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductInformation"`
	InformationDate        time.Time              `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 InformationDate"`
	ProductTypeInformation ProductTypeInformation `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductTypeInformation"`
	ProductIdentification  ProductIdentification  `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductIdentification"`
	ProductMiscellaneous   ProductMiscellaneous   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductMiscellaneous"`
}

type ProductMiscellaneous struct {
	XMLName           xml.Name `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductMiscellaneous"`
	ProductVendorName string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductVendorName"`
	ProductName       string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductName"`
}

type ProductTypeInformation struct {
	XMLName            xml.Name `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductTypeInformation"`
	ProductType        string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductType"`
	ProductTypeVersion string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductTypeVersion"`
}

type ProductVersion struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductVersion"`
	Local   Local    `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 Local,omitempty"`
	Central string   `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 Central,omitempty"`
}
