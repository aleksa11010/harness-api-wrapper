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

type Connectors struct {
	Status        ConnectorStatusEnum `json:"status"`
	Data          ConnectorData       `json:"data"`
	MetaData      interface{}         `json:"metaData"`
	CorrelationID string              `json:"correlationId"`
}

type ConnectorData struct {
	TotalPages    int64              `json:"totalPages"`
	TotalItems    int64              `json:"totalItems"`
	PageItemCount int64              `json:"pageItemCount"`
	PageSize      int64              `json:"pageSize"`
	Content       []ConnectorContent `json:"content"`
	PageIndex     int64              `json:"pageIndex"`
	Empty         bool               `json:"empty"`
	PageToken     interface{}        `json:"pageToken"`
}

type ConnectorContent struct {
	Connector             Connector                      `json:"connector"`
	CreatedAt             int64                          `json:"createdAt"`
	LastModifiedAt        int64                          `json:"lastModifiedAt"`
	Status                ConnectorStatusClass           `json:"status"`
	ActivityDetails       ConnectorActivityDetails       `json:"activityDetails"`
	HarnessManaged        bool                           `json:"harnessManaged"`
	GitDetails            map[string]interface{}         `json:"gitDetails"`
	EntityValidityDetails ConnectorEntityValidityDetails `json:"entityValidityDetails"`
	GovernanceMetadata    interface{}                    `json:"governanceMetadata"`
}

type ConnectorActivityDetails struct {
	LastActivityTime int64 `json:"lastActivityTime"`
}

type Connector struct {
	Name              string        `json:"name"`
	Identifier        string        `json:"identifier"`
	Description       *string       `json:"description"`
	OrgIdentifier     interface{}   `json:"orgIdentifier"`
	ProjectIdentifier interface{}   `json:"projectIdentifier"`
	Tags              interface{}   `json:"tags"`
	Type              string        `json:"type"`
	Spec              ConnectorSpec `json:"spec"`
}

type ConnectorSpec struct {
	URL                                 *string                      `json:"url,omitempty"`
	ValidationRepo                      *string                      `json:"validationRepo"`
	Authentication                      *ConnectorAuthentication     `json:"authentication,omitempty"`
	APIAccess                           *ConnectorAPIAccess          `json:"apiAccess"`
	DelegateSelectors                   []string                     `json:"delegateSelectors"`
	ExecuteOnDelegate                   *bool                        `json:"executeOnDelegate,omitempty"`
	Type                                *ConnectorFluffyType         `json:"type,omitempty"`
	Credential                          *ConnectorCredential         `json:"credential,omitempty"`
	DockerRegistryURL                   *string                      `json:"dockerRegistryUrl,omitempty"`
	ProviderType                        *ConnectorProviderType       `json:"providerType,omitempty"`
	Auth                                *ConnectorPurpleAuth         `json:"auth,omitempty"`
	ConnectorRef                        *string                      `json:"connectorRef"`
	FeaturesEnabled                     []ConnectorFeaturesEnabled   `json:"featuresEnabled"`
	ArtifactoryServerURL                *string                      `json:"artifactoryServerUrl,omitempty"`
	ApplicationKeyRef                   *string                      `json:"applicationKeyRef,omitempty"`
	APIKeyRef                           *string                      `json:"apiKeyRef"`
	NewRelicAccountID                   *string                      `json:"newRelicAccountId,omitempty"`
	Username                            *string                      `json:"username"`
	Accountname                         *string                      `json:"accountname,omitempty"`
	ControllerURL                       *string                      `json:"controllerUrl,omitempty"`
	PasswordRef                         *string                      `json:"passwordRef"`
	ClientSecretRef                     *string                      `json:"clientSecretRef"`
	ClientID                            *string                      `json:"clientId"`
	AuthType                            *string                      `json:"authType,omitempty"`
	HelmRepoURL                         *string                      `json:"helmRepoUrl,omitempty"`
	JiraURL                             *string                      `json:"jiraUrl,omitempty"`
	UsernameRef                         *string                      `json:"usernameRef"`
	APITokenRef                         *string                      `json:"apiTokenRef,omitempty"`
	ServiceNowURL                       *string                      `json:"serviceNowUrl,omitempty"`
	Headers                             []interface{}                `json:"headers"`
	TenantID                            *string                      `json:"tenantId,omitempty"`
	SubscriptionID                      *string                      `json:"subscriptionId,omitempty"`
	BillingExportSpec                   *ConnectorBillingExportSpec  `json:"billingExportSpec"`
	CrossAccountAccess                  *ConnectorCrossAccountAccess `json:"crossAccountAccess,omitempty"`
	CurAttributes                       *ConnectorCurAttributes      `json:"curAttributes"`
	AwsAccountID                        *string                      `json:"awsAccountId,omitempty"`
	IsAWSGovCloudAccount                *bool                        `json:"isAWSGovCloudAccount"`
	ProjectID                           *string                      `json:"projectId,omitempty"`
	ServiceAccountEmail                 *string                      `json:"serviceAccountEmail,omitempty"`
	Credentials                         interface{}                  `json:"credentials"`
	Default                             *bool                        `json:"default,omitempty"`
	AwsSDKClientBackOffStrategyOverride interface{}                  `json:"awsSdkClientBackOffStrategyOverride"`
	Hosts                               []ConnectorHost              `json:"hosts"`
	OnDelegate                          *bool                        `json:"onDelegate,omitempty"`
	Host                                interface{}                  `json:"host"`
	WorkingDirectory                    interface{}                  `json:"workingDirectory"`
	Template                            *ConnectorTemplate           `json:"template,omitempty"`
	AzureEnvironmentType                *string                      `json:"azureEnvironmentType,omitempty"`
	APIKeyID                            interface{}                  `json:"apiKeyId"`
	Region                              *string                      `json:"region,omitempty"`
	SecretNamePrefix                    *string                      `json:"secretNamePrefix,omitempty"`
}

