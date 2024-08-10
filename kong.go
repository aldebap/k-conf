////////////////////////////////////////////////////////////////////////////////
//	kong.go  -  Jul-5-2024  -  aldebap
//
//	Kong server configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Kong server attributes
type KongServer struct {
	address string
	port    int
}

// create a new Kong server configuration
func NewKongServer(address string, port int) *KongServer {

	return &KongServer{
		address: address,
		port:    port,
	}
}

func (ks *KongServer) ServerURL() string {
	var kongUrl string = ks.address

	if ks.port != 0 {
		kongUrl = fmt.Sprintf("http://%s:%d", ks.address, ks.port)
	}

	return kongUrl
}

// check Kong status
func (ks *KongServer) CheckStatus(options Options) error {

	var serviceURL string = ks.ServerURL()

	//	send a request to Kong to check it's status
	resp, err := http.Get(serviceURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("error sending check status command to Kong: " + resp.Status)
	}

	if options.jsonOutput {
		var respPayload []byte

		respPayload, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n%s\n", resp.Status, string(respPayload))
	} else {
		fmt.Printf("%s\n", resp.Status)
	}

	return nil
}
