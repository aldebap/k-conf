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

// kong consumer list response payload
type KongConsumerListResponse struct {
	Data []KongConsumerResponse `json:"data"`
	Next string                 `json:"next"`
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

// query a consumer by Id
func (ks *KongServerDomain) QueryConsumer(id string, options Options) error {

	var consumerURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), consumersResource, id)

	//	send a request to Kong to query the consumer by id
	resp, err := http.Get(consumerURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("consumer not found")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending query consumer command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var consumerResp KongConsumerResponse

		err = json.Unmarshal(respPayload, &consumerResp)
		if err != nil {
			return err
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nconsumer: %s --> %s (%s)\n", resp.Status,
				consumerResp.CustomId, consumerResp.UserName, consumerResp.Tags)
		} else {
			fmt.Printf("consumer: %s --> %s (%s)\n",
				consumerResp.CustomId, consumerResp.UserName, consumerResp.Tags)
		}
	}

	return nil
}

// query a consumer by Id
func (ks *KongServerDomain) ListConsumers(options Options) error {

	var consumerURL string = fmt.Sprintf("%s/%s/", ks.ServerURL(), consumersResource)

	//	send a request to Kong to get a list of all consumers
	resp, err := http.Get(consumerURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending list consumers command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var consumerListResp KongConsumerListResponse

		err = json.Unmarshal(respPayload, &consumerListResp)
		if err != nil {
			return err
		}

		if len(consumerListResp.Data) == 0 {
			if options.verbose {
				fmt.Printf("%s\nNo consumers\n", resp.Status)
			} else {
				fmt.Printf("No consumers\n")
			}

			return nil
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nconsumer list\n", resp.Status)
		}

		for _, consumer := range consumerListResp.Data {
			fmt.Printf("%s: (%s) %s %s\n", consumer.Id,
				consumer.CustomId, consumer.UserName, consumer.Tags)
		}
	}

	return nil
}

// update a consumer in Kong
func (ks *KongServerDomain) UpdateConsumer(id string, updatedKongConsumer *KongConsumer, options Options) error {

	var consumerURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), consumersResource, id)

	payload, err := json.Marshal(KongConsumerRequest{
		CustomId: updatedKongConsumer.customId,
		UserName: updatedKongConsumer.userName,
		Tags:     updatedKongConsumer.tags,
	})
	if err != nil {
		return err
	}

	//	log.Printf("[debug] patch payload: %s", payload)

	req, err := http.NewRequest("PATCH", consumerURL, bytes.NewBuffer([]byte(payload)))
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

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending patch consumer command to Kong: " + resp.Status)
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
			fmt.Printf("http response status code: %s\nconsumer: %s --> %s (%s)\n", resp.Status,
				consumerResp.CustomId, consumerResp.UserName, consumerResp.Tags)
		} else {
			fmt.Printf("consumer: %s --> %s (%s)\n",
				consumerResp.CustomId, consumerResp.UserName, consumerResp.Tags)
		}
	}

	return nil
}

// delete a consumer in Kong
func (ks *KongServerDomain) DeleteConsumer(id string, options Options) error {

	var consumerURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), consumersResource, id)

	//	send a request to Kong to delete the consumer by id
	req, err := http.NewRequest("DELETE", consumerURL, bytes.NewBuffer([]byte("")))
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
		return errors.New("fail sending delete consumer command to Kong: " + resp.Status)
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
