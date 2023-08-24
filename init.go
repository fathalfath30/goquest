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
	"github.com/fathalfath30/goquest/exception"
	"net/http"
	"net/url"
)

func New(baseUrl string, cfg *Config) (IGoQuest, error) {
	gq := new(GoQuest)
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, exception.NewGeneralException("failed to parse base url: " + err.Error())
	}

	gq.BaseUrl = u

	if cfg != nil {
		// input custom transport
		if cfg.Transport != nil {
			gq.transport = cfg.Transport
		}

		// input custom header
		if cfg.Header != nil {
			gq.header = cfg.Header
		}

		// put user defined
		if cfg.Client != nil {
			gq.client = cfg.Client
		}
	}

	if gq.client == nil {
		gq.client = &http.Client{}
	}

	return gq, nil
}
