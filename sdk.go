package khb

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

const version = "1.0.1"

const (
	// TypeJpg result image type is jpg
	TypeJpg = "jpeg"
	// TypeJpeg result image type is jpg
	TypeJpeg = "jpeg"
	// TypePng result image type is png
	TypePng = "png"
)

// Screenshot is struct for api request paramaters
type Screenshot struct {
	Template       string                 `json:"template,omitempty"`
	URL            string                 `json:"url,omitempty"`
	HTML           string                 `json:"html,omitempty"`
	Data           map[string]interface{} `json:"data"`
	Headers        map[string]interface{} `json:"headers"`
	Device         string                 `json:"device,omitempty"`
	CustomDevice   *DeviceDescriptor      `json:"custom_device,omitempty"`
	Type           string                 `json:"type,omitempty"`
	FullPage       bool                   `json:"full_page"`
	Quality        int                    `json:"quality,omitempty"`
	OmitBackground bool                   `json:"omit_background"`
}

// DeviceViewport is device veiewport info
type DeviceViewport struct {
	Width             int     `json:"width"`
	Height            int     `json:"height"`
	DeviceScaleFactor float32 `json:"deviceScaleFactor"`
	IsMobile          bool    `json:"isMobile"`
	HasTouch          bool    `json:"hasTouch"`
	IsLandscape       bool    `json:"isLandscape"`
}

// DeviceDescriptor is device descript
type DeviceDescriptor struct {
	Name      string         `json:"name"`
	UserAgent string         `json:"userAgent"`
	Viewport  DeviceViewport `json:"viewport"`
}

// APIError is error of API Request
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (apiError *APIError) Error() string {
	return fmt.Sprintf("APIError{code: %d, message:%s}", apiError.Code, apiError.Message)
}

type basicResponse struct {
	Data *Result `json:"data"`
}

// Result is screenshot result
type Result struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (screenshot *Screenshot) invoke(token string) (*Result, error) {
	if strings.TrimSpace(token) == "" {
		return nil, &APIError{
			Code:    2001,
			Message: "token 不能为空",
		}
	}
	if strings.TrimSpace(screenshot.Template) == "" &&
		strings.TrimSpace(screenshot.URL) == "" &&
		strings.TrimSpace(screenshot.HTML) == "" {
		return nil, &APIError{
			Code:    9001,
			Message: "HTML, Template, URL 至少一个不为空",
		}
	}
	if strings.TrimSpace(screenshot.URL) != "" {
		if _, err := url.Parse(screenshot.URL); err != nil {
			return nil, err
		}
	}
	response, err := jsonRequest(screenshot, token)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		_reader, _err := gzip.NewReader(response.Body)
		if _err != nil {
			return nil, _err
		}
		reader = _reader
		defer reader.Close()
	default:
		reader = response.Body
	}
	if response.StatusCode >= 500 {
		return nil, &APIError{
			Code:    response.StatusCode,
			Message: response.Status,
		}
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		var errResp APIError
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, &errResp
		// reponse.Body
	}
	var br basicResponse
	if err := json.Unmarshal(body, &br); err != nil {
		return nil, err
	}
	return br.Data, nil
}
