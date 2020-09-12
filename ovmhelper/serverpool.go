package ovmhelper

import (
	"errors"
)

// ServerPoolService - interface for a ServerPool
type ServerPoolService struct {
	client *Client
}

// GetIDFromName - return the ID for a ServerPool from the name
func (s *ServerPoolService) GetIDFromName(name string) (*ID, error) {
	req, err := s.client.NewRequest("GET", "/ovm/core/wsapi/rest/ServerPool/id", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []ID{}
	_, err = s.client.Do(req, &m)

	if err != nil {
		return nil, err
	}

	for _, r := range m {
		if r.Name == name {
			returnID := r
			return &returnID, nil
		}
	}

	return nil, errors.New("[error] Failed to find id for " + name)
}

func (s *ServerPoolService) Read(id string) (*ID, error) {
	req, err := s.client.NewRequest("GET", "/ovm/core/wsapi/rest/ServerPool/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &ID{}
	_, err = s.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}
