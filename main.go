package main

import (
	"flag"
	"fmt"
	"os"

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

	harnessAccount := harness.Account{}

	fmt.Println(color.GreenString("Creating report directory..."))
	if _, err := os.Stat("report"); os.IsNotExist(err) {
		err = os.Mkdir("report", 0755)
		if err != nil {
			panic("Unable to create directory")
		}
	}
	if _, err := os.Stat("report/data"); os.IsNotExist(err) {
		err = os.Mkdir("report/data", 0755)
		if err != nil {
			panic("Unable to create directory")
		}
	}

	err := templates.CopyEmbeddedFile(templates.EmbeddedFiles, "index.html", "report/index.html")
	if err != nil {
		fmt.Println("Failed to copy index template to report folder: ", err)
	}

	apiCalls := []func(string, string) (harness.Entities, error){
		api.GetAllUserGroups,
		api.GetAllResourceGroups,
		api.GetAllRoles,
		api.GetAllRoleAssignments,
		api.GetAllUsers,
		api.GetAllConnectors,
		api.GetAllProjects,
		api.GetAllOrganizations,
	}
	entities, _ := api.GetAccountOverview(len(apiCalls), apiCalls, *formatArg, *accountArg)
	tmpl := `{{ green "Generating reports: " }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}} `
	bar := pb.ProgressBarTemplate(tmpl).Start(len(entities))
	for _, entity := range entities {
		color.Set(color.FgHiGreen, color.Bold)
		bar.Increment()
		switch entity.EntityType {
		case "UserGroups":
			_ = entity.EntityResult.(harness.UserGroups).FormatUserGroups()
			harnessAccount.UserGroups = entity.EntityResult.(harness.UserGroups)
			fmt.Println("Generated report for User Groups!")
		case "Users":
			_ = entity.EntityResult.(harness.Users).FormatUsers()
			harnessAccount.Users = entity.EntityResult.(harness.Users)
			fmt.Println("Generated report for Users!")
		case "Roles":
			_ = entity.EntityResult.(harness.Roles).FormatRoles()
			harnessAccount.Roles = entity.EntityResult.(harness.Roles)
			fmt.Println("Generated report for Roles!")
		case "RoleAssignments":
			_ = entity.EntityResult.(harness.RoleAssignments).FormatRoleAssingment()
			harnessAccount.RoleAssignments = entity.EntityResult.(harness.RoleAssignments)
			fmt.Println("Generated report for Role Assignments!")
		case "ResourceGroups":
			_ = entity.EntityResult.(harness.ResourceGroups).FormatResourceGroups()
			harnessAccount.ResourceGroups = entity.EntityResult.(harness.ResourceGroups)
			fmt.Println("Generated report for Resource Groups!")
		case "Connectors":
			_ = entity.EntityResult.(harness.Connectors).FormatConnectors()
			harnessAccount.Connectors = entity.EntityResult.(harness.Connectors)
			fmt.Println("Generated report for Connectors!")
		case "Projects":
			_ = entity.EntityResult.(harness.Projects).FormatProjects()
			harnessAccount.Projects = entity.EntityResult.(harness.Projects)
			fmt.Println("Generated report for Projects!")
		case "Organizations":
			_ = entity.EntityResult.(harness.Organizations).FormatOrganizations()
			harnessAccount.Organizations = entity.EntityResult.(harness.Organizations)
			fmt.Println("Generated report for Organizations!")
		default:
			fmt.Println(color.RedString("No format function found for entity type: "), color.HiRedString(entity.EntityType))
		}
	}
	color.Unset()
	bar.Finish()

	bgreen := color.New(color.Bold, color.FgGreen)
	bgreen.Println("Generating custom reports...")

	if *adminUsersArg {
		bgreen.Println("- Generating report for admin users...")
		ug, _ := api.GetAllAdminUsers(*accountArg, []string{"Field_Engineering"})
		userMap := make(map[string]harness.UsersContent)
		for _, user := range harnessAccount.Users.Data.Content {
			userMap[user.UUID] = user
		}
		for _, ugContent := range ug.EntityResult.(harness.UserGroupsFiltered).Data {
			for _, u := range ugContent.Users {
				if userMap[u].Email == "" {
					fmt.Println(color.RedString("User not found: "), color.HiRedString(u))
					continue
				}
				fmt.Println(color.YellowString("User is an admin: "), color.HiYellowString(userMap[u].Email))
			}
		}
		_ = ug.EntityResult.(harness.UserGroupsFiltered).ListAdminUsers(userMap)
	}

	bgreen.Println("Report generated successfully! Please navigate to report folder and open index.html in your browser.")
}
