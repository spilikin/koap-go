package main

import (
	"errors"
	"log"
	"log/slog"

	"github.com/spilikin/koap-go"
	"github.com/spilikin/koap-go/api/conn/CardService_v8_1"
	"github.com/spilikin/koap-go/api/conn/CertificateService_v6_0_1"
	"github.com/spilikin/koap-go/api/conn/ConnectorContext_v2_0"
	"github.com/spilikin/koap-go/api/conn/EventService_v7_2_0"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	cfg, err := koap.ReadConfigFromEnv("")
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Read config", "config", cfg)

	client, err := koap.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Connected to konnektor", "uri", client.BaseURL.String())

	for _, service := range client.Services.ServiceInformation.Service {
		for _, version := range service.Versions {
			slog.Info("Service", "service", service.Name, "version", version.Version)
		}
	}

	serviceProxy, err := client.CreateServiceProxy(koap.ServiceNameEventService, "7.2.0")
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Service proxy", "serviceProxy", serviceProxy.String())

	document := EventService_v7_2_0.GetCards{
		Context: Context(client.Context),
	}

	response, err := serviceProxy.CallWithDocument(EventService_v7_2_0.OperationGetCards, &document)
	if err != nil {
		log.Fatal(err)
	}

	getCardsResponse, ok := response.(*EventService_v7_2_0.GetCardsResponse)
	if !ok {
		log.Fatal("unexpected response type")
	}

	for _, card := range getCardsResponse.Cards.Card {
		slog.Info("Card", "cardType", card.CardType)
		err := loadCertificate(client, card)
		if err != nil {
			slog.Error("Unable to load certificate", "error", err, "cardType", card.CardType, "cardHandle", card.CardHandle)
		}
	}

}

func loadCertificate(client *koap.Client, card CardService_v8_1.Card) error {
	document := CertificateService_v6_0_1.ReadCardCertificate{
		Context:    Context(client.Context),
		CardHandle: card.CardHandle,
		CertRefList: CertificateService_v6_0_1.CertRefList{
			CertRef: []string{"C.AUT"}, // TODO: use strings until enum is generated
		},
		Crypt: "ECC",
	}

	serviceProxy, err := client.CreateServiceProxy(koap.ServiceNameCertificateService, "6.0.1")
	if err != nil {
		return err
	}

	response, err := serviceProxy.CallWithDocument(CertificateService_v6_0_1.OperationReadCardCertificate, &document)
	if err != nil {
		return err
	}

	readCardCertificateResponse, ok := response.(*CertificateService_v6_0_1.ReadCardCertificateResponse)
	if !ok {
		return errors.New("unexpected response type")
	}

	slog.Info("Certificate", "certificate", readCardCertificateResponse)

	return nil
}

func Context(ctx koap.ConnectorContext) ConnectorContext_v2_0.Context {
	return ConnectorContext_v2_0.Context{
		MandantId:      ctx.MandantId,
		ClientSystemId: ctx.ClientSystemId,
		WorkplaceId:    ctx.WorkplaceId,
		UserId:         ctx.UserId,
	}
}
