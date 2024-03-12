package CardTerminalService_v1_1_0

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationRequestCard = &koap.Operation{
	Name:       "RequestCard",
	SOAPAction: "http://ws.gematik.de/conn/CardTerminalService/v1.1#RequestCard",
	InputType:  reflect.TypeOf(RequestCard{}),
	OutputType: reflect.TypeOf(RequestCardResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationEjectCard = &koap.Operation{
	Name:       "EjectCard",
	SOAPAction: "http://ws.gematik.de/conn/CardTerminalService/v1.1#EjectCard",
	InputType:  reflect.TypeOf(EjectCard{}),
	OutputType: reflect.TypeOf(EjectCardResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
