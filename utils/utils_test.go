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
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

// UtilsTestSuite is used to simplify generating test case
type (
	UtilsTestSuite struct {
		suite.Suite

		t   *testing.T
		ctx context.Context
	}
)

// Test_RunUtilsTestSuite Running the test suite
func Test_RunUtilsTestSuite(t *testing.T) {
	suite.Run(t, &UtilsTestSuite{t: t})
}

// SetupSuite this function executes before the test suite begins
// execution
func (ts *UtilsTestSuite) SetupSuite() {
	// set StartingNumber to one
}

// TearDownSuite which will run after all the tests in the suite
// have been run.
func (ts *UtilsTestSuite) TearDownSuite() {
	// todo: TearDownSuite
}

// SetupTest SetupTestSuite has a SetupTest method, which will run
// before each test in the suite.
func (ts *UtilsTestSuite) SetupTest() {
	ts.ctx = context.Background()
}

// TearDownTest TearDownTestSuite has a TearDownTest method, which
// will run after each test in the suite.
func (ts *UtilsTestSuite) TearDownTest() {
	// todo: TearDownTest
}
