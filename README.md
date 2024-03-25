# API Query Library

The `apiquery` library is a Go package designed to simplify the process of querying multiple RESTful APIs concurrently. It provides functionality to fetch data from multiple endpoints concurrently while allowing customization of API keys for each endpoint.

## Features

- Fetch data from multiple URLs concurrently.
- Supports providing custom API keys for each endpoint.
- Simple and easy-to-use API.
- Handles HTTP GET requests and JSON decoding.

## Usage

### Importing the Library

```go
import "apiquery"

# Fetching Data

The FetchDataConcurrently function fetches data from multiple URLs concurrently using the provided API keys.
```go
results, err := apiquery.FetchDataConcurrently(urlAPIKeyMap)
if err != nil {
    // Handle error
    return
}

// Process the results
for i, result := range results {
    fmt.Printf("Result %d: %+v\n", i+1, result)
}

Command-line Usage

The library also supports command-line usage for providing URLs and API keys as arguments.
go run main.go -urlapikey="https://api.example.com:API_KEY"

##Example
package main

import (
    "apiquery"
    "fmt"
)

func main() {
    // Parse command-line flags
    urlAPIKeyMap := apiquery.ParseFlags()
    if len(urlAPIKeyMap) == 0 {
        fmt.Println("Please provide URLs and API keys as arguments.")
        return
    }

    // Fetch data concurrently
    results, err := apiquery.FetchDataConcurrently(urlAPIKeyMap)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Process the results
    for i, result := range results {
        fmt.Printf("Result %d: %+v\n", i+1, result)
    }
}

License

This library is open-source and distributed under the MIT License. See the LICENSE file for details.