type ConnectorAPIAccess struct {
	Type ConnectorAPIAccessType `json:"type"`
	Spec ConnectorAPIAccessSpec `json:"spec"`
}

type ConnectorAPIAccessSpec struct {
	TokenRef          *string     `json:"tokenRef,omitempty"`
	InstallationID    *string     `json:"installationId,omitempty"`
	ApplicationID     *string     `json:"applicationId,omitempty"`
	InstallationIDRef interface{} `json:"installationIdRef"`
	ApplicationIDRef  interface{} `json:"applicationIdRef"`
	PrivateKeyRef     *string     `json:"privateKeyRef,omitempty"`
}

type ConnectorPurpleAuth struct {
	Type ConnectorAuthType    `json:"type"`
	Spec *ConnectorPurpleSpec `json:"spec,omitempty"`
}

type ConnectorPurpleSpec struct {
	Username    *string `json:"username"`
	UsernameRef *string `json:"usernameRef"`
	PasswordRef string  `json:"passwordRef"`
}

type ConnectorAuthentication struct {
	Type ConnectorAuthenticationType `json:"type"`
	Spec ConnectorAuthenticationSpec `json:"spec"`
}

type ConnectorAuthenticationSpec struct {
	Type      *ConnectorPurpleType `json:"type,omitempty"`
	Spec      *ConnectorSpecSpec   `json:"spec,omitempty"`
	SSHKeyRef *string              `json:"sshKeyRef,omitempty"`
}

type ConnectorSpecSpec struct {
	Username     *string     `json:"username,omitempty"`
	UsernameRef  interface{} `json:"usernameRef"`
	TokenRef     *string     `json:"tokenRef,omitempty"`
	PasswordRef  *string     `json:"passwordRef,omitempty"`
	AccessKey    *string     `json:"accessKey,omitempty"`
	AccessKeyRef interface{} `json:"accessKeyRef"`
	SecretKeyRef *string     `json:"secretKeyRef,omitempty"`
}

type ConnectorBillingExportSpec struct {
	StorageAccountName *string `json:"storageAccountName,omitempty"`
	ContainerName      *string `json:"containerName,omitempty"`
	DirectoryName      *string `json:"directoryName,omitempty"`
	ReportName         *string `json:"reportName,omitempty"`
	SubscriptionID     *string `json:"subscriptionId,omitempty"`
	DatasetID          *string `json:"datasetId,omitempty"`
	TableID            *string `json:"tableId,omitempty"`
}

type ConnectorCredential struct {
	Type               ConnectorCredentialType  `json:"type"`
	Spec               *ConnectorCredentialSpec `json:"spec"`
	CrossAccountAccess interface{}              `json:"crossAccountAccess"`
	Region             *string                  `json:"region"`
}

type ConnectorCredentialSpec struct {
	MasterURL     *string              `json:"masterUrl,omitempty"`
	Auth          *ConnectorFluffyAuth `json:"auth,omitempty"`
	SecretKeyRef  *string              `json:"secretKeyRef,omitempty"`
	AccessKey     *string              `json:"accessKey,omitempty"`
	AccessKeyRef  interface{}          `json:"accessKeyRef"`
	ApplicationID *string              `json:"applicationId,omitempty"`
	TenantID      *string              `json:"tenantId,omitempty"`
	SecretKey     *string              `json:"secretKey,omitempty"`
}

type ConnectorFluffyAuth struct {
	Type string               `json:"type"`
	Spec *ConnectorFluffySpec `json:"spec,omitempty"`
}

type ConnectorFluffySpec struct {
	Username               *string     `json:"username,omitempty"`
	UsernameRef            interface{} `json:"usernameRef"`
	PasswordRef            *string     `json:"passwordRef,omitempty"`
	ServiceAccountTokenRef *string     `json:"serviceAccountTokenRef,omitempty"`
	CACERTRef              interface{} `json:"caCertRef"`
	SecretRef              *string     `json:"secretRef,omitempty"`
}

