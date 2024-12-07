package AuthSignatureService_v7_4_0

import (
	"github.com/spilikin/koap-go/api/conn/SignatureService_v7_4"
	"github.com/spilikin/koap-go/api/tel/error_v2_0"
	"reflect"
)

type operation struct {
	name       string
	soapAction string
	inputType  reflect.Type
	outputType reflect.Type
	faultType  reflect.Type
}

func (o *operation) Name() string { return o.name }

func (o *operation) SOAPAction() string { return o.soapAction }

func (o *operation) InputType() reflect.Type { return o.inputType }

func (o *operation) OutputType() reflect.Type { return o.outputType }

func (o *operation) FaultType() reflect.Type { return o.faultType }

var OperationExternalAuthenticate = &operation{
	name:       "ExternalAuthenticate",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.4#ExternalAuthenticate",
	inputType:  reflect.TypeOf(SignatureService_v7_4.ExternalAuthenticate{}),
	outputType: reflect.TypeOf(SignatureService_v7_4.ExternalAuthenticateResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
