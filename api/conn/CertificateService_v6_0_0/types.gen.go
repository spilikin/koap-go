package CertificateService_v6_0_0

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/CertificateServiceCommon_v2_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/api/tel/error_v2_0"
	"time"
)

type CertRefList struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CertRefList"`
	CertRef []string `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CertRef"`
}

type CertificateExpiration struct {
	XMLName           xml.Name    `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CertificateExpiration"`
	CtID              string      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CtID"`
	CardHandle        string      `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	ICCSN             string      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 ICCSN"`
	SubjectCommonname string      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 subject_commonName"`
	SerialNumber      string      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 serialNumber"`
	Validity          interface{} `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 validity"`
}

type CheckCertificateExpiration struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CheckCertificateExpiration"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle,omitempty"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
}

type CheckCertificateExpirationResponse struct {
	XMLName               xml.Name                    `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CheckCertificateExpirationResponse"`
	Status                ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	CertificateExpiration []CertificateExpiration     `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CertificateExpiration,omitempty"`
}

type ReadCardCertificate struct {
	XMLName     xml.Name                      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 ReadCardCertificate"`
	CardHandle  string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	Context     ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CertRefList CertRefList                   `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CertRefList"`
}

type ReadCardCertificateResponse struct {
	XMLName          xml.Name                                       `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 ReadCardCertificateResponse"`
	Status           ConnectorCommon_v5_0.Status                    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	X509DataInfoList CertificateServiceCommon_v2_0.X509DataInfoList `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509DataInfoList"`
}

type RoleList struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 RoleList"`
	Role    []string `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 Role"`
}

type VerificationStatus struct {
	XMLName            xml.Name         `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 VerificationStatus"`
	VerificationResult string           `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 VerificationResult"`
	Error              error_v2_0.Error `xml:"http://ws.gematik.de/tel/error/v2.0 Error,omitempty"`
}

type VerifyCertificate struct {
	XMLName          xml.Name                      `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 VerifyCertificate"`
	Context          ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	X509Certificate  string                        `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509Certificate"`
	VerificationTime time.Time                     `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 VerificationTime,omitempty"`
}

type VerifyCertificateResponse struct {
	XMLName            xml.Name                    `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 VerifyCertificateResponse"`
	Status             ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	VerificationStatus VerificationStatus          `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 VerificationStatus"`
	RoleList           RoleList                    `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 RoleList"`
}
