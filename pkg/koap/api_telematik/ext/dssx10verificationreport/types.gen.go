package dssx10verificationreport

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/dss10core"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/etsi01903v132"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/etsi02232v2"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/xmldsig"
	"time"
)

type AllDataObjectsTimeStamp struct {
	XMLName                 xml.Name                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AllDataObjectsTimeStamp"`
	Id                      string                  `xml:"Id,attr"`
	FormatOK                FormatOK                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# FormatOK"`
	TimeStampContent        TimeStampContent        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TimeStampContent,omitempty"`
	MessageHashAlgorithm    MessageHashAlgorithm    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# MessageHashAlgorithm,omitempty"`
	SignatureOK             SignatureOK             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	CertificatePathValidity CertificatePathValidity `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificatePathValidity"`
}

type AttCertValidityPeriod struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttCertValidityPeriod"`
	NotBefore time.Time `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# NotBefore"`
	NotAfter  time.Time `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# NotAfter"`
}

type Attribute struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Attribute"`
	Type    Type     `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Type"`
	Value   []Value  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Value,omitempty"`
}

type AttributeCertificateContent struct {
	XMLName               xml.Name              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateContent"`
	Version               int                   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Version,omitempty"`
	Holder                Holder                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Holder"`
	Issuer                Issuer                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Issuer"`
	SignatureAlgorithm    string                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureAlgorithm"`
	SerialNumber          int                   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SerialNumber"`
	AttCertValidityPeriod AttCertValidityPeriod `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttCertValidityPeriod"`
	Attributes            Attributes            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Attributes"`
	IssuerUniqueID        string                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IssuerUniqueID,omitempty"`
	Extensions            Extensions            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extensions,omitempty"`
}

type AttributeCertificateIdentifier struct {
	XMLName      xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateIdentifier"`
	Holder       Holder   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Holder,omitempty"`
	Issuer       Issuer   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Issuer"`
	SerialNumber int      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SerialNumber"`
}

type AttributeCertificateValidity struct {
	XMLName                        xml.Name                       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateValidity"`
	AttributeCertificateIdentifier AttributeCertificateIdentifier `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateIdentifier"`
	AttributeCertificateValue      string                         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateValue,omitempty"`
	AttributeCertificateContent    AttributeCertificateContent    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateContent,omitempty"`
	SignatureOK                    SignatureOK                    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	CertificatePathValidity        CertificatePathValidity        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificatePathValidity"`
}

type Attributes struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Attributes"`
	Attribute Attribute `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Attribute"`
}

type BaseCertificateID struct {
	XMLName          xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# BaseCertificateID"`
	X509IssuerName   string   `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerName"`
	X509SerialNumber int      `xml:"http://www.w3.org/2000/09/xmldsig# X509SerialNumber"`
}

type CRLContent struct {
	XMLName             xml.Name            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLContent"`
	Version             int                 `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Version,omitempty"`
	Signature           Signature           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Signature"`
	Issuer              string              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Issuer"`
	ThisUpdate          time.Time           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ThisUpdate"`
	NextUpdate          time.Time           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# NextUpdate,omitempty"`
	RevokedCertificates RevokedCertificates `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevokedCertificates,omitempty"`
	CrlExtensions       CrlExtensions       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CrlExtensions,omitempty"`
}

type CRLIdentifier struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLIdentifier"`
	URI       string    `xml:"URI,attr"`
	Issuer    string    `xml:"http://uri.etsi.org/01903/v1.3.2# Issuer"`
	IssueTime time.Time `xml:"http://uri.etsi.org/01903/v1.3.2# IssueTime"`
	Number    int       `xml:"http://uri.etsi.org/01903/v1.3.2# Number,omitempty"`
}

type CRLReference struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLReference"`
	URI       string    `xml:"URI,attr"`
	Issuer    string    `xml:"http://uri.etsi.org/01903/v1.3.2# Issuer"`
	IssueTime time.Time `xml:"http://uri.etsi.org/01903/v1.3.2# IssueTime"`
	Number    int       `xml:"http://uri.etsi.org/01903/v1.3.2# Number,omitempty"`
}

