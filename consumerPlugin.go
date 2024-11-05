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

	var consumerPluginURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), consumersResource, id, basicAuthPlugins)

	payload, err := json.Marshal(KongConsumerBasicAuthRequest{
		UserName: newKongBasicAuthConfig.userName,
		Password: newKongBasicAuthConfig.password,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerPluginURL, bytes.NewBuffer([]byte(payload)))
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

	var consumerPluginURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), consumersResource, id, keyAuthPlugins)

	payload, err := json.Marshal(KongConsumerKeyAuthRequest{
		Key: newKongKeyAuthConfig.key,
		Ttl: newKongKeyAuthConfig.ttl,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerPluginURL, bytes.NewBuffer([]byte(payload)))
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

// kong JWT config attributes
type KongJWTConfig struct {
	algorithm string
	key       string
	secret    string
}

// create a new kong JWT config
func NewKongJWTConfig(algorithm string, key string, secret string) *KongJWTConfig {

	return &KongJWTConfig{
		algorithm: algorithm,
		key:       key,
		secret:    secret,
	}
}

// kong consumer ID payload
type KongConsumerID struct {
	Id string `json:"id"`
}

// kong consumer JWT request payload
type KongConsumerJWTRequest struct {
	Algorithm string `json:"algorithm,omitempty"`
	Key       string `json:"key,omitempty"`
	Secret    string `json:"secret,omitempty"`
}

// kong consumer JWT response payload
type KongConsumerJWTResponse struct {
	Id        string          `json:"id"`
	Consumer  *KongConsumerID `json:"service,omitempty"`
	Algorithm string          `json:"algorithm,omitempty"`
	Key       string          `json:"key,omitempty"`
	Secret    string          `json:"secret,omitempty"`
	Tags      []string        `json:"tags"`
}

func (ks *KongServerDomain) AddConsumerJWT(id string, newKongJWTConfig *KongJWTConfig, options Options) error {

	var consumerPluginURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), consumersResource, id, jwtPlugins)

	payload, err := json.Marshal(KongConsumerJWTRequest{
		Algorithm: newKongJWTConfig.algorithm,
		Key:       newKongJWTConfig.key,
		Secret:    newKongJWTConfig.secret,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerPluginURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("fail sending add consumer JWT command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var consumerJWTResp KongConsumerJWTResponse

	err = json.Unmarshal(respPayload, &consumerJWTResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew plugin ID: %s\n", resp.Status, consumerJWTResp.Id)
		} else {
			fmt.Printf("%s\n", consumerJWTResp.Id)
		}
	}

	return nil
}

// kong plugin config attributes
type KongIPRestrictionConfigRules struct {
	allow []string
	deny  []string
}

// kong IP Restriction config attributes
type KongIPRestrictionConfig struct {
	name   string
	config *KongIPRestrictionConfigRules
}

// create a new kong JWT config
func NewKongIPRestrictionConfig(name string, allow []string, deny []string) *KongIPRestrictionConfig {

	return &KongIPRestrictionConfig{
		name: name,
		config: &KongIPRestrictionConfigRules{
			allow: allow,
			deny:  deny,
		},
	}
}

// kong plugin config attributes
type KongIPRestrictionPluginConfig struct {
	Allow []string `json:"allow,omitempty"`
	Deny  []string `json:"deny,omitempty"`
}

// kong consumer JWT request payload
type KongConsumerIPRestrictionRequest struct {
	Name   string                         `json:"name,omitempty"`
	Config *KongIPRestrictionPluginConfig `json:"config,omitempty"`
}

// kong consumer JWT response payload
type KongConsumerIPRestrictionResponse struct {
	Id     string                        `json:"id"`
	Name   string                        `json:"name,omitempty"`
	Config KongIPRestrictionPluginConfig `json:"config,omitempty"`
}

func (ks *KongServerDomain) AddConsumerIPRestriction(id string, newKongIPRestrictionConfig *KongIPRestrictionConfig, options Options) error {

	var consumerPluginURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), consumersResource, id, IPRestrictionPlugins)

	payload, err := json.Marshal(KongConsumerIPRestrictionRequest{
		Name: newKongIPRestrictionConfig.name,
		Config: &KongIPRestrictionPluginConfig{
			Allow: newKongIPRestrictionConfig.config.allow,
			Deny:  newKongIPRestrictionConfig.config.deny,
		},
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerPluginURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("fail sending add consumer IP Restriction command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var consumerIPRestrictionResp KongConsumerIPRestrictionResponse

	err = json.Unmarshal(respPayload, &consumerIPRestrictionResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew plugin ID: %s\n", resp.Status, consumerIPRestrictionResp.Id)
		} else {
			fmt.Printf("%s\n", consumerIPRestrictionResp.Id)
		}
	}

	return nil
}
