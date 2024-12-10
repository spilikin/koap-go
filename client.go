package koap

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

type ConnectorContext struct {
	MandantId      string
	ClientSystemId string
	WorkplaceId    string
	UserId         string
}

type Client struct {
	httpClient *http.Client
	BaseURL    *url.URL
	Context    ConnectorContext
	Services   *ConnectorServices
}

func NewClient(config *Config) (*Client, error) {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			ServerName: config.TrustedSan,
		},
	}

	if len(config.TrustedCertificates) > 0 {
		certPool := x509.NewCertPool()
		for _, cert := range config.TrustedCertificates {
			slog.Debug("adding trusted certificate", "subject", cert.Subject)
			certPool.AddCert(cert)
		}

		transport.TLSClientConfig.RootCAs = certPool

	}

	client := &Client{
		Context: ConnectorContext{
			MandantId:      config.MandantId,
			ClientSystemId: config.ClientSystemId,
			WorkplaceId:    config.WorkplaceId,
			UserId:         config.UserId,
		},
	}

	var err error

	if client.BaseURL, err = url.Parse(config.URL); err != nil {
		return nil, fmt.Errorf("parsing base URL: %w", err)
	}

	if config.Credentials != nil {
		switch cred := config.Credentials.(type) {
		case CredentialsConfigBasic:
			client.httpClient.Transport = &basicAuthTransport{
				username: cred.Username,
				password: cred.Password,
				T:        client.httpClient.Transport,
			}
		case CredentialsCertDER:
			transport.TLSClientConfig.Certificates = []tls.Certificate{cred.Certificate}
		default:
			return nil, fmt.Errorf("unsupported credentials")
		}
	}

	client.httpClient = &http.Client{
		Transport: transport,
	}

	client.Services, err = LoadConnectorServices(context.TODO(), client.httpClient, client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("loading connector services: %w", err)
	}

	return client, nil
}

type basicAuthTransport struct {
	username string
	password string
	T        http.RoundTripper
}

func (t *basicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.T.RoundTrip(req)
}

func ReadConfigFromEnv(prefix string) (*Config, error) {
	if prefix == "" {
		prefix = "KONNEKTOR_"
	}

	var creds CredentialsConfig

	if os.Getenv(prefix+"AUTH_METHOD") == "basic" {
		creds = CredentialsConfigBasic{
			Type:     "basic",
			Username: os.Getenv(prefix + "AUTH_BASIC_USERNAME"),
			Password: os.Getenv(prefix + "AUTH_BASIC_PASSWORD"),
		}
	}

	return &Config{
		URL:            os.Getenv(prefix + "BASE_URL"),
		MandantId:      os.Getenv(prefix + "MANDANT_ID"),
		ClientSystemId: os.Getenv(prefix + "CLIENT_SYSTEM_ID"),
		WorkplaceId:    os.Getenv(prefix + "WORKPLACE_ID"),
		UserId:         os.Getenv(prefix + "USER_ID"),
		Credentials:    creds,
	}, nil
}

func (c *Client) CreateServiceProxy(serviceName ServiceName, version string) (*serviceProxy, error) {
	var service *Service
	var serviceVersion *ServiceVersion
	for _, s := range c.Services.ServiceInformation.Service {
		if s.Name == serviceName {
			for _, v := range s.Versions {
				if v.Version == version {
					service = &s
					serviceVersion = &v
					break
				}
			}
		}
	}
	if serviceVersion == nil {
		return nil, fmt.Errorf("service version not found: %s %s", serviceName, version)
	}
	slog.Debug("creating service proxy", "service", serviceName, "version", version, "endpointTLS", serviceVersion.EndpointTLS.Location)
	return &serviceProxy{
		endpoint:       serviceVersion.EndpointTLS.Location,
		client:         c,
		service:        service,
		serviceVersion: serviceVersion,
	}, nil
}

// converts a sem version string to a number for sorting and comparing
// errors are ignored, if the version string is not a valid semver string, 0 is returned
func semverAsNumber(version string) int {
	// parse version string, get major, minor and patch
	// convert to number by multiplying major by 10000, minor by 100 and adding patch
	// e.g. 1.2.3 -> 10000*1 + 100*2 + 3 = 10203

	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)$`)
	matches := re.FindStringSubmatch(version)
	if len(matches) != 4 {
		return 0
	}

	toInt := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}

	major := 10000 * toInt(matches[1])
	minor := 100 * toInt(matches[2])
	patch := toInt(matches[3])

	return major + minor + patch
}
