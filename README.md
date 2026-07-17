# sj-se-go

a simple go library to interact with sj.se's undocumented web booking api. \
so far, only search & offers (getting price of a search result) are implemented. \
this might be particularly useful to build a price/availability tracker :)

## installation

```bash
go get github.com/espcaa/sj-se-go
```

## usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/espcaa/sj-se-go"
)

var baseUrl = "https://prod-api.adp.sj.se/public/sales/booking/v3"
var subscriptionKey = "d6625619def348d38be070027fd24ff6"

func main() {
	client := sj.NewClient(baseUrl, subscriptionKey)

	// context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get the config to get train stations ids
	config, err := client.GetConfig(ctx)
	if err != nil {
		fmt.Printf("Error getting config: %v\n", err)
		return
	}

  for _, station := range config.Stations {
    fmt.Printf("Station: %s, ID: %s\n", station.Name, station.Id)
  }
}
```

_simple example of how to get train stations ids for search_ \
_you can find more examples in the `examples` folder_
