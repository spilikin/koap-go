package CardService_v8_1_2

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

var OperationEnablePin = &operation{
	name:       "EnablePin",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#EnablePin",
	inputType:  reflect.TypeOf(EnablePin{}),
	outputType: reflect.TypeOf(EnablePinResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationDisablePin = &operation{
	name:       "DisablePin",
	soapAction: "http://ws.gematik.de/conn/CardService/v8.1#DisablePin",
	inputType:  reflect.TypeOf(DisablePin{}),
	outputType: reflect.TypeOf(DisablePinResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