type CRLValidity struct {
	XMLName                 xml.Name                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLValidity"`
	Id                      string                  `xml:"Id,attr"`
	CRLIdentifier           CRLIdentifier           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLIdentifier"`
	CRLValue                string                  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLValue,omitempty"`
	CRLContent              CRLContent              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLContent,omitempty"`
	SignatureOK             SignatureOK             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	CertificatePathValidity CertificatePathValidity `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificatePathValidity"`
}

type CertID struct {
	XMLName        xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertID"`
	HashAlgorithm  string   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# HashAlgorithm"`
	IssuerNameHash string   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IssuerNameHash"`
	IssuerKeyHash  string   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IssuerKeyHash"`
	SerialNumber   int      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SerialNumber"`
}

type CertStatus struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertStatus"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type CertStatusOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertStatusOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type CertificateContent struct {
	XMLName            xml.Name       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateContent"`
	Version            int            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Version,omitempty"`
	SerialNumber       int            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SerialNumber"`
	SignatureAlgorithm string         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureAlgorithm"`
	Issuer             string         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Issuer"`
	ValidityPeriod     ValidityPeriod `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ValidityPeriod"`
	Subject            string         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Subject"`
	Extensions         Extensions     `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extensions,omitempty"`
}

type CertificateIdentifier struct {
	XMLName          xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateIdentifier"`
	X509IssuerName   string   `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerName"`
	X509SerialNumber int      `xml:"http://www.w3.org/2000/09/xmldsig# X509SerialNumber"`
}

type CertificatePathValidity struct {
	XMLName               xml.Name              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificatePathValidity"`
	PathValiditySummary   PathValiditySummary   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# PathValiditySummary"`
	CertificateIdentifier CertificateIdentifier `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateIdentifier"`
	PathValidityDetail    PathValidityDetail    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# PathValidityDetail,omitempty"`
}

type CertificateStatus struct {
	XMLName            xml.Name           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateStatus"`
	CertStatusOK       CertStatusOK       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertStatusOK"`
	RevocationInfo     RevocationInfo     `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationInfo,omitempty"`
	RevocationEvidence RevocationEvidence `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationEvidence,omitempty"`
}

type CertificateValidity struct {
	XMLName               xml.Name              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateValidity"`
	CertificateIdentifier CertificateIdentifier `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateIdentifier"`
	Subject               string                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Subject"`
	ChainingOK            ChainingOK            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ChainingOK"`
	ValidityPeriodOK      ValidityPeriodOK      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ValidityPeriodOK"`
	ExtensionsOK          ExtensionsOK          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtensionsOK"`
	CertificateValue      string                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateValue,omitempty"`
	CertificateContent    CertificateContent    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateContent,omitempty"`
	SignatureOK           SignatureOK           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	CertificateStatus     CertificateStatus     `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateStatus"`
}

type CertifiedRoles struct {
	XMLName                      xml.Name                       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertifiedRoles"`
	AttributeCertificateValidity []AttributeCertificateValidity `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AttributeCertificateValidity"`
}

type ChainingOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ChainingOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type ClaimedRoles struct {
	XMLName     xml.Name                    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ClaimedRoles"`
	ClaimedRole []etsi01903v132.ClaimedRole `xml:"http://uri.etsi.org/01903/v1.3.2# ClaimedRole"`
}

type CrlEntryExtensions struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CrlEntryExtensions"`
	Extension Extension `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extension"`
}

type CrlExtensions struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CrlExtensions"`
	Extension Extension `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extension"`
}

type Details struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Details"`
}

type DigestAlgAndValue struct {
	XMLName      xml.Name             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# DigestAlgAndValue"`
	DigestMethod xmldsig.DigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	DigestValue  string               `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue"`
}

type Extension struct {
	XMLName     xml.Name    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extension"`
	ExtnId      ExtnId      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtnId"`
	Critical    bool        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Critical"`
	ExtnValue   ExtnValue   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtnValue,omitempty"`
	ExtensionOK ExtensionOK `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtensionOK"`
}

type ExtensionOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtensionOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type Extensions struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extensions"`
	Extension Extension `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extension"`
}

type ExtensionsOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtensionsOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type ExtnId struct {
	XMLName                 xml.Name                              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtnId"`
	Identifier              etsi01903v132.Identifier              `xml:"http://uri.etsi.org/01903/v1.3.2# Identifier"`
	Description             string                                `xml:"http://uri.etsi.org/01903/v1.3.2# Description,omitempty"`
	DocumentationReferences etsi01903v132.DocumentationReferences `xml:"http://uri.etsi.org/01903/v1.3.2# DocumentationReferences,omitempty"`
}

type ExtnValue struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ExtnValue"`
}

type FormatOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# FormatOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type Holder struct {
	XMLName           xml.Name          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Holder"`
	BaseCertificateID BaseCertificateID `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# BaseCertificateID,omitempty"`
	Name              string            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Name,omitempty"`
	Other             Other             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}

type IndividualDataObjectsTimeStamp struct {
	XMLName                 xml.Name                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IndividualDataObjectsTimeStamp"`
	Id                      string                  `xml:"Id,attr"`
	FormatOK                FormatOK                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# FormatOK"`
	TimeStampContent        TimeStampContent        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TimeStampContent,omitempty"`
	MessageHashAlgorithm    MessageHashAlgorithm    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# MessageHashAlgorithm,omitempty"`
	SignatureOK             SignatureOK             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	CertificatePathValidity CertificatePathValidity `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificatePathValidity"`
}

type IndividualReport struct {
	XMLName                xml.Name               `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IndividualReport"`
	SignedObjectIdentifier SignedObjectIdentifier `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedObjectIdentifier"`
	Result                 dss10core.Result       `xml:"urn:oasis:names:tc:dss:1.0:core:schema Result"`
	Details                Details                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Details,omitempty"`
}

type Issuer struct {
	XMLName           xml.Name          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Issuer"`
	BaseCertificateID BaseCertificateID `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# BaseCertificateID,omitempty"`
	Name              string            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Name,omitempty"`
	Other             Other             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}

type MessageHashAlgorithm struct {
	XMLName     xml.Name    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# MessageHashAlgorithm"`
	Algorithm   string      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Algorithm"`
	Parameters  Parameters  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Parameters,omitempty"`
	Suitability Suitability `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Suitability,omitempty"`
}

type OCSPContent struct {
	XMLName            xml.Name           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPContent"`
	Version            int                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Version"`
	ResponderID        string             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResponderID"`
	ProducedAt         time.Time          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# producedAt"`
	Responses          Responses          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Responses"`
	ResponseExtensions ResponseExtensions `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResponseExtensions,omitempty"`
}

type OCSPIdentifier struct {
	XMLName     xml.Name                  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPIdentifier"`
	URI         string                    `xml:"URI,attr"`
	ResponderID etsi01903v132.ResponderID `xml:"http://uri.etsi.org/01903/v1.3.2# ResponderID"`
	ProducedAt  time.Time                 `xml:"http://uri.etsi.org/01903/v1.3.2# ProducedAt"`
}

type OCSPReference struct {
	XMLName     xml.Name                  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPReference"`
	URI         string                    `xml:"URI,attr"`
	ResponderID etsi01903v132.ResponderID `xml:"http://uri.etsi.org/01903/v1.3.2# ResponderID"`
	ProducedAt  time.Time                 `xml:"http://uri.etsi.org/01903/v1.3.2# ProducedAt"`
}

type OCSPValidity struct {
	XMLName                 xml.Name                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPValidity"`
	Id                      string                  `xml:"Id,attr"`
	OCSPIdentifier          OCSPIdentifier          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPIdentifier"`
	OCSPValue               string                  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPValue,omitempty"`
	OCSPContent             OCSPContent             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPContent,omitempty"`
	SignatureOK             SignatureOK             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	CertificatePathValidity CertificatePathValidity `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificatePathValidity"`
}

type Other struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other"`
}

type Parameters struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Parameters"`
}

type PathValidityDetail struct {
	XMLName             xml.Name              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# PathValidityDetail"`
	CertificateValidity []CertificateValidity `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertificateValidity,omitempty"`
	TSLValidity         TSLValidity           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TSLValidity,omitempty"`
	TrustAnchor         TrustAnchor           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TrustAnchor"`
}

