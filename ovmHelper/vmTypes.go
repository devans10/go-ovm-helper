package ovmHelper

type Vm struct {
	Id   *Id    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	//The following fields of the new Vm are optional:
	AffinityGroupIds         []*Id          `json:"affinityGroupIds,omitempty"`
	Architecture             string         `json:"architecture,omitempty"`
	BootOrder                []string       `json:"bootOrder,omitempty"`
	CpuCount                 int            `json:"cpuCount,omitempty"`
	CpuCountLimit            int            `json:"cpuCountLimit,omitempty"`
	CpuPriority              int            `json:"cpuPriority,omitempty"`
	CpuUtilizationCap        int            `json:"cpuUtilizationCap,omitempty"`
	CurrentMemory            int            `json:"currentMemory,omitempty"`
	CurrentCpuCount          int            `json:"currentCpuCount,omitempty"`
	CurrentDomainId          int            `json:"currentDomainId,omitempty"`
	Description              string         `json:"description,omitempty"`
	DiskLimit                int            `json:"diskLimit,omitempty"`
	Generation               int            `json:"generation,omitempty"`
	GuestDriverVersion       string         `json:"guestDriverVersion,omitempty"`
	HighAvailability         bool           `json:"highAvailability,omitempty"`
	HugePagesEnabled         bool           `json:"hugePagesEnabled,omitempty"`
	KernelVersion            string         `json:"kernelVersion,omitempty"`
	KeymapName               string         `json:"keymapName,omitempty"`
	Locked                   bool           `json:"locked,omitempty"`
	Memory                   int            `json:"memory,omitempty"`
	MemoryLimit              int            `json:"memoryLimit,omitempty"`
	NetworkInstallPath       string         `json:"networkInstallPath,omitempty"`
	Origin                   string         `json:"origin,omitempty"`
	OsType                   string         `json:"osType,omitempty"`
	OsVersion                string         `json:"osVersion,omitempty"`
	PinnedCpus               string         `json:"pinnedCpus,omitempty"`
	ReadOnly                 bool           `json:"readOnly,omitempty"`
	RepositoryId             *Id            `json:"repositoryId,omitempty"`
	ResourceGroupIds         []*Id          `json:"resourceGroupIds,omitempty"`
	RestartActionOnCrash     string         `json:"restartActionOnCrash,omitempty"`
	RestartActionOnPowerOff  string         `json:"restartActionOnPowerOff,omitempty"`
	RestartActionOnRestart   string         `json:"restartActionOnRestart,omitempty"`
	ServerId                 *Id            `json:"serverId,omitempty"`
	ServerPoolId             *Id            `json:"serverPoolId,omitempty"`
	SslVncPort               string         `json:"sslVncPort,omitempty"`
	SslTtyPort               string         `json:"sslTtyPort,omitempty"`
	UserData                 []KeyValuePair `json:"userData,omitempty"`
	VirtualNicIds            []*Id          `json:"virtualNicIds,omitempty"`
	VmApiVersion             string         `json:"vmApiVersion,omitempty"`
	VmCloneDefinitionIds     []*Id          `json:"vmCloneDefinitionIds,omitempty"`
	VmConfigFileAbsolutePath string         `json:"vmConfigFileAbsolutePath,omitempty"`
	VmConfigFileMountedPath  string         `json:"vmConfigFileMountedPath,omitempty"`
	VmDiskMappingIds         []*Id          `json:"vmDiskMappingIds,omitempty"`
	VmDomainType             string         `json:"vmDomainType,omitempty"`
	VmMouseType              string         `json:"vmMouseType,omitempty"`
	VmRunState               string         `json:"vmRunState,omitempty"`
	VmStartPolicy            string         `json:"vmStartPolicy,omitempty"`
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
