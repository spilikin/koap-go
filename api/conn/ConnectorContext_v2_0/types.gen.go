package ConnectorContext_v2_0

import (
	"encoding/xml"
)

type Context struct {
	XMLName        xml.Name `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	MandantId      string   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 MandantId"`
	ClientSystemId string   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 ClientSystemId"`
	WorkplaceId    string   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 WorkplaceId"`
	UserId         string   `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 UserId,omitempty"`
}
