# koap 🧼 - Konnektor SOAP

Convert legacy SOAP/WSDL services into modern, language-agnostic OpenAPI definitions.

Go library and CLI tool for interacting with the German Telematik Infrastructure (TI), specifically the Gematik Konnektor SOAP web services.

![CLI demo](images/cli.gif)

## Installation

```bash
go install github.com/spilikin/koap-go/cmd/ti@latest
```

## CLI Usage

### Configuration

The CLI uses `.kon` files (JSON) to configure Konnektor connections. Specify one with:

- `-k` / `--kon` flag
- `DOTKON_FILE` environment variable
- Falls back to `default.kon`

File resolution order: exact path, `<name>.kon` in current directory, then `$XDG_CONFIG_HOME/telematik/kon/`. Tilde (`~`) expansion is supported.

Example `.kon` file:

```json
{
  "version": "1.0.0",
  "url": "https://konnektor.example.com:8443",
  "mandantId": "M1",
  "workplaceId": "W1",
  "clientSystemId": "C1",
  "credentials": {
    "type": "basic",
    "username": "user",
    "password": "${MY_SECRET}"
  },
  "env": "ru",
  "insecureSkipVerify": true,
  "rewriteServiceEndpoints": true
}
```

Environment variables can be used with `${VAR_NAME}` syntax.

### Credential types

**Basic authentication:**
```json
{ "type": "basic", "username": "user", "password": "secret" }
```

**PKCS#12 client certificate (mTLS):**
```json
{ "type": "pkcs12", "data": "<base64>", "password": "secret" }
```

**macOS Keychain:**
```json
{ "type": "system", "name": "my-konnektor" }
```

### Commands

All Konnektor commands follow the pattern:

```
ti kon -k <config-name> <verb> <resource> [flags]
```

**Konnektor resources:**

```
ti kon -k <name> get info                              # Product information
ti kon -k <name> get services                          # List services (table)
ti kon -k <name> get services --raw                    # Raw service directory XML
ti kon -k <name> get cards                             # List inserted cards
ti kon -k <name> get certificates <card-handle>        # List certificates of a card
ti kon -k <name> describe certificate <card-handle> <cert-ref>  # Certificate details
```

Cert refs: `C.AUT`, `C.ENC`, `C.SIG`, `C.QES`

**PKCS#12 tools:**

```
ti pkcs12 inspect <file>              # Show PKCS#12 contents
ti pkcs12 convert <input> <output>    # Convert legacy BER to modern DER
ti pkcs12 encode <file>               # Output as .kon credentials JSON
```

**Global flags:**

- `-v` / `--verbose` - Enable debug logging (full HTTP request/response dumps)

**Output format:** Use `-o json` or `-o table` (default) on `get` commands. PKCS#12 password is taken from `--password` flag, `PKCS12_PASSWORD` env var, or interactive prompt. JSON and XML output is syntax-highlighted in terminals.

## Library Usage

```go
import koap "github.com/spilikin/koap-go"

// Parse .kon configuration
config, err := koap.ParseDotkon(data)

// Create a client (loads SDS automatically)
client, err := koap.NewClient(config)

// Get inserted cards
cards, err := client.GetCards()

// Read certificates from a card
certs, err := client.ReadCardCertificates(cardHandle, koap.CertRefAUT, koap.CertRefENC)

// Read all certificates (C.AUT)
certs, err := client.ReadAllCardCertificates(cardHandle)

// Access admission info (registration number, profession)
for _, c := range certs {
    if c.Admission != nil {
        fmt.Println(c.Admission.RegistrationNumber)
    }
}

// Or just get an HTTP client for manual requests
httpClient, baseURL, err := koap.NewHTTPClient(config)
```

## Features

- `.kon` configuration format with validation and environment variable expansion
- Three credential types: Basic, PKCS#12, macOS Keychain
- PKCS#12 support with legacy BER-to-DER auto-conversion
- Brainpool elliptic curve support for certificate parsing
- Admission statement extraction (registration number, profession OIDs)
- Service directory (SDS) loading with optional endpoint rewriting
- Typed Go bindings for Konnektor SOAP services (EventService, CertificateService)
- Full HTTP request/response debug logging with `-v`
- Syntax-highlighted XML and JSON output in terminals

## License

See [LICENSE](LICENSE) file.
