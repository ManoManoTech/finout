# Introduction

**This package allows you to quickly and easily use the Finout API via Go.**

This package provides full support for all Finout [API](https://docs.finout.io/en/collections/166488-api) endpoints.

# Installation

## Install Package

`go get github.com/ManoManoTech/finout`

## Dependencies

- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen), this package is basically the generated code from oapi-codegen for the forged Finout API spec.

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
    client, err := finout.NewSecuredClientWithResponses("YOUR_CLIENT_ID", "YOUR_SECRET_KEY")
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.GetViewWithResponse(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode() == http.StatusOK && resp.JSON200 != nil && resp.JSON200.Data != nil {
        for _, view := range *resp.JSON200.Data {
            log.Println(*view.Id, *view.Name)
        }
    } else {
        log.Fatalf("Status code: %d", resp.StatusCode())
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
