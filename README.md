# docker-secrets

[![GoDoc](https://godoc.org/github.com/ijustfool/docker-secrets?status.png)](http://godoc.org/github.com/ijustfool/docker-secrets)
[![Go Report](https://goreportcard.com/badge/github.com/ijustfool/docker-secrets)](https://goreportcard.com/report/github.com/ijustfool/docker-secrets)

## Requirements

Go 1.2 or above.

## Installation

Run the following command to install the package:

```
go get github.com/ijustfool/docker-secrets
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/ijustfool/docker-secrets"
)

func main() {
    dockerSecrets, _ := secrets.NewDockerSecrets("")

    // Get all secrets
    fmt.Println(dockerSecrets.GetAll())
    // print: map[secret_1:val_1 secret_2:val_2]

    // Get a single secret
    secret, _ := dockerSecrets.Get("secret_1")
    fmt.Println(secret)
    // print: val_1

    // Custom location
    dsCustomLoc, _ := secrets.NewDockerSecrets("/run/myCustomLocation")
}
```