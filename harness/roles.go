package harness

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"text/template"

	"github.com/aleksa11010/harness-api-wrapper/templates"
)

// Generated by https://quicktype.io

type Roles struct {
	Status        string      `json:"status"`
	Data          RolesData   `json:"data"`
	MetaData      interface{} `json:"metaData"`
	CorrelationID string      `json:"correlationId"`
}

type RolesData struct {
	TotalPages    int64          `json:"totalPages"`
	TotalItems    int64          `json:"totalItems"`
	PageItemCount int64          `json:"pageItemCount"`
	PageSize      int64          `json:"pageSize"`
	Content       []RolesContent `json:"content"`
	PageIndex     int64          `json:"pageIndex"`
	Empty         bool           `json:"empty"`
	PageToken     interface{}    `json:"pageToken"`
}

type RolesContent struct {
	Role           Role        `json:"role"`
	Scope          *RolesScope `json:"scope"`
	HarnessManaged bool        `json:"harnessManaged"`
	CreatedAt      int64       `json:"createdAt"`
	LastModifiedAt int64       `json:"lastModifiedAt"`
}

type Role struct {
	Identifier         string                   `json:"identifier"`
	Name               string                   `json:"name"`
	Permissions        []string                 `json:"permissions"`
	AllowedScopeLevels []RolesAllowedScopeLevel `json:"allowedScopeLevels"`
	Description        string                   `json:"description"`
	Tags               interface{}              `json:"tags"`
}

type RolesScope struct {
	AccountIdentifier RolesAccountIdentifier `json:"accountIdentifier"`
	OrgIdentifier     interface{}            `json:"orgIdentifier"`
	ProjectIdentifier interface{}            `json:"projectIdentifier"`
}

type RolesAllowedScopeLevel string

const (
	RolesAccount      RolesAllowedScopeLevel = "account"
	RolesOrganization RolesAllowedScopeLevel = "organization"
	RolesProject      RolesAllowedScopeLevel = "project"
)

type RolesAccountIdentifier string

type RolesAPI interface {
	GetAllRoles(format string, account string) (Entities, error)
}

func (r Roles) FormatRoles() error {

	content := make([]json.RawMessage, 0)
	for _, role := range r.Data.Content {
		role, err := json.Marshal(role)
		if err != nil {
			fmt.Printf("error marshalling user group: %s", err)
		}
		content = append(content, json.RawMessage(role))
	}

	reportData, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Unable to marshal json", err)
		return err
	}
	err = createRolesReport(reportData)
	if err != nil {
		fmt.Println("Unable to create a HTML output :", err)
		return err
	}
	return nil
}

func createRolesReport(content []byte) error {
	type ReportData struct {
		Header  string
		Content string
	}

	reportData := ReportData{
		Header:  "Roles",
		Content: string(content),
	}

	userGroupsTemplate, err := fs.ReadFile(templates.EmbeddedFiles, "report.html")
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

	file, err := os.OpenFile("./report/data/roles.html", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
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
