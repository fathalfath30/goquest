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

package exception_test

import (
	"github.com/fathalfath30/goquest/exception"
)

func (ts *ExceptionTestSuite) Test_HttpRequestException() {
	ts.Run("it can create new HttpRequestException and return default value if message are empty", func() {
		actual := exception.NewHttpRequestException("lorem ipsum")

		ts.Require().NotNil(actual)
		ts.Require().IsType(&exception.HttpRequestException{}, actual)
		ts.Require().Equal("lorem ipsum", actual.Error())
	})
}
