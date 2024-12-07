package etsi01903v132

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/ext/xmldsig"
)

type Cert struct {
	XMLName      xml.Name     `xml:"http://uri.etsi.org/01903/v1.3.2# Cert"`
	URI          string       `xml:"URI,attr"`
	CertDigest   CertDigest   `xml:"http://uri.etsi.org/01903/v1.3.2# CertDigest"`
	IssuerSerial IssuerSerial `xml:"http://uri.etsi.org/01903/v1.3.2# IssuerSerial"`
}

type CertDigest struct {
	XMLName      xml.Name             `xml:"http://uri.etsi.org/01903/v1.3.2# CertDigest"`
	DigestMethod xmldsig.DigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	DigestValue  string               `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue"`
}

type ClaimedRole struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# ClaimedRole"`
}

type CommitmentTypeId struct {
	XMLName                 xml.Name                `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeId"`
	Identifier              Identifier              `xml:"http://uri.etsi.org/01903/v1.3.2# Identifier"`
	Description             string                  `xml:"http://uri.etsi.org/01903/v1.3.2# Description,omitempty"`
	DocumentationReferences DocumentationReferences `xml:"http://uri.etsi.org/01903/v1.3.2# DocumentationReferences,omitempty"`
}

type CommitmentTypeIndication struct {
	XMLName                  xml.Name                 `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeIndication"`
	CommitmentTypeId         CommitmentTypeId         `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeId"`
	ObjectReference          []string                 `xml:"http://uri.etsi.org/01903/v1.3.2# ObjectReference,omitempty"`
	AllSignedDataObjects     interface{}              `xml:"http://uri.etsi.org/01903/v1.3.2# AllSignedDataObjects,omitempty"`
	CommitmentTypeQualifiers CommitmentTypeQualifiers `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeQualifiers,omitempty"`
}

type CommitmentTypeQualifier struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeQualifier"`
}

type CommitmentTypeQualifiers struct {
	XMLName                 xml.Name                  `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeQualifiers"`
	CommitmentTypeQualifier []CommitmentTypeQualifier `xml:"http://uri.etsi.org/01903/v1.3.2# CommitmentTypeQualifier,omitempty"`
}

type DataObjectFormat struct {
	XMLName          xml.Name         `xml:"http://uri.etsi.org/01903/v1.3.2# DataObjectFormat"`
	ObjectReference  string           `xml:"ObjectReference,attr"`
	Description      string           `xml:"http://uri.etsi.org/01903/v1.3.2# Description,omitempty"`
	ObjectIdentifier ObjectIdentifier `xml:"http://uri.etsi.org/01903/v1.3.2# ObjectIdentifier,omitempty"`
	MimeType         string           `xml:"http://uri.etsi.org/01903/v1.3.2# MimeType,omitempty"`
	Encoding         string           `xml:"http://uri.etsi.org/01903/v1.3.2# Encoding,omitempty"`
}

type DocumentationReferences struct {
	XMLName                xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# DocumentationReferences"`
	DocumentationReference string   `xml:"http://uri.etsi.org/01903/v1.3.2# DocumentationReference"`
}

type Identifier struct {
	XMLName   xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# Identifier"`
	Qualifier string   `xml:"Qualifier,attr"`
	Value     string   `xml:",chardata"`
}

type IssuerSerial struct {
	XMLName          xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# IssuerSerial"`
	X509IssuerName   string   `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerName"`
	X509SerialNumber int      `xml:"http://www.w3.org/2000/09/xmldsig# X509SerialNumber"`
}

type ObjectIdentifier struct {
	XMLName                 xml.Name                `xml:"http://uri.etsi.org/01903/v1.3.2# ObjectIdentifier"`
	Identifier              Identifier              `xml:"http://uri.etsi.org/01903/v1.3.2# Identifier"`
	Description             string                  `xml:"http://uri.etsi.org/01903/v1.3.2# Description,omitempty"`
	DocumentationReferences DocumentationReferences `xml:"http://uri.etsi.org/01903/v1.3.2# DocumentationReferences,omitempty"`
}

type ResponderID struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# ResponderID"`
	ByName  string   `xml:"http://uri.etsi.org/01903/v1.3.2# ByName,omitempty"`
	ByKey   string   `xml:"http://uri.etsi.org/01903/v1.3.2# ByKey,omitempty"`
}

type SigPolicyHash struct {
	XMLName      xml.Name             `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyHash"`
	DigestMethod xmldsig.DigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	DigestValue  string               `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue"`
}

type SigPolicyId struct {
	XMLName                 xml.Name                `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyId"`
	Identifier              Identifier              `xml:"http://uri.etsi.org/01903/v1.3.2# Identifier"`
	Description             string                  `xml:"http://uri.etsi.org/01903/v1.3.2# Description,omitempty"`
	DocumentationReferences DocumentationReferences `xml:"http://uri.etsi.org/01903/v1.3.2# DocumentationReferences,omitempty"`
}

type SigPolicyQualifier struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyQualifier"`
}

type SigPolicyQualifiers struct {
	XMLName            xml.Name             `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyQualifiers"`
	SigPolicyQualifier []SigPolicyQualifier `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyQualifier"`
}

type SignaturePolicyId struct {
	XMLName             xml.Name            `xml:"http://uri.etsi.org/01903/v1.3.2# SignaturePolicyId"`
	SigPolicyId         SigPolicyId         `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyId"`
	Transforms          xmldsig.Transforms  `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty"`
	SigPolicyHash       SigPolicyHash       `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyHash"`
	SigPolicyQualifiers SigPolicyQualifiers `xml:"http://uri.etsi.org/01903/v1.3.2# SigPolicyQualifiers,omitempty"`
}

type SignaturePolicyIdentifier struct {
	XMLName                xml.Name          `xml:"http://uri.etsi.org/01903/v1.3.2# SignaturePolicyIdentifier"`
	SignaturePolicyId      SignaturePolicyId `xml:"http://uri.etsi.org/01903/v1.3.2# SignaturePolicyId,omitempty"`
	SignaturePolicyImplied interface{}       `xml:"http://uri.etsi.org/01903/v1.3.2# SignaturePolicyImplied,omitempty"`
}

type SignatureProductionPlace struct {
	XMLName         xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# SignatureProductionPlace"`
	City            string   `xml:"http://uri.etsi.org/01903/v1.3.2# City,omitempty"`
	StateOrProvince string   `xml:"http://uri.etsi.org/01903/v1.3.2# StateOrProvince,omitempty"`
	PostalCode      string   `xml:"http://uri.etsi.org/01903/v1.3.2# PostalCode,omitempty"`
	CountryName     string   `xml:"http://uri.etsi.org/01903/v1.3.2# CountryName,omitempty"`
}

type SigningCertificate struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/01903/v1.3.2# SigningCertificate"`
	Cert    []Cert   `xml:"http://uri.etsi.org/01903/v1.3.2# Cert"`
}
