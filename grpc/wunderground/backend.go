package wunderground

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewApiClient(wundergroundApiKeyId string) (*WundergroundApiClient, error) {
	return &WundergroundApiClient{ApiKeyId: wundergroundApiKeyId}, nil
}

type WundergroundApiClient struct {
	ApiKeyId string
}

func (api WundergroundApiClient) Endpoint() string {
	return fmt.Sprintf("http://api.wunderground.com/api/%s/forecast10day/q/IL/Chicago.json", api.ApiKeyId)
}

func (api *WundergroundApiClient) WundergroundForecast() (WundergroundForecast, error) {
	var data WundergroundForecast

	log.Print("calling wunderground API...")
	response, err := http.Get(api.Endpoint())
	if err != nil {
		return data, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(body), &data)

	log.Print("successfully retrieved wunderground API data!")
	return data, err
}
