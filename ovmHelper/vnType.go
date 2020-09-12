package ovmhelper

// Vn - Virtual Nic Interface
type Vn struct {
	MacAddress string `json:"macAddress,omitempty"`
	//IpAddresses   *[]Id  `json:"ipAddresses,omitempty"`
	//InterfaceName string `json:"interfaceName,omitempty"`
	VMID      *ID    `json:"vmId,omitempty"`
	NetworkID *ID    `json:"networkId,omitempty"`
	ID        *ID    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	//Description      string `json:description,omitempty"`
	//Locked           bool   `json:locked,omitempty"`
	//ReadOnly         bool   `json:readOnly,omitempty"`
	Generation int `json:"generation,omitempty"`
	//ResourceGroupIds *[]Id  `json:resourceGroupIds,omitempty"`
}
