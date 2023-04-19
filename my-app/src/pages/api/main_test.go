package main

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"

)

func TestCreateUser(t *testing.T) {
    reqBody := bytes.NewBufferString(`{"name":"John Doe","email":"johndoe@example.com"}`) // create a request body with a valid JSON data
    w := httptest.NewRecorder() // create a ResponseRecorder for testing HTTP responses
    r, err := http.NewRequest("POST", "/users", reqBody) // create a new request to the CreateUser handler with the request body
    if err != nil {
        t.Fatalf("could not create request: %v", err)
    }

    CreateUser(w, r) // call the CreateUser function with the request

    // check the HTTP response status code is 200
    if status := w.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // check the HTTP response body is a valid ObjectID string
	expectedID := "64402085fb1d4391ae1f5b98"
	if w.Body.String() != expectedID {
		t.Errorf("handler returned unexpected body: got %v want %s", w.Body.String(), expectedID)
	}
	
}
