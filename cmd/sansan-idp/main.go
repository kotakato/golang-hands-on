package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kotakato/golang-hands-on/sansan-openapi/idp"
)

const usageTemplate = `Usage:

%s [OPTIONS] (users|departments|check)

Options:

`

func main() {
	var (
		endpoint  string
		apiKey    string
		isSpace   bool
		delimiter string
	)

	flag.StringVar(&endpoint, "endpoint", idp.DefaultEndpoint, "Endpoint of Sansan OpenAPI")
	flag.StringVar(&apiKey, "api-key", "", "API key of Sansan OpenAPI")
	flag.BoolVar(&isSpace, "space", false, "Use space as a delimiter")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageTemplate, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "-api-key is required\n")
		showUsageAndExit()
	}

	switch {
	case isSpace:
		delimiter = "space"
	default:
		delimiter = "comma"
	}

	client := idp.NewClient(endpoint, apiKey)

	switch flag.Arg(0) {
	case "users":
		executeUsers(client, delimiter)
	default:
		showUsageAndExit()
	}
}

func showUsageAndExit() {
	flag.Usage()
	os.Exit(1)
}

func executeUsers(client *idp.Client, delimiter string) {
	result, err := client.GetUsers(delimiter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}
