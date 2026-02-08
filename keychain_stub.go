//go:build !darwin

package koap

import (
	"crypto/tls"
	"fmt"
)

// LoadSystemCredential is not supported on this platform.
func LoadSystemCredential(name string) (tls.Certificate, error) {
	return tls.Certificate{}, fmt.Errorf("system credentials are not supported on this platform")
}
