package SignatureService_v7_4

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/api/ext/dss10core"
)

type BinaryString struct {
	XMLName    xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 BinaryString"`
	ID         string   `xml:"ID,attr"`
	RefURI     string   `xml:"RefURI,attr"`
	RefType    string   `xml:"RefType,attr"`
	SchemaRefs string   `xml:"SchemaRefs,attr"`
}

type ExternalAuthenticate struct {
	XMLName        xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 ExternalAuthenticate"`
	CardHandle     string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	Context        ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	OptionalInputs OptionalInputs                `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalInputs,omitempty"`
	BinaryString   BinaryString                  `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 BinaryString"`
}

type ExternalAuthenticateResponse struct {
	XMLName         xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 ExternalAuthenticateResponse"`
	Status          ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	SignatureObject dss10core.SignatureObject   `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject,omitempty"`
}

type OptionalInputs struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalInputs"`
}
