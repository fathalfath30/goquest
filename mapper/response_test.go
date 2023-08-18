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

package mapper_test

import (
	"bytes"
	"github.com/fathalfath30/goquest/mapper"
	"github.com/fathalfath30/goquest/mocks"
	"github.com/fathalfath30/goquest/testdata"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
)

func (ts *MapperTestSuite) Test_ResponseMapper_ErrorHandling() {
	readCloserMock := mocks.NewMockReadCloser()
	// if Read is called, it will return error
	readCloserMock.On("Read", mock.Anything).
		Return(0, testdata.ErrorReadingResponseBody).Once()
	readCloserMock.On("Close").Return(nil)

	ts.Run("fetch response body", func() {
		mp, err := mapper.JsonResponse[*testdata.Response](
			&http.Response{
				StatusCode: http.StatusOK,
				Body:       readCloserMock,
			},
		)

		ts.Require().Nil(mp)
		ts.Require().NotNil(err)
		ts.Require().Equal(testdata.ErrorReadingResponseBody, err)
	})
	ts.Run("unmarshall json", func() {
		mp, err := mapper.JsonResponse[*testdata.Response](
			&http.Response{
				StatusCode: http.StatusOK,
				Header:     testdata.SampleJsonHeader,
				Body:       testdata.InvalidJsonResponse,
			},
		)

		ts.Require().Nil(mp)
		ts.Require().NotNil(err)
		ts.Require().Equal("invalid character '<' looking for beginning of value", err.Error())
	})
}
func (ts *MapperTestSuite) Test_ItCanMappingResponseBodyToStruct() {
	ts.Run("json response", func() {
		mp, err := mapper.JsonResponse[*testdata.Response](
			&http.Response{
				StatusCode: http.StatusOK,
				Header:     testdata.SampleJsonHeader,
				Body:       io.NopCloser(bytes.NewReader([]byte(testdata.SampleSuccessJson))),
			},
		)

		ts.Require().NotNil(mp)
		ts.Require().Nil(err)
		ts.Require().IsType(&testdata.Response{}, mp)

		ts.Require().IsType(&testdata.Status{}, mp.Status)
		ts.Require().Equal(http.StatusOK, mp.Status.Code)
		ts.Require().Equal(testdata.GetOke, mp.Status.Message)

		ts.Require().NotNil(mp.Data)
		ts.Require().IsType(&testdata.Data{}, mp.Data)
		ts.Require().Equal(testdata.Ipsum, mp.Data.Lorem)
	})
}
