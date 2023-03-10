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

//----------------------------------------problems-------------------------------------------

type problemsSuite struct {
	suite.Suite
	utilsMock   *IUtilMock
	queriesMock *IQueryMock
}

func TestProblemsService(t *testing.T) {
	suite.Run(t, &problemsSuite{})
}

func (s *problemsSuite) SetupSubTest() {
	s.utilsMock = new(IUtilMock)
	s.queriesMock = new(IQueryMock)

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadAllProblems",
		&problemsetListResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*problemsetListResponseBody)
		arg.Data.ProblemsetQuestionList.Problems = []Problem{{}, {}}
		arg.Data.ProblemsetQuestionList.Total = 2
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadProblemContent",
		&problemContentResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*problemContentResponseBody)
		arg.Data.Question = ProblemContent{
			Content: "what a content",
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadProblemsByTopic",
		&problemsByTopicResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*problemsByTopicResponseBody)
		arg.Data.TopicTag.TopicName = "topic is unnecessary"
		arg.Data.TopicTag.Questions = []Problem{{}, {}}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadTopInterviewProblems",
		&problemsetListResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*problemsetListResponseBody)
		arg.Data.ProblemsetQuestionList.Problems = []Problem{{}, {}}
		arg.Data.ProblemsetQuestionList.Total = 2
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"takeError",
		mock.Anything,
	).Return(errors.New("ha ha ha"))
}

func (s *problemsSuite) TestGetAllProblems() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadAllProblems").Return("getGraphQLPayloadAllProblems")
		result, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getAllProblems()
		s.Assert().Nil(err)
		s.Assert().IsType(ProblemList{}, result)
		expected := ProblemList{
			Problems: []Problem{{}, {}},
			Total:    2,
		}
		s.Assert().Equal(expected.Total, result.Total)
		s.Assert().ElementsMatch(expected.Problems, result.Problems)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadAllProblems").Return("takeError")
		_, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getAllProblems()
		s.Assert().NotNil(err)
		s.Assert().Error(err, "ha ha ha")
	})
}

func (s *problemsSuite) TestGetProblemContentByTitleSlug() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadProblemContent", "schizophrenia").Return("getGraphQLPayloadProblemContent")
		result, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getProblemContentByTitleSlug("schizophrenia")
		s.Assert().Nil(err)
		s.Assert().IsType(ProblemContent{}, result)
		expected := ProblemContent{
			Content: "what a content",
		}
		s.Assert().Equal(expected, result)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadProblemContent", "schizophrenia").Return("takeError")
		_, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getProblemContentByTitleSlug("schizophrenia")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "ha ha ha")
	})
}

func (s *problemsSuite) TestGetProblemsByTopic() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadProblemsByTopic", "schizophrenia").Return("getGraphQLPayloadProblemsByTopic")
		result, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getProblemsByTopic("schizophrenia")
		s.Assert().Nil(err)
		s.Assert().IsType(ProblemsByTopic{}, result)
		expected := ProblemsByTopic{
			TopicName: "topic is unnecessary",
			Questions: []Problem{{}, {}},
		}
		s.Assert().Equal(expected.TopicName, result.TopicName)
		s.Assert().ElementsMatch(expected.Questions, result.Questions)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadProblemsByTopic", "schizophrenia").Return("takeError")
		_, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getProblemsByTopic("schizophrenia")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "ha ha ha")
	})
}

func (s *problemsSuite) TestGetTopInterviewProblems() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadTopInterviewProblems").Return("getGraphQLPayloadTopInterviewProblems")
		result, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getTopInterviewProblems()
		s.Assert().Nil(err)
		s.Assert().IsType(ProblemList{}, result)
		expected := ProblemList{
			Problems: []Problem{{}, {}},
			Total:    2,
		}
		s.Assert().Equal(expected.Total, result.Total)
		s.Assert().ElementsMatch(expected.Problems, result.Problems)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadTopInterviewProblems").Return("takeError")
		_, err := (&problemsService{utils: s.utilsMock, queries: s.queriesMock}).getTopInterviewProblems()
		s.Assert().NotNil(err)
		s.Assert().Error(err, "ha ha ha")
	})
}

//----------------------------------------users-------------------------------------------

type usersSuite struct {
	suite.Suite
	utilsMock   *IUtilMock
	queriesMock *IQueryMock
}

