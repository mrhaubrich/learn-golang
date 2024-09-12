package main

import (
	"fmt"
	"log"

	"github.com/mrhaubrich/learn-golang/system"
)

func main() {
	// Instantiate a WindowsFetcher (could be switched with other fetchers)
	fetcher, err := system.GetFetcher()
	if err != nil {
		log.Fatalf("Error getting fetcher: %v", err)
	}

	// Fetch system data using the fetcher and print the result
	systemData := fetcher.FetchSystemData()
	fmt.Printf("Logged User: %s\n", systemData.LoggedUser)
}
