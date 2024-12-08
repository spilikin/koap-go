package koap_test

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"testing"

	"github.com/spilikin/koap-go"
)

func init() {
	// set slog level to debug
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestLoadConnectorServices(t *testing.T) {
	httpClient := &http.Client{}
	ctx := context.TODO()
	_, err := koap.LoadConnectorServices(ctx, httpClient, &url.URL{Scheme: "https", Host: "tig.spilikin.dev"})
	if err != nil {
		t.Errorf("error loading service directory: %v", err)
	}

}
