////////////////////////////////////////////////////////////////////////////////
//	service.go  -  Jul-5-2024  -  aldebap
//
//	Kong service configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// kong service attributes
type KongService struct {
	name    string
	url     string
	enabled bool
}

// create a new Kong service
func NewKongService(name string, url string, enabled bool) *KongService {

	return &KongService{
		name:    name,
		url:     url,
		enabled: enabled,
	}
}

// kong service request payload
type KongServiceRequest struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Enabled bool   `json:"enabled"`
}

// kong service response payload
type KongServiceResponse struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Protocol           string `json:"protocol"`
	Port               int    `json:"port"`
	Host               string `json:"host"`
	Path               string `json:"path"`
	CACertificates     string `json:"ca_certificates"`
	ClientCertificates string `json:"client_certificates"`
	Tags               string `json:"tags"`
	Enabled            bool   `json:"enabled"`
}

// add a new service to Kong
func (ks *KongServer) AddService(newKongService *KongService, jsonOutput bool) error {

	var serviceURL string = fmt.Sprintf("%s/services", ks.ServerURL())

	payload, err := json.Marshal(KongServiceRequest{
		Name:    newKongService.name,
		Url:     newKongService.url,
		Enabled: newKongService.enabled,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("error sending add service command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var serviceResp KongServiceResponse

		err = json.Unmarshal(respPayload, &serviceResp)
		if err != nil {
			return err
		}

		fmt.Printf("%s\nnew service ID: %s\n", resp.Status, serviceResp.Id)
	}

	return nil
}

// query a service by Id
func (ks *KongServer) QueryService(id string, jsonOutput bool) error {

	var serviceURL string = fmt.Sprintf("%s/services/%s", ks.ServerURL(), id)

	//	send a request to Kong to query the service by id
	resp, err := http.Get(serviceURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("service not found for the id: " + id)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error sending query service command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var serviceResp KongServiceResponse

		err = json.Unmarshal(respPayload, &serviceResp)
		if err != nil {
			return err
		}

		fmt.Printf("%s\nservice: %s --> %s://%s:%d%s\n", resp.Status,
			serviceResp.Name, serviceResp.Protocol, serviceResp.Host, serviceResp.Port, serviceResp.Path)
	}

	return nil
}
