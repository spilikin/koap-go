package CertificateService_v6_0_0

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

var OperationCheckCertificateExpiration = &operation{
	name:       "CheckCertificateExpiration",
	soapAction: "http://ws.gematik.de/conn/CertificateService/v6.0#CheckCertificateExpiration",
	inputType:  reflect.TypeOf(CheckCertificateExpiration{}),
	outputType: reflect.TypeOf(CheckCertificateExpirationResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationReadCardCertificate = &operation{
	name:       "ReadCardCertificate",
	soapAction: "http://ws.gematik.de/conn/CertificateService/v6.0#ReadCardCertificate",
	inputType:  reflect.TypeOf(ReadCardCertificate{}),
	outputType: reflect.TypeOf(ReadCardCertificateResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationVerifyCertificate = &operation{
	name:       "VerifyCertificate",
	soapAction: "http://ws.gematik.de/conn/CertificateService/v6.0#VerifyCertificate",
	inputType:  reflect.TypeOf(VerifyCertificate{}),
	outputType: reflect.TypeOf(VerifyCertificateResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
