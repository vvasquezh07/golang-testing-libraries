package suite

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
// Defines general variables for the context
type TestSuite struct {
	suite.Suite
	path  string
	token string
}

// Initializes the general variables
func (suite *TestSuite) SetupTest() {
	suite.path = "http://demo-meli-ceiba-test.demo-meli-ceiba.melifrontends.com/ping"
	suite.token = "3a0cc3db478ff18d6400ca68d26ea71868cb07b286c2d1862fb462766c11baa6"
}

func (suite *TestSuite) TestHttpServer() {
	expectedResponse := "{ \"status\": \"expected service response\"}"
	// This method is to handle the expected response
	handler := func(w http.ResponseWriter, r *http.Request) {
		// The expected response is from external file at /testdata/get_expected_response.json
		io.WriteString(w, expectedResponse)
	}

	req := httptest.NewRequest("GET", suite.path, nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(suite.T(), resp.StatusCode, 200)
	suite.Equal(resp.Header.Get("Content-Type"), "text/plain; charset=utf-8")
	suite.Equal(string(body), expectedResponse)
}

func (suite *TestSuite) TestHttpClient() {
	client := &http.Client{
		CheckRedirect: nil,
	}
	req, err := http.NewRequest("GET", suite.path, nil)
	req.Header.Add("x-auth-token", suite.token)
	
	resp, err := client.Do(req)
	suite.NoError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	suite.Equal(resp.StatusCode, 200)
	suite.Equal(resp.Header.Get("Content-Type"), "text/plain; charset=utf-8")
	suite.Equal(string(body), "pong")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTheSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
