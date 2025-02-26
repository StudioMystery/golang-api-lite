package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-api-lite/library"
	"golang-api-lite/routes"
)

// Sample test
func TestExample(t *testing.T) {
	fooBar := library.FooBar{
		Foo: 123,
		Bar: "test",
	}

	if fooBar.ID != 1 ||
		fooBar.Foo != 123 ||
		fooBar.Bar != "test" {
		t.Fatalf("fooBar failed")
	}
}

func TestFoobarGet(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("GET", "/foobar", nil)
	req.SetBasicAuth("username", "password")
	req.URL.RawQuery = "Id=1"
	if err != nil {
		t.Fatalf("request creation failed")
	}
	resp := httptest.NewRecorder()
	mux := routes.AddRoutes()

	// Act
	mux.ServeHTTP(resp, req)

	// Assert
	if resp.Result().StatusCode != http.StatusOK {
		t.Fatalf("fooBar get failed. Received [%s]", resp.Result().Status)
	}
}

func TestFoobarPost(t *testing.T) {
	// Arrange
	body := library.FooBar{}
	body.Foo = 1234
	body.Bar = "test"

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, err := http.NewRequest("POST", "/foobar", payloadBuf)
	req.SetBasicAuth("username", "password")
	if err != nil {
		t.Fatalf("request creation failed")
	}
	resp := httptest.NewRecorder()
	mux := routes.AddRoutes()

	// Act
	mux.ServeHTTP(resp, req)

	// Assert
	if resp.Result().StatusCode != http.StatusOK {
		t.Fatalf("fooBar get failed. Received [%s]", resp.Result().Status)
	}
}
