////////////////////////////////////////////////////////////////////////////////
//	route_test.go  -  Aug-22-2024  -  aldebap
//
//	Test cases for Kong Route Configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_AddRoute unit tests for AddRoute() method
func Test_AddRoute(t *testing.T) {

	t.Run(">>> AddRoute: scenario 1 - error with the request", func(t *testing.T) {

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

		want := errors.New("fail sending add route command to Kong: 400 Bad Request")
		got := kongServer.AddRoute(&KongRoute{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> AddRoute: scenario 2 - route created successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"regex_priority": 0,
				"request_buffering": true,
				"response_buffering": true,
				"strip_path": true,
				"service": {
					"id":"1343894e-404a-4f9e-a982-9e5c0e9d1733"
				},
				"path_handling": "v0",
				"preserve_host": false,
				"snis": null,
				"name": "Produto",
				"tags": null,
				"hosts": null,
				"https_redirect_status_code": 426,
				"updated_at": 1724380393,
				"protocols": [
					"http"
				],
				"created_at": 1724380393,
				"paths": [
					"/gwa/v1/produtos"
				],
				"sources": null,
				"destinations": null,
				"headers": null,
				"id": "6154be97-ce77-47ea-bf46-7b0266b2c054",
				"methods": [
					"GET",
					"POST"
				]
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.AddRoute(&KongRoute{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_QueryRoute unit tests for QueryRoute() method
func Test_QueryRoute(t *testing.T) {

	t.Run(">>> QueryRoute: scenario 1 - route not found", func(t *testing.T) {

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

		want := errors.New("route not found")
		got := kongServer.QueryRoute("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryRoute: scenario 2 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending query route command to Kong: 500 Internal Server Error")
		got := kongServer.QueryRoute("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryRoute: scenario 3 - route returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"regex_priority": 0,
				"request_buffering": true,
				"response_buffering": true,
				"strip_path": true,
				"service": {
					"id":"1343894e-404a-4f9e-a982-9e5c0e9d1733"
				},
				"path_handling": "v0",
				"preserve_host": false,
				"snis": null,
				"name": "Produto",
				"tags": null,
				"hosts": null,
				"https_redirect_status_code": 426,
				"updated_at": 1724380393,
				"protocols": [
					"http"
				],
				"created_at": 1724380393,
				"paths": [
					"/gwa/v1/produtos"
				],
				"sources": null,
				"destinations": null,
				"headers": null,
				"id": "6154be97-ce77-47ea-bf46-7b0266b2c054",
				"methods": [
					"GET",
					"POST"
				]
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.QueryRoute("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_ListRoute unit tests for ListRoutes() method
func Test_ListRoute(t *testing.T) {

	t.Run(">>> ListRoute: scenario 1 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending list route command to Kong: 500 Internal Server Error")
		got := kongServer.ListRoutes(Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> ListRoute: scenario 2 - route returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"data": [
					{
						"regex_priority": 0,
						"request_buffering": true,
						"response_buffering": true,
						"strip_path": true,
						"service": {
							"id":"1343894e-404a-4f9e-a982-9e5c0e9d1733"
						},
						"path_handling": "v0",
						"preserve_host": false,
						"snis": null,
						"name": "Produto",
						"tags": null,
						"hosts": null,
						"https_redirect_status_code": 426,
						"updated_at": 1724380393,
						"protocols": [
							"http"
						],
						"created_at": 1724380393,
						"paths": [
							"/gwa/v1/produtos"
						],
						"sources": null,
						"destinations": null,
						"headers": null,
						"id": "6154be97-ce77-47ea-bf46-7b0266b2c054",
						"methods": [
							"GET",
							"POST"
						]
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
		got := kongServer.ListRoutes(Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_DeleteRoute unit tests for DeleteRoute() method
func Test_DeleteRoute(t *testing.T) {

	t.Run(">>> DeleteRoute: scenario 1 - route not found", func(t *testing.T) {

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

		want := errors.New("fail sending delete route command to Kong: 404 Not Found")
		got := kongServer.DeleteRoute("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> DeleteRoute: scenario 2 - route returned successfuly", func(t *testing.T) {

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
		got := kongServer.DeleteRoute("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}
