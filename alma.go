package almaclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	v1 "gitlab.com/fahecom/platform/release-eng/service-catalog-app/api/service-catalog/v1"
	sc "gitlab.com/fahecom/platform/release-eng/service-catalog-app/pkg/model"
)

// GetAllServices returns list of Services
func (c Client) GetAllServices() ([]*sc.Service, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	// var services []*Service
	var serviceCatalogue sc.ServiceCatalog
	err = json.Unmarshal(body, &serviceCatalogue)
	if err != nil {
		return nil, err
	}

	return serviceCatalogue.Services, nil
}

// GetService returns meta of Service
func (c Client) GetService(service string) (*v1.GetServiceReply, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/%s", c.HostURL, service), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	// var services []*Service
	// var serviceCatalogue sc.Service
	var serviceCatalogue v1.GetServiceReply
	err = json.Unmarshal(body, &serviceCatalogue)
	if err != nil {
		return nil, err
	}

	return &serviceCatalogue, nil
}
