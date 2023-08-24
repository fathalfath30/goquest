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
	"github.com/fathalfath30/goquest/mapper"
	"io"
	"net/http"
)

func (gq *GoQuest) Send(method, endpoint string, requestOption *RequestOption) (*Response, error) {
	var reader io.Reader = nil
	result := &Response{}

	// initialize default http header
	header := http.Header{}

	// get default header from GoQuest
	if gq.header != nil {
		header = *gq.header
	}

	req, err := http.NewRequest(method, gq.BaseUrl, reader)
	if err != nil {
		return result, err
	}

	req.Header = header
	res, err := gq.client.Do(req)
	if err != nil {
		return result, err
	}

	result.Status = res.Status
	result.StatusCode = res.StatusCode

	return result, mapper.StatusCode(res)
}
