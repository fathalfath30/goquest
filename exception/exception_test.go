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
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ExceptionTestSuite is used to simplify generating test case
type (
	ExceptionTestSuite struct {
		suite.Suite
		t   *testing.T
		ctx context.Context
	}
)

// Test_RunExceptionTestSuite Running the test suite
func Test_RunExceptionTestSuite(t *testing.T) {
	suite.Run(t, &ExceptionTestSuite{t: t})
}

// SetupTest SetupTestSuite has a SetupTest method, which will run
// before each test in the suite.
func (ts *ExceptionTestSuite) SetupTest() {
	ts.ctx = context.Background()
}
