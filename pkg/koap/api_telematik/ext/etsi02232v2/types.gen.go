package etsi02232v2

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/xmldsig"
	"time"
)

type AdditionalInformation struct {
	XMLName            xml.Name           `xml:"http://uri.etsi.org/02231/v2# AdditionalInformation"`
	TextualInformation TextualInformation `xml:"http://uri.etsi.org/02231/v2# TextualInformation,omitempty"`
	OtherInformation   OtherInformation   `xml:"http://uri.etsi.org/02231/v2# OtherInformation,omitempty"`
}

type DigitalId struct {
	XMLName         xml.Name         `xml:"http://uri.etsi.org/02231/v2# DigitalId"`
	X509Certificate string           `xml:"http://uri.etsi.org/02231/v2# X509Certificate,omitempty"`
	X509SubjectName string           `xml:"http://uri.etsi.org/02231/v2# X509SubjectName,omitempty"`
	KeyValue        xmldsig.KeyValue `xml:"http://www.w3.org/2000/09/xmldsig# KeyValue,omitempty"`
	X509SKI         string           `xml:"http://uri.etsi.org/02231/v2# X509SKI,omitempty"`
	Other           Other            `xml:"http://uri.etsi.org/02231/v2# Other,omitempty"`
}

type DistributionPoints struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# DistributionPoints"`
	URI     []string `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type ElectronicAddress struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# ElectronicAddress"`
	URI     []string `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type Extension struct {
	XMLName  xml.Name `xml:"http://uri.etsi.org/02231/v2# Extension"`
	Critical bool     `xml:"Critical,attr"`
}

type Name struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# Name"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}

type NextUpdate struct {
	XMLName  xml.Name  `xml:"http://uri.etsi.org/02231/v2# NextUpdate"`
	DateTime time.Time `xml:"http://uri.etsi.org/02231/v2# dateTime,omitempty"`
}

type Other struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# Other"`
}

type OtherInformation struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# OtherInformation"`
}

type OtherTSLPointer struct {
	XMLName                  xml.Name                 `xml:"http://uri.etsi.org/02231/v2# OtherTSLPointer"`
	ServiceDigitalIdentities ServiceDigitalIdentities `xml:"http://uri.etsi.org/02231/v2# ServiceDigitalIdentities,omitempty"`
	TSLLocation              string                   `xml:"http://uri.etsi.org/02231/v2# TSLLocation"`
	AdditionalInformation    AdditionalInformation    `xml:"http://uri.etsi.org/02231/v2# AdditionalInformation,omitempty"`
}

type PointersToOtherTSL struct {
	XMLName         xml.Name          `xml:"http://uri.etsi.org/02231/v2# PointersToOtherTSL"`
	OtherTSLPointer []OtherTSLPointer `xml:"http://uri.etsi.org/02231/v2# OtherTSLPointer"`
}

type PolicyOrLegalNotice struct {
	XMLName        xml.Name         `xml:"http://uri.etsi.org/02231/v2# PolicyOrLegalNotice"`
	TSLPolicy      []TSLPolicy      `xml:"http://uri.etsi.org/02231/v2# TSLPolicy,omitempty"`
	TSLLegalNotice []TSLLegalNotice `xml:"http://uri.etsi.org/02231/v2# TSLLegalNotice,omitempty"`
}

type PostalAddress struct {
	XMLName         xml.Name `xml:"http://uri.etsi.org/02231/v2# PostalAddress"`
	Lang            string   `xml:"lang,attr"`
	StreetAddress   string   `xml:"http://uri.etsi.org/02231/v2# StreetAddress"`
	Locality        string   `xml:"http://uri.etsi.org/02231/v2# Locality"`
	StateOrProvince string   `xml:"http://uri.etsi.org/02231/v2# StateOrProvince,omitempty"`
	PostalCode      string   `xml:"http://uri.etsi.org/02231/v2# PostalCode,omitempty"`
	CountryName     string   `xml:"http://uri.etsi.org/02231/v2# CountryName"`
}

