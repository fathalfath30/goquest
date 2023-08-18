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
	"errors"
	"net/http"
)

var (
	ErrorUnsupportedContentType = errors.New("unsupported Content-Type")
)

//go:generate mockery --name IHttpClient --filename http_client.mock.go --structname HttpClientMock
//go:generate mockery --name IGoQuest --filename goquest.mock.go --structname GoQuestMock
type (
	IHttpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	IGoQuest interface {
		AddHeader(key, value string) IGoQuest

		Post() ([]byte, error)
		Get() ([]byte, error)
		Patch() ([]byte, error)
		Put() ([]byte, error)
		Delete() ([]byte, error)
	}

	Config struct {
		// Http transport
		Transport *http.Transport

		// default http header
		Header *http.Header
	}

	GoQuest struct {
		config *Config
		client IHttpClient
	}
)
