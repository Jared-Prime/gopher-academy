package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := Router()
	test_server := httptest.NewServer(r)
	defer test_server.Close()

	result, err := http.Get(test_server.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result.StatusCode, http.StatusOK)

	result, err = http.Get(test_server.URL + "/erewhon")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result.StatusCode, http.StatusNotFound)
}
