package CardService_v8_1_2

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationVerifyPin = &koap.Operation{
	Name:       "VerifyPin",
	SOAPAction: "http://ws.gematik.de/conn/CardService/v8.1#VerifyPin",
	InputType:  reflect.TypeOf(VerifyPin{}),
	OutputType: reflect.TypeOf(VerifyPinResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationChangePin = &koap.Operation{
	Name:       "ChangePin",
	SOAPAction: "http://ws.gematik.de/conn/CardService/v8.1#ChangePin",
	InputType:  reflect.TypeOf(ChangePin{}),
	OutputType: reflect.TypeOf(ChangePinResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationUnblockPin = &koap.Operation{
	Name:       "UnblockPin",
	SOAPAction: "http://ws.gematik.de/conn/CardService/v8.1#UnblockPin",
	InputType:  reflect.TypeOf(UnblockPin{}),
	OutputType: reflect.TypeOf(UnblockPinResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetPinStatus = &koap.Operation{
	Name:       "GetPinStatus",
	SOAPAction: "http://ws.gematik.de/conn/CardService/v8.1#GetPinStatus",
	InputType:  reflect.TypeOf(GetPinStatus{}),
	OutputType: reflect.TypeOf(GetPinStatusResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationEnablePin = &koap.Operation{
	Name:       "EnablePin",
	SOAPAction: "http://ws.gematik.de/conn/CardService/v8.1#EnablePin",
	InputType:  reflect.TypeOf(EnablePin{}),
	OutputType: reflect.TypeOf(EnablePinResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationDisablePin = &koap.Operation{
	Name:       "DisablePin",
	SOAPAction: "http://ws.gematik.de/conn/CardService/v8.1#DisablePin",
	InputType:  reflect.TypeOf(DisablePin{}),
	OutputType: reflect.TypeOf(DisablePinResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
