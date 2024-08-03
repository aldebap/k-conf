////////////////////////////////////////////////////////////////////////////////
//	service.go  -  Jul-5-2024  -  aldebap
//
//	Kong service configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
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

	fmt.Printf("[debug]: name = %s; URL = %s\n", newKongService.name, newKongService.url)
	payload, err := json.Marshal(KongServiceRequest{
		Name:    newKongService.name,
		Url:     newKongService.url,
		Enabled: newKongService.enabled,
	})
	if err != nil {
		return err
	}
	fmt.Printf("[debug]: payload = %s\n", payload)

	var kongUrl string = ks.address

	if ks.port != 0 {
		kongUrl = fmt.Sprintf("http://%s:%d/services", ks.address, ks.port)
	}

	//	json.NewEncoder(httpResponse).Encode(produtoResp)

	request, err := http.NewRequest("POST", kongUrl, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)

	return nil
}
