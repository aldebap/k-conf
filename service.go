////////////////////////////////////////////////////////////////////////////////
//	service.go  -  Jul-5-2024  -  aldebap
//
//	Kong service configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// kong service attributes
type KongService struct {
	name    string
	url     string
	enabled bool
}

// kong service request payload
type KongServiceRequest struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Enabled bool   `json:"enabled"`
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
func (ks *KongServer) AddService(newKongService *KongService) error {

	fmt.Printf("[debug]: name = %s; URL = %s", newKongService.name, newKongService.url)
	//var kongUrl string = ks.address

	//_ = kongUrl
	// http.Post(kongUrl, "text/json")

	var kongUrl string = ks.address

	if ks.port != 0 {
		kongUrl = fmt.Sprintf("http://%s:%d", ks.address, ks.port)
	}

	//	json.NewEncoder(httpResponse).Encode(produtoResp)

	request, error := http.NewRequest("POST", kongUrl, bytes.NewBuffer([]byte("{}")))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)

	return nil
}
