package EncryptionService_v6_1_0

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

var OperationDecryptDocument = &operation{
	name:       "DecryptDocument",
	soapAction: "http://ws.gematik.de/conn/crypt/EncryptionService/v6.1#DecryptDocument",
	inputType:  reflect.TypeOf(DecryptDocument{}),
	outputType: reflect.TypeOf(DecryptDocumentResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationEncryptDocument = &operation{
	name:       "EncryptDocument",
	soapAction: "http://ws.gematik.de/conn/EncryptionService/v6.1#EncryptDocument",
	inputType:  reflect.TypeOf(EncryptDocument{}),
	outputType: reflect.TypeOf(EncryptDocumentResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
