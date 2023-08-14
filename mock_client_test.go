/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package goquest_test

import (
	"bytes"
	"io"
	"net/http"
)

type (
	// MockClient is the mock client
	MockClient struct {
		DoFunc func(req *http.Request) (*http.Response, error)
	}
)

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func MockResponse(statusCode int, body string) *MockClient {
	r := io.NopCloser(bytes.NewReader([]byte(body)))
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: statusCode, Body: r}, nil
	}

	return &MockClient{}
}

func MockResponseError(statusCode int, body string, err error) *MockClient {
	r := io.NopCloser(bytes.NewReader([]byte(body)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: statusCode, Body: r}, err
	}

	return &MockClient{}
}