func TestUsersService(t *testing.T) {
	suite.Run(t, &usersSuite{})
}

func (s *usersSuite) SetupSubTest() {
	s.utilsMock = new(IUtilMock)
	s.queriesMock = new(IQueryMock)

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadUserPublicProfile",
		&userPublicProfileReponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*userPublicProfileReponseBody)
		arg.Data.MatchedUser = UserPublicProfile{
			Username: "tourist",
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadUserSolveCountByTag",
		&userSolveCountByTagResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*userSolveCountByTagResponseBody)
		arg.Data.MatchedUser.TagProblemCounts.Advanced = []TagCount{
			{ProblemsSolved: 100, TagName: "name", TagSlug: "slug"},
			{ProblemsSolved: 1, TagName: "n", TagSlug: "s"},
		}
		arg.Data.MatchedUser.TagProblemCounts.Fundamental = []TagCount{
			{ProblemsSolved: 23, TagName: "name", TagSlug: "slug"},
		}
		arg.Data.MatchedUser.TagProblemCounts.Intermediate = []TagCount{
			{ProblemsSolved: 253, TagName: "name", TagSlug: "slug"},
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadUserContestRankingHistory",
		&userContestRankingHistoryResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*userContestRankingHistoryResponseBody)
		arg.Data = UserContestRankingDetails{
			UserContestRanking:        UserContestRanking{GlobalRanking: 1},
			UserContestRankingHistory: []UserContestRankingHistory{{Ranking: 1}, {Ranking: 2}},
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadUserSolveCountByDifficulty",
		&userSolveCountByDifficultyResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*userSolveCountByDifficultyResponseBody)
		arg.Data.SolveCount.BeatsStats = []DifficultyPercentage{{}}
		arg.Data.SolveCount.SubmitStatsGlobal.AcSubmissionNum = []DifficultyCount{{}}
		arg.Data.AllQuestionsCount = []DifficultyCount{{}}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadUserProfileCalendar",
		&userProfileCalendarResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*userProfileCalendarResponseBody)
		arg.Data.MatchedUser.UserCalendar = UserCalendar{
			TotalActiveDays: 20,
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"getGraphQLPayloadUserRecentAcSubmissions",
		&userRecentAcSubmissionsResponseBody{},
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*userRecentAcSubmissionsResponseBody)
		arg.Data.RecentAcSubmissionList = []AcSubmission{
			{Title: "tor jonno akash theke peja"},
			{Title: "ek tukro megh enechi veja"},
		}
	})

	s.utilsMock.On(
		"MakeGraphQLRequest",
		"takeError",
		mock.Anything,
	).Return(errors.New("oniket prantor"))
}

func (s *usersSuite) TestGetUserPublicProfile() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserPublicProfile", "tourist").Return("getGraphQLPayloadUserPublicProfile")
		result, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserPublicProfile("tourist")
		s.Assert().Nil(err)
		s.Assert().IsType(UserPublicProfile{}, result)
		expected := UserPublicProfile{
			Username: "tourist",
		}
		s.Assert().Equal(expected, result)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserPublicProfile", "tourist").Return("takeError")
		_, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserPublicProfile("tourist")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "oniket prantor")
	})
}

func (s *usersSuite) TestGetUserSolveCountByProblemTag() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserSolveCountByTag", "tourist").Return("getGraphQLPayloadUserSolveCountByTag")
		result, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserSolveCountByProblemTag("tourist")
		s.Assert().Nil(err)
		s.Assert().IsType(TagProblemCounts{}, result)
		expected := TagProblemCounts{
			Advanced: []TagCount{
				{ProblemsSolved: 100, TagName: "name", TagSlug: "slug"},
				{ProblemsSolved: 1, TagName: "n", TagSlug: "s"},
			},
			Fundamental: []TagCount{
				{ProblemsSolved: 23, TagName: "name", TagSlug: "slug"},
			},
			Intermediate: []TagCount{
				{ProblemsSolved: 253, TagName: "name", TagSlug: "slug"},
			},
		}
		s.Assert().ElementsMatch(expected.Advanced, result.Advanced)
		s.Assert().ElementsMatch(expected.Fundamental, result.Fundamental)
		s.Assert().ElementsMatch(expected.Intermediate, result.Intermediate)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserSolveCountByTag", "tourist").Return("takeError")
		_, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserSolveCountByProblemTag("tourist")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "oniket prantor")
	})
}

