package xos

// EnvID ...
type EnvID int

// Env ...
type Env struct {
	Base

	Name  string `json:"name"`
	Value string `json:"value"`
}
