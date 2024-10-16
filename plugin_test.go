////////////////////////////////////////////////////////////////////////////////
//	plugin_test.go  -  Oct-16-2024  -  aldebap
//
//	Test cases for Kong Plugin Configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_AddPlugin unit tests for AddPlugin() method
func Test_AddPlugin(t *testing.T) {

	t.Run(">>> AddPlugin: scenario 1 - error with the request", func(t *testing.T) {

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

		want := errors.New("fail sending add plugin command to Kong: 400 Bad Request")
		got := kongServer.AddPlugin(&KongPlugin{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> AddPlugin: scenario 2 - plugin created successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"instanceName": "key-auth",
				"name": "key-auth",
				"serviceId": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"protocols": [
					"http"
				],
				"enabled": true
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.AddPlugin(&KongPlugin{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_QueryPlugin unit tests for QueryPlugin() method
func Test_QueryPlugin(t *testing.T) {

	t.Run(">>> QueryPlugin: scenario 1 - plugin not found", func(t *testing.T) {

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

		want := errors.New("plugin not found")
		got := kongServer.QueryPlugin("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryPlugin: scenario 2 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending query plugin command to Kong: 500 Internal Server Error")
		got := kongServer.QueryPlugin("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryPlugin: scenario 3 - plugin returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"instanceName": "key-auth",
				"name": "key-auth",
				"serviceId": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"protocols": [
					"http"
				],
				"enabled": true
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.QueryPlugin("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_ListPlugin unit tests for ListPlugins() method
func Test_ListPlugin(t *testing.T) {

	t.Run(">>> ListPlugin: scenario 1 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending list plugins command to Kong: 500 Internal Server Error")
		got := kongServer.ListPlugins(Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> ListPlugin: scenario 2 - plugin returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"data": [
					{
						"instanceName": "key-auth",
						"name": "key-auth",
						"serviceId": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
						"protocols": [
							"http"
						],
						"enabled": true
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
		got := kongServer.ListPlugins(Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_UpdatePlugin unit tests for UpdatePlugin() method
func Test_UpdatePlugin(t *testing.T) {

	t.Run(">>> UpdatePlugin: scenario 1 - plugin not found", func(t *testing.T) {

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

		want := errors.New("plugin not found")
		got := kongServer.UpdatePlugin("1234", &KongPlugin{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> UpdatePlugin: scenario 2 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending patch plugin command to Kong: 500 Internal Server Error")
		got := kongServer.UpdatePlugin("1234", &KongPlugin{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> UpdatePlugin: scenario 3 - plugin returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"instanceName": "key-auth",
				"name": "key-auth",
				"serviceId": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"protocols": [
					"http"
				],
				"enabled": true
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.UpdatePlugin("1234", &KongPlugin{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_DeletePlugin unit tests for DeletePlugin() method
func Test_DeletePlugin(t *testing.T) {

	t.Run(">>> DeletePlugin: scenario 1 - plugin not found", func(t *testing.T) {

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

		want := errors.New("fail sending delete plugin command to Kong: 404 Not Found")
		got := kongServer.DeletePlugin("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> DeletePlugin: scenario 2 - plugin returned successfuly", func(t *testing.T) {

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
		got := kongServer.DeletePlugin("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}
