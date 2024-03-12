package SignatureService_v7_4_2

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/dss10core"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/dssx10verificationreport"
	"time"
)

type Document struct {
	XMLName    xml.Name             `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 Document"`
	ID         string               `xml:"ID,attr"`
	RefURI     string               `xml:"RefURI,attr"`
	RefType    string               `xml:"RefType,attr"`
	SchemaRefs string               `xml:"SchemaRefs,attr"`
	ShortText  string               `xml:"ShortText,attr"`
	Base64XML  string               `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Base64XML,omitempty"`
	Base64Data dss10core.Base64Data `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Data,omitempty"`
}

type DocumentWithSignature struct {
	XMLName    xml.Name             `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 DocumentWithSignature"`
	ID         string               `xml:"ID,attr"`
	RefURI     string               `xml:"RefURI,attr"`
	RefType    string               `xml:"RefType,attr"`
	SchemaRefs string               `xml:"SchemaRefs,attr"`
	ShortText  string               `xml:"ShortText,attr"`
	Base64XML  string               `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Base64XML,omitempty"`
	Base64Data dss10core.Base64Data `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Data,omitempty"`
}

type GetJobNumber struct {
	XMLName xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 GetJobNumber"`
	Context ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
}

type GetJobNumberResponse struct {
	XMLName   xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 GetJobNumberResponse"`
	JobNumber string   `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 JobNumber"`
}

type OptionalInputs struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalInputs"`
}

type OptionalOutputs struct {
	XMLName               xml.Name                                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalOutputs"`
	VerifyManifestResults dss10core.VerifyManifestResults             `xml:"urn:oasis:names:tc:dss:1.0:core:schema VerifyManifestResults,omitempty"`
	DocumentWithSignature DocumentWithSignature                       `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 DocumentWithSignature,omitempty"`
	UpdatedSignature      dss10core.UpdatedSignature                  `xml:"urn:oasis:names:tc:dss:1.0:core:schema UpdatedSignature,omitempty"`
	VerificationReport    dssx10verificationreport.VerificationReport `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# VerificationReport,omitempty"`
}

type SignDocument struct {
	XMLName     xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 SignDocument"`
	CardHandle  string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	Crypt       string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 Crypt,omitempty"`
	Context     ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	TvMode      string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 TvMode"`
	JobNumber   string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 JobNumber,omitempty"`
	SignRequest []SignRequest                 `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 SignRequest"`
}

type SignDocumentResponse struct {
	XMLName      xml.Name       `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 SignDocumentResponse"`
	SignResponse []SignResponse `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 SignResponse"`
}

type SignRequest struct {
	XMLName               xml.Name       `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 SignRequest"`
	RequestID             string         `xml:"RequestID,attr"`
	OptionalInputs        OptionalInputs `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalInputs,omitempty"`
	Document              Document       `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 Document"`
	IncludeRevocationInfo bool           `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 IncludeRevocationInfo"`
}

type SignResponse struct {
	XMLName         xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 SignResponse"`
	RequestID       string                      `xml:"RequestID,attr"`
	Status          ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	OptionalOutputs OptionalOutputs             `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalOutputs,omitempty"`
	SignatureObject dss10core.SignatureObject   `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject,omitempty"`
}

type StopSignature struct {
	XMLName   xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 StopSignature"`
	Context   ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	JobNumber string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 JobNumber"`
}

type StopSignatureResponse struct {
	XMLName xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 StopSignatureResponse"`
	Status  ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
}

type VerificationResult struct {
	XMLName         xml.Name  `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 VerificationResult"`
	HighLevelResult string    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 HighLevelResult"`
	TimestampType   string    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 TimestampType"`
	Timestamp       time.Time `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 Timestamp"`
}

type VerifyDocument struct {
	XMLName               xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 VerifyDocument"`
	Context               ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	TvMode                string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 TvMode,omitempty"`
	OptionalInputs        OptionalInputs                `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalInputs,omitempty"`
	Document              Document                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 Document,omitempty"`
	SignatureObject       dss10core.SignatureObject     `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject,omitempty"`
	IncludeRevocationInfo bool                          `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 IncludeRevocationInfo"`
}

type VerifyDocumentResponse struct {
	XMLName            xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 VerifyDocumentResponse"`
	Status             ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	VerificationResult VerificationResult          `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 VerificationResult"`
	OptionalOutputs    OptionalOutputs             `xml:"http://ws.gematik.de/conn/SignatureService/v7.4 OptionalOutputs,omitempty"`
}
