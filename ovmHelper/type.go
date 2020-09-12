package ovmhelper

// ID - Interface for a SimpleID object
type ID struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	URI   string `json:"uri,omitempty"`
	Name  string `json:"name,omitempty"`
}
