# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**koap-go** is a Go library for communicating with the German Konnektor (healthcare connector) via SOAP web services. The Konnektor is part of the TI (Telematikinfrastruktur) and provides secure access to medical cards, digital signatures, certificates, and encryption services.

Module: `github.com/spilikin/koap-go`
Go version: 1.25

## Commands

### Tests
```bash
go test ./...                          # Run all tests
go test -run TestSerialize ./...       # Run a single test
go test -v ./...                       # Verbose output
```

Note: Some tests (e.g., `TestSOAP`, `TestParseConfigFile`) require a live Konnektor connection or external config files and will fail without them.

### Code Generation
```bash
just generate-kon
```
Runs `wsdl2openapi2go` to regenerate the `api/` directory from `Konnektor-OPB6.json` using `naming-kon.json` for package naming. The generator tool (`wsdl2openapi2go`) lives in a sibling directory (`../wsdl2openapi/`).

### Dependency Management
```bash
go mod tidy
```

## Architecture

### Core Library (root package `koap`)

- **`client.go`** — `Client` struct: entry point for Konnektor communication. Creates HTTP client with TLS/auth, discovers services via SDS, and creates service proxies.
- **`config.go`** — `Config` struct with polymorphic credentials (basic auth or mTLS certificate). Loaded from JSON (`ParseConfigJSON`) or environment variables (`ReadConfigFromEnv` with `KONNEKTOR_` prefix).
- **`sds.go`** — Service Directory Service: XML parsing of `connector.sds` to discover available services, versions, and endpoints at runtime.
- **`soap.go`** — SOAP protocol primitives: `Envelope`, `Body`, `Fault`, and the `Operation` interface (Name, SOAPAction, InputType, OutputType, FaultType).
- **`service_proxy.go`** — `serviceProxy` handles SOAP call execution: XML marshaling, envelope wrapping, HTTP POST with SOAPAction header, response parsing.
- **`bootstrap.go`** — TLS certificate utilities (load trust stores, fetch server certs, save PEM files).

### Generated API (`api/`)

All files under `api/` are **generated** — do not edit manually. They are organized by namespace:

- `api/gematik/conn/` — Konnektor services (CardService, EventService, SignatureService, CertificateService, EncryptionService, AuthSignatureService, CardTerminalService) with version-specific packages (e.g., `cardservice81/`, `cardservice812/`)
- `api/gematik/tel/` — Error types (`error20/`)
- `api/oasis/` — OASIS standards (SAML assertions, DSS)
- `api/etsi*/`, `api/w3*/` — ETSI and W3C XML standards

Each generated package contains:
- `types.go` — XML-marshallable structs with namespace-qualified tags
- `soap.go` — SOAP operation definitions and envelope wrappers

### Key Design Patterns

1. **Service Discovery**: The client fetches `connector.sds` at startup to discover service endpoints dynamically — URLs are never hardcoded.
2. **Version-specific packages**: Multiple versions of the same service coexist (e.g., CardService 8.1, 8.1.1, 8.1.2), selected at runtime via `CreateServiceProxy(serviceName, version)`.
3. **XML namespace fidelity**: All generated types carry full namespace URIs in their `xml.Name` fields for correct SOAP/XML serialization.
4. **Polymorphic config**: `Config.Credentials` unmarshals to different types based on a `"type"` discriminator field (`"basic"` or `"der"`).

### Usage Flow

```go
config, _ := koap.ReadConfigFromEnv("")        // or ParseConfigJSON()
client, _ := koap.NewClient(config)             // sets up TLS, discovers services
proxy, _ := client.CreateServiceProxy(koap.ServiceNameEventService, "7.2.0")
result, _ := proxy.CallWithDocument(operation, &request)
```