func (s *usersSuite) TestGetUserContestRankingHistory() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserContestRankingHistory", "tourist").Return("getGraphQLPayloadUserContestRankingHistory")
		result, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserContestRankingHistory("tourist")
		s.Assert().Nil(err)
		s.Assert().IsType(UserContestRankingDetails{}, result)
		expected := UserContestRankingDetails{
			UserContestRanking:        UserContestRanking{GlobalRanking: 1},
			UserContestRankingHistory: []UserContestRankingHistory{{Ranking: 2}, {Ranking: 1}},
		}
		s.Assert().Equal(expected.UserContestRanking, result.UserContestRanking)
		s.Assert().ElementsMatch(expected.UserContestRankingHistory, result.UserContestRankingHistory)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserContestRankingHistory", "tourist").Return("takeError")
		_, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserContestRankingHistory("tourist")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "oniket prantor")
	})
}

func (s *usersSuite) TestGetUserSolveCountByDifficulty() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserSolveCountByDifficulty", "tourist").Return("getGraphQLPayloadUserSolveCountByDifficulty")
		result, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserSolveCountByDifficulty("tourist")
		s.Assert().Nil(err)
		s.Assert().IsType(UserSolveCountByDifficultyDetails{}, result)
		expected := UserSolveCountByDifficultyDetails{
			SolveCount: UserSolveCountByDifficulty{
				BeatsStats: []DifficultyPercentage{{}},
			},
			AllQuestionsCount: []DifficultyCount{{}},
		}
		expected.SolveCount.SubmitStatsGlobal.AcSubmissionNum = []DifficultyCount{{}}
		s.Assert().ElementsMatch(expected.SolveCount.BeatsStats, result.SolveCount.BeatsStats)
		s.Assert().ElementsMatch(expected.SolveCount.SubmitStatsGlobal.AcSubmissionNum, result.SolveCount.SubmitStatsGlobal.AcSubmissionNum)
		s.Assert().ElementsMatch(expected.AllQuestionsCount, result.AllQuestionsCount)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserSolveCountByDifficulty", "tourist").Return("takeError")
		_, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserSolveCountByDifficulty("tourist")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "oniket prantor")
	})
}

func (s *usersSuite) TestGetUserProfileCalendar() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserProfileCalendar", "tourist").Return("getGraphQLPayloadUserProfileCalendar")
		result, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserProfileCalendar("tourist")
		s.Assert().Nil(err)
		s.Assert().IsType(UserCalendar{}, result)
		expected := UserCalendar{
			TotalActiveDays: 20,
		}
		s.Assert().Equal(expected, result)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserProfileCalendar", "tourist").Return("takeError")
		_, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserProfileCalendar("tourist")
		s.Assert().NotNil(err)
		s.Assert().Error(err, "oniket prantor")
	})
}

func (s *usersSuite) TestGetUserRecentAcSubmissions() {
	s.Run("should execute without an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserRecentAcSubmissions", "tourist", 2).Return("getGraphQLPayloadUserRecentAcSubmissions")
		result, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserRecentAcSubmissions("tourist", 2)
		s.Assert().Nil(err)
		s.Assert().IsType([]AcSubmission{}, result)
		expected := []AcSubmission{
			{Title: "tor jonno akash theke peja"},
			{Title: "ek tukro megh enechi veja"},
		}
		s.Assert().ElementsMatch(expected, result)
	})

	s.Run("should execute with an error", func() {
		s.queriesMock.On("getGraphQLPayloadUserRecentAcSubmissions", "tourist", 2).Return("takeError")
		_, err := (&usersService{utils: s.utilsMock, queries: s.queriesMock}).getUserRecentAcSubmissions("tourist", 2)
		s.Assert().NotNil(err)
		s.Assert().Error(err, "oniket prantor")
	})
}

//-------------------------------queries-------------------------------

type queriesSuite struct {
	suite.Suite
	utilsMock *IUtilMock
}

func TestQueryService(t *testing.T) {
	suite.Run(t, &queriesSuite{})
}

func (s *queriesSuite) SetupSubTest() {
	s.utilsMock = new(IUtilMock)
}

