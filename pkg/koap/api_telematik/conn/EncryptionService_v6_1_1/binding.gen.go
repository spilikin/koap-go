package EncryptionService_v6_1_1

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationDecryptDocument = &koap.Operation{
	Name:       "DecryptDocument",
	SOAPAction: "http://ws.gematik.de/conn/crypt/EncryptionService/v6.1#DecryptDocument",
	InputType:  reflect.TypeOf(DecryptDocument{}),
	OutputType: reflect.TypeOf(DecryptDocumentResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationEncryptDocument = &koap.Operation{
	Name:       "EncryptDocument",
	SOAPAction: "http://ws.gematik.de/conn/EncryptionService/v6.1#EncryptDocument",
	InputType:  reflect.TypeOf(EncryptDocument{}),
	OutputType: reflect.TypeOf(EncryptDocumentResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
