////////////////////////////////////////////////////////////////////////////////
//	upstream.go  -  Oct-23-2024  -  aldebap
//
//	Kong upstream configuration
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

// kong upstream attributes
type KongUpstream struct {
	name      string
	algorithm string
	tags      []string
}

// create a new Kong upstream
func NewKongUpstream(name string, algorithm string, tags []string) *KongUpstream {

	return &KongUpstream{
		name:      name,
		algorithm: algorithm,
		tags:      tags,
	}
}

// kong upstream request payload
type KongUpstreamRequest struct {
	Name      string   `json:"name,omitempty"`
	Algorithm string   `json:"algorithm,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// kong upstream response payload
type KongUpstreamResponse struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Algorithm string   `json:"algorithm,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// kong upstream list response payload
type KongUpstreamListResponse struct {
	Data []KongUpstreamResponse `json:"data"`
	Next string                 `json:"next"`
}

const (
	upstreamResource string = "upstreams"
)

// add a new upstream to Kong
func (ks *KongServerDomain) AddUpstream(newKongUpstream *KongUpstream, options Options) error {

	var upstreamURL string = fmt.Sprintf("%s/%s", ks.ServerURL(), upstreamResource)

	payload, err := json.Marshal(KongUpstreamRequest{
		Name:      newKongUpstream.name,
		Algorithm: newKongUpstream.algorithm,
		Tags:      newKongUpstream.tags,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", upstreamURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("fail sending add upstream command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var upstreamResp KongUpstreamResponse

	err = json.Unmarshal(respPayload, &upstreamResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew upstream ID: %s\n", resp.Status, upstreamResp.Id)
		} else {
			fmt.Printf("%s\n", upstreamResp.Id)
		}
	}

	return nil
}
