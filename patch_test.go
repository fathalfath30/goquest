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
)

func (ts *GoQuesTestSuite) Test_Patch() {
	gq, err := goquest.New(testdata.ValidSampleBaseUrl, &goquest.Config{
		BaseUrl: testdata.ValidSampleBaseUrl,
		Client:  testdata.ValidHttpOkJson(ts.t),
	})

	ts.Require().Nil(err)
	ts.Require().NotNil(gq)

	actual, err := gq.Patch(ts.ctx, testdata.ValidSampleEndpoint, nil)
	ts.Require().Nil(err)
	ts.Require().NotNil(actual)
}