func (s *queriesSuite) TestGetGraphQLPayloadAllProblems() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadAllProblems()
		expected := `{
		"query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        name\n        id\n        slug\n      }\n      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
		"variables": {
			"categorySlug": "",
			"skip": 0,
			"limit": 50,
			"filters": {}
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadProblemContent() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadProblemContent("Nemesis - kobe")
		expected := `{
		"query": "\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    mysqlSchemas\n  }\n}\n    ",
		"variables": {
			"titleSlug": "Nemesis - kobe"
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadProblemsByTopic() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadProblemsByTopic("Google")
		expected := `{
		"operationName": "getTopicTag",
		"variables": {
			"slug": "Google"
		},
		"query": "query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    slug\n    questions {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        name\n        id\n        slug\n      }\n     companyTags {\n        name\n        slug\n        }\n      }\n    frequencies\n      }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n          }\n    privateFavorites {\n      ...favoriteFields\n          }\n      }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n      }\n  }\n"
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadTopInterviewProblems() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadTopInterviewProblems()
		expected := `{
		"query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        name\n        id\n        slug\n      }\n      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
		"variables": {
			"categorySlug": "",
			"skip": 0,
			"limit": 50,
			"filters": {
				"listId": "top-interview-questions"
			}
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadDiscussionList() {
	s.Run("should return correct value", func() {
		s.utilsMock.On(
			"convertListToString",
			[]string{"item-1", "item-2"},
		).Return(`["item-1","item-2"]`)
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadDiscussionList(
			[]string{"item-1", "item-2"},
			[]string{"item-1", "item-2"},
			"top",
			"google",
			1,
		)
		expected := `{
		"operationName": "categoryTopicList",
		"variables": {
			"orderBy": "top",
			"query": "google",
			"skip": 1,
			"first": 15,
			"tags": ["item-1","item-2"],
			"categories": ["item-1","item-2"]
		},
		"query": "query categoryTopicList($categories: [String!]!, $first: Int!, $orderBy: TopicSortingOption, $skip: Int, $query: String, $tags: [String!]) {\n  categoryTopicList(categories: $categories, orderBy: $orderBy, skip: $skip, query: $query, first: $first, tags: $tags) {\n    ...TopicsList\n    __typename\n  }\n}\n\nfragment TopicsList on TopicConnection {\n  totalNum\n  edges {\n    node {\n      id\n      title\n      commentCount\n      viewCount\n      pinned\n      tags {\n        name\n        slug\n        __typename\n      }\n      post {\n        id\n        voteCount\n        creationDate\n        isHidden\n        author {\n          username\n          isActive\n          nameColor\n          activeBadge {\n            displayName\n            icon\n            __typename\n          }\n          profile {\n            userAvatar\n            __typename\n          }\n          __typename\n        }\n        status\n        coinRewards {\n          ...CoinReward\n          __typename\n        }\n        __typename\n      }\n      lastComment {\n        id\n        post {\n          id\n          author {\n            isActive\n            username\n            __typename\n          }\n          peek\n          creationDate\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    cursor\n    __typename\n  }\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadDiscussion() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadDiscussion(20)
		expected := `{
		"operationName": "DiscussTopic",
		"variables": {
			"topicId": 20
		},
		"query": "query DiscussTopic($topicId: Int!) {\n  topic(id: $topicId) {\n    id\n    viewCount\n    topLevelCommentCount\n    subscribed\n    title\n    pinned\n    tags\n    hideFromTrending\n    post {\n      ...DiscussPost\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadDiscussionComments() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadDiscussionComments(
			20, "big brother", 2, 10,
		)
		expected := `{
		"operationName": "discussComments",
		"variables": {
			"orderBy": "big brother",
			"pageNo": 2,
			"numPerPage": 10,
			"topicId": 20
		},
		"query": "query discussComments($topicId: Int!, $orderBy: String = \"newest_to_oldest\", $pageNo: Int = 1, $numPerPage: Int = 10) {\n  topicComments(topicId: $topicId, orderBy: $orderBy, pageNo: $pageNo, numPerPage: $numPerPage) {\n    data {\n      id\n      pinned\n      pinnedBy {\n        username\n        __typename\n      }\n      post {\n        ...DiscussPost\n        __typename\n      }\n      numChildren\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadCommentReplies() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadCommentReplies(20)
		expected := `{
		"operationName": "fetchCommentReplies",
		"variables": {
			"commentId": 20
		},
		"query": "query fetchCommentReplies($commentId: Int!) {\n  commentReplies(commentId: $commentId) {\n    id\n    pinned\n    pinnedBy {\n      username\n      __typename\n    }\n    post {\n      ...DiscussPost\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadUserPublicProfile() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadUserPublicProfile("dustyRAIN")
		expected := `{
		"query": "\n    query userPublicProfile($username: String!) {\n  matchedUser(username: $username) {\n    contestBadge {\n      name\n      expired\n      hoverText\n      icon\n    }\n    username\n    githubUrl\n    twitterUrl\n    linkedinUrl\n    profile {\n      ranking\n      userAvatar\n      realName\n      aboutMe\n      school\n      websites\n      countryName\n      company\n      jobTitle\n      skillTags\n      postViewCount\n      postViewCountDiff\n      reputation\n      reputationDiff\n      solutionCount\n      solutionCountDiff\n      categoryDiscussCount\n      categoryDiscussCountDiff\n    }\n  }\n}\n    ",
		"variables": {
			"username": "dustyRAIN"
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadUserSolveCountByTag() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadUserSolveCountByTag("dustyRAIN")
		expected := `{
		"query": "\n    query skillStats($username: String!) {\n  matchedUser(username: $username) {\n    tagProblemCounts {\n      advanced {\n        tagName\n        tagSlug\n        problemsSolved\n      }\n      intermediate {\n        tagName\n        tagSlug\n        problemsSolved\n      }\n      fundamental {\n        tagName\n        tagSlug\n        problemsSolved\n      }\n    }\n  }\n}\n    ",
		"variables": {
			"username": "dustyRAIN"
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadUserContestRankingHistory() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadUserContestRankingHistory("dustyRAIN")
		expected := `{
		"query": "\n    query userContestRankingInfo($username: String!) {\n  userContestRanking(username: $username) {\n    attendedContestsCount\n    rating\n    globalRanking\n    totalParticipants\n    topPercentage\n    badge {\n      name\n    }\n  }\n  userContestRankingHistory(username: $username) {\n    attended\n    trendDirection\n    problemsSolved\n    totalProblems\n    finishTimeInSeconds\n    rating\n    ranking\n    contest {\n      title\n      startTime\n    }\n  }\n}\n    ",
		"variables": {
			"username": "dustyRAIN"
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadUserSolveCountByDifficulty() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadUserSolveCountByDifficulty("dustyRAIN")
		expected := `{
		"query": "\n    query userProblemsSolved($username: String!) {\n  allQuestionsCount {\n    difficulty\n    count\n  }\n  matchedUser(username: $username) {\n    problemsSolvedBeatsStats {\n      difficulty\n      percentage\n    }\n    submitStatsGlobal {\n      acSubmissionNum {\n        difficulty\n        count\n      }\n    }\n  }\n}\n    ",
		"variables": {
			"username": "dustyRAIN"
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadUserProfileCalendar() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadUserProfileCalendar("dustyRAIN")
		expected := `{
		"query": "\n    query userProfileCalendar($username: String!, $year: Int) {\n  matchedUser(username: $username) {\n    userCalendar(year: $year) {\n      activeYears\n      streak\n      totalActiveDays\n      dccBadges {\n        timestamp\n        badge {\n          name\n          icon\n        }\n      }\n      submissionCalendar\n    }\n  }\n}\n    ",
		"variables": {
			"username": "dustyRAIN"
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}

func (s *queriesSuite) TestGetGraphQLPayloadUserRecentAcSubmissions() {
	s.Run("should return correct value", func() {
		actual := (&queryService{utils: s.utilsMock}).getGraphQLPayloadUserRecentAcSubmissions("dustyRAIN", 10)
		expected := `{
		"query": "\n    query recentAcSubmissions($username: String!, $limit: Int!) {\n  recentAcSubmissionList(username: $username, limit: $limit) {\n    id\n    title\n    titleSlug\n    timestamp\n  }\n}\n    ",
		"variables": {
			"username": "dustyRAIN",
			"limit": 10
		}
	}`
		s.Assert().Equal(expected, actual)
	})
}
