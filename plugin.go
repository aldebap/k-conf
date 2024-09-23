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
	routeId      string
	service      string
	consumer     string
	config       []KongPluginConfig
	protocols    []string
	enabled      bool
	tags         []string
}

// create a new Kong plugin
func NewKongPlugin(name string, routeId string, config []KongPluginConfig, enabled bool) *KongPlugin {

	return &KongPlugin{
		name:    name,
		routeId: routeId,
		config:  config,
		enabled: enabled,
	}
}

// kong plugin Id payload
type KongPluginEntityId struct {
	Id string `json:"id"`
}

// kong plugin request payload
type KongPluginRequest struct {
	Name    string             `json:"name,omitempty"`
	Route   KongPluginEntityId `json:"route,omitempty"`
	Enabled bool               `json:"enabled"`
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
)

// add a new plugin to Kong
func (ks *KongServerDomain) AddPlugin(newKongPlugin *KongPlugin, options Options) error {

	var pluginURL string = fmt.Sprintf("%s/%s", ks.ServerURL(), pluginsResource)

	payload, err := json.Marshal(KongPluginRequest{
		Name: newKongPlugin.name,
		Route: KongPluginEntityId{
			Id: newKongPlugin.routeId,
		},
		Enabled: newKongPlugin.enabled,
	})
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
