package error_v2_0

import (
	"encoding/xml"
	"time"
)

type Detail struct {
	XMLName  xml.Name `xml:"http://ws.gematik.de/tel/error/v2.0 Detail"`
	Encoding string   `xml:"Encoding,attr"`
	Value    string   `xml:",chardata"`
}

type Error struct {
	XMLName   xml.Name  `xml:"http://ws.gematik.de/tel/error/v2.0 Error"`
	MessageID string    `xml:"http://ws.gematik.de/tel/error/v2.0 MessageID"`
	Timestamp time.Time `xml:"http://ws.gematik.de/tel/error/v2.0 Timestamp"`
	Trace     []Trace   `xml:"http://ws.gematik.de/tel/error/v2.0 Trace"`
}

type Trace struct {
	XMLName      xml.Name `xml:"http://ws.gematik.de/tel/error/v2.0 Trace"`
	EventID      string   `xml:"http://ws.gematik.de/tel/error/v2.0 EventID"`
	Instance     string   `xml:"http://ws.gematik.de/tel/error/v2.0 Instance"`
	LogReference string   `xml:"http://ws.gematik.de/tel/error/v2.0 LogReference"`
	CompType     string   `xml:"http://ws.gematik.de/tel/error/v2.0 CompType"`
	Code         int      `xml:"http://ws.gematik.de/tel/error/v2.0 Code"`
	Severity     string   `xml:"http://ws.gematik.de/tel/error/v2.0 Severity"`
	ErrorType    string   `xml:"http://ws.gematik.de/tel/error/v2.0 ErrorType"`
	ErrorText    string   `xml:"http://ws.gematik.de/tel/error/v2.0 ErrorText"`
	Detail       Detail   `xml:"http://ws.gematik.de/tel/error/v2.0 Detail,omitempty"`
}
