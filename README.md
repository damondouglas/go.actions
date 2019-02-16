[![godoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/damondouglas/go.actions)

# About

go.actions encodes and decodes dialogflow fulfillment requests and responses sent through Google Actions.

# Install

`go get github.com:damondouglas/go.actions`

# Usage

```golang
import (
    "github.com/damondouglas/go.actions/v2/dialogflow"
    "encoding/json"
)

func HandleAction(w http.ResponseWriter, r *http.Request) {
    var req *dialogflow.Request
	if err := json.NewDecoder(r).Decode(&req); err != nil {
        // Handle err
    }

    // Do something with req
}
```