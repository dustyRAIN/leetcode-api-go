package leetcodeapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type contestsTestSuite struct {
	suite.Suite
	utilsMock *IUtilMock
}

func TestContestsSuite(t *testing.T) {
	suite.Run(t, &contestsTestSuite{})
}

func (s *contestsTestSuite) SetupSubTest() {
	s.utilsMock = new(IUtilMock)
	s.utilsMock.On(
		"makeHttpRequest",
		"GET",
		"https://leetcode.com/contest/api/info/contest-12/",
		"application/json",
		"",
		&Contest{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(4).(*Contest)
		arg.Company.Description = "description"
		arg.Company.Logo = "logo"
		arg.Company.Name = "name"
		arg.ContainsPremium = false
	})

	s.utilsMock.On(
		"makeHttpRequest",
		"GET",
		"https://leetcode.com/contest/api/info/gimme-error/",
		"application/json",
		"",
		&Contest{},
	).Return(errors.New("some error"))

	s.utilsMock.On(
		"makeHttpRequest",
		"GET",
		"https://leetcode.com/contest/api/ranking/contest-12/?pagination=2&region=global",
		"application/json",
		"",
		&ContestRanking{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(4).(*ContestRanking)
		arg.IsPast = true
		arg.TotalUser = 120
	})

	s.utilsMock.On(
		"makeHttpRequest",
		"GET",
		"https://leetcode.com/contest/api/ranking/gimme-error/?pagination=2&region=global",
		"application/json",
		"",
		&ContestRanking{},
	).Return(errors.New("ow no"))
}

func (s *contestsTestSuite) TestGetContestInfo() {
	s.Run("should execute getContestInfo without an error", func() {
		result, err := (&contestService{utils: s.utilsMock}).getContestInfo("contest-12")
		s.Assert().Nil(err)
		s.Assert().IsType(Contest{}, result)
		expected := Contest{}
		expected.Company.Description = "description"
		expected.Company.Logo = "logo"
		expected.Company.Name = "name"
		expected.ContainsPremium = false
		s.Assert().Equal(expected, result)
	})

	s.Run("should execute getContestInfo returning with an error", func() {
		_, err := (&contestService{utils: s.utilsMock}).getContestInfo("gimme-error")
		s.Assert().EqualError(err, "some error")
	})
}

func (s *contestsTestSuite) TestGetContestRanking() {
	s.Run("should execute getContestRanking without an error", func() {
		result, err := (&contestService{utils: s.utilsMock}).getContestRanking("contest-12", 2)
		s.Assert().Nil(err)
		s.Assert().IsType(ContestRanking{}, result)
		expected := ContestRanking{}
		expected.IsPast = true
		expected.TotalUser = 120
		expected.TotalPage = 5
		s.Assert().Equal(expected, result)
	})

	s.Run("should execute getContestRanking returning with an error", func() {
		_, err := (&contestService{utils: s.utilsMock}).getContestRanking("gimme-error", 2)
		s.Assert().EqualError(err, "ow no")
	})
}
