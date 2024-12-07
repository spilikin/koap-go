package EventService_v7_2_0

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

var OperationSubscribe = &operation{
	name:       "Subscribe",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#Subscribe",
	inputType:  reflect.TypeOf(Subscribe{}),
	outputType: reflect.TypeOf(SubscribeResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationUnsubscribe = &operation{
	name:       "Unsubscribe",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#Unsubscribe",
	inputType:  reflect.TypeOf(Unsubscribe{}),
	outputType: reflect.TypeOf(UnsubscribeResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetSubscription = &operation{
	name:       "GetSubscription",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#GetSubscription",
	inputType:  reflect.TypeOf(GetSubscription{}),
	outputType: reflect.TypeOf(GetSubscriptionResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetResourceInformation = &operation{
	name:       "GetResourceInformation",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#GetResourceInformation",
	inputType:  reflect.TypeOf(GetResourceInformation{}),
	outputType: reflect.TypeOf(GetResourceInformationResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetCardTerminals = &operation{
	name:       "GetCardTerminals",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#GetCardTerminals",
	inputType:  reflect.TypeOf(GetCardTerminals{}),
	outputType: reflect.TypeOf(GetCardTerminalsResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationGetCards = &operation{
	name:       "GetCards",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#GetCards",
	inputType:  reflect.TypeOf(GetCards{}),
	outputType: reflect.TypeOf(GetCardsResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}

var OperationRenewSubscriptions = &operation{
	name:       "RenewSubscriptions",
	soapAction: "http://ws.gematik.de/conn/EventService/v7.2#RenewSubscriptions",
	inputType:  reflect.TypeOf(RenewSubscriptions{}),
	outputType: reflect.TypeOf(RenewSubscriptionsResponse{}),
	faultType:  reflect.TypeOf(error_v2_0.Error{}),
}
