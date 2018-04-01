package code

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) SetupSuite() {}

func (s *Suite) TearDownSuite() {}

func (s *Suite) SetupTest() {}

func (s *Suite) TearDownTest() {}

func (s *Suite) TestF1() {
	s.T().Log("F1")
	s.Nil(nil)
}

func (s *Suite) TestF2() {
	s.T().Log("F2")
	s.Nil(nil)
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}