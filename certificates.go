package koap

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"

	"github.com/gematik/zero-lab/go/brainpool"
	"github.com/gematik/zero-lab/go/gempki"
	"github.com/spilikin/koap-go/api/gematik/conn/certificateservice60"
)

type CertRef string

const (
	CertRefAUT CertRef = "C.AUT"
	CertRefENC CertRef = "C.ENC"
	CertRefSIG CertRef = "C.SIG"
	CertRefQES CertRef = "C.QES"
)

type AdmissionInfo struct {
	ProfessionItems    []string `json:"professionItems,omitempty"`
	ProfessionOids     []string `json:"professionOids,omitempty"`
	RegistrationNumber string   `json:"registrationNumber,omitempty"`
}

type CardCertificate struct {
	CertRef     CertRef           `json:"certRef"`
	X509        *x509.Certificate `json:"-"`
	IssuerName  string            `json:"issuerName"`
	SubjectName string            `json:"subjectName"`
	Admission   *AdmissionInfo    `json:"admission,omitempty"`
}

func (c *Client) ReadCardCertificates(cardHandle string, certRefs ...CertRef) ([]CardCertificate, error) {
	proxy, err := c.createLatestServiceProxy(ServiceNameCertificateService)
	if err != nil {
		return nil, err
	}

	refs := make([]string, len(certRefs))
	for i, r := range certRefs {
		refs[i] = string(r)
	}

	envelope := &certificateservice60.ReadCardCertificateEnvelope{
		ReadCardCertificate: &certificateservice60.ReadCardCertificate{
			CardHandle: cardHandle,
			Context:    c.connectorContext(),
			CertRefList: struct {
				XMLName xml.Name `xml:"http://ws.gematik.de/conn/CertificateService/v6.0 CertRefList"`
				CertRef []string `xml:"CertRef"`
			}{CertRef: refs},
		},
	}

	var resp certificateservice60.ReadCardCertificateResponseEnvelope
	if err := proxy.Call(&certificateservice60.OperationReadCardCertificate, envelope, &resp); err != nil {
		return nil, fmt.Errorf("ReadCardCertificate: %w", err)
	}

	if resp.Fault != nil {
		return nil, fmt.Errorf("ReadCardCertificate SOAP fault: %s", resp.Fault.String)
	}
	if resp.ReadCardCertificateResponse == nil {
		return nil, fmt.Errorf("ReadCardCertificate: empty response")
	}

	var certs []CardCertificate
	for _, info := range resp.ReadCardCertificateResponse.X509DataInfoList.X509DataInfo {
		if info.X509Data == nil || info.X509Data.X509Certificate == "" {
			continue
		}
		der, err := base64.StdEncoding.DecodeString(info.X509Data.X509Certificate)
		if err != nil {
			return nil, fmt.Errorf("decoding certificate %s: %w", info.CertRef, err)
		}
		cert, err := brainpool.ParseCertificate(der)
		if err != nil {
			return nil, fmt.Errorf("parsing certificate %s: %w", info.CertRef, err)
		}
		cc := CardCertificate{
			CertRef:     CertRef(info.CertRef),
			X509:        cert,
			IssuerName:  info.X509Data.X509IssuerSerial.X509IssuerName,
			SubjectName: info.X509Data.X509SubjectName,
		}
		if adm, err := gempki.ParseAdmissionStatement(cert); err == nil {
			cc.Admission = &AdmissionInfo{
				ProfessionItems:    adm.ProfessionItems,
				ProfessionOids:     adm.ProfessionOids,
				RegistrationNumber: adm.RegistrationNumber,
			}
		}
		certs = append(certs, cc)
	}

	return certs, nil
}

func (c *Client) ReadAllCardCertificates(cardHandle string) ([]CardCertificate, error) {
	return c.ReadCardCertificates(cardHandle, CertRefAUT)
}
