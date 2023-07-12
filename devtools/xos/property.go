package xos

// PropertyID ...
type PropertyID int

// Property ...
type Property struct {
	Base

	Name  string `json:"name"`
	Value string `json:"value"`
}
