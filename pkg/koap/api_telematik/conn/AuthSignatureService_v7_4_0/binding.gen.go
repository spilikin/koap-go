package AuthSignatureService_v7_4_0

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/SignatureService_v7_4"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationExternalAuthenticate = &koap.Operation{
	Name:       "ExternalAuthenticate",
	SOAPAction: "http://ws.gematik.de/conn/SignatureService/v7.4#ExternalAuthenticate",
	InputType:  reflect.TypeOf(SignatureService_v7_4.ExternalAuthenticate{}),
	OutputType: reflect.TypeOf(SignatureService_v7_4.ExternalAuthenticateResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
