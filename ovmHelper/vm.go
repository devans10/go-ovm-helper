package ovmHelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type VmService struct {
	client *Client
}

func (v *VmService) GetIdFromName(name string) (*Id, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/id", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Id{}
	_, err = v.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	for _, id := range m {
		if id.Name == name {
			returnId := id
			return &returnId, nil
		}
	}

	return nil, errors.New("[error] Failed to find id for " + name)
}

func (v *VmService) Read(id string) (*Vm, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/Vm/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vm{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

func (v *VmService) Stop(id string) error {
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+id+"/stop", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}

func (v *VmService) Start(id string) error {
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+id+"/start", nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}

	return err
}

func (v *VmService) CreateVm(vm Vm, cfgVm CfgVm) (*string, error) {
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
	v.client.Jobs.WaitForJob(m.Id.Value)

	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return nil, j.Error
	}

	if cfgVm.NetworkId != "" {
		vn := Vn{NetworkId: &Id{Type: "com.oracle.ovm.mgr.ws.model.Network",
			Value: cfgVm.NetworkId},
			VmId: &Id{Type: "com.oracle.ovm.mgr.ws.model.Vm",
				Value: j.ResultId.Value,
			},
		}
		_, err = v.client.Vns.Create(j.ResultId.Value, vn)
		if err != nil {
			return &j.ResultId.Value, err
		}

		//log.Printf("[DEBUG] cfgvm: %s", cfgVm)

	}

	return &j.ResultId.Value, err
}

func (v *VmService) CloneVm(cloneVmId string, vmCloneDefinitionId string, vm Vm, cfgVm CfgVm) (*string, error) {

	params := make(map[string]string)
	params["vmId"] = cloneVmId
	params["serverPoolId"] = vm.ServerPoolId.Value
	params["repositoryId"] = vm.RepositoryId.Value
	if vmCloneDefinitionId != "" {
		params["vmCloneDefinitionId"] = vmCloneDefinitionId
	}

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/clone", cloneVmId)
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

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if err != nil {
		return nil, err
	}

	jsonJobResult, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jsonJobResult))
	vm.Id = &Id{Value: j.ResultId.Value,
		Type: "com.oracle.ovm.mgr.ws.model.Vm"}

	log.Printf("[INFO] Update vmId: %s", j.ResultId.Value)
	err = v.client.Vms.UpdateVm(j.ResultId.Value, vm)
	if err != nil {
		return nil, err
	}

	err = v.client.Vms.Start(j.ResultId.Value)
	if err != nil {
		return nil, err
	}

	err = v.client.Vms.SendMessageToVm(j.ResultId.Value, cfgVm)
	if err != nil {
		fmt.Println(err)
	}

	err = v.client.Vms.SendRootPasswordToVm(j.ResultId.Value, cfgVm)
	if err != nil {
		fmt.Println(err)
	}
	return &j.ResultId.Value, err
}

func (v *VmService) UpdateVm(vmId string, vm Vm) error {
	p := make(map[string]string)

	p["vmId"] = vmId
	rVm, _ := v.client.Vms.Read(vmId)

	rVm.Name = vm.Name
	rVm.Description = vm.Description
	rVm.BootOrder = vm.BootOrder
	rVm.CpuCount = vm.CpuCount
	rVm.CpuCountLimit = vm.CpuCountLimit
	rVm.CpuPriority = vm.CpuPriority
	rVm.CpuUtilizationCap = vm.CpuUtilizationCap
	rVm.HighAvailability = vm.HighAvailability
	rVm.HugePagesEnabled = vm.HugePagesEnabled
	rVm.KeymapName = vm.KeymapName
	rVm.Memory = vm.Memory
	rVm.MemoryLimit = vm.MemoryLimit
	rVm.NetworkInstallPath = vm.NetworkInstallPath
	rVm.OsType = vm.OsType
	rVm.ServerId = vm.ServerId
	rVm.VmDomainType = vm.VmDomainType
	rVm.VmMouseType = vm.VmMouseType
	rVm.VmRunState = vm.VmRunState
	rVm.VmStartPolicy = vm.VmStartPolicy
	rVm.RestartActionOnCrash = vm.RestartActionOnCrash	
	
	req, err := v.client.NewRequest("PUT", "/ovm/core/wsapi/rest/Vm/"+vmId, p, rVm)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)

	jobJson, _ := json.Marshal(j)
	log.Printf("[DEBUG] %v", string(jobJson))
	return err
}

func (v *VmService) DeleteVm(vmId string) error {

	req, err := v.client.NewRequest("DELETE", "/ovm/core/wsapi/rest/Vm/"+vmId, nil, nil)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println("inside error")
		return err
	}

	v.client.Jobs.WaitForJob(m.Id.Value)
	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}

	return err
}

func (v *VmService) SendMessageToVm(vmId string, cfgVm CfgVm) error {
	time.Sleep(30 * time.Second)
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/sendMessage", vmId)
	req, err := v.client.NewRequest("PUT", url, nil, cfgVm.SendMessages)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)

	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}

func (v *VmService) SendRootPasswordToVm(vmId string, cfgVm CfgVm) error {
	time.Sleep(5 * time.Second)
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/sendMessage", vmId)
	req, err := v.client.NewRequest("PUT", url, nil, cfgVm.RootPassword)
	if err != nil {
		return err
	}

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return err
	}
	v.client.Jobs.WaitForJob(m.Id.Value)

	j, _ := v.client.Jobs.Read(m.Id.Value)
	if !j.succeed() {
		return j.Error
	}
	return err
}
