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

package utils_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// UtilsTestSuite is used to simplify generating test case
type (
	UtilsTestSuite struct {
		suite.Suite
		t *testing.T
	}
)

// Test_RunUtilsTestSuite Running the test suite
func Test_RunUtilsTestSuite(t *testing.T) {
	suite.Run(t, &UtilsTestSuite{t: t})
}
