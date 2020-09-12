package ovmhelper

import (
	"fmt"
	"log"
)

// VdmService - Virtual Disk Mapping interface
type VdmService struct {
	client *Client
}

// Read - Read the Virtual Disk Mapping object
func (v *VdmService) Read(vmID string, vdmID string) (*Vdm, error) {

	listOfVdm, err := v.List(vmID)
	if err != nil {
		return nil, err
	}

	for _, v := range *listOfVdm {
		if vdmID == v.ID.Value {
			log.Printf("[DEBUG] Find VDM id: %v", vdmID)
			return &v, nil
		}
	}

	log.Printf("[DEBUG] Read VDM found no mapping")
	return nil, nil

}

// List - List the Virtual Disk Mappings
func (v *VdmService) List(vmID string) (*[]Vdm, error) {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmDiskMapping", vmID)
	req, err := v.client.NewRequest("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &ListVdm
	_, err = v.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Create - Create a Virtual Disk Mapping
func (v *VdmService) Create(vdm Vdm) (*string, error) {

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmDiskMapping", vdm.VMID.Value)
	req, err := v.client.NewRequest("POST", url, nil, vdm)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	log.Printf("[DEBUG] %v", req)
	m := &JobResponse{}
	//m := JobResponse{}

	_, err = v.client.Do(req, &m)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return nil, j.Error
	}
	return &j.ResultID.Value, nil
}

// Delete - Delete a Virtual Disk Mapping
func (v *VdmService) Delete(vmID string, vdmID string) error {

	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VmDiskMapping/%s", vmID, vdmID)
	req, err := v.client.NewRequest("DELETE", url, nil, nil)
	if err != nil {
		fmt.Println("error")
		return err
	}

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
