////////////////////////////////////////////////////////////////////////////////
//	kong.go  -  Jul-5-2024  -  aldebap
//
//	Kong server configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
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

// check Kong status
func (ks *KongServer) CheckStatus() error {

	var kongUrl string = ks.address

	if ks.port != 0 {
		kongUrl = fmt.Sprintf("http://%s:%d", ks.address, ks.port)
	}

	//	send a request to Kong to check it's status
	resp, err := http.Get(kongUrl)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", resp.Status)

	return nil
}