type PostalAddresses struct {
	XMLName       xml.Name        `xml:"http://uri.etsi.org/02231/v2# PostalAddresses"`
	PostalAddress []PostalAddress `xml:"http://uri.etsi.org/02231/v2# PostalAddress"`
}

type SchemeExtensions struct {
	XMLName   xml.Name    `xml:"http://uri.etsi.org/02231/v2# SchemeExtensions"`
	Extension []Extension `xml:"http://uri.etsi.org/02231/v2# Extension"`
}

type SchemeInformation struct {
	XMLName                     xml.Name                 `xml:"http://uri.etsi.org/02231/v2# SchemeInformation"`
	TSLVersionIdentifier        int                      `xml:"http://uri.etsi.org/02231/v2# TSLVersionIdentifier"`
	TSLSequenceNumber           int                      `xml:"http://uri.etsi.org/02231/v2# TSLSequenceNumber"`
	TSLType                     string                   `xml:"http://uri.etsi.org/02231/v2# TSLType"`
	SchemeOperatorName          SchemeOperatorName       `xml:"http://uri.etsi.org/02231/v2# SchemeOperatorName"`
	SchemeOperatorAddress       SchemeOperatorAddress    `xml:"http://uri.etsi.org/02231/v2# SchemeOperatorAddress"`
	SchemeName                  SchemeName               `xml:"http://uri.etsi.org/02231/v2# SchemeName"`
	SchemeInformationURI        SchemeInformationURI     `xml:"http://uri.etsi.org/02231/v2# SchemeInformationURI"`
	StatusDeterminationApproach string                   `xml:"http://uri.etsi.org/02231/v2# StatusDeterminationApproach"`
	SchemeTypeCommunityRules    SchemeTypeCommunityRules `xml:"http://uri.etsi.org/02231/v2# SchemeTypeCommunityRules,omitempty"`
	SchemeTerritory             string                   `xml:"http://uri.etsi.org/02231/v2# SchemeTerritory,omitempty"`
	PolicyOrLegalNotice         PolicyOrLegalNotice      `xml:"http://uri.etsi.org/02231/v2# PolicyOrLegalNotice,omitempty"`
	HistoricalInformationPeriod int                      `xml:"http://uri.etsi.org/02231/v2# HistoricalInformationPeriod"`
	PointersToOtherTSL          PointersToOtherTSL       `xml:"http://uri.etsi.org/02231/v2# PointersToOtherTSL,omitempty"`
	ListIssueDateTime           time.Time                `xml:"http://uri.etsi.org/02231/v2# ListIssueDateTime"`
	NextUpdate                  NextUpdate               `xml:"http://uri.etsi.org/02231/v2# NextUpdate"`
	DistributionPoints          DistributionPoints       `xml:"http://uri.etsi.org/02231/v2# DistributionPoints,omitempty"`
	SchemeExtensions            SchemeExtensions         `xml:"http://uri.etsi.org/02231/v2# SchemeExtensions,omitempty"`
}

type SchemeInformationURI struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# SchemeInformationURI"`
	URI     []URI    `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type SchemeName struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# SchemeName"`
	Name    []Name   `xml:"http://uri.etsi.org/02231/v2# Name"`
}

type SchemeOperatorAddress struct {
	XMLName           xml.Name          `xml:"http://uri.etsi.org/02231/v2# SchemeOperatorAddress"`
	PostalAddresses   PostalAddresses   `xml:"http://uri.etsi.org/02231/v2# PostalAddresses"`
	ElectronicAddress ElectronicAddress `xml:"http://uri.etsi.org/02231/v2# ElectronicAddress"`
}

type SchemeOperatorName struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# SchemeOperatorName"`
	Name    []Name   `xml:"http://uri.etsi.org/02231/v2# Name"`
}

