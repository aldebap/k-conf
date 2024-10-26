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

// query a upstream by Id
func (ks *KongServerDomain) QueryUpstream(id string, options Options) error {

	var upstreamURL string = fmt.Sprintf("%s/%s/%s", ks.ServerURL(), upstreamResource, id)

	//	send a request to Kong to query the upstream by id
	resp, err := http.Get(upstreamURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("upstream not found")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending query upstream command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var upstreamResp KongUpstreamResponse

		err = json.Unmarshal(respPayload, &upstreamResp)
		if err != nil {
			return err
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nupstream: %s --> %s (%s)\n", resp.Status,
				upstreamResp.Name, upstreamResp.Algorithm, upstreamResp.Tags)
		} else {
			fmt.Printf("upstream: %s --> %s (%s)\n",
				upstreamResp.Name, upstreamResp.Algorithm, upstreamResp.Tags)
		}
	}

	return nil
}

// list all upstreams
func (ks *KongServerDomain) ListUpstreams(options Options) error {

	var upstreamURL string = fmt.Sprintf("%s/%s/", ks.ServerURL(), upstreamResource)

	//	send a request to Kong to get a list of all upstreams
	resp, err := http.Get(upstreamURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("fail sending list upstreams command to Kong: " + resp.Status)
	}

	var respPayload []byte

	respPayload, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if options.jsonOutput {
		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		var upstreamListResp KongUpstreamListResponse

		err = json.Unmarshal(respPayload, &upstreamListResp)
		if err != nil {
			return err
		}

		if len(upstreamListResp.Data) == 0 {
			if options.verbose {
				fmt.Printf("%s\nNo upstreams\n", resp.Status)
			} else {
				fmt.Printf("No upstreams\n")
			}

			return nil
		}

		if options.verbose {
			fmt.Printf("http response status code: %s\nupstream list\n", resp.Status)
		}

		for _, upstream := range upstreamListResp.Data {
			fmt.Printf("%s: %s --> %s (%s)\n", upstream.Id,
				upstream.Name, upstream.Algorithm, upstream.Tags)
		}
	}

	return nil
}
