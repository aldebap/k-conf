////////////////////////////////////////////////////////////////////////////////
//	upstream_test.go  -  Oct-23-2024  -  aldebap
//
//	Test cases for Kong Upstream Configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_AddUpstream unit tests for AddUpstream() method
func Test_AddUpstream(t *testing.T) {

	t.Run(">>> AddUpstream: scenario 1 - error with the request", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		want := errors.New("fail sending add upstream command to Kong: 400 Bad Request")
		got := kongServer.AddUpstream(&KongUpstream{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> AddUpstream: scenario 2 - upstream created successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"name": "Pedidos",
				"algorithm": "round-robin",
				"tags": null
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.AddUpstream(&KongUpstream{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}
