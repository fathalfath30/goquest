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

func (ts *GoQuesTestSuite) Test_Send_PositiveJourney() {
	ts.Run("it can set request", func() {
		gq, err := goquest.New(testdata.ValidSampleEndpoint, &goquest.Config{
			BaseUrl: testdata.ValidSampleBaseUrl,
			Client:  testdata.ValidHttpOkJson(ts.t),
		})

		ts.Require().Nil(err)
		ts.Require().NotNil(gq)

		actual, err := gq.Send(ts.ctx, http.MethodGet, testdata.ValidSampleEndpoint, nil)
		ts.Require().NotNil(actual)
		ts.Require().Nil(err)
	})
}
