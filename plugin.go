////////////////////////////////////////////////////////////////////////////////
//	plugin.go  -  Sep-18-2024  -  aldebap
//
//	Kong plugin configuration
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

// kong plugin config attributes
type KongPluginConfig struct {
	key   string
	value string
}

// kong plugin attributes
type KongPlugin struct {
	instanceName string
	name         string
	serviceId    string
	routeId      string
	consumer     string
	config       []KongPluginConfig
	protocols    []string
	enabled      bool
	tags         []string
}

// create a new Kong plugin
func NewKongPlugin(name string, serviceId string, routeId string, config []KongPluginConfig, enabled bool) *KongPlugin {

	return &KongPlugin{
		name:      name,
		serviceId: serviceId,
		routeId:   routeId,
		config:    config,
		enabled:   enabled,
	}
}

// kong plugin Id payload
type KongPluginEntityId struct {
	Id string `json:"id,omitempty"`
}

// kong plugin request payload
type KongPluginRequest struct {
	Name    string              `json:"name,omitempty"`
	Service *KongPluginEntityId `json:"service,omitempty"`
	Route   *KongPluginEntityId `json:"route,omitempty"`
	Enabled bool                `json:"enabled"`
}

// kong plugin response payload
type KongPluginResponse struct {
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	InstanceName string             `json:"instance_name"`
	Protocols    []string           `json:"protocols"`
	Service      KongPluginEntityId `json:"service,omitempty"`
	Route        KongPluginEntityId `json:"route,omitempty"`
	Consumer     KongPluginEntityId `json:"consumer,omitempty"`
	Tags         string             `json:"tags"`
	CreatedAt    uint64             `json:"created_at"`
	UpdatedAt    uint64             `json:"updated_at"`
	Ordering     string             `json:"ordering"`
	Enabled      bool               `json:"enabled"`
}

//    "config": {
//        "hide_credentials": false,
//        "key_in_query": true,
//        "key_in_header": true,
//        "key_in_body": false,
//        "anonymous": null,
//        "run_on_preflight": true,
//        "key_names": [
//            "apikey"
//        ]
//    }

// kong plugin list response payload
type KongPluginListResponse struct {
	Data []KongPluginResponse `json:"data"`
	Next string               `json:"next"`
}

const (
	pluginsResource string = "plugins"

	basicAuthPlugins           string = "basic-auth"
	keyAuthPlugins             string = "key-auth"
	jwtPlugins                 string = "jwt"
	IPRestrictionPlugins       string = "ip-restriction"
	RateLimitingPlugins        string = "rate-limiting"
	RequestSizeLimitingPlugins string = "request-size-limiting"
)

// add a new plugin to Kong
func (ks *KongServerDomain) AddPlugin(newKongPlugin *KongPlugin, options Options) error {

	var pluginURL string = fmt.Sprintf("%s/%s", ks.ServerURL(), pluginsResource)

	pluginReq := KongPluginRequest{
		Name:    newKongPlugin.name,
		Enabled: newKongPlugin.enabled,
	}

	if len(newKongPlugin.serviceId) > 0 {

		pluginReq.Service = &KongPluginEntityId{
			Id: newKongPlugin.serviceId,
		}
	}

	if len(newKongPlugin.routeId) > 0 {

		pluginReq.Route = &KongPluginEntityId{
			Id: newKongPlugin.routeId,
		}
	}

	payload, err := json.Marshal(pluginReq)
	if err != nil {
		return err
	}

	//log.Printf("[debug] post payload: %s\n", payload)

	req, err := http.NewRequest("POST", pluginURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("fail sending add plugin command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var pluginResp KongPluginResponse

	err = json.Unmarshal(respPayload, &pluginResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew plugin ID: %s\n", resp.Status, pluginResp.Id)
		} else {
			fmt.Printf("%s\n", pluginResp.Id)
		}
	}

	return nil
}

// query a plugin by Id
func (ks *KongServerDomain) QueryPlugin(id string, options Options) error {

	var pluginURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), pluginsResource, id)

	//	send a request to Kong to query the plugin by id
	resp, err := http.Get(pluginURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("plugin not found")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending query plugin command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var pluginResp KongPluginResponse

		err = json.Unmarshal(respPayload, &pluginResp)
		if err != nil {
			return err
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\n%s: %s - %s: serviceId: %s ; routeId: %s ; consumerId: %s\n", resp.Status,
				pluginResp.Id, pluginResp.Name, pluginResp.Protocols, pluginResp.Service.Id, pluginResp.Route.Id, pluginResp.Consumer.Id)
		} else {
			fmt.Printf("%s: %s - %s: serviceId: %s ; routeId: %s ; consumerId: %s\n",
				pluginResp.Id, pluginResp.Name, pluginResp.Protocols, pluginResp.Service.Id, pluginResp.Route.Id, pluginResp.Consumer.Id)
		}
	}

	return nil
}

