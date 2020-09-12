package ovmhelper

import (
	"fmt"
	"log"
)

// VmcdService - Virtual Machice Clone Definition interface
type VmcdService struct {
	client *Client
}

// Read - Read the Virtual Machine Clone Definition
func (v *VmcdService) Read(vmcdID string) (*Vmcd, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/VmCloneDefinition/%s", vmcdID)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	m := &Vmcd{}
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create - Create a Virtual Machine Clone Definition
func (v *VmcdService) Create(VMID string, vmcd Vmcd) (*string, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmCloneDefinition", VMID)
	req, err := v.client.NewRequest("POST", url, nil, vmcd)
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
	return &j.ResultID.Value, err
}

// Delete - Delete a Virtual Machine Clone Definition
func (v *VmcdService) Delete(vmID string, vmCloneDefinitionID string) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmCloneDefinition/%s", vmID, vmCloneDefinitionID)
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