type PathValiditySummary struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# PathValiditySummary"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type ResponseExtensions struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResponseExtensions"`
	Extension Extension `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extension"`
}

type Responses struct {
	XMLName        xml.Name       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Responses"`
	SingleResponse SingleResponse `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SingleResponse"`
}

type ResultMessage struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}

type RevocationEvidence struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationEvidence"`
	CRLValidity   CRLValidity   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLValidity,omitempty"`
	CRLReference  CRLReference  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CRLReference,omitempty"`
	OCSPValidity  OCSPValidity  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPValidity,omitempty"`
	OCSPReference OCSPReference `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# OCSPReference,omitempty"`
	Other         Other         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}

type RevocationInfo struct {
	XMLName          xml.Name         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationInfo"`
	RevocationDate   time.Time        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationDate"`
	RevocationReason RevocationReason `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationReason"`
}

type RevocationReason struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationReason"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type RevokedCertificates struct {
	XMLName            xml.Name           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevokedCertificates"`
	UserCertificate    int                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# UserCertificate"`
	RevocationDate     time.Time          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# RevocationDate"`
	CrlEntryExtensions CrlEntryExtensions `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CrlEntryExtensions,omitempty"`
}

type SAMLv1Identifier struct {
	XMLName       xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SAMLv1Identifier"`
	NameQualifier string   `xml:"NameQualifier,attr"`
	Format        string   `xml:"Format,attr"`
	Value         string   `xml:",chardata"`
}

type SAMLv2Identifier struct {
	XMLName         xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SAMLv2Identifier"`
	NameQualifier   string   `xml:"NameQualifier,attr"`
	SPNameQualifier string   `xml:"SPNameQualifier,attr"`
	Format          string   `xml:"Format,attr"`
	SPProvidedID    string   `xml:"SPProvidedID,attr"`
	Value           string   `xml:",chardata"`
}

type SigMathOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SigMathOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type Signature struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Signature"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type SignatureAlgorithm struct {
	XMLName     xml.Name    `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureAlgorithm"`
	Algorithm   string      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Algorithm"`
	Parameters  Parameters  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Parameters,omitempty"`
	Suitability Suitability `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Suitability,omitempty"`
}

type SignatureOK struct {
	XMLName            xml.Name           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
	SigMathOK          SigMathOK          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SigMathOK"`
	SignatureAlgorithm SignatureAlgorithm `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureAlgorithm,omitempty"`
}

type SignedDataObjectProperties struct {
	XMLName                        xml.Name                                 `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedDataObjectProperties"`
	Id                             string                                   `xml:"Id,attr"`
	DataObjectFormat               []etsi01903v132.DataObjectFormat         `xml:"http://uri.etsi.org/01903/v1.3.2# DataObjectFormat,omitempty"`
	CommitmentTypeIndication       []etsi01903v132.CommitmentTypeIndication `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeIndication,omitempty"`
	Reason                         string                                   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Reason,omitempty"`
	AllDataObjectsTimeStamp        []AllDataObjectsTimeStamp                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# AllDataObjectsTimeStamp,omitempty"`
	IndividualDataObjectsTimeStamp []IndividualDataObjectsTimeStamp         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IndividualDataObjectsTimeStamp,omitempty"`
}

type SignedObjectIdentifier struct {
	XMLName                xml.Name                       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedObjectIdentifier"`
	WhichDocument          string                         `xml:"WhichDocument,attr"`
	XPath                  string                         `xml:"XPath,attr"`
	Offset                 int                            `xml:"Offset,attr"`
	FieldName              string                         `xml:"FieldName,attr"`
	DigestAlgAndValue      DigestAlgAndValue              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# DigestAlgAndValue,omitempty"`
	CanonicalizationMethod xmldsig.CanonicalizationMethod `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod,omitempty"`
	SignedProperties       SignedProperties               `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedProperties,omitempty"`
	SignatureValue         xmldsig.SignatureValue         `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue,omitempty"`
	Other                  Other                          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}

type SignedProperties struct {
	XMLName                    xml.Name                   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedProperties"`
	Id                         string                     `xml:"Id,attr"`
	SignedSignatureProperties  SignedSignatureProperties  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedSignatureProperties,omitempty"`
	SignedDataObjectProperties SignedDataObjectProperties `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedDataObjectProperties,omitempty"`
	Other                      Other                      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}

