package SignatureService_v7_4_2

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationVerifyDocument = &koap.Operation{
	Name:       "VerifyDocument",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.4#VerifyDocument",
	InputType:  reflect.TypeOf(VerifyDocument{}),
	OutputType: reflect.TypeOf(VerifyDocumentResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationSignDocument = &koap.Operation{
	Name:       "SignDocument",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.4#SignDocument",
	InputType:  reflect.TypeOf(SignDocument{}),
	OutputType: reflect.TypeOf(SignDocumentResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetJobNumber = &koap.Operation{
	Name:       "GetJobNumber",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.4#GetJobNumber",
	InputType:  reflect.TypeOf(GetJobNumber{}),
	OutputType: reflect.TypeOf(GetJobNumberResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationStopSignature = &koap.Operation{
	Name:       "StopSignature",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.4#StopSignature",
	InputType:  reflect.TypeOf(StopSignature{}),
	OutputType: reflect.TypeOf(StopSignatureResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
