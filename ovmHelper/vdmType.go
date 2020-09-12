package ovmhelper

// Vdm - vmDiskMapping element
// The VmDiskMapping object contains the relationships of the VirtualDisk and StorageElement for a Vm.
type Vdm struct {
	VMID                *ID    `json:"vmId,omitempty"`
	VirtualDiskID       *ID    `json:"virtualDiskId,omitempty"`
	DiskTarget          int    `json:"diskTarget"`
	EmulatedBlockDevice bool   `json:"emulatedBlockDevice,omitempty"` // false,
	StorageElementID    *ID    `json:"storageElementId,omitempty"`
	DiskWriteMode       string `json:"diskWriteMode,omitempty"` // "READ_ONLY",READ_WRITE
	ID                  *ID    `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`        //: "...",
	Description         string `json:"description,omitempty"` //  : "...",
	Locked              bool   `json:"locked,omitempty"`      // false,
	ReadOnly            bool   `json:"readOnly,omitempty"`    //false,
	Generation          int    `json:"generation,omitempty"`
	ResourceGroupIds    *[]ID  `json:"resourceGroupIds,omitempty"` //" : [ {  }, ... ]
	//	userData            // [ {"key" : "...", "value" : "..." }, ... ],
}

// ListVdm - slice of virtual disk mappings
var ListVdm []Vdm
