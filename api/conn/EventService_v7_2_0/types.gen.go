package EventService_v7_2_0

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/CardService_v8_1"
	"github.com/spilikin/koap-go/api/conn/CardTerminalInfo_v8_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
	"time"
)

type GetCardTerminals struct {
	XMLName     xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetCardTerminals"`
	MandantWide bool                          `xml:"mandant-wide,attr"`
	Context     ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
}

type GetCardTerminalsResponse struct {
	XMLName       xml.Name                            `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetCardTerminalsResponse"`
	Status        ConnectorCommon_v5_0.Status         `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	CardTerminals CardTerminalInfo_v8_0.CardTerminals `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 CardTerminals"`
}

type GetCards struct {
	XMLName     xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetCards"`
	MandantWide bool                          `xml:"mandant-wide,attr"`
	Context     ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CtId        string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CtId,omitempty"`
	SlotId      int                           `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 SlotId,omitempty"`
	CardType    string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CardType,omitempty"`
}

type GetCardsResponse struct {
	XMLName xml.Name                    `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetCardsResponse"`
	Status  ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	Cards   CardService_v8_1.Cards      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Cards"`
}

type GetResourceInformation struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetResourceInformation"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CtId       string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CtId,omitempty"`
	SlotId     int                           `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 SlotId,omitempty"`
	Iccsn      string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 Iccsn,omitempty"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle,omitempty"`
}

type GetResourceInformationResponse struct {
	XMLName      xml.Name                           `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetResourceInformationResponse"`
	Status       ConnectorCommon_v5_0.Status        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	Card         CardService_v8_1.Card              `xml:"http://ws.gematik.de/conn/CardService/v8.1 Card,omitempty"`
	CardTerminal CardTerminalInfo_v8_0.CardTerminal `xml:"http://ws.gematik.de/conn/CardTerminalInfo/v8.0 CardTerminal,omitempty"`
	Connector    ConnectorCommon_v5_0.Connector     `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Connector,omitempty"`
}

type GetSubscription struct {
	XMLName        xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetSubscription"`
	MandantWide    bool                          `xml:"mandant-wide,attr"`
	Context        ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	SubscriptionID string                        `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionID,omitempty"`
}

type GetSubscriptionResponse struct {
	XMLName       xml.Name                    `xml:"http://ws.gematik.de/conn/EventService/v7.2 GetSubscriptionResponse"`
	Status        ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	Subscriptions Subscriptions               `xml:"http://ws.gematik.de/conn/EventService/v7.2 Subscriptions"`
}

type RenewSubscriptions struct {
	XMLName        xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 RenewSubscriptions"`
	Context        ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	SubscriptionID []string                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionID"`
}

type RenewSubscriptionsResponse struct {
	XMLName           xml.Name                    `xml:"http://ws.gematik.de/conn/EventService/v7.2 RenewSubscriptionsResponse"`
	Status            ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	SubscribeRenewals SubscribeRenewals           `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscribeRenewals"`
}

type Subscribe struct {
	XMLName      xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 Subscribe"`
	Context      ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	Subscription Subscription                  `xml:"http://ws.gematik.de/conn/EventService/v7.2 Subscription"`
}

type SubscribeRenewals struct {
	XMLName             xml.Name              `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscribeRenewals"`
	SubscriptionRenewal []SubscriptionRenewal `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionRenewal"`
}

type SubscribeResponse struct {
	XMLName         xml.Name                    `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscribeResponse"`
	Status          ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	SubscriptionID  string                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionID"`
	TerminationTime time.Time                   `xml:"http://ws.gematik.de/conn/EventService/v7.2 TerminationTime"`
}

type Subscription struct {
	XMLName         xml.Name  `xml:"http://ws.gematik.de/conn/EventService/v7.2 Subscription"`
	SubscriptionID  string    `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionID,omitempty"`
	TerminationTime time.Time `xml:"http://ws.gematik.de/conn/EventService/v7.2 TerminationTime,omitempty"`
	EventTo         string    `xml:"http://ws.gematik.de/conn/EventService/v7.2 EventTo"`
	Topic           string    `xml:"http://ws.gematik.de/conn/EventService/v7.2 Topic"`
	Filter          string    `xml:"http://ws.gematik.de/conn/EventService/v7.2 Filter,omitempty"`
}

type SubscriptionRenewal struct {
	XMLName         xml.Name  `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionRenewal"`
	SubscriptionID  string    `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionID"`
	TerminationTime time.Time `xml:"http://ws.gematik.de/conn/EventService/v7.2 TerminationTime"`
}

type Subscriptions struct {
	XMLName      xml.Name       `xml:"http://ws.gematik.de/conn/EventService/v7.2 Subscriptions"`
	Subscription []Subscription `xml:"http://ws.gematik.de/conn/EventService/v7.2 Subscription,omitempty"`
}

type Unsubscribe struct {
	XMLName        xml.Name                      `xml:"http://ws.gematik.de/conn/EventService/v7.2 Unsubscribe"`
	Context        ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	SubscriptionID string                        `xml:"http://ws.gematik.de/conn/EventService/v7.2 SubscriptionID,omitempty"`
	EventTo        string                        `xml:"http://ws.gematik.de/conn/EventService/v7.2 EventTo,omitempty"`
}

type UnsubscribeResponse struct {
	XMLName xml.Name                    `xml:"http://ws.gematik.de/conn/EventService/v7.2 UnsubscribeResponse"`
	Status  ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
}
