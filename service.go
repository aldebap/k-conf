////////////////////////////////////////////////////////////////////////////////
//	service.go  -  Jul-5-2024  -  aldebap
//
//	Kong service configuration
////////////////////////////////////////////////////////////////////////////////

package main

// kong service attributes
type KongService struct {
	name    string
	url     string
	enabled bool
}

// kong service request payload
type KongServiceRequest struct {
	name    string `json:"name"`
	url     string `json:"url"`
	enabled bool   `json:"enabled"`
}

// create a new Kong service
func NewKongService(name string, url string, enabled bool) *KongService {

	return &KongService{
		name:    name,
		url:     url,
		enabled: enabled,
	}
}

// add a new service to Kong
func (ks *KongServer) AddService(newKongService KongService) {

	var kongUrl string = ks.address

	_ = kongUrl
	// http.Post(kongUrl, "text/json")
}
