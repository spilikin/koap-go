package xmldsig

import (
	"encoding/xml"
)

type CanonicalizationMethod struct {
	XMLName   xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

type DSAKeyValue struct {
	XMLName     xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DSAKeyValue"`
	P           string   `xml:"http://www.w3.org/2000/09/xmldsig# P"`
	Q           string   `xml:"http://www.w3.org/2000/09/xmldsig# Q"`
	G           string   `xml:"http://www.w3.org/2000/09/xmldsig# G,omitempty"`
	Y           string   `xml:"http://www.w3.org/2000/09/xmldsig# Y"`
	J           string   `xml:"http://www.w3.org/2000/09/xmldsig# J,omitempty"`
	Seed        string   `xml:"http://www.w3.org/2000/09/xmldsig# Seed"`
	PgenCounter string   `xml:"http://www.w3.org/2000/09/xmldsig# PgenCounter"`
}

type DigestMethod struct {
	XMLName   xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

type KeyInfo struct {
	XMLName         xml.Name        `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo"`
	Id              string          `xml:"Id,attr"`
	KeyName         string          `xml:"http://www.w3.org/2000/09/xmldsig# KeyName,omitempty"`
	KeyValue        KeyValue        `xml:"http://www.w3.org/2000/09/xmldsig# KeyValue,omitempty"`
	RetrievalMethod RetrievalMethod `xml:"http://www.w3.org/2000/09/xmldsig# RetrievalMethod,omitempty"`
	X509Data        X509Data        `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
	PGPData         PGPData         `xml:"http://www.w3.org/2000/09/xmldsig# PGPData,omitempty"`
	SPKIData        SPKIData        `xml:"http://www.w3.org/2000/09/xmldsig# SPKIData,omitempty"`
	MgmtData        string          `xml:"http://www.w3.org/2000/09/xmldsig# MgmtData,omitempty"`
}

type KeyValue struct {
	XMLName     xml.Name    `xml:"http://www.w3.org/2000/09/xmldsig# KeyValue"`
	DSAKeyValue DSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# DSAKeyValue,omitempty"`
	RSAKeyValue RSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue,omitempty"`
}

type Object struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Object"`
	Id       string   `xml:"Id,attr"`
	MimeType string   `xml:"MimeType,attr"`
	Encoding string   `xml:"Encoding,attr"`
}

type PGPData struct {
	XMLName      xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# PGPData"`
	PGPKeyID     string   `xml:"http://www.w3.org/2000/09/xmldsig# PGPKeyID"`
	PGPKeyPacket string   `xml:"http://www.w3.org/2000/09/xmldsig# PGPKeyPacket,omitempty"`
}

type RSAKeyValue struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue"`
	Modulus  string   `xml:"http://www.w3.org/2000/09/xmldsig# Modulus"`
	Exponent string   `xml:"http://www.w3.org/2000/09/xmldsig# Exponent"`
}

type Reference struct {
	XMLName      xml.Name     `xml:"http://www.w3.org/2000/09/xmldsig# Reference"`
	Id           string       `xml:"Id,attr"`
	URI          string       `xml:"URI,attr"`
	Type         string       `xml:"Type,attr"`
	Transforms   Transforms   `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty"`
	DigestMethod DigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod"`
	DigestValue  string       `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue"`
}

type RetrievalMethod struct {
	XMLName    xml.Name   `xml:"http://www.w3.org/2000/09/xmldsig# RetrievalMethod"`
	URI        string     `xml:"URI,attr"`
	Type       string     `xml:"Type,attr"`
	Transforms Transforms `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty"`
}

type SPKIData struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SPKIData"`
	SPKISexp string   `xml:"http://www.w3.org/2000/09/xmldsig# SPKISexp"`
}

type Signature struct {
	XMLName        xml.Name       `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	Id             string         `xml:"Id,attr"`
	SignedInfo     SignedInfo     `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo"`
	SignatureValue SignatureValue `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue"`
	KeyInfo        KeyInfo        `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo,omitempty"`
	Object         []Object       `xml:"http://www.w3.org/2000/09/xmldsig# Object,omitempty"`
}

type SignatureMethod struct {
	XMLName          xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod"`
	Algorithm        string   `xml:"Algorithm,attr"`
	HMACOutputLength int      `xml:"http://www.w3.org/2000/09/xmldsig# HMACOutputLength,omitempty"`
}

type SignatureValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue"`
	Id      string   `xml:"Id,attr"`
	Value   string   `xml:",chardata"`
}

type SignedInfo struct {
	XMLName                xml.Name               `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo"`
	Id                     string                 `xml:"Id,attr"`
	CanonicalizationMethod CanonicalizationMethod `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod"`
	Reference              []Reference            `xml:"http://www.w3.org/2000/09/xmldsig# Reference"`
}

type Transform struct {
	XMLName   xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transform"`
	Algorithm string   `xml:"Algorithm,attr"`
	XPath     string   `xml:"http://www.w3.org/2000/09/xmldsig# XPath,omitempty"`
}

type Transforms struct {
	XMLName   xml.Name    `xml:"http://www.w3.org/2000/09/xmldsig# Transforms"`
	Transform []Transform `xml:"http://www.w3.org/2000/09/xmldsig# Transform"`
}

type X509Data struct {
	XMLName          xml.Name         `xml:"http://www.w3.org/2000/09/xmldsig# X509Data"`
	X509IssuerSerial X509IssuerSerial `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerSerial,omitempty"`
	X509SKI          string           `xml:"http://www.w3.org/2000/09/xmldsig# X509SKI,omitempty"`
	X509SubjectName  string           `xml:"http://www.w3.org/2000/09/xmldsig# X509SubjectName,omitempty"`
	X509Certificate  string           `xml:"http://www.w3.org/2000/09/xmldsig# X509Certificate,omitempty"`
	X509CRL          string           `xml:"http://www.w3.org/2000/09/xmldsig# X509CRL,omitempty"`
}

type X509IssuerSerial struct {
	XMLName          xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerSerial"`
	X509IssuerName   string   `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerName"`
	X509SerialNumber int      `xml:"http://www.w3.org/2000/09/xmldsig# X509SerialNumber"`
}
