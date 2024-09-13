////////////////////////////////////////////////////////////////////////////////
//	consumer.go  -  Sep-12-2024  -  aldebap
//
//	Kong consumer configuration
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

// kong consumer attributes
type KongConsumer struct {
	customId string
	userName string
	tags     []string
}

// create a new Kong consumer
func NewKongConsumer(customId string, userName string, tags []string) *KongConsumer {

	return &KongConsumer{
		customId: customId,
		userName: userName,
		tags:     tags,
	}
}

// kong consumer request payload
type KongConsumerRequest struct {
	CustomId string   `json:"custom_id,omitempty"`
	UserName string   `json:"username,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

// kong consumer response payload
type KongConsumerResponse struct {
	Id       string   `json:"id"`
	CustomId string   `json:"custom_id,omitempty"`
	UserName string   `json:"username,omitempty"`
	Tags     []string `json:"tags"`
}

const (
	consumersResource string = "consumers"
)

// add a new consumer to Kong
func (ks *KongServerDomain) AddConsumer(newKongConsumer *KongConsumer, options Options) error {

	var consumerURL string = fmt.Sprintf("%s/%s", ks.ServerURL(), consumersResource)

	payload, err := json.Marshal(KongConsumerRequest{
		CustomId: newKongConsumer.customId,
		UserName: newKongConsumer.userName,
		Tags:     newKongConsumer.tags,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", consumerURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("fail sending add consumer command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var consumerResp KongConsumerResponse

	err = json.Unmarshal(respPayload, &consumerResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew consumer ID: %s\n", resp.Status, consumerResp.Id)
		} else {
			fmt.Printf("%s\n", consumerResp.Id)
		}
	}

	return nil
}
