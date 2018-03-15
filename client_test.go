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

func getTestPair() (*api.UserContext, *api.Client) {
	return api.NewUserContext(context.Background(), username, password, api.DefaultRequestTimeout),
		api.NewClient(api.DefaultConfig(addr),
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

}

func TestClient(t *testing.T) {
	ctx, client := getTestPair()

	_, _, err := client.RuckusZones().RuckusWirelessApZoneRetrieveListGet(ctx, nil)
	if err != nil {
		t.Logf("Error: %s", err)
		t.FailNow()
	}

}
