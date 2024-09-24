# Introduction

**This library allows you to quickly and easily use the Finout API via Go.**

This library provides full support for all Finout [API](https://docs.finout.io/en/collections/166488-api) endpoints.

# Installation

## Install Package

`go get github.com/ManoManoTech/finout`

## Dependencies

- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen), this library is basically the generated code from oapi-codent for the forged Finout API spec.

# Quick Start

## Hello Finout

The following is the minimum needed code to call the Finout API.

```go
package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/ManoManoTech/finout"
)

func main() {
    client, err := finout.NewSecuredClient("YOUR_CLIENT_ID", "YOUR_SECRET_KEY")
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    resp, err := client.GetView(ctx)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        bodyString := string(bodyBytes)
        log.Println(bodyString)
    } else {
        log.Fatalf("Status code: %d", resp.StatusCode)
    }
}
```

## Regenerate the client

If you need to regenerate the client, you can use the following command:

```bash
go generate
```

# Documentation

If you need to check the spec, you can use the [Swagger Editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/ManoManoTech/finout/refs/heads/main/finout.yaml) with the spec in `finout.yaml`.
