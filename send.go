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
	"bytes"
	"context"
	"fmt"
	"github.com/fathalfath30/goquest/mapper"
	"github.com/fathalfath30/goquest/utils"
	"io"
	"net/http"
	"net/http/httputil"
)

func (gq *GoQuest) Send(ctx context.Context, method, endpoint string, requestOption *RequestOption) (*Response, error) {
	var body io.Reader = nil
	result := &Response{}

	// initialize default http header
	header := http.Header{}

	// get default header from GoQuest
	if gq.header != nil {
		header = *gq.header
	}

	if requestOption != nil {
		if requestOption.Json != nil {
			// set body data to the reader
			body = bytes.NewReader(requestOption.Json)

			// force Content-Type to 'application/json'
			if header.Get("Content-Type") != utils.ContentTypeAppJson {
				header.Set("Content-Type", utils.ContentTypeAppJson)
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, gq.BaseUrl.JoinPath(endpoint).String(), body)
	if err != nil {
		return result, err
	}

	req.URL.RawQuery = gq.queryParam.Encode()
	req.Header = header
	if gq.dumpRequest {
		dump, errDump := httputil.DumpRequest(req, true)
		if errDump != nil {
			fmt.Println("cannot dump request", err)
		}

		fmt.Println("Request")
		fmt.Println(string(dump))
	}

	res, err := gq.client.Do(req)
	if err != nil {
		return result, err
	}

	result.Status = res.Status
	result.StatusCode = res.StatusCode
	if gq.dumpResponse {
		dump, errDump := httputil.DumpResponse(res, true)
		if errDump != nil {
			fmt.Println("cannot dump response", err)
		}

		fmt.Println("Response")
		fmt.Println(string(dump))
	}

	return result, mapper.StatusCode(res)
}
