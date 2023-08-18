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
	"errors"
	"github.com/fathalfath30/goquest/mapper"
	"net/http"
)

func (ts *MapperTestSuite) Test_StatusCode_ShouldReturnNil() {
	ts.Run("informational message", func() {
		ts.Require().Nil(
			mapper.StatusCode(&http.Response{
				StatusCode: http.StatusContinue,
			}),
		)
	})
	ts.Run("success message", func() {
		ts.Require().Nil(
			mapper.StatusCode(&http.Response{
				StatusCode: http.StatusOK,
			}),
		)
	})
}

func (ts *MapperTestSuite) Test_StatusCode_ShouldReturnError() {
	ts.Run("user error", func() {
		ts.Require().Error(
			errors.New(http.StatusText(http.StatusBadRequest)),
			mapper.StatusCode(&http.Response{
				StatusCode: http.StatusBadRequest,
			}),
		)
	})
	ts.Run("server error", func() {
		ts.Require().Error(
			errors.New(http.StatusText(http.StatusInternalServerError)),
			mapper.StatusCode(&http.Response{
				StatusCode: http.StatusInternalServerError,
			}),
		)
	})
}
