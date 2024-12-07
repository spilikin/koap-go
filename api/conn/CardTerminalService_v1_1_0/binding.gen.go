package CardTerminalService_v1_1_0

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

var OperationRequestCard = &operation{
	name:       "RequestCard",
	soapAction: "http://ws.gematik.de/conn/CardTerminalService/v1.1#RequestCard",
	inputType:  reflect.TypeOf(RequestCard{}),
	outputType: reflect.TypeOf(RequestCardResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationEjectCard = &operation{
	name:       "EjectCard",
	soapAction: "http://ws.gematik.de/conn/CardTerminalService/v1.1#EjectCard",
	inputType:  reflect.TypeOf(EjectCard{}),
	outputType: reflect.TypeOf(EjectCardResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
