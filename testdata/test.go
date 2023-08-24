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

package testdata

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

type (
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	Data struct {
		Lorem string `json:"lorem"`
	}

	Response struct {
		Status *Status `json:"status"`
		Data   *Data   `json:"data"`
	}
)

var (
	GetOke = "get oke"
	Ipsum  = "ipsum"

	ErrorReadingResponseBody = errors.New("error reading body")

	InvalidJsonResponse = io.NopCloser(bytes.NewReader([]byte("<string>")))

	SampleJsonHeader = func() http.Header {
		h := http.Header{}
		h.Set("Content-Type", "application/json")
		return h
	}()
)
