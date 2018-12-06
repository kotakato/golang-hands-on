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
		endpoint  string
		apiKey    string
		usersPath string
	)

	flag.StringVar(&endpoint, "endpoint", idp.DefaultEndpoint, "Endpoint of Sansan OpenAPI")
	flag.StringVar(&apiKey, "api-key", "", "API key of Sansan OpenAPI")
	flag.StringVar(&usersPath, "users", "", "Path to users CSV file to import")
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
	case "departments":
		executeDepartments(client)
	case "check":
		executeCheck(client, usersPath)
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

func executeDepartments(client *idp.Client) {
	result, err := client.GetDepartments()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func executeCheck(client *idp.Client, usersPath string) {
	if usersPath == "" {
		fmt.Fprintf(os.Stderr, "-users is required.\n")
		showUsageAndExit()
	}

	usersFile, err := os.Open(usersPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file '%s': %s\n", usersPath, err)
		os.Exit(1)
	}
	defer usersFile.Close()

	result, err := client.Check(usersFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}