// query a plugin by Id
func (ks *KongServerDomain) ListPlugins(options Options) error {

	var pluginURL string = fmt.Sprintf("%s/%s", ks.ServerURL(), pluginsResource)

	//	send a request to Kong to get a list of all plugins
	resp, err := http.Get(pluginURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending list plugins command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var pluginListResp KongPluginListResponse

		err = json.Unmarshal(respPayload, &pluginListResp)
		if err != nil {
			return err
		}

		if len(pluginListResp.Data) == 0 {
			if options.verbose {
				fmt.Printf("%s\nNo plugins\n", resp.Status)
			} else {
				fmt.Printf("No plugins\n")
			}

			return nil
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nplugin list\n", resp.Status)
		}

		for _, plugin := range pluginListResp.Data {
			fmt.Printf("plugin: %s: %s - %s: serviceId: %s ; routeId: %s ; consumerId: %s\n",
				plugin.Id, plugin.Name, plugin.Protocols, plugin.Service.Id, plugin.Route.Id, plugin.Consumer.Id)
		}
	}

	return nil
}

// update a plugin in Kong
func (ks *KongServerDomain) UpdatePlugin(id string, updatedKongPlugin *KongPlugin, options Options) error {

	var pluginURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), pluginsResource, id)

	pluginReq := KongPluginRequest{
		Name:    updatedKongPlugin.name,
		Enabled: updatedKongPlugin.enabled,
	}

	if len(updatedKongPlugin.serviceId) > 0 {

		pluginReq.Service = &KongPluginEntityId{
			Id: updatedKongPlugin.serviceId,
		}
	}

	if len(updatedKongPlugin.routeId) > 0 {

		pluginReq.Route = &KongPluginEntityId{
			Id: updatedKongPlugin.routeId,
		}
	}

	payload, err := json.Marshal(pluginReq)
	if err != nil {
		return err
	}

	//	log.Printf("[debug] patch payload: %s", payload)

	req, err := http.NewRequest("PATCH", pluginURL, bytes.NewBuffer([]byte(payload)))
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

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("plugin not found")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending patch plugin command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var pluginResp KongPluginResponse

	err = json.Unmarshal(respPayload, &pluginResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\n%s: %s - %s: serviceId: %s ; routeId: %s ; consumerId: %s\n", resp.Status,
				pluginResp.Id, pluginResp.Name, pluginResp.Protocols, pluginResp.Service.Id, pluginResp.Route.Id, pluginResp.Consumer.Id)
		} else {
			fmt.Printf("%s: %s - %s: serviceId: %s ; routeId: %s ; consumerId: %s\n",
				pluginResp.Id, pluginResp.Name, pluginResp.Protocols, pluginResp.Service.Id, pluginResp.Route.Id, pluginResp.Consumer.Id)
		}
	}

	return nil
}

// delete a plugin in Kong
func (ks *KongServerDomain) DeletePlugin(id string, options Options) error {

	var pluginURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), pluginsResource, id)

	//	send a request to Kong to delete the plugin by id
	req, err := http.NewRequest("DELETE", pluginURL, bytes.NewBuffer([]byte("")))
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

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("fail sending delete plugin command to Kong: " + resp.Status)
	}

	if options.jsonOutput {
		fmt.Printf("%s\n{}\n", resp.Status)
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\n", resp.Status)
		}
	}

	return nil
}
