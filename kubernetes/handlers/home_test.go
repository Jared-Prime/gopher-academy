package handlers

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	recorder := httptest.NewRecorder()
	home(recorder, nil)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusOK)

	response_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(response_body), "Hello! Your request was processed.\n")
}
