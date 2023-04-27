package harness

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/cheggaaa/pb/v3"
	resty "github.com/go-resty/resty/v2"
)

type HarnessAPI interface {
	// GetAllConnectors() ([]Connectors, error)
	GetAllUserGroups() ([]UserGroups, error)
}

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
		Get(api.BaseURL + "/ng/api/user-groups?accountIdentifier=" + account + "&filterType=INCLUDE_CHILD_SCOPE_GROUPS&pageSize=1000")
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
		Get(api.BaseURL + "/authz/api/roleassignments?accountIdentifier=" + account)
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
		Get(api.BaseURL + "/authz/api/roleassignments?accountIdentifier=" + account)
	if err != nil {
		return Entities{}, err
	}

	resourceGroups := ResourceGroup{}
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
		Get(api.BaseURL + "/authz/api/roles?accountIdentifier=" + account + "&pageSize=500")
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

func (api *APIRequest) GetAllConnectors(format string) (Entities, error) {
	resp, err := api.Client.R().Post(api.BaseURL + "/gateway")
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
