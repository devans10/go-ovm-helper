package ovmhelper

// VM - Virtual Machine interface
type VM struct {
	ID                       *ID             `json:"id,omitempty"`
	Name                     string          `json:"name,omitempty"`
	AffinityGroupIDs         []*ID           `json:"affinityGroupIds,omitempty"`
	Architecture             string          `json:"architecture,omitempty"`
	BootOrder                []string        `json:"bootOrder,omitempty"`
	CPUCount                 int             `json:"cpuCount,omitempty"`
	CPUCountLimit            int             `json:"cpuCountLimit,omitempty"`
	CPUPriority              int             `json:"cpuPriority,omitempty"`
	CPUUtilizationCap        int             `json:"cpuUtilizationCap,omitempty"`
	CurrentMemory            int             `json:"currentMemory,omitempty"`
	CurrentCPUCount          int             `json:"currentCpuCount,omitempty"`
	CurrentDomainID          int             `json:"currentDomainId,omitempty"`
	Description              string          `json:"description,omitempty"`
	DiskLimit                int             `json:"diskLimit,omitempty"`
	Generation               int             `json:"generation,omitempty"`
	GuestDriverVersion       string          `json:"guestDriverVersion,omitempty"`
	HighAvailability         bool            `json:"highAvailability,omitempty"`
	HugePagesEnabled         bool            `json:"hugePagesEnabled,omitempty"`
	KernelVersion            string          `json:"kernelVersion,omitempty"`
	KeymapName               string          `json:"keymapName,omitempty"`
	Locked                   bool            `json:"locked,omitempty"`
	Memory                   int             `json:"memory,omitempty"`
	MemoryLimit              int             `json:"memoryLimit,omitempty"`
	NetworkInstallPath       string          `json:"networkInstallPath,omitempty"`
	Origin                   string          `json:"origin,omitempty"`
	OsType                   string          `json:"osType,omitempty"`
	OsVersion                string          `json:"osVersion,omitempty"`
	PinnedCpus               string          `json:"pinnedCpus,omitempty"`
	ReadOnly                 bool            `json:"readOnly,omitempty"`
	RepositoryID             *ID             `json:"repositoryId,omitempty"`
	ResourceGroupIds         []*ID           `json:"resourceGroupIds,omitempty"`
	RestartActionOnCrash     string          `json:"restartActionOnCrash,omitempty"`
	RestartActionOnPowerOff  string          `json:"restartActionOnPowerOff,omitempty"`
	RestartActionOnRestart   string          `json:"restartActionOnRestart,omitempty"`
	ServerID                 *ID             `json:"serverId,omitempty"`
	ServerPoolID             *ID             `json:"serverPoolId,omitempty"`
	SslVncPort               string          `json:"sslVncPort,omitempty"`
	SslTtyPort               string          `json:"sslTtyPort,omitempty"`
	UserData                 []*KeyValuePair `json:"userData,omitempty"`
	VirtualNicIDs            []*ID           `json:"virtualNicIds,omitempty"`
	VMApiVersion             string          `json:"vmApiVersion,omitempty"`
	VMCloneDefinitionIDs     []*ID           `json:"vmCloneDefinitionIds,omitempty"`
	VMConfigFileAbsolutePath string          `json:"vmConfigFileAbsolutePath,omitempty"`
	VMConfigFileMountedPath  string          `json:"vmConfigFileMountedPath,omitempty"`
	VMDiskMappingIds         []*ID           `json:"vmDiskMappingIds,omitempty"`
	VMDomainType             string          `json:"vmDomainType,omitempty"`
	VMMouseType              string          `json:"vmMouseType,omitempty"`
	VMRunState               string          `json:"vmRunState,omitempty"`
	VMStartPolicy            string          `json:"vmStartPolicy,omitempty"`
}

// CfgVM - Configuration interface for VM
type CfgVM struct {
	NetworkID    string
	SendMessages *[]KeyValuePair
	RootPassword *[]KeyValuePair
}

// KeyValuePair - interface for a key-value pair
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
