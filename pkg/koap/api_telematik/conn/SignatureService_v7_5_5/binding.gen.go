package SignatureService_v7_5_5

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationVerifyDocument = &koap.Operation{
	Name:       "VerifyDocument",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#VerifyDocument",
	InputType:  reflect.TypeOf(VerifyDocument{}),
	OutputType: reflect.TypeOf(VerifyDocumentResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationSignDocument = &koap.Operation{
	Name:       "SignDocument",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#SignDocument",
	InputType:  reflect.TypeOf(SignDocument{}),
	OutputType: reflect.TypeOf(SignDocumentResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetJobNumber = &koap.Operation{
	Name:       "GetJobNumber",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#GetJobNumber",
	InputType:  reflect.TypeOf(GetJobNumber{}),
	OutputType: reflect.TypeOf(GetJobNumberResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationStopSignature = &koap.Operation{
	Name:       "StopSignature",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#StopSignature",
	InputType:  reflect.TypeOf(StopSignature{}),
	OutputType: reflect.TypeOf(StopSignatureResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationActivateComfortSignature = &koap.Operation{
	Name:       "ActivateComfortSignature",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#ActivateComfortSignature",
	InputType:  reflect.TypeOf(ActivateComfortSignature{}),
	OutputType: reflect.TypeOf(ActivateComfortSignatureResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationDeactivateComfortSignature = &koap.Operation{
	Name:       "DeactivateComfortSignature",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#DeactivateComfortSignature",
	InputType:  reflect.TypeOf(DeactivateComfortSignature{}),
	OutputType: reflect.TypeOf(DeactivateComfortSignatureResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetSignatureMode = &koap.Operation{
	Name:       "GetSignatureMode",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.5#GetSignatureMode",
	InputType:  reflect.TypeOf(GetSignatureMode{}),
	OutputType: reflect.TypeOf(GetSignatureModeResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
