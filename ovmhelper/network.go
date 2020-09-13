package ovmhelper

import (
	"errors"
)

// NetworkService - interface for a ServerPool
type NetworkService struct {
	client *Client
}

// GetIDFromName - return the ID for a Network from the name
func (s *NetworkService) GetIDFromName(name string) (*ID, error) {
	req, err := s.client.NewRequest("GET", "/ovm/core/wsapi/rest/Network/id", nil, nil)
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

// Read - Read the Network object
func (s *NetworkService) Read(id string) (*ID, error) {
	req, err := s.client.NewRequest("GET", "/ovm/core/wsapi/rest/Network/"+id, nil, nil)
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
