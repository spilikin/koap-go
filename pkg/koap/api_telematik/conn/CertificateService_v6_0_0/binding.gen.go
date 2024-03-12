package CertificateService_v6_0_0

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationCheckCertificateExpiration = &koap.Operation{
	Name:       "CheckCertificateExpiration",
	SOAPAction: "http://ws.gematik.de/conn/CertificateService/v6.0#CheckCertificateExpiration",
	InputType:  reflect.TypeOf(CheckCertificateExpiration{}),
	OutputType: reflect.TypeOf(CheckCertificateExpirationResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationReadCardCertificate = &koap.Operation{
	Name:       "ReadCardCertificate",
	SOAPAction: "http://ws.gematik.de/conn/CertificateService/v6.0#ReadCardCertificate",
	InputType:  reflect.TypeOf(ReadCardCertificate{}),
	OutputType: reflect.TypeOf(ReadCardCertificateResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationVerifyCertificate = &koap.Operation{
	Name:       "VerifyCertificate",
	SOAPAction: "http://ws.gematik.de/conn/CertificateService/v6.0#VerifyCertificate",
	InputType:  reflect.TypeOf(VerifyCertificate{}),
	OutputType: reflect.TypeOf(VerifyCertificateResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
