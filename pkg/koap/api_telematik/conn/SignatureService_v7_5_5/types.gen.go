package SignatureService_v7_5_5

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/dss10core"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/dssx10verificationreport"
	"time"
)

type ActivateComfortSignature struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 ActivateComfortSignature"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
}

type ActivateComfortSignatureResponse struct {
	XMLName       xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 ActivateComfortSignatureResponse"`
	Status        ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	SignatureMode string                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignatureMode"`
}

type DeactivateComfortSignature struct {
	XMLName    xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 DeactivateComfortSignature"`
	CardHandle []string `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
}

type DeactivateComfortSignatureResponse struct {
	XMLName xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 DeactivateComfortSignatureResponse"`
	Status  ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
}

type Document struct {
	XMLName    xml.Name             `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 Document"`
	ID         string               `xml:"ID,attr"`
	RefURI     string               `xml:"RefURI,attr"`
	RefType    string               `xml:"RefType,attr"`
	SchemaRefs string               `xml:"SchemaRefs,attr"`
	ShortText  string               `xml:"ShortText,attr"`
	Base64XML  string               `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Base64XML,omitempty"`
	Base64Data dss10core.Base64Data `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Data,omitempty"`
}

type DocumentWithSignature struct {
	XMLName    xml.Name             `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 DocumentWithSignature"`
	ID         string               `xml:"ID,attr"`
	RefURI     string               `xml:"RefURI,attr"`
	RefType    string               `xml:"RefType,attr"`
	SchemaRefs string               `xml:"SchemaRefs,attr"`
	ShortText  string               `xml:"ShortText,attr"`
	Base64XML  string               `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Base64XML,omitempty"`
	Base64Data dss10core.Base64Data `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Data,omitempty"`
}

type GetJobNumber struct {
	XMLName xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 GetJobNumber"`
	Context ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
}

type GetJobNumberResponse struct {
	XMLName   xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 GetJobNumberResponse"`
	JobNumber string   `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 JobNumber"`
}

type GetSignatureMode struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 GetSignatureMode"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
}

type GetSignatureModeResponse struct {
	XMLName                xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 GetSignatureModeResponse"`
	Status                 ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	ComfortSignatureStatus string                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 ComfortSignatureStatus"`
	ComfortSignatureMax    int                         `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 ComfortSignatureMax"`
	ComfortSignatureTimer  interface{}                 `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 ComfortSignatureTimer"`
	SessionInfo            SessionInfo                 `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SessionInfo,omitempty"`
}

type OptionalInputs struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 OptionalInputs"`
}

type OptionalOutputs struct {
	XMLName               xml.Name                                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 OptionalOutputs"`
	VerifyManifestResults dss10core.VerifyManifestResults             `xml:"urn:oasis:names:tc:dss:1.0:core:schema VerifyManifestResults,omitempty"`
	DocumentWithSignature DocumentWithSignature                       `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 DocumentWithSignature,omitempty"`
	UpdatedSignature      dss10core.UpdatedSignature                  `xml:"urn:oasis:names:tc:dss:1.0:core:schema UpdatedSignature,omitempty"`
	VerificationReport    dssx10verificationreport.VerificationReport `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# VerificationReport,omitempty"`
}

type SessionInfo struct {
	XMLName        xml.Name    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SessionInfo"`
	SignatureMode  string      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignatureMode"`
	CountRemaining int         `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 CountRemaining"`
	TimeRemaining  interface{} `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 TimeRemaining"`
}

type SignDocument struct {
	XMLName     xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignDocument"`
	CardHandle  string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	Crypt       string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 Crypt,omitempty"`
	Context     ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	TvMode      string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 TvMode"`
	JobNumber   string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 JobNumber,omitempty"`
	SignRequest []SignRequest                 `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignRequest"`
}

type SignDocumentResponse struct {
	XMLName      xml.Name       `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignDocumentResponse"`
	SignResponse []SignResponse `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignResponse"`
}

type SignRequest struct {
	XMLName               xml.Name       `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignRequest"`
	RequestID             string         `xml:"RequestID,attr"`
	OptionalInputs        OptionalInputs `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 OptionalInputs,omitempty"`
	Document              Document       `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 Document"`
	IncludeRevocationInfo bool           `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 IncludeRevocationInfo"`
}

type SignResponse struct {
	XMLName         xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 SignResponse"`
	RequestID       string                      `xml:"RequestID,attr"`
	Status          ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	OptionalOutputs OptionalOutputs             `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 OptionalOutputs,omitempty"`
	SignatureObject dss10core.SignatureObject   `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject,omitempty"`
}

type StopSignature struct {
	XMLName   xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 StopSignature"`
	Context   ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	JobNumber string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 JobNumber"`
}

type StopSignatureResponse struct {
	XMLName xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 StopSignatureResponse"`
	Status  ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
}

type VerificationResult struct {
	XMLName         xml.Name  `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 VerificationResult"`
	HighLevelResult string    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 HighLevelResult"`
	TimestampType   string    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 TimestampType"`
	Timestamp       time.Time `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 Timestamp"`
}

type VerifyDocument struct {
	XMLName               xml.Name                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 VerifyDocument"`
	Context               ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	TvMode                string                        `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 TvMode,omitempty"`
	OptionalInputs        OptionalInputs                `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 OptionalInputs,omitempty"`
	Document              Document                      `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 Document,omitempty"`
	SignatureObject       dss10core.SignatureObject     `xml:"urn:oasis:names:tc:dss:1.0:core:schema SignatureObject,omitempty"`
	IncludeRevocationInfo bool                          `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 IncludeRevocationInfo"`
}

type VerifyDocumentResponse struct {
	XMLName            xml.Name                    `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 VerifyDocumentResponse"`
	Status             ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	VerificationResult VerificationResult          `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 VerificationResult"`
	OptionalOutputs    OptionalOutputs             `xml:"http://ws.gematik.de/conn/SignatureService/v7.5 OptionalOutputs,omitempty"`
}
