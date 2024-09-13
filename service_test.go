////////////////////////////////////////////////////////////////////////////////
//	service_test.go  -  Aug-20-2024  -  aldebap
//
//	Test cases for Kong Service Configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_AddService unit tests for AddService() method
func Test_AddService(t *testing.T) {

	t.Run(">>> AddService: scenario 1 - error with the request", func(t *testing.T) {

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

		want := errors.New("fail sending add service command to Kong: 400 Bad Request")
		got := kongServer.AddService(&KongService{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> AddService: scenario 2 - service created successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"host": "192.168.68.107",
				"write_timeout": 60000,
				"retries": 5,
				"tls_verify":null,
				"protocol": "http",
				"tls_verify_depth": null,
				"name": "Produtos",
				"client_certificate": null,
				"updated_at": 1724293955,
				"enabled": true,
				"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"created_at": 1724293955,
				"path": "/api/v1/produto",
				"connect_timeout": 60000,
				"port": 8080,
				"tags": null,
				"ca_certificates": null,
				"read_timeout": 60000
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.AddService(&KongService{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_QueryService unit tests for QueryService() method
func Test_QueryService(t *testing.T) {

	t.Run(">>> QueryService: scenario 1 - service not found", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		want := errors.New("service not found")
		got := kongServer.QueryService("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryService: scenario 2 - internal server error", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		want := errors.New("fail sending query service command to Kong: 500 Internal Server Error")
		got := kongServer.QueryService("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryService: scenario 3 - service returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"host": "192.168.68.107",
				"write_timeout": 60000,
				"retries": 5,
				"tls_verify":null,
				"protocol": "http",
				"tls_verify_depth": null,
				"name": "Produtos",
				"client_certificate": null,
				"updated_at": 1724293955,
				"enabled": true,
				"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"created_at": 1724293955,
				"path": "/api/v1/produto",
				"connect_timeout": 60000,
				"port": 8080,
				"tags": null,
				"ca_certificates": null,
				"read_timeout": 60000
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.QueryService("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_ListService unit tests for ListServices() method
func Test_ListService(t *testing.T) {

	t.Run(">>> ListService: scenario 1 - internal server error", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		want := errors.New("fail sending list service command to Kong: 500 Internal Server Error")
		got := kongServer.ListServices(Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> ListService: scenario 2 - service returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"data": [
					{
						"host": "192.168.68.107",
						"write_timeout": 60000,
						"retries": 5,
						"tls_verify":null,
						"protocol": "http",
						"tls_verify_depth": null,
						"name": "Produtos",
						"client_certificate": null,
						"updated_at": 1724293955,
						"enabled": true,
						"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
						"created_at": 1724293955,
						"path": "/api/v1/produto",
						"connect_timeout": 60000,
						"port": 8080,
						"tags": null,
						"ca_certificates": null,
						"read_timeout": 60000
					}
				],
				"next": "null"
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.ListServices(Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_DeleteService unit tests for DeleteService() method
func Test_DeleteService(t *testing.T) {

	t.Run(">>> DeleteService: scenario 1 - service not found", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		want := errors.New("fail sending delete service command to Kong: 404 Not Found")
		got := kongServer.DeleteService("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> DeleteService: scenario 2 - service returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(""))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.DeleteService("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

//	TODO: add test cases for update service
