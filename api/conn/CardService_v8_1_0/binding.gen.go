package CardService_v8_1_0

import (
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

var OperationVerifyPin = &operation{
	name:       "VerifyPin",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#VerifyPin",
	inputType:  reflect.TypeOf(VerifyPin{}),
	outputType: reflect.TypeOf(VerifyPinResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationChangePin = &operation{
	name:       "ChangePin",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#ChangePin",
	inputType:  reflect.TypeOf(ChangePin{}),
	outputType: reflect.TypeOf(ChangePinResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationUnblockPin = &operation{
	name:       "UnblockPin",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#UnblockPin",
	inputType:  reflect.TypeOf(UnblockPin{}),
	outputType: reflect.TypeOf(UnblockPinResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetPinStatus = &operation{
	name:       "GetPinStatus",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#GetPinStatus",
	inputType:  reflect.TypeOf(GetPinStatus{}),
	outputType: reflect.TypeOf(GetPinStatusResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationAuthorizeSMC = &operation{
	name:       "AuthorizeSMC",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#AuthorizeSMC",
	inputType:  reflect.TypeOf(AuthorizeSmc{}),
	outputType: reflect.TypeOf(AuthorizeSmcResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
