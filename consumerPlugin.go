////////////////////////////////////////////////////////////////////////////////
//	consumerPlugin.go  -  Sep-25-2024  -  aldebap
//
//	Kong consumer Plugins
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

// kong Basic Auth config attributes
type KongBasicAuthConfig struct {
	userName string
	password string
}

// create a new kong Basic Auth config
func NewKongBasicAuthConfig(userName string, password string) *KongBasicAuthConfig {

	return &KongBasicAuthConfig{
		userName: userName,
		password: password,
	}
}

// kong consumer basic auth request payload
type KongConsumerBasicAuthRequest struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// kong consumer response payload
type KongConsumerBasicAuthResponse struct {
	Id string `json:"id"`
}

// add a new consumer to Kong
func (ks *KongServerDomain) AddConsumerBasicAuth(id string, newKongBasicAuthConfig *KongBasicAuthConfig, options Options) error {

	var consumerBasicAuthURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), consumersResource, id, basicAuthPlugins)

	payload, err := json.Marshal(KongConsumerBasicAuthRequest{
		UserName: newKongBasicAuthConfig.userName,
		Password: newKongBasicAuthConfig.password,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerBasicAuthURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("consumer not found")
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New("fail sending add consumer basic auth command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var consumerBasicAuthResp KongConsumerBasicAuthResponse

	err = json.Unmarshal(respPayload, &consumerBasicAuthResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew plugin ID: %s\n", resp.Status, consumerBasicAuthResp.Id)
		} else {
			fmt.Printf("%s\n", consumerBasicAuthResp.Id)
		}
	}

	return nil
}

// kong KeyAuth config attributes
type KongKeyAuthConfig struct {
	key string
	ttl int64
}

// create a new kong KeyAuth config
func NewKongKeyAuthConfig(key string, ttl int64) *KongKeyAuthConfig {

	return &KongKeyAuthConfig{
		key: key,
		ttl: ttl,
	}
}

// kong consumer KeyAuth request payload
type KongConsumerKeyAuthRequest struct {
	Key string `json:"key,omitempty"`
	Ttl int64  `json:"ttl,omitempty"`
}

// kong consumer KeyAuth response payload
type KongConsumerKeyAuthResponse struct {
	Id string `json:"id"`
}

func (ks *KongServerDomain) AddConsumerKeyAuth(id string, newKongKeyAuthConfig *KongKeyAuthConfig, options Options) error {

	var consumerBasicAuthURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), consumersResource, id, keyAuthPlugins)

	payload, err := json.Marshal(KongConsumerKeyAuthRequest{
		Key: newKongKeyAuthConfig.key,
		Ttl: newKongKeyAuthConfig.ttl,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerBasicAuthURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("consumer not found")
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New("fail sending add consumer keyAuth command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var consumerKeyAuthResp KongConsumerKeyAuthResponse

	err = json.Unmarshal(respPayload, &consumerKeyAuthResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew plugin ID: %s\n", resp.Status, consumerKeyAuthResp.Id)
		} else {
			fmt.Printf("%s\n", consumerKeyAuthResp.Id)
		}
	}

	return nil
}
