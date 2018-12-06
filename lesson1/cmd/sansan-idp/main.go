package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kotakato/golang-hands-on/lesson1/sansan-openapi/idp"
)

const usageTemplate = `Usage:

%s [OPTIONS] (users|departments|check)

Options:

`

func main() {
	var (
		endpoint        string
		apiKey          string
		usersPath       string
		departmentsPath string
	)

	flag.StringVar(&endpoint, "endpoint", idp.DefaultEndpoint, "Endpoint of Sansan OpenAPI")
	flag.StringVar(&apiKey, "api-key", "", "API key of Sansan OpenAPI")
	flag.StringVar(&usersPath, "users", "", "Path to users CSV file to import")
	flag.StringVar(&departmentsPath, "departments", "", "Path to departments CSV file to import")
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
		executeCheck(client, usersPath, departmentsPath)
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

func executeCheck(client *idp.Client, usersPath, departmentsPath string) {
	if usersPath == "" && departmentsPath == "" {
		fmt.Fprintf(os.Stderr, "Either -users or -departments is required.\n")
		showUsageAndExit()
	}

	files := make(map[string]io.Reader)
	if usersPath != "" {
		usersFile, err := os.Open(usersPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file '%s': %s\n", usersPath, err)
			os.Exit(1)
		}
		defer usersFile.Close()
		files["user"] = usersFile
	}
	if departmentsPath != "" {
		departmentsFile, err := os.Open(departmentsPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file '%s': %s\n", departmentsPath, err)
			os.Exit(1)
		}
		defer departmentsFile.Close()
		files["department"] = departmentsFile
	}

	result, err := client.Check(files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}
