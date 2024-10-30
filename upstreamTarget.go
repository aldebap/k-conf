////////////////////////////////////////////////////////////////////////////////
//	upstreamTarget.go  -  Oct-29-2024  -  aldebap
//
//	Kong upstream target configuration
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

// kong upstream target attributes
type KongUpstreamTarget struct {
	target string
}

// create a new Kong upstreamTarget
func NewKongUpstreamTarget(target string) *KongUpstreamTarget {

	return &KongUpstreamTarget{
		target: target,
	}
}

// kong upstream target request payload
type KongUpstreamTargetRequest struct {
	Target string `json:"target,omitempty"`
}

// kong upstream target response payload
type KongUpstreamTargetResponse struct {
	Id     string `json:"id"`
	Target string `json:"target,omitempty"`
}

// kong upstream target list response payload
type KongUpstreamTargetListResponse struct {
	Data []KongUpstreamTargetResponse `json:"data"`
	Next string                       `json:"next"`
}

const (
	upstreamTargetResource string = "targets"
)

// add a new upstreamTarget to Kong
func (ks *KongServerDomain) AddUpstreamTarget(id string, newKongUpstreamTarget *KongUpstreamTarget, options Options) error {

	var upstreamTargetURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), upstreamResource, id, upstreamTargetResource)

	payload, err := json.Marshal(KongUpstreamTargetRequest{
		Target: newKongUpstreamTarget.target,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", upstreamTargetURL, bytes.NewBuffer([]byte(payload)))
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
		return errors.New("fail sending add upstream target command to Kong: " + resp.Status)
	}

	//	parse response payload
	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var upstreamTargetResp KongUpstreamResponse

	err = json.Unmarshal(respPayload, &upstreamTargetResp)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		if options.verbose {
			fmt.Printf("http response status code: %s\nnew upstream target ID: %s\n", resp.Status, upstreamTargetResp.Id)
		} else {
			fmt.Printf("%s\n", upstreamTargetResp.Id)
		}
	}

	return nil
}
