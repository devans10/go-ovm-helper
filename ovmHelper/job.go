package ovmhelper

import (
	"time"
)

// JobService - client to check OVM Manager Jobs
type JobService struct {
	client *Client
}

//  Read - Read the job Uri
func (j *JobService) Read(id string) (*JobResponse, error) {
	req, err := j.client.NewRequest("GET", "/ovm/core/wsapi/rest/Job/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &JobResponse{}
	_, err = j.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

// Running - check if a job is still running
func (j *JobService) Running(id string) bool {

	job, _ := j.Read(id)

	if !job.Done {
		return true
	}
	return false
}

// WaitForJob - function to wait for job to complete
func (j *JobService) WaitForJob(id string) {
	time.Sleep(1 * time.Second)
	for j.Running(id) {
		time.Sleep(5 * time.Second)
	}
}
