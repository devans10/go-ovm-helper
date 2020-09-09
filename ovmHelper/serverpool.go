package ovmHelper

import (
	"errors"
)

type ServerPoolService struct {
	client *Client
}

func (s *ServerPoolService) GetIdFromName(name string) (*Id, error) {
	req, err := s.client.NewRequest("GET", "/ovm/core/wsapi/rest/ServerPool/id", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []Id{}
	_, err = s.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	for _, r := range m {
		if r.Name == name {
			returnId := r
			return &returnId, nil
		}
	}

	return nil, errors.New("[error] Failed to find id for " + name)
}

func (s *ServerPoolService) Read(id string) (*Id, error) {
	req, err := s.client.NewRequest("GET", "/ovm/core/wsapi/rest/ServerPool/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &Id{}
	_, err = s.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}
