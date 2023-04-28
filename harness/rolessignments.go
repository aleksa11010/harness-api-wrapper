package harness

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type RoleAssignments struct {
	Status        string              `json:"status"`
	Data          RoleAssignmentsData `json:"data"`
	MetaData      interface{}         `json:"metaData"`
	CorrelationID string              `json:"correlationId"`
}

type RoleAssignmentsData struct {
	TotalPages    int64                    `json:"totalPages"`
	TotalItems    int64                    `json:"totalItems"`
	PageItemCount int64                    `json:"pageItemCount"`
	PageSize      int64                    `json:"pageSize"`
	Content       []RoleAssignmentsContent `json:"content"`
	PageIndex     int64                    `json:"pageIndex"`
	Empty         bool                     `json:"empty"`
	PageToken     interface{}              `json:"pageToken"`
}

type RoleAssignmentsContent struct {
	RoleAssignment RoleAssignment       `json:"roleAssignment"`
	Scope          RoleAssignmentsScope `json:"scope"`
	CreatedAt      int64                `json:"createdAt"`
	LastModifiedAt int64                `json:"lastModifiedAt"`
	HarnessManaged bool                 `json:"harnessManaged"`
}

type RoleAssignment struct {
	Identifier              string                   `json:"identifier"`
	ResourceGroupIdentifier ResourceGroupIdentifier  `json:"resourceGroupIdentifier"`
	RoleIdentifier          string                   `json:"roleIdentifier"`
	Principal               RoleAssignmentsPrincipal `json:"principal"`
	Disabled                bool                     `json:"disabled"`
	Managed                 bool                     `json:"managed"`
	Internal                bool                     `json:"internal"`
}

type RoleAssignmentsPrincipal struct {
	ScopeLevel *string `json:"scopeLevel"`
	Identifier string  `json:"identifier"`
	Type       Type    `json:"type"`
}

type RoleAssignmentsScope struct {
	AccountIdentifier ResourceGroupsAccountIdentifier `json:"accountIdentifier"`
	OrgIdentifier     interface{}                     `json:"orgIdentifier"`
	ProjectIdentifier interface{}                     `json:"projectIdentifier"`
}

type Type string

const (
	ServiceAccount Type = "SERVICE_ACCOUNT"
	User           Type = "USER"
	UserGroup      Type = "USER_GROUP"
)

type ResourceGroupIdentifier string

type RoleAssignmentsAccountIdentifier string

type RoleAssignmentAPI interface {
	GetAllRoleAssignments(format string, account string) (Entities, error)
}

func (ra RoleAssignments) FormatRoleAssingment() error {

	content := make([]json.RawMessage, 0)
	for _, ras := range ra.Data.Content {
		r, err := json.Marshal(ras)
		if err != nil {
			fmt.Printf("error marshalling : %s", err)
		}
		content = append(content, json.RawMessage(r))
	}

	reportData, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Unable to marshal json", err)
		return err
	}
	err = createRoleAssingmentsReport(reportData)
	if err != nil {
		fmt.Println("Unable to create a HTML output :", err)
		return err
	}
	return nil
}

func createRoleAssingmentsReport(content []byte) error {
	type ReportData struct {
		Header  string
		Content string
	}

	reportData := ReportData{
		Header:  "Role Assignments",
		Content: string(content),
	}

	userGroupsTemplate, err := os.ReadFile("templates/report.html")
	if err != nil {
		fmt.Printf("error reading template: %s", err)
	}

	t, err := template.New("roles").Parse(string(userGroupsTemplate))
	if err != nil {
		fmt.Printf("error parsing template: %s", err)
		return err
	}

	var output bytes.Buffer
	err = t.Execute(&output, reportData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	file, err := os.OpenFile("./report/data/rolessignments.html", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(output.Bytes())
	if err != nil {
		panic(err)
	}

	return nil
}
