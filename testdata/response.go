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
	"github.com/fathalfath30/goquest"
	"github.com/fathalfath30/goquest/mocks"
	"github.com/fathalfath30/goquest/utils"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

var (
	ProtoHttp1 = "HTTP/1.0"

	ResponseBody = func(value string) io.ReadCloser {
		return io.NopCloser(bytes.NewReader([]byte(value)))
	}
	Client = func(t *testing.T, resp *http.Response, err error) goquest.IHttpClient {
		cl := mocks.NewHttpClientMock(t)
		cl.On("Do", mock.Anything).
			Return(resp, err).
			Once()

		return cl
	}

	ValidJsonHeader = func() http.Header {
		hd := http.Header{}
		hd.Add("Content-Type", utils.ContentTypeAppJson)

		return hd
	}

	ValidHttpOkJson = func(t *testing.T) goquest.IHttpClient {
		return Client(t, &http.Response{
			Status:        http.StatusText(http.StatusOK),
			StatusCode:    http.StatusOK,
			Proto:         ProtoHttp1,
			ProtoMajor:    1,
			ProtoMinor:    0,
			Header:        ValidJsonHeader(),
			Body:          ResponseBody(SampleSuccessJson),
			ContentLength: int64(len(SampleSuccessJson)),
		}, nil)
	}
)
