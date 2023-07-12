package libgin

// Connector ...
type Connector interface {
	Registration() *ConnectorRegistration
}

// ConnectorRegistration ...
type ConnectorRegistration struct {
	Name      string
	Connector Connector
}
