package ovmHelper

import (
	"errors"
)

type RepoService struct {
	client *Client
}

func (r *RepoService) GetIdFromName(name string) (*Id, error) {
	req, err := r.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository/id", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Id{}
	_, err = r.client.Do(req, m)

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

func (r *RepoService) Read(id string) (*Id, error) {
	req, err := r.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Id{}
	_, err = r.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}
