package ovmhelper

import (
	"fmt"
	"log"
)

// VnService - Virtual NIC interface
type VnService struct {
	client *Client
}

// Read - Read the virtual nic object
func (v *VnService) Read(id string) (*Vn, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/VirtualNic/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vn{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

// Create - Create the virtual nic object
func (v *VnService) Create(vmID string, vn Vn) (*string, error) {
	params := make(map[string]string)
	params["vmid"] = vmID
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VirtualNic", vmID)
	req, err := v.client.NewRequest("POST", url, params, vn)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] req: %v \n", req)

	m := &JobResponse{}

	_, err = v.client.Do(req, m)
	if err != nil {
		log.Printf("[ERROR] err: %v", err)
		return nil, err
	}

	v.client.Jobs.WaitForJob(m.ID.Value)
	j, _ := v.client.Jobs.Read(m.ID.Value)
	if !j.succeed() {
		return nil, j.Error
	}
	return &j.ResultID.Value, err
}

// Delete - Delete the virtual nic object
func (v *VnService) Delete(vnID string, vn Vn) error {
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Vm/%s/VirtualNic/%s", vn.VMID.Value, vnID)
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
	return err
}
