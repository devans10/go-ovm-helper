package ovmhelper

import (
	"errors"
)

// RepoService - Interface for a Repository
type RepoService struct {
	client *Client
}

// GetIDFromName - return the repository ID from the name
func (r *RepoService) GetIDFromName(name string) (*ID, error) {
	req, err := r.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository/id", nil, nil)
	if err != nil {
		return nil, err
	}

	m := []ID{}
	_, err = r.client.Do(req, &m)

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

func (r *RepoService) Read(id string) (*ID, error) {
	req, err := r.client.NewRequest("GET", "/ovm/core/wsapi/rest/Repository/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	m := &ID{}
	_, err = r.client.Do(req, m)

	if err != nil {
		return nil, err
	}

	return m, err
}
