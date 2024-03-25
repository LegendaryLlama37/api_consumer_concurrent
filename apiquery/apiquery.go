// apiquery.go
package apiquery

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
  "strings"
)

// Credentials represents authentication credentials for API requests.
type Credentials struct {
	APIKey string
}

// fetchData fetches data from the given URL using HTTP GET request with provided API key.
func fetchData(url string, credentials Credentials, wg *sync.WaitGroup) (interface{}, error) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if credentials.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+credentials.APIKey)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// FetchDataConcurrently fetches data from multiple URLs concurrently using the provided API key.
func FetchDataConcurrently(urlAPIKeyMap map[string]string) ([]interface{}, error) {
	var wg sync.WaitGroup
	var results []interface{}

	// Increment the wait group counter
	wg.Add(len(urlAPIKeyMap))

	// Fetch data for each URL concurrently
	for url, apiKey := range urlAPIKeyMap {
		go func(url string, apiKey string) {
			defer wg.Done()
			credentials := Credentials{APIKey: apiKey}
			data, err := fetchData(url, credentials, &wg)
			if err != nil {
				fmt.Println(err) // Handle error
				return
			}
			results = append(results, data)
		}(url, apiKey)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return results, nil
}

// ParseFlags parses command-line flags to extract URLs and API keys.
func ParseFlags() map[string]string {
	flag.Parse()
	args := flag.Args()
	urlAPIKeyMap := make(map[string]string)
	for _, arg := range args {
		parts := strings.Split(arg, ":")
		if len(parts) != 2 {
			fmt.Println("Each argument should be in the form: url:api_key")
			continue
		}
		urlAPIKeyMap[parts[0]] = parts[1]
	}
	return urlAPIKeyMap
}