type ConnectorCrossAccountAccess struct {
	CrossAccountRoleArn string `json:"crossAccountRoleArn"`
	ExternalID          string `json:"externalId"`
}

type ConnectorCurAttributes struct {
	ReportName   string `json:"reportName"`
	S3BucketName string `json:"s3BucketName"`
	Region       string `json:"region"`
	S3Prefix     string `json:"s3Prefix"`
}

type ConnectorHost struct {
	Hostname       string                  `json:"hostname"`
	HostAttributes ConnectorTemplateInputs `json:"hostAttributes"`
}

type ConnectorTemplateInputs struct {
}

type ConnectorTemplate struct {
	TemplateRef    string                  `json:"templateRef"`
	VersionLabel   string                  `json:"versionLabel"`
	TemplateInputs ConnectorTemplateInputs `json:"templateInputs"`
}

type ConnectorEntityValidityDetails struct {
	Valid       bool        `json:"valid"`
	InvalidYAML interface{} `json:"invalidYaml"`
}

type ConnectorStatusClass struct {
	Status          ConnectorStatusEnum `json:"status"`
	ErrorSummary    *string             `json:"errorSummary"`
	Errors          []ConnectorError    `json:"errors"`
	TestedAt        int64               `json:"testedAt"`
	LastTestedAt    int64               `json:"lastTestedAt"`
	LastConnectedAt int64               `json:"lastConnectedAt"`
}

type ConnectorError struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type ConnectorAPIAccessType string

const (
	GithubApp   ConnectorAPIAccessType = "GithubApp"
	PurpleOAuth ConnectorAPIAccessType = "OAuth"
	Token       ConnectorAPIAccessType = "Token"
)

type ConnectorAuthType string

const (
	ConnectorAnonymous              ConnectorAuthType = "Anonymous"
	ConnectorPurpleUsernamePassword ConnectorAuthType = "UsernamePassword"
)

type ConnectorPurpleType string

const (
	ConnectorAWSCredentials         ConnectorPurpleType = "AWSCredentials"
	ConnectorFluffyOAuth            ConnectorPurpleType = "OAuth"
	ConnectorFluffyUsernamePassword ConnectorPurpleType = "UsernamePassword"
	ConnectorUsernameToken          ConnectorPurpleType = "UsernameToken"
)

type ConnectorAuthenticationType string

const (
	ConnectorHTTP  ConnectorAuthenticationType = "Http"
	ConnectorHTTPS ConnectorAuthenticationType = "HTTPS"
	ConnectorSSH   ConnectorAuthenticationType = "Ssh"
)

type ConnectorCredentialType string

const (
	ConnectorInheritFromDelegate ConnectorCredentialType = "InheritFromDelegate"
	ConnectorManualConfig        ConnectorCredentialType = "ManualConfig"
)

type ConnectorFeaturesEnabled string

const (
	ConnectorBilling      ConnectorFeaturesEnabled = "BILLING"
	ConnectorGovernance   ConnectorFeaturesEnabled = "GOVERNANCE"
	ConnectorOptimization ConnectorFeaturesEnabled = "OPTIMIZATION"
	ConnectorVisibility   ConnectorFeaturesEnabled = "VISIBILITY"
)

type ConnectorProviderType string

const (
	ConnectorDockerHub ConnectorProviderType = "DockerHub"
	ConnectorOther     ConnectorProviderType = "Other"
)

type ConnectorFluffyType string

const (
	ConnectorAccount ConnectorFluffyType = "Account"
	ConnectorRepo    ConnectorFluffyType = "Repo"
)

type ConnectorStatusEnum string

const (
	ConnectorFailure ConnectorStatusEnum = "FAILURE"
	ConnectorSuccess ConnectorStatusEnum = "SUCCESS"
)

func (c Connector) String() string {
	return c.Name
}

type ConnectorAPI interface {
	GetAllConnectors() ([]Connector, error)
}

func (c Connectors) FormatConnectors() error {

	content := make([]json.RawMessage, 0)
	for _, con := range c.Data.Content {
		c, err := json.Marshal(con)
		if err != nil {
			fmt.Printf("error marshalling user group: %s", err)
		}
		content = append(content, json.RawMessage(c))
	}

	reportData, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Unable to marshal json", err)
		return err
	}
	err = createConnectorsReport(reportData)
	if err != nil {
		fmt.Println("Unable to create a HTML output :", err)
		return err
	}
	return nil
}

func createConnectorsReport(content []byte) error {
	type ReportData struct {
		Header  string
		Content string
	}

	reportData := ReportData{
		Header:  "Connectors",
		Content: string(content),
	}

	rgTemplate, err := fs.ReadFile(templates.EmbeddedFiles, "report.html")
	if err != nil {
		fmt.Printf("error reading template: %s", err)
	}

	t, err := template.New("connectors").Parse(string(rgTemplate))
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

	file, err := os.OpenFile("./report/data/connectors.html", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
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
