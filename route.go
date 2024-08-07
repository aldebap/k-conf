////////////////////////////////////////////////////////////////////////////////
//	route.go  -  Ago-6-2024  -  aldebap
//
//	Kong route configuration
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

// kong route attributes
type KongRoute struct {
	name      string
	protocols []string
	methods   []string
	paths     []string
	serviceId string
}

// create a new Kong route
func NewKongRoute(name string, protocols []string, methods []string, paths []string, serviceId string) *KongRoute {

	return &KongRoute{
		name:      name,
		protocols: protocols,
		methods:   methods,
		paths:     paths,
		serviceId: serviceId,
	}
}

// kong route request payload
type serviceId struct {
	Id string `json:"id"`
}

type KongRouteRequest struct {
	Name      string    `json:"name"`
	Protocols []string  `json:"protocols"`
	Methods   []string  `json:"methods"`
	Paths     []string  `json:"paths"`
	Service   serviceId `json:"service"`
}

// kong route response payload
type KongRouteResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Path     string `json:"path"`
}

// kong route list response payload
type KongRouteListResponse struct {
	Data []KongServiceResponse `json:"data"`
	Next string                `json:"next"`
}

// add a new route to Kong
func (ks *KongServer) AddRoute(newKongRoute *KongRoute, jsonOutput bool) error {

	var serviceURL string = fmt.Sprintf("%s/routes", ks.ServerURL())

	payload, err := json.Marshal(KongRouteRequest{
		Name:      newKongRoute.name,
		Protocols: newKongRoute.protocols,
		Methods:   newKongRoute.methods,
		Paths:     newKongRoute.paths,
		Service: serviceId{
			Id: newKongRoute.serviceId,
		},
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
		return errors.New("error sending add route command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var routeResp KongRouteResponse

		err = json.Unmarshal(respPayload, &routeResp)
		if err != nil {
			return err
		}

		fmt.Printf("%s\nnew route ID: %s\n", resp.Status, routeResp.Id)
	}

	return nil
}
