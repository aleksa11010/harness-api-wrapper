package harness

type Connectors struct {
	Status        StatusEnum  `json:"status"`
	Data          Data        `json:"data"`
	MetaData      interface{} `json:"metaData"`
	CorrelationID string      `json:"correlationId"`
}

type Data struct {
	TotalPages    int64       `json:"totalPages"`
	TotalItems    int64       `json:"totalItems"`
	PageItemCount int64       `json:"pageItemCount"`
	PageSize      int64       `json:"pageSize"`
	Content       []Content   `json:"content"`
	PageIndex     int64       `json:"pageIndex"`
	Empty         bool        `json:"empty"`
	PageToken     interface{} `json:"pageToken"`
}

type Content struct {
	Connector             Connector              `json:"connector"`
	CreatedAt             int64                  `json:"createdAt"`
	LastModifiedAt        int64                  `json:"lastModifiedAt"`
	Status                StatusClass            `json:"status"`
	ActivityDetails       ActivityDetails        `json:"activityDetails"`
	HarnessManaged        bool                   `json:"harnessManaged"`
	GitDetails            map[string]interface{} `json:"gitDetails"`
	EntityValidityDetails EntityValidityDetails  `json:"entityValidityDetails"`
	GovernanceMetadata    interface{}            `json:"governanceMetadata"`
}

type ActivityDetails struct {
	LastActivityTime int64 `json:"lastActivityTime"`
}

type Connector struct {
	Name              string        `json:"name"`
	Identifier        string        `json:"identifier"`
	Description       *string       `json:"description"`
	OrgIdentifier     interface{}   `json:"orgIdentifier"`
	ProjectIdentifier interface{}   `json:"projectIdentifier"`
	Tags              Tags          `json:"tags"`
	Type              string        `json:"type"`
	Spec              ConnectorSpec `json:"spec"`
}

type ConnectorSpec struct {
	URL                                 *string             `json:"url,omitempty"`
	ValidationRepo                      *string             `json:"validationRepo"`
	Authentication                      *Authentication     `json:"authentication,omitempty"`
	APIAccess                           *APIAccess          `json:"apiAccess"`
	DelegateSelectors                   []string            `json:"delegateSelectors"`
	ExecuteOnDelegate                   *bool               `json:"executeOnDelegate,omitempty"`
	Type                                *FluffyType         `json:"type,omitempty"`
	Credential                          *Credential         `json:"credential,omitempty"`
	DockerRegistryURL                   *string             `json:"dockerRegistryUrl,omitempty"`
	ProviderType                        *ProviderType       `json:"providerType,omitempty"`
	Auth                                *PurpleAuth         `json:"auth,omitempty"`
	ConnectorRef                        *string             `json:"connectorRef"`
	FeaturesEnabled                     []FeaturesEnabled   `json:"featuresEnabled"`
	ArtifactoryServerURL                *string             `json:"artifactoryServerUrl,omitempty"`
	ApplicationKeyRef                   *string             `json:"applicationKeyRef,omitempty"`
	APIKeyRef                           *string             `json:"apiKeyRef"`
	NewRelicAccountID                   *string             `json:"newRelicAccountId,omitempty"`
	Username                            *string             `json:"username"`
	Accountname                         *string             `json:"accountname,omitempty"`
	ControllerURL                       *string             `json:"controllerUrl,omitempty"`
	PasswordRef                         *string             `json:"passwordRef"`
	ClientSecretRef                     *string             `json:"clientSecretRef"`
	ClientID                            *string             `json:"clientId"`
	AuthType                            *string             `json:"authType,omitempty"`
	HelmRepoURL                         *string             `json:"helmRepoUrl,omitempty"`
	JiraURL                             *string             `json:"jiraUrl,omitempty"`
	UsernameRef                         *string             `json:"usernameRef"`
	APITokenRef                         *string             `json:"apiTokenRef,omitempty"`
	ServiceNowURL                       *string             `json:"serviceNowUrl,omitempty"`
	Headers                             []interface{}       `json:"headers"`
	TenantID                            *string             `json:"tenantId,omitempty"`
	SubscriptionID                      *string             `json:"subscriptionId,omitempty"`
	BillingExportSpec                   *BillingExportSpec  `json:"billingExportSpec"`
	CrossAccountAccess                  *CrossAccountAccess `json:"crossAccountAccess,omitempty"`
	CurAttributes                       *CurAttributes      `json:"curAttributes"`
	AwsAccountID                        *string             `json:"awsAccountId,omitempty"`
	IsAWSGovCloudAccount                *bool               `json:"isAWSGovCloudAccount"`
	ProjectID                           *string             `json:"projectId,omitempty"`
	ServiceAccountEmail                 *string             `json:"serviceAccountEmail,omitempty"`
	Credentials                         interface{}         `json:"credentials"`
	Default                             *bool               `json:"default,omitempty"`
	AwsSDKClientBackOffStrategyOverride interface{}         `json:"awsSdkClientBackOffStrategyOverride"`
	Hosts                               []Host              `json:"hosts"`
	OnDelegate                          *bool               `json:"onDelegate,omitempty"`
	Host                                interface{}         `json:"host"`
	WorkingDirectory                    interface{}         `json:"workingDirectory"`
	Template                            *Template           `json:"template,omitempty"`
	AzureEnvironmentType                *string             `json:"azureEnvironmentType,omitempty"`
	APIKeyID                            interface{}         `json:"apiKeyId"`
	Region                              *string             `json:"region,omitempty"`
	SecretNamePrefix                    *string             `json:"secretNamePrefix,omitempty"`
}

