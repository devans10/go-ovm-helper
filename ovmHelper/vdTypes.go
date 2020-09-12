package ovmhelper

// Vd - Virtual Disk interface
type Vd struct {
	DiskType              string `json:"diskType,omitempty"` // "VIRTUAL_DISK",
	Size                  int    `json:"size,omitempty"`     //Bytes multiple of 4096
	OnDiskSize            int    `json:"onDiskSize,omitempty"`
	Path                  string `json:"path,omitempty"`
	VMDiskMappingIDs      []*ID  `json:"vmDiskMappingIds,omitempty"`
	RepositoryID          *ID    `json:"repositoryId,omitempty"`
	Shareable             bool   `json:"shareable,omitempty"`
	ImportFileName        string `json:"importFileName,omitempty"`
	AbsolutePath          string `json:"absolutePath,omitempty"`
	MountedPath           string `json:"mountedPath,omitempty"`
	AssemblyVirtualDiskID *ID    `json:"assemblyVirtualDiskId,omitempty"`
	ID                    *ID    `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	Description           string `json:"description,omitempty"`
	Locked                bool   `json:"locked,omitempty"`   //false,
	ReadOnly              bool   `json:"readOnly,omitempty"` //: false,
	Generation            int    `json:"generation,omitempty"`
	/*"userData" : [ {
	  "key" : "...",
	  "value" : "..."
	}, ... ],*/
	ResourceGroupIds []*ID `json:"resourceGroupIds,omitempty"`
}
