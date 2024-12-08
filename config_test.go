package koap_test

import (
	"os"
	"testing"

	"github.com/spilikin/koap-go"
)

func TestParseConfigJSON(t *testing.T) {
	tests := []struct {
		name          string
		jsonData      string
		expectedError bool
	}{
		{
			name: "Valid Basic Credentials",
			jsonData: `{
				"url": "https://konnektor.example.com:8443",
				"mandantId": "M1",
				"workplaceId": "W1",
				"clientSystemId": "C1",
				"userId": "U1",
				"credentials": {
					"type": "basic",
					"username": "test_user",
					"password": "test_pass"
				}
			}`,
			expectedError: false,
		},
		{
			name: "Valid Cert-DER Credentials",
			jsonData: `{
				"url": "https://konnektor.example.com:8443",
				"mandantId": "M1",
				"workplaceId": "W1",
				"clientSystemId": "C1",
				"userId": "U1",
				"credentials": {
					"type": "der",
					"cert": "Y2VydGlmaWNhdGU=", 
					"key": "a2V5"
				}
			}`,
			expectedError: false,
		},
		{
			name: "Unsupported Credentials Type",
			jsonData: `{
				"url": "https://konnektor.example.com:8443",
				"mandantId": "M1",
				"workplaceId": "W1",
				"clientSystemId": "C1",
				"userId": "U1",
				"credentials": {
					"type": "unsupported",
					"someField": "value"
				}
			}`,
			expectedError: true,
		},
		{
			name: "Missing URL",
			jsonData: `{
				"mandantId": "M1",
				"workplaceId": "W1",
				"clientSystemId": "C1",
				"userId": "U1",
				"credentials": {
					"type": "basic",
					"username": "test_user",
					"password": "test_pass"
				}
			}`,
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := koap.ParseConfigJSON([]byte(test.jsonData))
			if (err != nil) != test.expectedError {
				t.Errorf("expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

func TestParseConfigFile(t *testing.T) {
	filename := "/tmp/konnektor.kon"
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	cfg, err := koap.ParseConfigJSON(data)
	if err != nil {
		t.Errorf("error parsing config: %v", err)
	}

	if len(cfg.TrustedCertificatesRaw) == 0 {
		t.Errorf("no trusted certificates found")
	}

	client, err := koap.NewClient(cfg)
	if err != nil {
		t.Errorf("error creating client: %v", err)
	}

	t.Logf("client is: %+v", client)
}
