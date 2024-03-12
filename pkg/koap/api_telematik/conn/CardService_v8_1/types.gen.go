package CardService_v8_1

import (
	"encoding/xml"
	"time"
)

type ATRVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 ATRVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type COSVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 COSVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type Card struct {
	XMLName                   xml.Name    `xml:"http://ws.gematik.de/conn/CardService/v8.1 Card"`
	CardHandle                string      `xml:"http://ws.gematik.de/conn/ConnectorCommon/v5.0 CardHandle"`
	CardType                  string      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CardType"`
	CardVersion               CardVersion `xml:"http://ws.gematik.de/conn/CardService/v8.1 CardVersion,omitempty"`
	Iccsn                     string      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 Iccsn,omitempty"`
	CtId                      string      `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 CtId"`
	SlotId                    int         `xml:"http://ws.gematik.de/conn/CardServiceCommon/v2.0 SlotId"`
	InsertTime                time.Time   `xml:"http://ws.gematik.de/conn/CardService/v8.1 InsertTime"`
	CardHolderName            string      `xml:"http://ws.gematik.de/conn/CardService/v8.1 CardHolderName,omitempty"`
	Kvnr                      string      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Kvnr,omitempty"`
	CertificateExpirationDate interface{} `xml:"http://ws.gematik.de/conn/CardService/v8.1 CertificateExpirationDate,omitempty"`
}

type CardPTPersVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 CardPTPersVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type CardVersion struct {
	XMLName              xml.Name             `xml:"http://ws.gematik.de/conn/CardService/v8.1 CardVersion"`
	COSVersion           COSVersion           `xml:"http://ws.gematik.de/conn/CardService/v8.1 COSVersion"`
	ObjectSystemVersion  ObjectSystemVersion  `xml:"http://ws.gematik.de/conn/CardService/v8.1 ObjectSystemVersion"`
	CardPTPersVersion    CardPTPersVersion    `xml:"http://ws.gematik.de/conn/CardService/v8.1 CardPTPersVersion,omitempty"`
	DataStructureVersion DataStructureVersion `xml:"http://ws.gematik.de/conn/CardService/v8.1 DataStructureVersion,omitempty"`
	LoggingVersion       LoggingVersion       `xml:"http://ws.gematik.de/conn/CardService/v8.1 LoggingVersion,omitempty"`
	ATRVersion           ATRVersion           `xml:"http://ws.gematik.de/conn/CardService/v8.1 ATRVersion,omitempty"`
	GDOVersion           GDOVersion           `xml:"http://ws.gematik.de/conn/CardService/v8.1 GDOVersion,omitempty"`
	KeyInfoVersion       KeyInfoVersion       `xml:"http://ws.gematik.de/conn/CardService/v8.1 KeyInfoVersion,omitempty"`
}

type Cards struct {
	XMLName xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 Cards"`
	Card    []Card   `xml:"http://ws.gematik.de/conn/CardService/v8.1 Card,omitempty"`
}

type DataStructureVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 DataStructureVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type GDOVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 GDOVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type KeyInfoVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 KeyInfoVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type LoggingVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 LoggingVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}

type ObjectSystemVersion struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/conn/CardService/v8.1 ObjectSystemVersion"`
	Major    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Major"`
	Minor    int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Minor"`
	Revision int      `xml:"http://ws.gematik.de/conn/CardService/v8.1 Revision"`
}
