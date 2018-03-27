# VSZ API Package

[![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/myENA/vsz-api)

Helps you ruckus vsz-ize stuff.

# Version Map

`master` will always be the most up-to-date version of our generated client.  Differing VSZ versions always have subtle
(and sometimes not-so-subtle) differences, so please have a look at the various branches and pick the one that matches
your VSZ version.  We currently will not guarantee any degree of backwards compatibility between versions.

|VSZ Version|Package Version|
|---|---|
|3.5.0-3.5.1|0.5.*|

## Basic Usage

```go
package main

import (
	"github.com/myENA/vsz-api"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {
	conf := api.DefaultConfig("yourhost.whatever")
	
	// create authenticator
	auth := api.NewPasswordAuthenticator("username", "password", 30 * time.Minute)
	
	// nil as a 3rd argument will use a default non-pooled http client
	client, err := api.NewClient(conf, auth, nil)
	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}
	
	// create context for this request
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	
	// execute
	httpResponse, responseData, err := client.Session().LoginSessionRetrieveGet(ctx)
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

### Future Development
Some things we'd like to have in the future:

- Request struct field value validation
- Better error typing
- Better doc comments on api's
- Cleanup of user session on shutdown
- Tests... 