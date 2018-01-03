# VSZ API Package

Helps you ruckus vsz-ize stuff.

# Version Map

`master` will always be the most up-to-date version of our generated client.  Differing VSZ versions always have subtle
(and sometimes not-so-subtle) differences, so please have a look at the various branches and pick the one that matches
your VSZ version.  We currently will not guarantee any degree of backwards compatibility between versions.

## Basic Usage

```go
package main

import (
	"github.com/myENA/vsz-api"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	httpClient := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	conf := api.DefaultConfig("yourhost.whatever", "username", "password")
	client := api.NewClient(conf, httpClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	httpResponse, responseData, err := client.Session().LoginSessionRetrieveGet(ctx)
	
	cancel()
	 
	if err != nil {
		log.Printf("Unable to retrieve current login session: %s", err)
		os.Exit(1)
	}

	if httpResponse.StatusCode != 200 {
		log.Printf("Received non-200 response. Code: %d; Status: %s", httpResponse.StatusCode, httpResponse.Status)
		os.Exit(1)
	}

	encoded, err := json.MarshalIndent(responseData, "", "\t")
	if err != nil {
		log.Printf("Unable to turn response into JSON: %s", err)
		os.Exit(1)
	}

	log.Printf("Response:\n%s", encoded)
	os.Exit(0)
}
```

## Setting a Session Cookie Refresher

If a 401 unauthorized response is seen the client will attempt to refresh it's session using a 
[SessionCookieRefreshFunc](./client.go#L28).  There is a default implementation of this that simply attempts to re-use
the password sent in at init.

To specify your own set [CookieRefresher](./client.go#L70) in your [Config](./client.go#L60).

## Setting a Session Password Refresher

In the event that a 401 is seen even AFTER a session refresh is attempted, the client will attempt to use a
[SessionPasswordRefreshFunc](./client.go#L54).  This MUST return a now-valid password or error.  If the returned
password STILL results in a 401, the execution chain will exit with an error.


### Future Development
Some things we'd like to have in the future:

- Request struct field value validation
- Better error typing
- Better doc comments on api's
- Tests... 