type SchemeServiceDefinitionURI struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# SchemeServiceDefinitionURI"`
	URI     []URI    `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type SchemeTypeCommunityRules struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# SchemeTypeCommunityRules"`
	URI     []string `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type ServiceDigitalIdentities struct {
	XMLName                xml.Name                 `xml:"http://uri.etsi.org/02231/v2# ServiceDigitalIdentities"`
	ServiceDigitalIdentity []ServiceDigitalIdentity `xml:"http://uri.etsi.org/02231/v2# ServiceDigitalIdentity"`
}

type ServiceDigitalIdentity struct {
	XMLName   xml.Name    `xml:"http://uri.etsi.org/02231/v2# ServiceDigitalIdentity"`
	DigitalId []DigitalId `xml:"http://uri.etsi.org/02231/v2# DigitalId,omitempty"`
}

type ServiceHistory struct {
	XMLName                xml.Name                 `xml:"http://uri.etsi.org/02231/v2# ServiceHistory"`
	ServiceHistoryInstance []ServiceHistoryInstance `xml:"http://uri.etsi.org/02231/v2# ServiceHistoryInstance,omitempty"`
}

type ServiceHistoryInstance struct {
	XMLName                      xml.Name                     `xml:"http://uri.etsi.org/02231/v2# ServiceHistoryInstance"`
	ServiceTypeIdentifier        string                       `xml:"http://uri.etsi.org/02231/v2# ServiceTypeIdentifier"`
	ServiceName                  ServiceName                  `xml:"http://uri.etsi.org/02231/v2# ServiceName"`
	ServiceDigitalIdentity       ServiceDigitalIdentity       `xml:"http://uri.etsi.org/02231/v2# ServiceDigitalIdentity"`
	ServiceStatus                string                       `xml:"http://uri.etsi.org/02231/v2# ServiceStatus"`
	StatusStartingTime           time.Time                    `xml:"http://uri.etsi.org/02231/v2# StatusStartingTime"`
	ServiceInformationExtensions ServiceInformationExtensions `xml:"http://uri.etsi.org/02231/v2# ServiceInformationExtensions,omitempty"`
}

type ServiceInformation struct {
	XMLName                      xml.Name                     `xml:"http://uri.etsi.org/02231/v2# ServiceInformation"`
	ServiceTypeIdentifier        string                       `xml:"http://uri.etsi.org/02231/v2# ServiceTypeIdentifier"`
	ServiceName                  ServiceName                  `xml:"http://uri.etsi.org/02231/v2# ServiceName"`
	ServiceDigitalIdentity       ServiceDigitalIdentity       `xml:"http://uri.etsi.org/02231/v2# ServiceDigitalIdentity"`
	ServiceStatus                string                       `xml:"http://uri.etsi.org/02231/v2# ServiceStatus"`
	StatusStartingTime           time.Time                    `xml:"http://uri.etsi.org/02231/v2# StatusStartingTime"`
	SchemeServiceDefinitionURI   SchemeServiceDefinitionURI   `xml:"http://uri.etsi.org/02231/v2# SchemeServiceDefinitionURI,omitempty"`
	ServiceSupplyPoints          ServiceSupplyPoints          `xml:"http://uri.etsi.org/02231/v2# ServiceSupplyPoints,omitempty"`
	TSPServiceDefinitionURI      TSPServiceDefinitionURI      `xml:"http://uri.etsi.org/02231/v2# TSPServiceDefinitionURI,omitempty"`
	ServiceInformationExtensions ServiceInformationExtensions `xml:"http://uri.etsi.org/02231/v2# ServiceInformationExtensions,omitempty"`
}

type ServiceInformationExtensions struct {
	XMLName   xml.Name    `xml:"http://uri.etsi.org/02231/v2# ServiceInformationExtensions"`
	Extension []Extension `xml:"http://uri.etsi.org/02231/v2# Extension"`
}

type ServiceName struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# ServiceName"`
	Name    []Name   `xml:"http://uri.etsi.org/02231/v2# Name"`
}

