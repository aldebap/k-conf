////////////////////////////////////////////////////////////////////////////////
//	upstreamTarget_test.go  -  Oct-29-2024  -  aldebap
//
//	Test cases for Kong Upstream Target Configuration
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test_AddUpstreamTarget unit tests for AddUpstreamTarget() method
func Test_AddUpstreamTarget(t *testing.T) {

	t.Run(">>> AddUpstreamTarget: scenario 1 - error with the request", func(t *testing.T) {

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

		want := errors.New("fail sending add upstream target command to Kong: 400 Bad Request")
		got := kongServer.AddUpstreamTarget("1234", &KongUpstreamTarget{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> AddUpstreamTarget: scenario 2 - upstreamTarget created successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{
				"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"target": "192.168.68.107:8080"
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.AddUpstreamTarget("1234", &KongUpstreamTarget{}, Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_QueryUpstreamTarget unit tests for QueryUpstreamTarget() method
func Test_QueryUpstreamTarget(t *testing.T) {

	t.Run(">>> QueryUpstreamTarget: scenario 1 - upstream target not found", func(t *testing.T) {

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

		want := errors.New("upstream target not found")
		got := kongServer.QueryUpstreamTarget("1234", "5678", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryUpstreamTarget: scenario 2 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending query upstream target command to Kong: 500 Internal Server Error")
		got := kongServer.QueryUpstreamTarget("1234", "5678", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> QueryUpstreamTarget: scenario 3 - upstream target returned successfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
				"target": "192.168.40.27:8080"
			}`))
		}))
		defer mockKongAdmin.Close()

		//	connect to mock server
		kongServer := NewKongServer(mockKongAdmin.URL, 0)
		if kongServer == nil {
			t.Errorf("fail connectring to mock Kong Admin")
		}

		var want error = nil
		got := kongServer.QueryUpstreamTarget("1234", "5678", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}

// Test_ListUpstreamTarget unit tests for ListUpstreamTargets() method
func Test_ListUpstreamTarget(t *testing.T) {

	t.Run(">>> ListUpstreamTarget: scenario 1 - internal server error", func(t *testing.T) {

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

		want := errors.New("fail sending list upstream targets command to Kong: 500 Internal Server Error")
		got := kongServer.ListUpstreamTargets("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want.Error() != got.Error() {
			t.Errorf("failed checking kong status: error expected: %d result: %d", want, got)
		}
	})

	t.Run(">>> ListUpstreamTarget: scenario 2 - upstream targets uccessfuly", func(t *testing.T) {

		//	mock for Kong Admin
		var mockKongAdmin *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"data": [
					{
						"id": "1343894e-404a-4f9e-a982-9e5c0e9d1733",
						"target": "192.168.40.27:8080"
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
		got := kongServer.ListUpstreamTargets("1234", Options{
			verbose:    false,
			jsonOutput: false,
		})

		//	check the invocation result
		if want != got {
			t.Errorf("failed checking kong status: success expected: result: %s", got.Error())
		}
	})
}
