package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHttpServer(t *testing.T) {
	// This method is to handle the expected response
	handler := func(w http.ResponseWriter, r *http.Request) {
		// The expected response is from external file at /testdata/get_expected_response.json
		io.WriteString(w, expectedResponse)
	}

	req := httptest.NewRequest("GET", "http://demo-meli-ceiba-test.demo-meli-ceiba.melifrontends.com/pong", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, resp.Header.Get("Content-Type"), "text/plain; charset=utf-8")
	assert.Equal(t, string(body), expectedResponse)
}

func TestHttpClient(t *testing.T) {

	client := &http.Client{
		CheckRedirect: nil,
	}
	req, err := http.NewRequest("GET", "http://demo-meli-ceiba-test.demo-meli-ceiba.melifrontends.com/ping", nil)
	req.Header.Add("x-auth-token", "3a0cc3db478ff18d6400ca68d26ea71868cb07b286c2d1862fb462766c11baa6")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("There was an error, try again")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, resp.Header.Get("Content-Type"), "text/plain; charset=utf-8")
	assert.Equal(t, string(body), "pong")
}