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

package goquest

import (
	"encoding/json"
	"github.com/fathalfath30/goquest/exception"
	"net/http"
)

var (
	ErrorUnsupportedContentType = exception.NewGeneralException("unsupported Content-Type")
)

//go:generate mockery --name IHttpClient --filename http_client.mock.go --structname HttpClientMock
//go:generate mockery --name IGoQuest --filename goquest.mock.go --structname GoQuestMock
type (
	IHttpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	IGoQuest interface {
		AddHeader(key, value string) IGoQuest

		Post(endpoint string, option *RequestOption) (*Response, error)
		Get(endpoint string, option *RequestOption) (*Response, error)
		Patch(endpoint string, option *RequestOption) (*Response, error)
		Put(endpoint string, option *RequestOption) (*Response, error)
		Delete(endpoint string, option *RequestOption) (*Response, error)

		Send(method, endpoint string, requestOption *RequestOption) (*Response, error)
	}

	Config struct {
		// base url
		BaseUrl string

		// Http transport
		Transport *http.Transport

		// default http header
		Header *http.Header

		// main client
		Client IHttpClient
	}

	GoQuest struct {
		config *Config
		client IHttpClient

		// base url
		BaseUrl string

		// Http transport
		transport *http.Transport

		// default http header
		header *http.Header

		// main client
		Client IHttpClient
	}

	RequestOption struct {
		// Json is used to send json encoded data as the part of body request, and also
		// if this option is not empty Content-Type "application/json" will be
		// automatically added if there is no "application/json" present on header
		Json json.RawMessage
	}

	Response struct {
		StatusCode int
	}
)
