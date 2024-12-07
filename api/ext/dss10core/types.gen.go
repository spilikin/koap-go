package dss10core

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/ext/xmldsig"
	"time"
)

type AdditionalTimeInfo struct {
	XMLName xml.Name  `xml:"urn:oasis:names:tc:dss:1.0:core:schema AdditionalTimeInfo"`
	Type    string    `xml:"Type,attr"`
	Ref     string    `xml:"Ref,attr"`
	Value   time.Time `xml:",chardata"`
}

type Base64Data struct {
	XMLName  xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Data"`
	MimeType string   `xml:"MimeType,attr"`
	Value    string   `xml:",chardata"`
}

type Base64Signature struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Signature"`
	Type    string   `xml:"Type,attr"`
	Value   string   `xml:",chardata"`
}

type ManifestResult struct {
	XMLName        xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema ManifestResult"`
	ReferenceXpath string   `xml:"urn:oasis:names:tc:dss:1.0:core:schema ReferenceXpath"`
	Status         string   `xml:"urn:oasis:names:tc:dss:1.0:core:schema Status"`
}

type Other struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema Other"`
}

type Property struct {
	XMLName    xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema Property"`
	Identifier string   `xml:"urn:oasis:names:tc:dss:1.0:core:schema Identifier"`
	Value      Value    `xml:"urn:oasis:names:tc:dss:1.0:core:schema Value,omitempty"`
}

type Result struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss:1.0:core:schema Result"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss:1.0:core:schema ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss:1.0:core:schema ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss:1.0:core:schema ResultMessage,omitempty"`
}

type ResultMessage struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema ResultMessage"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}

type SignatureObject struct {
	XMLName         xml.Name          `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject"`
	SchemaRefs      string            `xml:"SchemaRefs,attr"`
	Signature       xmldsig.Signature `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty"`
	Timestamp       Timestamp         `xml:"urn:oasis:names:tc:dss:1.0:core:schema Timestamp,omitempty"`
	Base64Signature Base64Signature   `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Signature,omitempty"`
	SignaturePtr    SignaturePtr      `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignaturePtr,omitempty"`
	Other           Other             `xml:"urn:oasis:names:tc:dss:1.0:core:schema Other,omitempty"`
}

type SignaturePtr struct {
	XMLName       xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignaturePtr"`
	WhichDocument string   `xml:"WhichDocument,attr"`
	XPath         string   `xml:"XPath,attr"`
}

type TSA struct {
	XMLName       xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema TSA"`
	NameQualifier string   `xml:"NameQualifier,attr"`
	Format        string   `xml:"Format,attr"`
	Value         string   `xml:",chardata"`
}

type Timestamp struct {
	XMLName               xml.Name          `xml:"urn:oasis:names:tc:dss:1.0:core:schema Timestamp"`
	Signature             xmldsig.Signature `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty"`
	RFC3161TimeStampToken string            `xml:"urn:oasis:names:tc:dss:1.0:core:schema RFC3161TimeStampToken,omitempty"`
	Other                 Other             `xml:"urn:oasis:names:tc:dss:1.0:core:schema Other,omitempty"`
}

type TstInfo struct {
	XMLName      xml.Name    `xml:"urn:oasis:names:tc:dss:1.0:core:schema TstInfo"`
	SerialNumber int         `xml:"urn:oasis:names:tc:dss:1.0:core:schema SerialNumber"`
	CreationTime time.Time   `xml:"urn:oasis:names:tc:dss:1.0:core:schema CreationTime"`
	Policy       string      `xml:"urn:oasis:names:tc:dss:1.0:core:schema Policy,omitempty"`
	ErrorBound   interface{} `xml:"urn:oasis:names:tc:dss:1.0:core:schema ErrorBound,omitempty"`
	Ordered      bool        `xml:"urn:oasis:names:tc:dss:1.0:core:schema Ordered,omitempty"`
	TSA          TSA         `xml:"urn:oasis:names:tc:dss:1.0:core:schema TSA,omitempty"`
}

type UpdatedSignature struct {
	XMLName         xml.Name        `xml:"urn:oasis:names:tc:dss:1.0:core:schema UpdatedSignature"`
	Type            string          `xml:"Type,attr"`
	SignatureObject SignatureObject `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject"`
}

type Value struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss:1.0:core:schema Value"`
}

type VerificationTimeInfo struct {
	XMLName            xml.Name             `xml:"urn:oasis:names:tc:dss:1.0:core:schema VerificationTimeInfo"`
	VerificationTime   time.Time            `xml:"urn:oasis:names:tc:dss:1.0:core:schema VerificationTime"`
	AdditionalTimeInfo []AdditionalTimeInfo `xml:"urn:oasis:names:tc:dss:1.0:core:schema AdditionalTimeInfo,omitempty"`
}

type VerifyManifestResults struct {
	XMLName        xml.Name         `xml:"urn:oasis:names:tc:dss:1.0:core:schema VerifyManifestResults"`
	ManifestResult []ManifestResult `xml:"urn:oasis:names:tc:dss:1.0:core:schema ManifestResult"`
}
