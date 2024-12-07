package CardService_v8_1_2

import (
	"encoding/xml"
	"github.com/spilikin/koap-go/api/conn/ConnectorCommon_v5_0"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
)

type ChangePin struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 ChangePin"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	PinTyp     string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinTyp"`
}

type ChangePinResponse struct {
	XMLName   xml.Name                    `xml:"http://ws.gematik.de/conn/CardService/v8.1 ChangePinResponse"`
	Status    ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	PinResult string                      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinResult"`
	LeftTries int                         `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 LeftTries,omitempty"`
}

type DisablePin struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 DisablePin"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	PinTyp     string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinTyp"`
}

type DisablePinResponse struct {
	XMLName   xml.Name                    `xml:"http://ws.gematik.de/conn/CardService/v8.1 DisablePinResponse"`
	Status    ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	PinResult string                      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinResult"`
	LeftTries int                         `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 LeftTries,omitempty"`
}

type EnablePin struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 EnablePin"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	PinTyp     string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinTyp"`
}

type EnablePinResponse struct {
	XMLName   xml.Name                    `xml:"http://ws.gematik.de/conn/CardService/v8.1 EnablePinResponse"`
	Status    ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	PinResult string                      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinResult"`
	LeftTries int                         `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 LeftTries,omitempty"`
}

type GetPinStatus struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 GetPinStatus"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	PinTyp     string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinTyp"`
}

type GetPinStatusResponse struct {
	XMLName   xml.Name                    `xml:"http://ws.gematik.de/conn/CardService/v8.1 GetPinStatusResponse"`
	Status    ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	PinStatus string                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 PinStatus,omitempty"`
	LeftTries int                         `xml:"http://ws.gematik.de/conn/CardService/v8.1 LeftTries,omitempty"`
}

type UnblockPin struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 UnblockPin"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	PinTyp     string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinTyp"`
	SetNewPin  bool                          `xml:"http://ws.gematik.de/conn/CardService/v8.1 SetNewPin,omitempty"`
}

type UnblockPinResponse struct {
	XMLName   xml.Name                    `xml:"http://ws.gematik.de/conn/CardService/v8.1 UnblockPinResponse"`
	Status    ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	PinResult string                      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinResult"`
	LeftTries int                         `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 LeftTries,omitempty"`
}

type VerifyPin struct {
	XMLName    xml.Name                      `xml:"http://ws.gematik.de/conn/CardService/v8.1 VerifyPin"`
	Context    ConnectorContext_v2_0.Context `xml:"http://ws.gematik.de/conn/ConnectorContext/v2.0 Context"`
	CardHandle string                        `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	PinTyp     string                        `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinTyp"`
}

type VerifyPinResponse struct {
	XMLName   xml.Name                    `xml:"http://ws.gematik.de/conn/CardService/v8.1 VerifyPinResponse"`
	Status    ConnectorCommon_v5_0.Status `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 Status"`
	PinResult string                      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 PinResult"`
	LeftTries int                         `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 LeftTries,omitempty"`
}
