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
func (ks *KongServerDomain) AddUpstreamTarget(upstreamId string, newKongUpstreamTarget *KongUpstreamTarget, options Options) error {

	var upstreamTargetURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), upstreamResource, upstreamId, upstreamTargetResource)

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

// query an upstream target by Id
func (ks *KongServerDomain) QueryUpstreamTarget(upstreamId string, id string, options Options) error {

	var upstreamTargetURL string = fmt.Sprintf("%s/%s/%s/%s/%s", ks.ServerURL(), upstreamResource, upstreamId, upstreamTargetResource, id)

	//	send a request to Kong to query the upstream target by id
	resp, err := http.Get(upstreamTargetURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("upstream target not found")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending query upstream target command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var upstreamTargetResp KongUpstreamTargetResponse

		err = json.Unmarshal(respPayload, &upstreamTargetResp)
		if err != nil {
			return err
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nupstream target: %s\n", resp.Status,
				upstreamTargetResp.Target)
		} else {
			fmt.Printf("upstream target: %s\n",
				upstreamTargetResp.Target)
		}
	}

	return nil
}

// list all upstreams
func (ks *KongServerDomain) ListUpstreamTargets(upstreamId string, options Options) error {

	var upstreamTargetURL string = fmt.Sprintf("%s/%s/%s/%s", ks.ServerURL(), upstreamResource, upstreamId, upstreamTargetResource)

	//	send a request to Kong to get a list of all upstream targets
	resp, err := http.Get(upstreamTargetURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending list upstream targets command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var upstreamTargetListResp KongUpstreamTargetListResponse

		err = json.Unmarshal(respPayload, &upstreamTargetListResp)
		if err != nil {
			return err
		}

		if len(upstreamTargetListResp.Data) == 0 {
			if options.verbose {
				fmt.Printf("%s\nNo upstream targets\n", resp.Status)
			} else {
				fmt.Printf("No upstream targets\n")
			}

			return nil
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nupstream target list\n", resp.Status)
		}

		for _, upstreamTarget := range upstreamTargetListResp.Data {
			fmt.Printf("%s: %s\n", upstreamTarget.Id,
				upstreamTarget.Target)
		}
	}

	return nil
}
