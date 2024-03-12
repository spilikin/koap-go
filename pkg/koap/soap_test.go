package koap_test

import (
	"encoding/xml"
	"log/slog"
	"testing"

	"github.com/spilikin/koap-go/pkg/koap"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/pkg/koap/api_telematik/conn/EventService_v7_2_0"
)

func init() {
	// set slog level to debug
	// slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestSerialize(t *testing.T) {
	message := EventService_v7_2_0.GetCards{
		Context: ConnectorContext_v2_0.Context{
			MandantId:      "m1",
			ClientSystemId: "c1",
			WorkplaceId:    "w1",
			UserId:         "u1",
		},
	}

	messageXML, err := xml.Marshal(message)
	if err != nil {
		t.Fatalf("error marshalling: %v", err)
	}

	envelope := koap.Envelope{
		Body: koap.Body{
			Content: messageXML,
		},
	}
	enc, err := xml.Marshal(envelope)
	if err != nil {
		t.Errorf("error marshalling: %v", err)
	}
	t.Logf("envelope: %s", enc)
}

func TestSOAP(t *testing.T) {
	config, err := koap.ReadConfigFromEnv("")
	if err != nil {
		t.Errorf("error reading config from env: %v", err)
	}

	client, _ := koap.NewClient(config)
	//ep := "https://tig.spilikin.dev/ws/EventService"
	//soapAction := "http://ws.gematik.de/conn/EventService/v7.2#GetCards"

	message := EventService_v7_2_0.GetCards{
		Context: ConnectorContext_v2_0.Context{
			MandantId:      "m1",
			ClientSystemId: "c1",
			WorkplaceId:    "w1",
			UserId:         "u1",
		},
	}

	svc, err := client.CreateServiceProxy(koap.ServiceNameEventService, "7.2.0")
	if err != nil {
		t.Fatalf("error creating service proxy: %v", err)
	}

	r, err := svc.CallWithDocument(EventService_v7_2_0.OperationGetCards, &message)
	if err != nil {
		t.Fatalf("error calling document: %v", err)
	}

	slog.Info("response", "response", r)

}
