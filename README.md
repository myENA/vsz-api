# VSZ API Package

[![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/myENA/vsz-api)

Helps you ruckus vsz-ize stuff.

# Version Map

`master` will always be the most up-to-date version of our generated client.  Differing VSZ versions always have subtle
(and sometimes not-so-subtle) differences, so please have a look at the various branches and pick the one that matches
your VSZ version.  We currently will not guarantee any degree of backwards compatibility between versions.

|VSZ Version|Package Version|
|---|---|
|3.5.0|0.1-0.2|

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
	
	// nil as a 2nd argument will use a default non-pooled http client
	client := api.NewClient(conf, nil)
	
	userCtx := api.NewUserContext(context.Background(), "username", "password", 2 * time.Second)
	defer userCtx.Cancel()
    
	httpResponse, responseData, err := client.Session().LoginSessionRetrieveGet(userCtx)
	
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
- Tests... 