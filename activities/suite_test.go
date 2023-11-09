package activities

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestActivityEnvironment
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestActivityEnvironment()
}

func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
