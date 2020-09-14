package ovmhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

// VMService - Virtual Machine Interface
type VMService struct {
	client *Client
}

// GetIDFromName - Return the ID of a virtual machine from its name
func (v *VMService) GetIDFromName(name string) (*ID, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/id", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []ID{}
	_, err = v.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	for _, id := range m {
		if id.Name == name {
			returnID := id
			return &returnID, nil
		}
	}

	return nil, errors.New("[error] Failed to find id for " + name)
}

func (v *VMService) Read(id string) (*VM, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &VM{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

// Stop - Stop a virtual machine
func (v *VMService) Stop(id string) error {
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+id+"/stop", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}

// Start - Start a virtual machine
func (v *VMService) Start(id string) error {
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+id+"/start", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return err
	}
	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return j.Error
	}

	return err
}

// CreateVM - Create a new virtual machine
func (v *VMService) CreateVM(vm VM, cfgVM CfgVM) (*string, error) {
	req, err := v.client.NewRequest("POST", "/ovm/core/wsapi/rest/Vm", nil, vm)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] vmrequest: %v", req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}
	v.client.Jobs.WaitForJob(m.ID.Value)

	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return nil, j.Error
	}

	if cfgVM.NetworkID != "" {
		vn := Vn{NetworkID: &ID{Type: "com.oracle.ovm.mgr.ws.model.Network",
			Value: cfgVM.NetworkID},
			VMID: &ID{Type: "com.oracle.ovm.mgr.ws.model.Vm",
				Value: j.ResultID.Value,
			},
		}
		_, err = v.client.Vns.Create(j.ResultID.Value, vn)
		if err != nil {
			return &j.ResultID.Value, err
		}
	}

	return &j.ResultID.Value, err
}

// CloneVM - Clone a VM from another VM or Template
func (v *VMService) CloneVM(cloneVMID string, vmCloneDefinitionID string, vm VM, cfgVM CfgVM) (*string, error) {

	params := make(map[string]string)
	params["vmId"] = cloneVMID
	params["serverPoolId"] = vm.ServerPoolID.Value
	params["repositoryId"] = vm.RepositoryID.Value
	if vmCloneDefinitionID != "" {
		params["vmCloneDefinitionId"] = vmCloneDefinitionID
	}

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/clone", cloneVMID)
	req, err := v.client.NewRequest("PUT", url, params, nil)
	if err != nil {
		return nil, err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if err != nil {
		return nil, err
	}

	jsonJobResult, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jsonJobResult))
	vm.ID = &ID{Value: j.ResultID.Value,
		Type: "com.oracle.ovm.mgr.ws.model.Vm"}

	log.Printf("[INFO] Update vmId: %s", j.ResultID.Value)
	err = v.client.Vms.UpdateVM(j.ResultID.Value, vm)
	if err != nil {
		return nil, err
	}

	return &j.ResultID.Value, err
}

// UpdateVM - Update the settings of a VM
func (v *VMService) UpdateVM(vmID string, vm VM) error {
	p := make(map[string]string)

	p["vmId"] = vmID
	rVM, _ := v.client.Vms.Read(vmID)

	rVM.Name = vm.Name
	rVM.Description = vm.Description
	rVM.BootOrder = vm.BootOrder
	rVM.CPUCount = vm.CPUCount
	rVM.CPUCountLimit = vm.CPUCountLimit
	rVM.CPUPriority = vm.CPUPriority
	rVM.CPUUtilizationCap = vm.CPUUtilizationCap
	rVM.HighAvailability = vm.HighAvailability
	rVM.HugePagesEnabled = vm.HugePagesEnabled
	rVM.KeymapName = vm.KeymapName
	rVM.Memory = vm.Memory
	rVM.MemoryLimit = vm.MemoryLimit
	rVM.NetworkInstallPath = vm.NetworkInstallPath
	rVM.OsType = vm.OsType
	rVM.ServerID = vm.ServerID
	rVM.VMDomainType = vm.VMDomainType
	rVM.VMMouseType = vm.VMMouseType
	rVM.VMRunState = vm.VMRunState
	rVM.VMStartPolicy = vm.VMStartPolicy
	rVM.RestartActionOnCrash = vm.RestartActionOnCrash

	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+vmID, p, rVM)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)

	jobJSON, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJSON))
	return err
}

// DeleteVM - Delete a Virtual Machine
func (v *VMService) DeleteVM(vmID string) error {

	req, err := v.client.NewRequest("DELETE", "/ovm/core/wsapi/rest/Vm/"+vmID, nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return j.Error
	}

	return err
}

// SendMessageToVM - Send a message to a virtual machine via OVMD daemon
func (v *VMService) SendMessageToVM(vmID string, cfgVM CfgVM) error {
	time.Sleep(30 * time.Second)
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/sendMessage", vmID)
	req, err := v.client.NewRequest("PUT", url, nil, cfgVM.SendMessages)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return err
	}
	v.client.Jobs.WaitForJob(m.ID.Value)

	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}

// SendRootPasswordToVM - Send the root password to a vm
func (v *VMService) SendRootPasswordToVM(vmID string, cfgVM CfgVM) error {
	time.Sleep(5 * time.Second)
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/sendMessage", vmID)
	req, err := v.client.NewRequest("PUT", url, nil, cfgVM.RootPassword)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return err
	}
	v.client.Jobs.WaitForJob(m.ID.Value)

	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}
