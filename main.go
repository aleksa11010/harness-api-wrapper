package main

import (
	"flag"
	"fmt"

	"github.com/aleksa11010/harness-api-wrapper/harness"
	"github.com/aleksa11010/harness-api-wrapper/templates"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
)

func main() {

	accountArg := flag.String("account", "", "Provide your account ID.")
	apiKeyArg := flag.String("api-key", "", "Provide your API Key.")
	formatArg := flag.String("format", "json", "Provide the output format, defaults to json. Options: json, csv, yaml ")
	adminUsersArg := flag.Bool("admin-users", false, "Custome report to return admin users.")

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

	apiCalls := []func(string, string) (harness.Entities, error){
		api.GetAllUserGroups,
		api.GetAllResourceGroups,
		api.GetAllRoles,
		api.GetAllRoleAssignments,
		api.GetAllUsers,
		api.GetAllConnectors,
	}
	entities, _ := api.GetAccountOverview(len(apiCalls), apiCalls, *formatArg, *accountArg)
	tmpl := `{{ green "Generating report: " }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}} `
	bar := pb.ProgressBarTemplate(tmpl).Start(len(entities))
	for _, entity := range entities {
		color.Set(color.FgHiGreen, color.Bold)
		bar.Increment()
		switch entity.EntityType {
		case "UserGroups":
			_ = entity.EntityResult.(harness.UserGroups).FormatUserGroups()
			fmt.Println("Generated report for User Groups!")
		case "Users":
			_ = entity.EntityResult.(harness.Users).FormatUsers()
			if *adminUsersArg {
				fmt.Println(color.RedString("Report for Admin Users not implemented yet!"))
			}
			fmt.Println("Generated report for Users!")
		case "Roles":
			_ = entity.EntityResult.(harness.Roles).FormatRoles()
			fmt.Println("Generated report for Roles!")
		case "RoleAssignments":
			_ = entity.EntityResult.(harness.RoleAssignments).FormatRoleAssingment()
			fmt.Println("Generated report for Role Assignments!")
		case "ResourceGroups":
			_ = entity.EntityResult.(harness.ResourceGroups).FormatResourceGroups()
			fmt.Println("Generated report for Resource Groups!")
		case "Connectors":
			_ = entity.EntityResult.(harness.Connectors).FormatConnectors()
			fmt.Println("Generated report for Connectors!")
		default:
			fmt.Println(color.RedString("No format function found for entity type: "), color.HiRedString(entity.EntityType))

		}
	}
	color.Unset()
	err := templates.CopyEmbeddedFile(templates.EmbeddedFiles, "index.html", "report/index.html")
	if err != nil {
		fmt.Println("Failed to copy index template to report folder: ", err)
	}
	bar.Finish()

	bgreen := color.New(color.Bold, color.FgGreen)
	bgreen.Println("Report generated successfully! Please navigate to report folder and open index.html in your browser.")
}
