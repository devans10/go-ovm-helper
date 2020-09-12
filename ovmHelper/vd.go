package ovmhelper

import (
	"fmt"
	"log"
	"strconv"
)

// VdService - Virtual Disk interface
type VdService struct {
	client *Client
}

// Read - Read the Virtual disk URI
func (v *VdService) Read(id string) (*Vd, error) {
	req, err := v.client.NewRequest("GET", "/ovm/core/wsapi/rest/VirtualDisk/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Vd{}
	_, err = v.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}

// Create - Create a Virtual Disk
func (v *VdService) Create(repositoryID string, sparse bool, vd Vd) (*string, error) {
	params := make(map[string]string)
	params["repositoryId"] = repositoryID
	params["sparse"] = strconv.FormatBool(sparse)
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Repository/%s/VirtualDisk", repositoryID)
	req, err := v.client.NewRequest("POST", url, params, vd)
	if err != nil {
		fmt.Println("error")
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

// Update - Update a Virtual Disk
func (v *VdService) Update(vdID string, vd Vd) error {
	params := make(map[string]string)
	params["virtualDiskId"] = vdID

	rVd, _ := v.client.Vds.Read(vdID)
	if rVd.Name != vd.Name {
		rVd.Name = vd.Name
	}
	if rVd.Description != vd.Description {
		rVd.Description = vd.Description
	}
	if rVd.Shareable != vd.Shareable {
		rVd.Shareable = vd.Shareable
	}

	url := fmt.Sprintf("/ovm/core/wsapi/rest/VirtualDisk/%s", vdID)
	req, err := v.client.NewRequest("PUT", url, params, rVd)
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

// Delete - Delete a Virtual disk
func (v *VdService) Delete(repositoryID string, vdID string) error {
	params := make(map[string]string)
	params["repositoryId"] = repositoryID
	params["virtualDiskId"] = vdID
	url := fmt.Sprintf("/ovm/core/wsapi/rest/Repository/%s/VirtualDisk/%s", repositoryID, vdID)
	req, err := v.client.NewRequest("DELETE", url, params, nil)
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
