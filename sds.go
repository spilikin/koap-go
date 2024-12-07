package koap

import (
	"context"
	"encoding/xml"
	"log/slog"
	"net/http"
	"net/url"
)

type LocalProductVersion struct {
	HWVersion string `xml:"HWVersion"`
	FWVersion string `xml:"FWVersion"`
}

type ProductVersion struct {
	Local LocalProductVersion `xml:"Local"`
}

type ProductIdentification struct {
	ProductVendorID string         `xml:"ProductVendorID"`
	ProductCode     string         `xml:"ProductCode"`
	ProductVersion  ProductVersion `xml:"ProductVersion"`
}

type ProductTypeInformation struct {
	ProductType        string `xml:"ProductType"`
	ProductTypeVersion string `xml:"ProductTypeVersion"`
}

type ProductInformation struct {
	ProductTypeInformation ProductTypeInformation `xml:"ProductTypeInformation"`
	ProductIdentification  ProductIdentification  `xml:"ProductIdentification"`
}

type ServiceVersionEndpoint struct {
	Location string `xml:"Location,attr"`
}

type ServiceVersion struct {
	TargetNamespace string                 `xml:"TargetNamespace,attr"`
	Version         string                 `xml:"Version,attr"`
	Abstract        string                 `xml:"Abstract"`
	EndpointTLS     ServiceVersionEndpoint `xml:"EndpointTLS"`
	Endpoint        ServiceVersionEndpoint `xml:"Endpoint"`
}

type ServiceName string

const (
	ServiceNameCardService          ServiceName = "CardService"
	ServiceNameEventService         ServiceName = "EventService"
	ServiceNameAuthSignatureService ServiceName = "AuthSignatureService"
	ServiceNameCartTerminalService  ServiceName = "CardTerminalService"
	ServiceNameCertificateService   ServiceName = "CertificateService"
)

type Service struct {
	XMLName  xml.Name         `xml:"Service"`
	Name     ServiceName      `xml:"Name,attr"`
	Abstract string           `xml:"Abstract"`
	Versions []ServiceVersion `xml:"Versions>Version"`
}

type ServiceInformation struct {
	Service []Service `xml:"Service"`
}

type ConnectorServices struct {
	ProductInformation ProductInformation
	ServiceInformation ServiceInformation
}

func LoadConnectorServices(ctx context.Context, httpClient *http.Client, baseUrl url.URL) (*ConnectorServices, error) {
	var services ConnectorServices
	url := baseUrl.ResolveReference(&url.URL{Path: "./connector.sds"})

	slog.Debug("Loading service directory", "url", url.String())

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := xml.NewDecoder(resp.Body).Decode(&services); err != nil {
		return nil, err
	}

	slog.Debug("Loaded service directory", "services", services)

	for _, s := range services.ServiceInformation.Service {
		slog.Debug("Service", "name", s.Name, "versionsCount", len(s.Versions))
		for _, v := range s.Versions {
			slog.Debug("Version", "version", v)
		}
	}

	return &services, nil
}
