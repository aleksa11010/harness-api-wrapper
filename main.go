package main

import (
	"flag"
	"fmt"
)

func main() {
	accountArg := flag.String("account", "", "Provide your account ID.")
	apiKeyArg := flag.String("api-key", "", "Provide your API Key.")
	formatArg := flag.String("format", "json", "Provide the output format, defaults to json. Options: json, csv, yaml ")

	flag.Parse()

	fmt.Println("Account:", *accountArg)
	fmt.Println("API Key:", *apiKeyArg)
	fmt.Println("Format:", *formatArg)
}
