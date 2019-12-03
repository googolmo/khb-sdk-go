package khb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.kuaihaibao.com/services/screenshot"

var userAgent = fmt.Sprintf("khb-sdk/go verion/%s", version)

var client = &http.Client{}

func jsonRequest(body interface{}, token string) (*http.Response, error) {
	bByte, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(bByte))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Authorization", "Bearer "+token)
	return client.Do(request)
}
