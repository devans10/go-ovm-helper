package ovmhelper

import (
	"fmt"
	"log"
)

// VmcsmService - Virtual Machine Clone Storage Mapping interface
type VmcsmService struct {
	client *Client
}

// Read - Read a virtual machine clone storage mapping object
func (v *VmcsmService) Read(vmcsmID string) (*Vmcsm, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneStorageMapping/%s", vmcsmID)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	m := &Vmcsm{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create - Create a virtual machine clone storage mapping
func (v *VmcsmService) Create(vmCloneDefinitionID string, vmcsm Vmcsm) (*string, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneStorageMapping", vmCloneDefinitionID)
	req, err := v.client.NewRequest("POST", url, nil, vmcsm)
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

// Delete - Delete a virtual machine clone storage mapping object
func (v *VmcsmService) Delete(vmCloneDefinitionID string, vmcsmID string) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s/VmCloneStorageMapping/%s", vmCloneDefinitionID, vmcsmID)
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
