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
	"context"
	"encoding/json"
	"github.com/fathalfath30/goquest/exception"
	"io"
	"net/http"
	"net/url"
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
		GetHeader(key string) string
		AddQueryParam(key, value string) IGoQuest

		Post(ctx context.Context, endpoint string, option *RequestOption) (*Response, error)
		Get(ctx context.Context, endpoint string, option *RequestOption) (*Response, error)
		Patch(ctx context.Context, endpoint string, option *RequestOption) (*Response, error)
		Put(ctx context.Context, endpoint string, option *RequestOption) (*Response, error)
		Delete(ctx context.Context, endpoint string, option *RequestOption) (*Response, error)

		Send(ctx context.Context, method, endpoint string, requestOption *RequestOption) (*Response, error)
	}

	Config struct {
		// base url
		BaseUrl string

		// Http transport
		Transport *http.Transport

		// default http header
		Header *http.Header

		DumpRequest  bool
		DumpResponse bool

		// main client
		Client IHttpClient
	}

	GoQuest struct {
		config *Config
		client IHttpClient

		// base url
		BaseUrl *url.URL

		// Http transport
		transport *http.Transport

		// default http header
		header *http.Header

		// Url Query Parameters
		queryParam url.Values

		dumpRequest  bool
		dumpResponse bool

		// main client
		Client IHttpClient
	}

	RequestOption struct {
		// Json is used to send json encoded data as the part of body request, and also
		// if this option is not empty Content-Type "application/json" will be
		// automatically added if there is no "application/json" present on header
		Json json.RawMessage

		//  Header represents the key-value pairs in an HTTP header.
		Header http.Header
	}

	Response struct {
		Status     string // e.g. "200 OK"
		StatusCode int    // e.g. 200

		// Header maps header keys to values.
		Header http.Header

		// Body represents the response body.
		Body io.ReadCloser
	}
)
