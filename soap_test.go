package koap_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/spilikin/koap-go"
	"github.com/spilikin/koap-go/api/gematik/conn/connectorcontext20"
	"github.com/spilikin/koap-go/api/gematik/conn/eventservice72"
)

func init() {
	// set slog level to debug
	// slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestSerialize(t *testing.T) {
	message := eventservice72.GetCards{
		Context: connectorcontext20.Context{
			MandantId:      "m1",
			ClientSystemId: "c1",
			WorkplaceId:    "w1",
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
	konFile := os.Getenv("KONNEKTOR_KON_FILE")
	if konFile == "" {
		t.Skip("KONNEKTOR_KON_FILE not set")
	}

	data, err := os.ReadFile(konFile)
	if err != nil {
		t.Fatalf("error reading .kon file: %v", err)
	}

	config, err := koap.ParseDotkon(data)
	if err != nil {
		t.Fatalf("error parsing .kon file: %v", err)
	}

	client, err := koap.NewClient(config)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	message := eventservice72.GetCards{
		Context: connectorcontext20.Context{
			MandantId:      "m1",
			ClientSystemId: "c1",
			WorkplaceId:    "w1",
		},
	}

	svc, err := client.CreateServiceProxy(koap.ServiceNameEventService, "7.2.0")
	if err != nil {
		t.Fatalf("error creating service proxy: %v", err)
	}

	t.Log(svc)
	t.Log(message)
}
