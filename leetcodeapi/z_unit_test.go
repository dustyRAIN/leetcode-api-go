package leetcodeapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

//----------------------------------------contests-------------------------------------------

type contestsTestSuite struct {
	suite.Suite
	utilsMock *IUtilMock
}

func TestContestService(t *testing.T) {
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
		s.Assert().NotNil(err)
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
		s.Assert().NotNil(err)
		s.Assert().EqualError(err, "ow no")
	})
}

//----------------------------------------discussions-------------------------------------------

type discussionsSuite struct {
	suite.Suite
	utilsMock   *IUtilMock
	queriesMock *IQueryMock
}

func TestDiscussionService(t *testing.T) {
	suite.Run(t, &discussionsSuite{})
}

func (s *discussionsSuite) SetupSubTest() {
	s.utilsMock = new(IUtilMock)
	s.queriesMock = new(IQueryMock)

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadDiscussionList",
		&discussionListResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*discussionListResponseBody)
		arg.Data.CategoryTopicList.Data = []DiscussionListItem{
			{
				Cursor: "cursor1",
				Node:   DiscussionNode{},
			},
			{
				Cursor: "cursor2",
				Node:   DiscussionNode{},
			},
		}
		arg.Data.CategoryTopicList.TotalNum = 2
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadDiscussion",
		&discussionResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*discussionResponseBody)
		arg.Data.Topic = Discussion{
			Title: "what can that be",
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadDiscussionComments",
		&discussionCommentsResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*discussionCommentsResponseBody)
		arg.Data.TopicComments.Data = []Comment{
			{},
			{},
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadCommentReplies",
		&commentRepliesResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*commentRepliesResponseBody)
		arg.Data.CommentReplies = []Comment{
			{},
			{},
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"takeError",
		mock.Anything,
	).Return(errors.New("error error"))

	s.queriesMock.On(
		"getGraphQLPayloadDiscussionList",
		[]string{}, []string{}, "", "", 0,
	).Return("getGraphQLPayloadDiscussionList")

	s.queriesMock.On(
		"getGraphQLPayloadDiscussionList",
		[]string{}, []string{}, "", "neederror", 0,
	).Return("takeError")

	s.queriesMock.On(
		"getGraphQLPayloadDiscussion",
		int64(12),
	).Return("getGraphQLPayloadDiscussion")

	s.queriesMock.On(
		"getGraphQLPayloadDiscussion",
		int64(-1),
	).Return("takeError")

	s.queriesMock.On(
		"getGraphQLPayloadDiscussionComments",
		int64(10), "", 1, 1,
	).Return("getGraphQLPayloadDiscussionComments")

	s.queriesMock.On(
		"getGraphQLPayloadDiscussionComments",
		int64(10), "neederror", 1, 1,
	).Return("takeError")

	s.queriesMock.On(
		"getGraphQLPayloadCommentReplies",
		int64(10),
	).Return("getGraphQLPayloadCommentReplies")

	s.queriesMock.On(
		"getGraphQLPayloadCommentReplies",
		int64(-1),
	).Return("takeError")
}

func (s *discussionsSuite) TestGetDiscussions() {
	s.Run("should execute getDiscussions without an error", func() {
		result, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getDiscussions([]string{}, []string{}, "", "", 0)
		s.Assert().Nil(err)
		s.Assert().IsType(result, DiscussionList{})
		expected := DiscussionList{
			Data: []DiscussionListItem{
				{
					Cursor: "cursor2",
					Node:   DiscussionNode{},
				},
				{
					Cursor: "cursor1",
					Node:   DiscussionNode{},
				},
			},
			TotalNum: 2,
		}
		s.Assert().Equal(expected.TotalNum, result.TotalNum)
		s.Assert().ElementsMatch(expected.Data, result.Data)
	})

	s.Run("should execute getDiscussions returning with an error", func() {
		_, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getDiscussions([]string{}, []string{}, "", "neederror", 0)
		s.Assert().NotNil(err)
		s.Assert().Error(err, "error error")
	})
}

func (s *discussionsSuite) TestGetDiscussion() {
	s.Run("should execute getDiscussion without an error", func() {
		result, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getDiscussion(12)
		s.Assert().Nil(err)
		s.Assert().IsType(Discussion{}, result)
		expected := Discussion{
			Title: "what can that be",
		}
		s.Assert().Equal(expected, result)
	})

	s.Run("should execute getDiscussion returning with an error", func() {
		_, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getDiscussion(-1)
		s.Assert().NotNil(err)
		s.Assert().Error(err, "error error")
	})
}

func (s *discussionsSuite) TestGetDiscussionComments() {
	s.Run("should execute getDiscussionComments without an error", func() {
		result, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getDiscussionComments(10, "", 1, 1)
		s.Assert().Nil(err)
		s.Assert().IsType([]Comment{}, result)
		expected := []Comment{{}, {}}
		s.Assert().ElementsMatch(expected, result)
	})

	s.Run("should execute getDiscussionComments with an error", func() {
		_, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getDiscussionComments(10, "neederror", 1, 1)
		s.Assert().NotNil(err)
		s.Assert().Error(err, "error error")
	})
}

func (s *discussionsSuite) TestGetCommentRepliess() {
	s.Run("should execute getCommentReplies without an error", func() {
		result, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getCommentReplies(10)
		s.Assert().Nil(err)
		s.Assert().IsType([]Comment{}, result)
		expected := []Comment{{}, {}}
		s.Assert().ElementsMatch(expected, result)
	})

	s.Run("should execute getCommentReplies with an error", func() {
		_, err := (&discussionService{utils: s.utilsMock, queries: s.queriesMock}).getCommentReplies(-1)
		s.Assert().NotNil(err)
		s.Assert().Error(err, "error error")
	})
}
