package ovmhelper

// Vmcsm - VM Clone Storage Mapping
type Vmcsm struct {
	VMDiskMappingID     *ID    `json:"vmDiskMappingId,omitempty"`
	VMCloneDefinitionID *ID    `json:"vmCloneDefinitionId,omitempty"`
	RepositoryID        *ID    `json:"repositoryId,omitempty"`
	StorageArrayID      *ID    `json:"storageArrayId,omitempty"`
	StorageElementID    *ID    `json:",omitempty"`
	ID                  *ID    `json:"id,omitempty"`
	CloneType           string `json:"cloneType,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Locked              bool   `json:"locked,imitempty"`
	ReadOnly            bool   `json:"readOnly,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	ResourceGroupIds    *[]ID  `json:"resourceGroupIds,omitempty"`
}
