package almaclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAllServices returns list of Avengers
func (c Client) GetAllServices() ([]*Service, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	// var services []*Service
	var serviceCatalogue ServiceCatalog
	// err = json.Unmarshal(body, &services)
	err = json.Unmarshal(body, &serviceCatalogue)
	if err != nil {
		return nil, err
	}

	return serviceCatalogue.Services, nil
}
