package CardTerminalService_v1_1_0

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/CardService_v8_1"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
)

type EjectCard struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 EjectCard"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle,omitempty"`
	Slot       Slot                          `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 Slot,omitempty"`
	DisplayMsg string                        `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 DisplayMsg,omitempty"`
	TimeOut    int                           `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 TimeOut,omitempty"`
}

type EjectCardResponse struct {
	XMLName xml.Name                    `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 EjectCardResponse"`
	Status  ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
}

type RequestCard struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 RequestCard"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	Slot       Slot                          `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 Slot"`
	CardType   string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CardType,omitempty"`
	DisplayMsg string                        `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 DisplayMsg,omitempty"`
	TimeOut    int                           `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 TimeOut,omitempty"`
}

type RequestCardResponse struct {
	XMLName         xml.Name                    `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 RequestCardResponse"`
	Status          ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	AlreadyInserted bool                        `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 AlreadyInserted,omitempty"`
	Card            CardService_v8_1.Card       `xml:"http://ws.gematik.de/conn/CardService/v8.1 Card,omitempty"`
}

type Slot struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/CardTerminalService/v1.1 Slot"`
	CtId    string   `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CtId"`
	SlotId  int      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 SlotId"`
}
