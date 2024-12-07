package SignatureService_v7_5_6

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

var OperationVerifyDocument = &operation{
	name:       "VerifyDocument",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#VerifyDocument",
	inputType:  reflect.TypeOf(VerifyDocument{}),
	outputType: reflect.TypeOf(VerifyDocumentResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationSignDocument = &operation{
	name:       "SignDocument",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#SignDocument",
	inputType:  reflect.TypeOf(SignDocument{}),
	outputType: reflect.TypeOf(SignDocumentResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetJobNumber = &operation{
	name:       "GetJobNumber",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#GetJobNumber",
	inputType:  reflect.TypeOf(GetJobNumber{}),
	outputType: reflect.TypeOf(GetJobNumberResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationStopSignature = &operation{
	name:       "StopSignature",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#StopSignature",
	inputType:  reflect.TypeOf(StopSignature{}),
	outputType: reflect.TypeOf(StopSignatureResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationActivateComfortSignature = &operation{
	name:       "ActivateComfortSignature",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#ActivateComfortSignature",
	inputType:  reflect.TypeOf(ActivateComfortSignature{}),
	outputType: reflect.TypeOf(ActivateComfortSignatureResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationDeactivateComfortSignature = &operation{
	name:       "DeactivateComfortSignature",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#DeactivateComfortSignature",
	inputType:  reflect.TypeOf(DeactivateComfortSignature{}),
	outputType: reflect.TypeOf(DeactivateComfortSignatureResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetSignatureMode = &operation{
	name:       "GetSignatureMode",
	soapAction: "http://ws.gematik.de/conn/SignatureService/v7.5#GetSignatureMode",
	inputType:  reflect.TypeOf(GetSignatureMode{}),
	outputType: reflect.TypeOf(GetSignatureModeResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
