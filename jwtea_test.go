package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server         *httptest.Server
	reader         io.Reader
	createTokenURL string
	decodeTokenURL string
	generatedToken string
)

func init() {
	server = httptest.NewServer(Handlers())
	createTokenURL = server.URL + "/create"
	decodeTokenURL = server.URL + "/decode"
}

func TestCreateToken(t *testing.T) {
	tokenJSON := `{"username": "dennis", "balance": 200}`
	reader = strings.NewReader(tokenJSON)

	request, err := http.NewRequest("POST", createTokenURL, reader)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	var tokenRes struct {
		Success bool
		Message string
		Payload TokenResponse
	}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &tokenRes)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
	generatedToken = tokenRes.Payload.Token
}

func TestDecodeToken(t *testing.T) {
	tokenJSON := `{"token": "` + generatedToken + `"}`
	reader = strings.NewReader(tokenJSON)

	request, err := http.NewRequest("POST", decodeTokenURL, reader)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	var decodeRes Message
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &decodeRes)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}
