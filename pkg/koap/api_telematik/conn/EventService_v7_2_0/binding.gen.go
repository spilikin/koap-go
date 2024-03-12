package EventService_v7_2_0

import (
	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/tel/error_v2_0"
	"reflect"
)

var OperationSubscribe = &koap.Operation{
	Name:       "Subscribe",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#Subscribe",
	InputType:  reflect.TypeOf(Subscribe{}),
	OutputType: reflect.TypeOf(SubscribeResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationUnsubscribe = &koap.Operation{
	Name:       "Unsubscribe",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#Unsubscribe",
	InputType:  reflect.TypeOf(Unsubscribe{}),
	OutputType: reflect.TypeOf(UnsubscribeResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetSubscription = &koap.Operation{
	Name:       "GetSubscription",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#GetSubscription",
	InputType:  reflect.TypeOf(GetSubscription{}),
	OutputType: reflect.TypeOf(GetSubscriptionResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetResourceInformation = &koap.Operation{
	Name:       "GetResourceInformation",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#GetResourceInformation",
	InputType:  reflect.TypeOf(GetResourceInformation{}),
	OutputType: reflect.TypeOf(GetResourceInformationResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetCardTerminals = &koap.Operation{
	Name:       "GetCardTerminals",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#GetCardTerminals",
	InputType:  reflect.TypeOf(GetCardTerminals{}),
	OutputType: reflect.TypeOf(GetCardTerminalsResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetCards = &koap.Operation{
	Name:       "GetCards",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#GetCards",
	InputType:  reflect.TypeOf(GetCards{}),
	OutputType: reflect.TypeOf(GetCardsResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationRenewSubscriptions = &koap.Operation{
	Name:       "RenewSubscriptions",
	SOAPAction: "http://ws.gematik.de/conn/EventService/v7.2#RenewSubscriptions",
	InputType:  reflect.TypeOf(RenewSubscriptions{}),
	OutputType: reflect.TypeOf(RenewSubscriptionsResponse{}),
	FaultType:  reflect.TypeOf(error_v2_0.Error{}),
}
