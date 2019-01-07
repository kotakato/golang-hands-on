package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kotakato/golang-hands-on/lesson1/sansan-openapi/idp"
)

const usageTemplate = `Usage:

%s [OPTIONS] (users|departments|check)

Options:

`

func main() {
	var (
		endpoint string
		apiKey   string
	)

	flag.StringVar(&endpoint, "endpoint", idp.DefaultEndpoint, "Endpoint of Sansan OpenAPI")
	flag.StringVar(&apiKey, "api-key", "", "API key of Sansan OpenAPI")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageTemplate, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "-api-key is required\n")
		showUsageAndExit()
	}

	client := idp.NewClient(endpoint, apiKey)

	switch flag.Arg(0) {
	case "users":
		executeUsers(client)
	default:
		showUsageAndExit()
	}
}

func showUsageAndExit() {
	flag.Usage()
	os.Exit(1)
}

func executeUsers(client *idp.Client) {
	result, err := client.GetUsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}
