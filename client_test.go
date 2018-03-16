package api_test

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/myENA/vsz-api"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	addr     string
	username string
	password string
)

func init() {
	flag.StringVar(&addr, "addr", "", "Address of VSZ to hit")
	flag.StringVar(&username, "username", "", "Username of user to use for tests")
	flag.StringVar(&password, "password", "", "Password of user to use for tests")
	flag.Parse()

	if addr == "" {
		log.Println("addr cannot be empty")
		os.Exit(1)
	}
	if username == "" {
		log.Println("username cannot be empty")
		os.Exit(1)
	}
	if password == "" {
		log.Println("password cannot be empty")
		os.Exit(1)
	}

	api.EnableDebug()
}

func testClient(t *testing.T) *api.Client {
	client, err := api.NewClient(
		api.DefaultConfig(addr),
		api.NewPasswordAuthenticator(username, password, 30*time.Minute, 2*time.Second),
		&http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true,
				MaxIdleConnsPerHost:   -1,
			},
		})
	if err != nil {
		t.Logf("Error creating client: %s", err)
		t.FailNow()
	}
	return client
}

func TestClient(t *testing.T) {
	var ctx context.Context
	var cancel context.CancelFunc

	client := testClient(t)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	_, _, err := client.RuckusZones().RuckusWirelessApZoneRetrieveListGet(ctx, nil)
	cancel()
	if err != nil {
		t.Logf("Error: %s", err)
		t.FailNow()
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	_, _, err = client.Session().LoginSessionLogoffDelete(ctx)
	cancel()
	if err != nil {
		t.Logf("Error: %s", err)
		t.FailNow()
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	_, errBody, err := client.Session().LoginSessionRetrieveGet(ctx)
	cancel()
	if err == nil {
		t.Logf("Expected error, saw success with %+v", errBody)
		t.FailNow()
	}
}
