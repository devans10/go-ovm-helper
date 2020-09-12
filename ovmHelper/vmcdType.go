package ovmhelper

// Vmcd - VM Clone Definition
type Vmcd struct {
	VMID                     *ID    `json:"vmId,omitempty"`
	VMCloneNetworkMappingIDs *[]ID  `json:"vmCloneNetworkMappingIds,omitempty"`
	VMCloneStorageMappingIDs *[]ID  `json:"vmCloneStorageMappingIds,omitempty"`
	ID                       *ID    `json:"id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	Locked                   bool   `json:"locked,imitempty"`
	ReadOnly                 bool   `json:"readOnly,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	ResourceGroupIds         *[]ID  `json:"resourceGroupIds,omitempty"`
	/*"userData" : [ {
	    "key" : "...",
	    "value" : "..."
	  }, ... ],
	*/

}
