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

package goquest_test

import (
	"github.com/fathalfath30/goquest"
	"github.com/fathalfath30/goquest/testdata"
	"net/http"
)

func (ts *GoQuesTestSuite) Test_InitialSetup() {
	ts.Run("it should return default value", func() {
		gq, err := goquest.New(testdata.ValidSampleBaseUrl, nil)

		ts.Require().NotNil(gq)
		ts.Require().IsType(&goquest.GoQuest{}, gq)
		ts.Require().Nil(err)
	})
	ts.Run("it can set default header from client", func() {
		gq, err := goquest.New(testdata.ValidSampleBaseUrl, &goquest.Config{
			Client:    &http.Client{},
			Header:    &http.Header{},
			Transport: &http.Transport{},
		})

		ts.Require().NotNil(gq)
		ts.Require().IsType(&goquest.GoQuest{}, gq)
		ts.Require().Nil(err)
	})
	ts.Run("it can set baseUrl", func() {
		gq, err := goquest.New(testdata.ValidSampleBaseUrl, &goquest.Config{
			BaseUrl: testdata.ValidSampleBaseUrl,
		})

		ts.Require().NotNil(gq)
		ts.Require().IsType(&goquest.GoQuest{}, gq)
		ts.Require().Nil(err)
	})
}
func (ts *GoQuesTestSuite) Test_InitialSetup_MustHandleError() {
	ts.Run("it should handle error when failed parse base url", func() {
		gq, err := goquest.New("-://localhost:we", nil)

		ts.Require().Nil(gq)
		ts.Require().Error(err)
		ts.Require().Contains(err.Error(), "failed to parse base url")
	})
}