type APIAccess struct {
	Type APIAccessType `json:"type"`
	Spec APIAccessSpec `json:"spec"`
}

type APIAccessSpec struct {
	TokenRef          *string     `json:"tokenRef,omitempty"`
	InstallationID    *string     `json:"installationId,omitempty"`
	ApplicationID     *string     `json:"applicationId,omitempty"`
	InstallationIDRef interface{} `json:"installationIdRef"`
	ApplicationIDRef  interface{} `json:"applicationIdRef"`
	PrivateKeyRef     *string     `json:"privateKeyRef,omitempty"`
}

type PurpleAuth struct {
	Type AuthType    `json:"type"`
	Spec *PurpleSpec `json:"spec,omitempty"`
}

type PurpleSpec struct {
	Username    *string `json:"username"`
	UsernameRef *string `json:"usernameRef"`
	PasswordRef string  `json:"passwordRef"`
}

type Authentication struct {
	Type AuthenticationType `json:"type"`
	Spec AuthenticationSpec `json:"spec"`
}

type AuthenticationSpec struct {
	Type      *PurpleType `json:"type,omitempty"`
	Spec      *SpecSpec   `json:"spec,omitempty"`
	SSHKeyRef *string     `json:"sshKeyRef,omitempty"`
}

type SpecSpec struct {
	Username     *string     `json:"username,omitempty"`
	UsernameRef  interface{} `json:"usernameRef"`
	TokenRef     *string     `json:"tokenRef,omitempty"`
	PasswordRef  *string     `json:"passwordRef,omitempty"`
	AccessKey    *string     `json:"accessKey,omitempty"`
	AccessKeyRef interface{} `json:"accessKeyRef"`
	SecretKeyRef *string     `json:"secretKeyRef,omitempty"`
}

type BillingExportSpec struct {
	StorageAccountName *string `json:"storageAccountName,omitempty"`
	ContainerName      *string `json:"containerName,omitempty"`
	DirectoryName      *string `json:"directoryName,omitempty"`
	ReportName         *string `json:"reportName,omitempty"`
	SubscriptionID     *string `json:"subscriptionId,omitempty"`
	DatasetID          *string `json:"datasetId,omitempty"`
	TableID            *string `json:"tableId,omitempty"`
}

type Credential struct {
	Type               CredentialType  `json:"type"`
	Spec               *CredentialSpec `json:"spec"`
	CrossAccountAccess interface{}     `json:"crossAccountAccess"`
	Region             *string         `json:"region"`
}

