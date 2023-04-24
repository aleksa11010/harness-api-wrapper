package harness

type Environment struct {
	ID              string `json:"connector_id"`
	Name            string `json:"name"`
	EnvironmentType string `json:"connector_type"`
}
