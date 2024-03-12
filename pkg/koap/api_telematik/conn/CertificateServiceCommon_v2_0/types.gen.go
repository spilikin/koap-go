package CertificateServiceCommon_v2_0

import (
	"encoding/xml"
)

type X509Data struct {
	XMLName          xml.Name         `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509Data"`
	X509IssuerSerial X509IssuerSerial `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509IssuerSerial"`
	X509SubjectName  string           `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509SubjectName"`
	X509Certificate  string           `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509Certificate"`
}

type X509DataInfo struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509DataInfo"`
	CertRef  string   `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 CertRef"`
	X509Data X509Data `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509Data,omitempty"`
}

type X509DataInfoList struct {
	XMLName      xml.Name       `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509DataInfoList"`
	X509DataInfo []X509DataInfo `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509DataInfo"`
}

type X509IssuerSerial struct {
	XMLName          xml.Name `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509IssuerSerial"`
	X509IssuerName   string   `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509IssuerName"`
	X509SerialNumber string   `xml:"http://ws.gematik.de/conn/CertificateServiceCommon/v2.0 X509SerialNumber"`
}
