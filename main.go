package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aleksa11010/harness-api-wrapper/harness"
	"github.com/cheggaaa/pb/v3"
	"github.com/go-resty/resty/v2"
)

func main() {

	tmpl := `{{ blue "Calling Harness API: " }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}} `
	bar := pb.ProgressBarTemplate(tmpl).Start(100)
	for i := 0; i < 100; i++ {
		bar.Increment()
		time.Sleep(time.Second * 2)
	}

	accountArg := flag.String("account", "", "Provide your account ID.")
	apiKeyArg := flag.String("api-key", "", "Provide your API Key.")
	formatArg := flag.String("format", "json", "Provide the output format, defaults to json. Options: json, csv, yaml ")

	flag.Parse()

	if *accountArg == "" || *apiKeyArg == "" {
		fmt.Println("Account ID and API Key are required!")
		return
	}

	api := harness.APIRequest{
		BaseURL: harness.BaseURL,
		Client:  resty.New(),
		APIKey:  *apiKeyArg,
	}

	apiCalls := []func(string) (harness.EntityResult, error){
		api.GetAllConnectors,
	}
	_, _ = api.GetAccountOverview(len(apiCalls), apiCalls, *formatArg)
}
