package ovmHelper

type Vm struct {
	Id   *Id    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	//The following fields of the new Vm are optional:
	//BootOrder          []string
	CpuCount                 int    `json:"cpuCount,omitempty"`
	CpuCountLimit            int    `json:"cpuCountLimit,omitempty"`
	CpuPriority              int    `json:"cpuPriority,omitempty"`
	CpuUtilizationCap        int    `json:"cpuUtilizationCap,omitempty"`
	HighAvailability         bool   `json:"highAvailability,omitempty"`
	HugePagesEnabled         bool   `json:"hugePagesEnabled,omitempty"`
	KeymapName               string `json:"keymapName,omitempty"`
	Memory                   int    `json:"memory,omitempty"`
	MemoryLimit              int    `json:"memoryLimit,omitempty"`
	NetworkInstallPath       string `json:"networkInstallPath,omitempty"`
	OsType                   string `json:"osType,omitempty"`
	VmDomainType             string `json:"vmDomainType,omitempty"`
	VmMouseType              string `json:"vmMouseType,omitempty"`
	VmRunState               string `json:"vmRunState,omitempty"`
	VmStartPolicy            string `json:"vmStartPolicy,omitempty"`
	SslVncPort               int    `json:"sslVncPort,omitempty"`
	SslTtyPort               string `json:"sslTtyPort,omitempty"`
	AffinityGroupIds         []int  `json:"affinityGroupIds,omitempty"`
	OsVersion                string `json:"osVersion,omitempty"`
	KernelVersion            string `json:"kernelVersion,omitempty"`
	GuestDriverVersion       string `json:"guestDriverVersion,omitempty"`
	VmApiVersion             string `json:"vmApiVersion,omitempty"`
	CurrentMemory            int    `json:"currentMemory,omitempty"`
	CurrentCpuCount          int    `json:"currentCpuCount,omitempty"`
	CurrentDomainId          int    `json:"currentDomainId,omitempty"`
	VmConfigFileAbsolutePath string `json:"vmConfigFileAbsolutePath,omitempty"`
	VmConfigFileMountedPath  string `json:"vmConfigFileMountedPath,omitempty"`
	Architecture             string `json:"architecture,omitempty"`
	VmDiskMappingIds         []*Id  `json:"vmDiskMappingIds,omitempty"`
	VirtualNicIds            []*Id  `json:"virtualNicIds,omitempty"`
	RepositoryId             *Id    `json:"repositoryId,omitempty"`
	ServerPoolId             *Id    `json:"serverPoolId,omitempty"`
	ServerId                 *Id    `json:"serverId,omitempty"`
	Generation               int    `json:"generation,omitempty"`
}

type CfgVm struct {
	NetworkId    string
	SendMessages *[]KeyValuePair
	RootPassword *[]KeyValuePair
}

type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
