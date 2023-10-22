package libgin

// Connector ...
type Connector interface {
	// Registration( ... ) *ConnectorRegistration
}

// ConnectorRegistry ...
type ConnectorRegistry interface {
	ListRegistrations() []*ConnectorRegistration
}

// ConnectorRegistration ...
type ConnectorRegistration struct {
	Name      string
	Connector Connector
}
