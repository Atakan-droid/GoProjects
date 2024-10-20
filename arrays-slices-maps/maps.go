package main

import "fmt"

// Maps
func maps() {
	webSites := map[string]string{"Google": "www.google.com", "AWS": "www.aws.com"}
	fmt.Println(webSites)

	fmt.Println(webSites["Google"])

	// Add new key value
	webSites["Azure"] = "www.azure.com"
	fmt.Println(webSites)

	// Delete key value
	delete(webSites, "AWS")
	fmt.Println(webSites)
}
