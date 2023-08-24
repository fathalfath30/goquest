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

import "net/http"

func New(cfg *Config) (IGoQuest, error) {
	gq := new(GoQuest)
	if cfg != nil {
		// input custom transport
		if cfg.Transport != nil {
			gq.transport = cfg.Transport
		}

		// input custom header
		if cfg.Header != nil {
			gq.header = cfg.Header
		}

		if cfg.Client != nil {
			gq.client = cfg.Client
		}
	}

	if gq.client == nil {
		gq.client = &http.Client{}
	}

	return gq, nil
}