type SignedSignatureProperties struct {
	XMLName                   xml.Name                                `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignedSignatureProperties"`
	SigningTime               time.Time                               `xml:"http://uri.etsi.org/01903/v1.3.2# SigningTime,omitempty"`
	SigningCertificate        etsi01903v132.SigningCertificate        `xml:"http://uri.etsi.org/01903/v1.3.2# SigningCertificate,omitempty"`
	SignaturePolicyIdentifier etsi01903v132.SignaturePolicyIdentifier `xml:"http://uri.etsi.org/01903/v1.3.2# SignaturePolicyIdentifier,omitempty"`
	SignatureProductionPlace  etsi01903v132.SignatureProductionPlace  `xml:"http://uri.etsi.org/01903/v1.3.2# SignatureProductionPlace,omitempty"`
	Location                  string                                  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Location,omitempty"`
	SignerRole                SignerRole                              `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignerRole,omitempty"`
}

type SignerRole struct {
	XMLName        xml.Name       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignerRole"`
	ClaimedRoles   ClaimedRoles   `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ClaimedRoles,omitempty"`
	CertifiedRoles CertifiedRoles `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertifiedRoles,omitempty"`
}

type SingleExtensions struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SingleExtensions"`
	Extension Extension `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Extension"`
}

type SingleResponse struct {
	XMLName          xml.Name         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SingleResponse"`
	CertID           CertID           `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertID"`
	CertStatus       CertStatus       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# CertStatus"`
	ThisUpdate       time.Time        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ThisUpdate"`
	NextUpdate       time.Time        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# NextUpdate,omitempty"`
	SingleExtensions SingleExtensions `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SingleExtensions,omitempty"`
}

type Suitability struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Suitability"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type TSLValidity struct {
	XMLName                  xml.Name                             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TSLValidity"`
	TSLTag                   string                               `xml:"TSLTag,attr"`
	Id                       string                               `xml:"Id,attr"`
	SchemeInformation        etsi02232v2.SchemeInformation        `xml:"http://uri.etsi.org/02231/v2# SchemeInformation"`
	TrustServiceProviderList etsi02232v2.TrustServiceProviderList `xml:"http://uri.etsi.org/02231/v2# TrustServiceProviderList,omitempty"`
	SignatureOK              SignatureOK                          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SignatureOK"`
}

type TimeStampContent struct {
	XMLName xml.Name          `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TimeStampContent"`
	TstInfo dss10core.TstInfo `xml:"urn:oasis:names:tc:dss:1.0:core:schema TstInfo,omitempty"`
	Other   Other             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}

type TrustAnchor struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# TrustAnchor"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type Type struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Type"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type ValidityPeriod struct {
	XMLName   xml.Name  `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ValidityPeriod"`
	NotBefore time.Time `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# NotBefore"`
	NotAfter  time.Time `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# NotAfter"`
}

type ValidityPeriodOK struct {
	XMLName       xml.Name      `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ValidityPeriodOK"`
	ResultMajor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMajor"`
	ResultMinor   string        `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMinor,omitempty"`
	ResultMessage ResultMessage `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# ResultMessage,omitempty"`
}

type Value struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Value"`
}

type VerificationReport struct {
	XMLName              xml.Name                       `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# VerificationReport"`
	VerificationTimeInfo dss10core.VerificationTimeInfo `xml:"urn:oasis:names:tc:dss:1.0:core:schema VerificationTimeInfo,omitempty"`
	VerifierIdentity     VerifierIdentity               `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# VerifierIdentity,omitempty"`
	IndividualReport     []IndividualReport             `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# IndividualReport,omitempty"`
}

type VerifierIdentity struct {
	XMLName          xml.Name         `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# VerifierIdentity"`
	X509Data         xmldsig.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
	SAMLv1Identifier SAMLv1Identifier `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SAMLv1Identifier,omitempty"`
	SAMLv2Identifier SAMLv2Identifier `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# SAMLv2Identifier,omitempty"`
	Other            Other            `xml:"urn:oasis:names:tc:dss-x:1.0:profiles:verificationreport:schema# Other,omitempty"`
}