type ServiceSupplyPoints struct {
	XMLName            xml.Name `xml:"http://uri.etsi.org/02231/v2# ServiceSupplyPoints"`
	ServiceSupplyPoint string   `xml:"http://uri.etsi.org/02231/v2# ServiceSupplyPoint"`
}

type TSLLegalNotice struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TSLLegalNotice"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}

type TSLPolicy struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TSLPolicy"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}

type TSPAddress struct {
	XMLName           xml.Name          `xml:"http://uri.etsi.org/02231/v2# TSPAddress"`
	PostalAddresses   PostalAddresses   `xml:"http://uri.etsi.org/02231/v2# PostalAddresses"`
	ElectronicAddress ElectronicAddress `xml:"http://uri.etsi.org/02231/v2# ElectronicAddress"`
}

type TSPInformation struct {
	XMLName                  xml.Name                 `xml:"http://uri.etsi.org/02231/v2# TSPInformation"`
	TSPName                  TSPName                  `xml:"http://uri.etsi.org/02231/v2# TSPName"`
	TSPTradeName             TSPTradeName             `xml:"http://uri.etsi.org/02231/v2# TSPTradeName,omitempty"`
	TSPAddress               TSPAddress               `xml:"http://uri.etsi.org/02231/v2# TSPAddress"`
	TSPInformationURI        TSPInformationURI        `xml:"http://uri.etsi.org/02231/v2# TSPInformationURI"`
	TSPInformationExtensions TSPInformationExtensions `xml:"http://uri.etsi.org/02231/v2# TSPInformationExtensions,omitempty"`
}

type TSPInformationExtensions struct {
	XMLName   xml.Name    `xml:"http://uri.etsi.org/02231/v2# TSPInformationExtensions"`
	Extension []Extension `xml:"http://uri.etsi.org/02231/v2# Extension"`
}

type TSPInformationURI struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TSPInformationURI"`
	URI     []URI    `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type TSPName struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TSPName"`
	Name    []Name   `xml:"http://uri.etsi.org/02231/v2# Name"`
}

type TSPService struct {
	XMLName            xml.Name           `xml:"http://uri.etsi.org/02231/v2# TSPService"`
	ServiceInformation ServiceInformation `xml:"http://uri.etsi.org/02231/v2# ServiceInformation"`
	ServiceHistory     ServiceHistory     `xml:"http://uri.etsi.org/02231/v2# ServiceHistory,omitempty"`
}

type TSPServiceDefinitionURI struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TSPServiceDefinitionURI"`
	URI     []URI    `xml:"http://uri.etsi.org/02231/v2# URI"`
}

type TSPServices struct {
	XMLName    xml.Name     `xml:"http://uri.etsi.org/02231/v2# TSPServices"`
	TSPService []TSPService `xml:"http://uri.etsi.org/02231/v2# TSPService"`
}

type TSPTradeName struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TSPTradeName"`
	Name    []Name   `xml:"http://uri.etsi.org/02231/v2# Name"`
}

type TextualInformation struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# TextualInformation"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}

type TrustServiceProvider struct {
	XMLName        xml.Name       `xml:"http://uri.etsi.org/02231/v2# TrustServiceProvider"`
	TSPInformation TSPInformation `xml:"http://uri.etsi.org/02231/v2# TSPInformation"`
	TSPServices    TSPServices    `xml:"http://uri.etsi.org/02231/v2# TSPServices"`
}

type TrustServiceProviderList struct {
	XMLName              xml.Name               `xml:"http://uri.etsi.org/02231/v2# TrustServiceProviderList"`
	TrustServiceProvider []TrustServiceProvider `xml:"http://uri.etsi.org/02231/v2# TrustServiceProvider"`
}

type URI struct {
	XMLName xml.Name `xml:"http://uri.etsi.org/02231/v2# URI"`
	Lang    string   `xml:"lang,attr"`
	Value   string   `xml:",chardata"`
}
