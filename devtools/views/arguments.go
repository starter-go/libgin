package views

// Arguments 是 Arguments 的 VO
type Arguments struct {
	Base

	// Items []*xos.Argument `json:"arguments"`
	Args []string `json:"args"`
}
