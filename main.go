package main

/*
 * main.go
 *
 * Command line utility to retrieve BTC conversion rates for various currencies.
 * Usage: go run main.go <symbol> <symbol> ...
 * Ex:    go run main.go USD EUR GBP
 *
 */

// Import libraries
import (
	"fmt" // Used to print & format strings
	"os" // Gets command line arguments
	"net/http" // Makes HTTP requests
	"io/ioutil" // Parses HTTP response
	"encoding/json" // Parses HTTP response as JSON data
)


// Main functions
func main() {
	// Get command line args without the program name
	args_without_exe_name := os.Args[1:]

	var currency_symbols = args_without_exe_name
	var urls []string

	// Iterate over symbols. Add url to array of URLs for each currency symbol
	for _, symbol := range currency_symbols {
		symbol_url := fmt.Sprintf("http://api.coindesk.com/v1/bpi/currentprice/%s.jsons", symbol)
		urls = append(urls, symbol_url)
	}

	// For each URL, make request to Coindesk
	for index, url := range urls {
		res, _ := http.Get(url) // Make GET request (ignoring errors)
		body, _ := ioutil.ReadAll(res.Body) // Read response (ignoring errors)
		var json_data map[string]interface{} // Set up variable to hold JSON data
		json.Unmarshal(body, &json_data) // Parse response into JSON data

		// Get current currency symbol, get array of currency info from json data
		currency_symbol := currency_symbols[index]
		currency_info_array, ok := json_data["bpi"].(map[string]interface{})

		// If the data didn't exist, we must have an invalid currency symbol
		if !ok {
			fmt.Printf("%s: invalid\n", currency_symbol)
			continue
		}

		// Finally, get the currency info for the current symbol and print
		currency_info := currency_info_array[currency_symbol].(map[string]interface{})
		fmt.Printf("%s: %f\n", currency_symbol, currency_info["rate_float"])
	}
}