type CredentialSpec struct {
	MasterURL     *string     `json:"masterUrl,omitempty"`
	Auth          *FluffyAuth `json:"auth,omitempty"`
	SecretKeyRef  *string     `json:"secretKeyRef,omitempty"`
	AccessKey     *string     `json:"accessKey,omitempty"`
	AccessKeyRef  interface{} `json:"accessKeyRef"`
	ApplicationID *string     `json:"applicationId,omitempty"`
	TenantID      *string     `json:"tenantId,omitempty"`
	SecretKey     *string     `json:"secretKey,omitempty"`
}

type FluffyAuth struct {
	Type string      `json:"type"`
	Spec *FluffySpec `json:"spec,omitempty"`
}

type FluffySpec struct {
	Username               *string     `json:"username,omitempty"`
	UsernameRef            interface{} `json:"usernameRef"`
	PasswordRef            *string     `json:"passwordRef,omitempty"`
	ServiceAccountTokenRef *string     `json:"serviceAccountTokenRef,omitempty"`
	CACERTRef              interface{} `json:"caCertRef"`
	SecretRef              *string     `json:"secretRef,omitempty"`
}

type CrossAccountAccess struct {
	CrossAccountRoleArn string `json:"crossAccountRoleArn"`
	ExternalID          string `json:"externalId"`
}

type CurAttributes struct {
	ReportName   string `json:"reportName"`
	S3BucketName string `json:"s3BucketName"`
	Region       string `json:"region"`
	S3Prefix     string `json:"s3Prefix"`
}

type Host struct {
	Hostname       string         `json:"hostname"`
	HostAttributes TemplateInputs `json:"hostAttributes"`
}

type TemplateInputs struct {
}

type Template struct {
	TemplateRef    string         `json:"templateRef"`
	VersionLabel   string         `json:"versionLabel"`
	TemplateInputs TemplateInputs `json:"templateInputs"`
}

type Tags struct {
	Ccm   *string `json:"ccm,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Robin *string `json:"robin,omitempty"`
}

type EntityValidityDetails struct {
	Valid       bool        `json:"valid"`
	InvalidYAML interface{} `json:"invalidYaml"`
}

type StatusClass struct {
	Status          StatusEnum `json:"status"`
	ErrorSummary    *string    `json:"errorSummary"`
	Errors          []Error    `json:"errors"`
	TestedAt        int64      `json:"testedAt"`
	LastTestedAt    int64      `json:"lastTestedAt"`
	LastConnectedAt int64      `json:"lastConnectedAt"`
}

type Error struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type APIAccessType string

const (
	GithubApp   APIAccessType = "GithubApp"
	PurpleOAuth APIAccessType = "OAuth"
	Token       APIAccessType = "Token"
)

type AuthType string

const (
	Anonymous              AuthType = "Anonymous"
	PurpleUsernamePassword AuthType = "UsernamePassword"
)

type PurpleType string

const (
	AWSCredentials         PurpleType = "AWSCredentials"
	FluffyOAuth            PurpleType = "OAuth"
	FluffyUsernamePassword PurpleType = "UsernamePassword"
	UsernameToken          PurpleType = "UsernameToken"
)

type AuthenticationType string

const (
	HTTP  AuthenticationType = "Http"
	HTTPS AuthenticationType = "HTTPS"
	SSH   AuthenticationType = "Ssh"
)

type CredentialType string

const (
	InheritFromDelegate CredentialType = "InheritFromDelegate"
	ManualConfig        CredentialType = "ManualConfig"
)

type FeaturesEnabled string

const (
	Billing      FeaturesEnabled = "BILLING"
	Governance   FeaturesEnabled = "GOVERNANCE"
	Optimization FeaturesEnabled = "OPTIMIZATION"
	Visibility   FeaturesEnabled = "VISIBILITY"
)

type ProviderType string

const (
	DockerHub ProviderType = "DockerHub"
	Other     ProviderType = "Other"
)

type FluffyType string

const (
	Account FluffyType = "Account"
	Repo    FluffyType = "Repo"
)

type StatusEnum string

const (
	Failure StatusEnum = "FAILURE"
	Success StatusEnum = "SUCCESS"
)

func (c Connector) String() string {
	return c.Name
}

type ConnectorAPI interface {
	GetAllConnectors() ([]Connector, error)
}
