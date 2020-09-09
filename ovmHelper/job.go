package ovmHelper

import (
	"time"
)

type JobService struct {
	client *Client
}

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

func (j *JobService) Running(id string) bool {

	job, _ := j.Read(id)

	if !job.Done {
		return true
	} else {
		return false
	}

}

func (j *JobService) WaitForJob(id string) {
	duration := time.Duration(500000)
	for j.Running(id) {
		time.Sleep(duration * time.Microsecond)
		if duration <= 5000000 {
			duration += 500000
		}
	}
}
