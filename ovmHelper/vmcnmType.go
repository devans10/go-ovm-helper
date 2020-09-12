package ovmhelper

// Vmcnm - VM Clone Network Mapping
type Vmcnm struct {
	NetworkID           *ID    `json:"networkId,omitempty"`
	VMCloneDefinitionID *ID    `json:"vmCloneDefinitionId,omitempty"`
	VirtualNicID        *ID    `json:"virtualNicId,omitempty"`
	ID                  *ID    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Locked              bool   `json:"locked,imitempty"`
	ReadOnly            bool   `json:"readOnly,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	ResourceGroupIDs    *[]ID  `json:"resourceGroupIds,omitempty"`
}
