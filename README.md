# pixelsparser (Go)


![GitHub issues](https://img.shields.io/github/issues/ptinosq/pixelsparser-go)
![GitHub last commit](https://img.shields.io/github/last-commit/ptinosq/pixelsparser-go)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/pTinosq/pixelsparser-go/go.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pTinosq/pixelsparser-go)
![Codecov](https://img.shields.io/codecov/c/github/pTinosq/pixelsparser-go)

A lightweight Go library for importing and handling Pixels JSON data. Offers a simple, type-safe API to parse and manipulate data efficiently, with no external dependencies.

## Installation

Install the library using Go modules:
```bash
go get github.com/ptinosq/pixelsparser-go
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/ptinosq/pixelsparser-go"
)

func main() {
	// Load Pixels JSON data from a file
	pixels, err := pixelsparser.Load("path/to/data.json")
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		return
	}

	// Access the mood of the first pixel
	fmt.Println(pixels[0].Mood)

	// Access the notes of the second pixel
	fmt.Println(pixels[1].Notes)
}
```

## Documentation

View full documentation on [pkg.go.dev](https://pkg.go.dev/github.com/pTinosq/pixelsparser-go)

## Credits

This library was created to parse data from the Pixels app by Teo Vogel, available on the [Google Play Store](https://play.google.com/store/apps/details?id=ar.teovogel.yip) and on the [App Store](https://apps.apple.com/sg/app/pixels-mental-health-and-mood/id1481910141)
