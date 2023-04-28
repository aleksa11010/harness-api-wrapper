package harness

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/cheggaaa/pb/v3"
	resty "github.com/go-resty/resty/v2"
)

type APIRequest struct {
	BaseURL string
	Client  *resty.Client
	APIKey  string
}

type Entities struct {
	EntityType   string
	EntityResult interface{}
}

type EntityResult interface{}

func (a *APIRequest) GetAccountOverview(callCount int, callFuncs []func(string, string) (Entities, error), format string, account string) ([]Entities, error) {
	type result struct {
		entityResponse Entities
		err            error
	}

	results := make(chan result)
	tmpl := `{{ blue "Calling Harness API: " }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}} `
	bar := pb.ProgressBarTemplate(tmpl).Start(callCount)

	var wg sync.WaitGroup
	wg.Add(callCount)

	for _, callFunc := range callFuncs {
		go func(callFunc func(string, string) (Entities, error)) {
			defer wg.Done()
			resp, err := callFunc(format, account)
			results <- result{entityResponse: resp, err: err}
			bar.Increment()
		}(callFunc)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	responses := make([]Entities, 0, callCount)
	for res := range results {
		if res.err != nil {
			bar.Finish()
			return nil, res.err
		}
		responses = append(responses, res.entityResponse)
	}

	bar.Finish()
	return responses, nil
}

func (api *APIRequest) GetAllUserGroups(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"filterType":        "INCLUDE_CHILD_SCOPE_GROUPS",
			"pageSize":          "1000"}).
		Get(api.BaseURL + "/ng/api/user-groups?")
	if err != nil {
		return Entities{}, err
	}

	userGroups := UserGroups{}
	err = json.Unmarshal([]byte(resp.String()), &userGroups)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "UserGroups",
		EntityResult: userGroups,
	}

	return entity, nil
}

func (api *APIRequest) GetAllRoleAssignments(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
		}).
		Get(api.BaseURL + "/authz/api/roleassignments")
	if err != nil {
		return Entities{}, err
	}

	roleAssignment := RoleAssignments{}
	err = json.Unmarshal([]byte(resp.String()), &roleAssignment)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return Entities{}, err
	}
	entity := Entities{
		EntityType:   "RoleAssignments",
		EntityResult: roleAssignment,
	}

	return entity, nil
}

func (api *APIRequest) GetAllResourceGroups(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"pageSize":          "500",
		}).
		Get(api.BaseURL + "/resourcegroup/api/v2/resourcegroup")
	if err != nil {
		return Entities{}, err
	}

	resourceGroups := ResourceGroups{}
	err = json.Unmarshal([]byte(resp.String()), &resourceGroups)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "ResourceGroups",
		EntityResult: resourceGroups,
	}

	return entity, nil
}

func (api *APIRequest) GetAllRoles(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"pageSize":          "500"}).
		Get(api.BaseURL + "/authz/api/roles")
	if err != nil {
		return Entities{}, err
	}

	roles := Roles{}
	err = json.Unmarshal([]byte(resp.String()), &roles)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "Roles",
		EntityResult: roles,
	}

	return entity, nil
}

func (api *APIRequest) GetAllUsers(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(`{"filterType": "INCLUDE_CHILD_SCOPE_GROUPS"}`).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"pageSize":          "100",
		}).
		Post(api.BaseURL + "/ng/api/user/batch")
	if err != nil {
		return Entities{}, err
	}

	users := Users{}
	err = json.Unmarshal([]byte(resp.String()), &users)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return Entities{}, err
	}
	userData := []UsersData{}
	if users.Data.TotalPages > 1 && users.Data.PageIndex < users.Data.TotalPages {
		userData = append(userData, users.Data)
		for i := users.Data.PageIndex + 1; i < users.Data.TotalPages; i++ {
			users = Users{}
			resp, err := api.Client.R().
				SetHeader("x-api-key", api.APIKey).
				SetHeader("Content-Type", "application/json").
				SetBody(`{"filterType": "INCLUDE_CHILD_SCOPE_GROUPS"}`).
				SetQueryParams(map[string]string{
					"accountIdentifier": account,
					"pageSize":          "100",
					"pageIndex":         strconv.FormatInt(i, 10),
				}).
				Post(api.BaseURL + "/ng/api/user/batch")
			if err != nil {
				return Entities{}, err
			}
			err = json.Unmarshal([]byte(resp.String()), &users)
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
				return Entities{}, err
			}
			userData = append(userData, users.Data)
		}

		combinedReport := []UsersContent{}
		for _, user := range userData {
			combinedReport = append(combinedReport, user.Content...)
		}

		users.Data.Content = combinedReport
		entity := Entities{
			EntityType:   "Users",
			EntityResult: users,
		}

		return entity, nil
	}

	entity := Entities{
		EntityType:   "Users",
		EntityResult: users,
	}

	return entity, nil
}

func (api *APIRequest) GetAllAdminUsers(account string, ug []string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"accountIdentifier": "%s","identifierFilter": ["%s"],"filterType": "EXCLUDE_INHERITED_GROUPS"}`, account, strings.Join(ug, "\",\""))).
		SetQueryParam("accountIdentifier", account).
		Post(api.BaseURL + "/ng/api/user-groups/batch")
	if err != nil {
		return Entities{}, err
	}

	userGroups := UserGroupsFiltered{}
	err = json.Unmarshal([]byte(resp.String()), &userGroups)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "UserGroups",
		EntityResult: userGroups,
	}

	return entity, nil
}

func (api *APIRequest) GetAllConnectors(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"pageSize":          "500",
		}).
		SetBody(`{"filterType": "Connector"}`).
		Post(api.BaseURL + "/ng/api/connectors/listV2")
	if err != nil {
		return Entities{}, err
	}
	connectors := Connectors{}
	err = json.Unmarshal(resp.Body(), &connectors)
	if err != nil {
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "Connectors",
		EntityResult: connectors,
	}

	return entity, nil
}

func (api *APIRequest) GetAllProjects(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"hasModule":         "true",
			"pageSize":          "500",
		}).
		Get(api.BaseURL + "/ng/api/projects")
	if err != nil {
		return Entities{}, err
	}
	projects := Projects{}
	err = json.Unmarshal(resp.Body(), &projects)
	if err != nil {
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "Projects",
		EntityResult: projects,
	}

	return entity, nil
}

func (api *APIRequest) GetAllOrganizations(format string, account string) (Entities, error) {
	resp, err := api.Client.R().
		SetHeader("x-api-key", api.APIKey).
		SetQueryParams(map[string]string{
			"accountIdentifier": account,
			"pageSize":          "500",
		}).
		Get(api.BaseURL + "/ng/api/organizations")
	if err != nil {
		return Entities{}, err
	}
	organizations := Organizations{}
	err = json.Unmarshal(resp.Body(), &organizations)
	if err != nil {
		return Entities{}, err
	}

	entity := Entities{
		EntityType:   "Organizations",
		EntityResult: organizations,
	}

	return entity, nil
}
