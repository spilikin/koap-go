package ConnectorCommon_v5_0

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/ext/dss10core"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"time"
)

type Connector struct {
	XMLName        xml.Name       `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Connector"`
	VPNTIStatus    VPNTIStatus    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 VPNTIStatus"`
	VPNSISStatus   VPNSISStatus   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 VPNSISStatus"`
	OperatingState OperatingState `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 OperatingState"`
}

type Document struct {
	XMLName    xml.Name             `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Document"`
	ID         string               `xml:"ID,attr"`
	RefURI     string               `xml:"RefURI,attr"`
	RefType    string               `xml:"RefType,attr"`
	SchemaRefs string               `xml:"SchemaRefs,attr"`
	Base64XML  string               `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Base64XML,omitempty"`
	Base64Data dss10core.Base64Data `xml:"urn:oasis:names:tc:dss:1.0:core:schema Base64Data,omitempty"`
}

type ErrorState struct {
	XMLName        xml.Name  `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ErrorState"`
	ErrorCondition string    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ErrorCondition"`
	Severity       string    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Severity"`
	Type           string    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Type"`
	Value          bool      `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Value"`
	ValidFrom      time.Time `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ValidFrom"`
}

type OperatingState struct {
	XMLName    xml.Name   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 OperatingState"`
	ErrorState ErrorState `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ErrorState"`
}

type Status struct {
	XMLName xml.Name         `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	Result  string           `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Result"`
	Error   error_v2_0.Error `xml:"http://ws.gematik.de/tel/error/v2.0 Error,omitempty"`
}

type VPNSISStatus struct {
	XMLName          xml.Name  `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 VPNSISStatus"`
	ConnectionStatus string    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ConnectionStatus"`
	Timestamp        time.Time `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Timestamp"`
}

type VPNTIStatus struct {
	XMLName          xml.Name  `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 VPNTIStatus"`
	ConnectionStatus string    `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ConnectionStatus"`
	Timestamp        time.Time `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Timestamp"`
}

type WorkplaceIds struct {
	XMLName     xml.Name `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 WorkplaceIds"`
	WorkplaceId []string `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 WorkplaceId,omitempty"`
}
