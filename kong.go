////////////////////////////////////////////////////////////////////////////////
//	kong.go  -  Jul-5-2024  -  aldebap
//
//	Kong server configuration
////////////////////////////////////////////////////////////////////////////////

package main

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
