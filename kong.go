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

// Kong server interface
type KongServer interface {
	ServerURL() string
	CheckStatus(options Options) error

	AddService(newKongService *KongService, options Options) error
	QueryService(id string, options Options) error
	ListServices(options Options) error
	UpdateService(id string, updatedKongService *KongService, options Options) error
	DeleteService(id string, options Options) error

	AddRoute(newKongRoute *KongRoute, options Options) error
	QueryRoute(id string, options Options) error
	ListRoutes(options Options) error
	UpdateRoute(id string, updatedKongRoute *KongRoute, options Options) error
	DeleteRoute(id string, options Options) error

	AddConsumer(newKongConsumer *KongConsumer, options Options) error
	QueryConsumer(id string, options Options) error
	ListConsumers(options Options) error
	UpdateConsumer(id string, updatedKongConsumer *KongConsumer, options Options) error
	DeleteConsumer(id string, options Options) error

	AddPlugin(newKongPlugin *KongPlugin, options Options) error
}

// Kong server attributes
type KongServerDomain struct {
	address string
	port    int
}

// create a new Kong server configuration
func NewKongServer(address string, port int) KongServer {

	return &KongServerDomain{
		address: address,
		port:    port,
	}
}

func (ks *KongServerDomain) ServerURL() string {
	var kongUrl string = ks.address

	if ks.port != 0 {
		kongUrl = fmt.Sprintf("http://%s:%d", ks.address, ks.port)
	}

	return kongUrl
}

// check Kong status
func (ks *KongServerDomain) CheckStatus(options Options) error {

	var checkStatusURL string = ks.ServerURL()

	//	send a request to Kong to check it's status
	resp, err := http.Get(checkStatusURL)
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
