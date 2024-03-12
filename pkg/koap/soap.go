package koap

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

type Operation struct {
	Name       string
	SOAPAction string
	InputType  reflect.Type
	OutputType reflect.Type
	FaultType  reflect.Type
}

type Binding struct {
	Style      string
	Transport  string
	Operations []*Operation
}

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
}

type Header struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
}

type Body struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Content []byte   `xml:",innerxml"`
}

type FaultDetail struct {
	DetailRaw string      `xml:",innerxml"`
	Detail    interface{} `xml:"-"`
}

type Fault struct {
	XMLName     xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string      `xml:"faultcode"`
	FaultString string      `xml:"faultstring"`
	Detail      FaultDetail `xml:"detail"`
}

func (f *Fault) Error() string {
	return fmt.Sprintf("%s: %s", f.FaultCode, f.FaultString)
}

type FaultEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  Header
	Body    struct {
		Fault Fault
	}
}
