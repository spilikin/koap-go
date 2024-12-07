package CardTerminalInfo_v8_0

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/int/version/ProductInformation_v1_1"
)

type CardTerminal struct {
	XMLName            xml.Name                                   `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 CardTerminal"`
	ProductInformation ProductInformation_v1_1.ProductInformation `xml:"http://ws.gematik.de/int/version/ProductInformation/v1.1 ProductInformation"`
	CtId               string                                     `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CtId"`
	WorkplaceIds       ConnectorCommon_v5_0.WorkplaceIds          `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 WorkplaceIds"`
	Name               string                                     `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 Name"`
	MacAddress         string                                     `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 MacAddress"`
	IPAddress          IPAddress                                  `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 IPAddress,omitempty"`
	Slots              int                                        `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 Slots"`
	ISPhysical         bool                                       `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 IS_PHYSICAL"`
	Connected          bool                                       `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 Connected"`
}

type CardTerminals struct {
	XMLName      xml.Name       `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 CardTerminals"`
	CardTerminal []CardTerminal `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 CardTerminal,omitempty"`
}

type IPAddress struct {
	XMLName     xml.Name `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 IPAddress"`
	IPV4Address string   `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 IPV4Address,omitempty"`
	IPV6Address string   `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 IPV6Address,omitempty"`
}
