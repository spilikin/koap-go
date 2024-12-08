package koap

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
)

// CredentialsConfig is an interface for different types of credentials
type CredentialsConfig interface {
}

// CredentialsConfigBasic is a struct for HTTP basic credentials
type CredentialsConfigBasic struct {
	Type     string `json:"type" yaml:"type"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

// CredentialsConfigCertDER is a struct for DER certificate credentials
type CredentialsCertDER struct {
	Type        string          `json:"type" yaml:"type"`
	CertRaw     []byte          `json:"cert" yaml:"cert"`
	KeyRaw      []byte          `json:"key" yaml:"key"`
	Certificate tls.Certificate `json:"-" yaml:"-"`
}

type Config struct {
	Version                string              `json:"version" yaml:"version" validate:"required"`
	Name                   string              `json:"name,omitempty" yaml:"name,omitempty"`
	URL                    string              `json:"url" yaml:"url" validate:"required"`
	MandantId              string              `json:"mandantId" yaml:"mandantId" validate:"required"`
	WorkplaceId            string              `json:"workplaceId" yaml:"workplaceId" validate:"required"`
	ClientSystemId         string              `json:"clientSystemId" yaml:"clientSystemId" validate:"required"`
	UserId                 string              `json:"userId,omitempty" yaml:"userId,omitempty"`
	RandomUserId           bool                `json:"randomUserId" yaml:"randomUserId"`
	Credentials            CredentialsConfig   `json:"credentials" yaml:"credentials"`
	Env                    string              `json:"env,omitempty" yaml:"env,omitempty" validate:"oneof=ru tu pu"`
	TrustedSAN             string              `json:"trustedSAN,omitempty" yaml:"trustedSAN,omitempty"`
	TrustedCertificatesRaw [][]byte            `json:"trustedCertificates,omitempty" yaml:"trustedCertificates,omitempty"`
	TrustedCertificates    []*x509.Certificate `json:"-" yaml:"-"`
}

// Intermediate struct for determining type during unmarshalling
type rawCredentials struct {
	Type string `json:"type"`
}

// Custom unmarshalling logic for credentials
func (c *Config) UnmarshalJSON(data []byte) error {
	type Alias Config // Alias to avoid infinite recursion
	var temp struct {
		Alias
		Credentials json.RawMessage `json:"credentials"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*c = Config(temp.Alias)

	var rawCred rawCredentials
	if err := json.Unmarshal(temp.Credentials, &rawCred); err != nil {
		return err
	}

	switch rawCred.Type {
	case "basic":
		var basicCred CredentialsConfigBasic
		if err := json.Unmarshal(temp.Credentials, &basicCred); err != nil {
			return err
		}
		c.Credentials = basicCred
	case "der":
		var certCred CredentialsCertDER
		if err := json.Unmarshal(temp.Credentials, &certCred); err != nil {
			return err
		}
		c.Credentials = certCred
	default:
		return errors.New("unsupported credential type")
	}

	// parse trusted certificates
	c.TrustedCertificates = make([]*x509.Certificate, 0, len(c.TrustedCertificatesRaw))
	for _, cert := range c.TrustedCertificatesRaw {
		parsedCert, err := x509.ParseCertificate(cert)
		if err != nil {
			return fmt.Errorf("parsing certificate: %w", err)
		}
		c.TrustedCertificates = append(c.TrustedCertificates, parsedCert)
	}

	return nil
}

func ParseConfigJSON(data []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling config: %w", err)
	}
	return &config, nil
}

func (d *CredentialsCertDER) UnmarshalJSON(data []byte) error {
	type Alias CredentialsCertDER
	var surrogate = new(Alias)
	if err := json.Unmarshal(data, surrogate); err != nil {
		return err
	}

	cert, err := x509.ParseCertificate(surrogate.CertRaw)
	if err != nil {
		return err
	}

	// try parsing ECDSA private key first
	var key crypto.PrivateKey
	key, err = x509.ParseECPrivateKey(surrogate.KeyRaw)
	if err != nil {
		// try parsing PKCS1 private key
		key, err = x509.ParsePKCS1PrivateKey(surrogate.KeyRaw)
		if err != nil {
			// try parsing PKCS8 private key
			key, err = x509.ParsePKCS8PrivateKey(surrogate.KeyRaw)
			if err != nil {
				return err
			}
		}
	}

	d.Certificate = tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
		Leaf:        cert,
	}

	return nil

}
