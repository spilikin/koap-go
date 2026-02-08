package koap

import (
	"crypto/tls"
	"fmt"

	"github.com/keybase/go-keychain"
)

const keychainService = "dotkon"

// LoadSystemCredential loads a TLS client certificate from the macOS Keychain.
// The credential is looked up by name using the service "dotkon".
// The stored data is expected to be a PKCS#12 container.
func LoadSystemCredential(name string) (tls.Certificate, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(keychainService)
	query.SetAccount(name)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("querying keychain: %w", err)
	}
	if len(results) == 0 {
		return tls.Certificate{}, fmt.Errorf("credential %q not found in keychain", name)
	}

	return parsePKCS12Data(results[0].Data, "")
}
