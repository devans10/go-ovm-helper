package ovmhelper

import (
	"fmt"
	"log"
)

// VmcnmService - Virtual Machine Clone Network Mapping interface
type VmcnmService struct {
	client *Client
}

// Read - Read a Virtual Machine Clone Network Mapping object
func (v *VmcnmService) Read(vmcnmID string) (*Vmcnm, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneNetworkMapping/%s", vmcnmID)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	m := &Vmcnm{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create - Create a virtual machine clone network mapping
func (v *VmcnmService) Create(vmCloneDefinitionID string, vmcnm Vmcnm) (*string, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneNetworkMapping", vmCloneDefinitionID)
	req, err := v.client.NewRequest("POST", url, nil, vmcnm)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] %v", req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return nil, j.Error
	}
	log.Printf("[DEBUG] %v", j)
	return &j.ResultID.Value, err
}

// Delete - Delete a virtual machine network mapping
func (v *VmcnmService) Delete(vmCloneDefinitionID string, vmcnmID string) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneNetworkMapping/%s", vmCloneDefinitionID, vmcnmID)
	req, err := v.client.NewRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] %v", req)

	m := &JobResponse{}
	_, err = v.client.Do(req, m)
	if err != nil {
		fmt.Println(err)
		return err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return j.Error
	}
	return nil
}
