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
	"net/http"
)

func (ts *GoQuesTestSuite) Test_InitialSetup() {
	ts.Run("it should return default value", func() {
		gq, err := goquest.New(nil)

		ts.Require().NotNil(gq)
		ts.Require().IsType(&goquest.GoQuest{}, gq)
		ts.Require().Nil(err)
	})

	ts.Run("it can set default header from client", func() {
		gq, err := goquest.New(&goquest.Config{
			Client:    &http.Client{},
			Header:    &http.Header{},
			Transport: &http.Transport{},
		})

		ts.Require().NotNil(gq)
		ts.Require().IsType(&goquest.GoQuest{}, gq)
		ts.Require().Nil(err)
	})
}
