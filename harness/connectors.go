package harness

type Connector struct {
	ID            string `json:"connector_id"`
	Name          string `json:"name"`
	ConnectorType string `json:"connector_type"`
}

func (c *Connector) GetID() string {
	return c.ID
}

func (c *Connector) GetName() string {
	return c.Name
}

func (c *Connector) GetConnectorType() string {
	return c.ConnectorType
}
