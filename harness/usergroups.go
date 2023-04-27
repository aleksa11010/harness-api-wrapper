package harness

type UserGroup struct {
	CorrelationID string `json:"correlationId"`
	Data          struct {
		Content []struct {
			AccountIdentifier   string        `json:"accountIdentifier"`
			Description         string        `json:"description"`
			ExternallyManaged   bool          `json:"externallyManaged"`
			HarnessManaged      bool          `json:"harnessManaged"`
			Identifier          string        `json:"identifier"`
			Name                string        `json:"name"`
			NotificationConfigs []interface{} `json:"notificationConfigs"`
			SsoLinked           bool          `json:"ssoLinked"`
			Tags                struct{}      `json:"tags"`
			Users               []string      `json:"users"`
		} `json:"content"`
		Empty         bool        `json:"empty"`
		PageIndex     int64       `json:"pageIndex"`
		PageItemCount int64       `json:"pageItemCount"`
		PageSize      int64       `json:"pageSize"`
		PageToken     interface{} `json:"pageToken"`
		TotalItems    int64       `json:"totalItems"`
		TotalPages    int64       `json:"totalPages"`
	} `json:"data"`
	MetaData interface{} `json:"metaData"`
	Status   string      `json:"status"`
}

type UserGroupDetails struct {
	CorrelationID string `json:"correlationId"`
	Data          struct {
		Content []struct {
			Disabled                       bool   `json:"disabled"`
			Email                          string `json:"email"`
			ExternallyManaged              bool   `json:"externallyManaged"`
			Locked                         bool   `json:"locked"`
			Name                           string `json:"name"`
			TwoFactorAuthenticationEnabled bool   `json:"twoFactorAuthenticationEnabled"`
			UUID                           string `json:"uuid"`
		} `json:"content"`
		Empty         bool        `json:"empty"`
		PageIndex     int64       `json:"pageIndex"`
		PageItemCount int64       `json:"pageItemCount"`
		PageSize      int64       `json:"pageSize"`
		PageToken     interface{} `json:"pageToken"`
		TotalItems    int64       `json:"totalItems"`
		TotalPages    int64       `json:"totalPages"`
	} `json:"data"`
	MetaData interface{} `json:"metaData"`
	Status   string      `json:"status"`
}
