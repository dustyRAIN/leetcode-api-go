package leetcodeapi

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type contestsTestSuite struct {
	suite.Suite
}

func TestContestsSuite(t *testing.T) {
	suite.Run(t, &contestsTestSuite{})
}

func (s *contestsTestSuite) TestGetContestInfo() {
	s.Run("should execute getContestInfo without an error", func() {
		utilsMock := new(IUtilMock)
		utilsMock.On(
			"makeHttpRequest",
			"GET",
			"https://leetcode.com/contest/api/info/contest-12/",
			"application/json",
			"",
			mock.Anything,
		).Return(nil)

		_, err := (&contestService{utils: utilsMock}).getContestInfo("contest-12")

		s.Assert().Nil(err)
	})
}
