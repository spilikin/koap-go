package EncryptionService_v6_1_0

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/api/ext/dss10core"
)

type CertificateOnCard struct {
	XMLName      xml.Name `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 CertificateOnCard"`
	CardHandle   string   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	KeyReference string   `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 KeyReference,omitempty"`
}

type DecryptDocument struct {
	XMLName          xml.Name                      `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 DecryptDocument"`
	Context          ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	PrivateKeyOnCard PrivateKeyOnCard              `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 PrivateKeyOnCard"`
	Document         ConnectorCommon_v5_0.Document `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Document"`
	OptionalInputs   interface{}                   `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 OptionalInputs,omitempty"`
}

type DecryptDocumentResponse struct {
	XMLName         xml.Name                      `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 DecryptDocumentResponse"`
	Status          ConnectorCommon_v5_0.Status   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	Document        ConnectorCommon_v5_0.Document `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Document,omitempty"`
	OptionalOutputs interface{}                   `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 OptionalOutputs,omitempty"`
}

type Element struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 Element"`
	Type    string   `xml:"Type,attr"`
	Value   string   `xml:",chardata"`
}

type EncryptDocument struct {
	XMLName        xml.Name                      `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 EncryptDocument"`
	Context        ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	RecipientKeys  RecipientKeys                 `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 RecipientKeys"`
	Document       ConnectorCommon_v5_0.Document `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Document"`
	OptionalInputs OptionalInputs                `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 OptionalInputs,omitempty"`
}

type EncryptDocumentResponse struct {
	XMLName         xml.Name                      `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 EncryptDocumentResponse"`
	Status          ConnectorCommon_v5_0.Status   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	OptionalOutputs interface{}                   `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 OptionalOutputs,omitempty"`
	Document        ConnectorCommon_v5_0.Document `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Document,omitempty"`
}

type OptionalInputs struct {
	XMLName               xml.Name              `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 OptionalInputs"`
	EncryptionType        string                `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 EncryptionType,omitempty"`
	Element               []Element             `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 Element,omitempty"`
	UnprotectedProperties UnprotectedProperties `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 UnprotectedProperties,omitempty"`
}

type PrivateKeyOnCard struct {
	XMLName      xml.Name `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 PrivateKeyOnCard"`
	CardHandle   string   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	KeyReference string   `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 KeyReference,omitempty"`
}

type RecipientKeys struct {
	XMLName           xml.Name          `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 RecipientKeys"`
	CertificateOnCard CertificateOnCard `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 CertificateOnCard,omitempty"`
	Certificate       []string          `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 Certificate,omitempty"`
}

type UnprotectedProperties struct {
	XMLName  xml.Name             `xml:"http://ws.gematik.de/conn/EncryptionService/v6.1 UnprotectedProperties"`
	Property []dss10core.Property `xml:"urn:oasis:names:tc:dss:1.0:core:schema Property"`
}
