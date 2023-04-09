package leetcodeapi_test

import (
	"os"
	"testing"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
	"github.com/stretchr/testify/suite"
)

//--------------------------------------contest---------------------------------------------

type contestsSuite struct {
	suite.Suite
}

func TestContestsSuite(t *testing.T) {
	suite.Run(t, &contestsSuite{})
}

func (s *contestsSuite) SetupTest() {
	err := os.Setenv("LEETCODEAPI_ENV", "test")
	s.Assert().Nil(err)
}

func (s *contestsSuite) TearDownTest() {
	err := os.Unsetenv("LEETCODEAPI_ENV")
	s.Assert().Nil(err)
}

func (s *contestsSuite) TestGetContestInfo() {
	responseBody := []byte(`{
		"company": {
			"description":	"description",
			"logo":	"logo",
			"name":	"name"
		},
		"containsPremium": false,
		"contest": {
			"description":	"description",
			"discuss_topic_id":	12,
			"duration":	12,
			"id":	12,
			"is_private":	false,
			"is_virtual":	false,
			"origin_start_time":	12,
			"start_time":	12,
			"title":	"title",
			"title_slug":	"title_slug"
		},
		"questions": [],
		"registered": false
	}`)

	expected := leetcodeapi.Contest{
		Company: leetcodeapi.Company{
			Description: "description",
			Logo:        "logo",
			Name:        "name",
		},
		ContainsPremium: false,
		ContestMeta: leetcodeapi.ContestMeta{
			Description:     "description",
			DiscussTopicId:  12,
			Duration:        12,
			Id:              12,
			IsPrivate:       false,
			IsVirtual:       false,
			OriginStartTime: 12,
			StartTime:       12,
			Title:           "title",
			TitleSlug:       "title_slug",
		},
		Questions:  []leetcodeapi.ContestProblemInfo{},
		Registered: false,
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetContestInfo("contest-10")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *contestsSuite) TestGetContestRanking() {
	responseBody := []byte(`{
		"is_past": false,
		"questions": [],
		"total_rank": [],
		"submissions": [],
		"time": 12.12,
		"user_num": 398
	}`)

	expected := leetcodeapi.ContestRanking{
		IsPast:      false,
		Questions:   []leetcodeapi.ContestProblemInfo{},
		Ranks:       []leetcodeapi.ParticipantDetails{},
		Submissions: []map[string]leetcodeapi.SubmissionInContest{},
		Time:        12.12,
		TotalUser:   398,
		TotalPage:   16,
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetContestRanking("contest-10", 2)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

//--------------------------------------discussion---------------------------------------------

type discussionsSuite struct {
	suite.Suite
}

func TestDiscussionsSuite(t *testing.T) {
	suite.Run(t, &discussionsSuite{})
}

func (s *discussionsSuite) SetupTest() {
	err := os.Setenv("LEETCODEAPI_ENV", "test")
	s.Assert().Nil(err)
}

func (s *discussionsSuite) TearDownTest() {
	err := os.Unsetenv("LEETCODEAPI_ENV")
	s.Assert().Nil(err)
}

func (s *discussionsSuite) TestGetDiscussions() {
	responseBody := []byte(`{
		"data": {
			"categoryTopicList": {
				"edges": [{
					"cursor": "cursor",
					"node": {
						"commentCount": 12,
						"id": "id",
						"pinned": false
					}
				}],
				"totalNum": 1
			}
		}
	}`)
	expected := leetcodeapi.DiscussionList{
		Data: []leetcodeapi.DiscussionListItem{
			{
				Cursor: "cursor",
				Node: leetcodeapi.DiscussionNode{
					CommentCount: 12,
					Id:           "id",
					Pinned:       false,
				},
			},
		},
		TotalNum: 1,
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetDiscussions([]string{}, []string{}, "", "", 0)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *discussionsSuite) TestGetDiscussion() {
	responseBody := []byte(`{
		"data": {
			"topic": {
				"hideFromTrending": false,
				"id": 321321,
				"pinned": false,
				"subscribed": false,
				"title": "title",
				"viewCount": 32
			}
		}
	}`)

	expected := leetcodeapi.Discussion{
		HideFromTrending: false,
		Id:               321321,
		Pinned:           false,
		Subscribed:       false,
		Title:            "title",
		ViewCount:        32,
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetDiscussion(10)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *discussionsSuite) TestGetDiscussionComments() {
	responseBody := []byte(`{
		"data": {
			"topicComments": {
				"data": [{
					"id": 123123,
					"pinned": false,
					"post": {
						"content": "this is a comment"
					}
				}]
			}
		}
	}`)

	expected := []leetcodeapi.Comment{{
		Id:     123123,
		Pinned: false,
		Post: leetcodeapi.Post{
			Content: "this is a comment",
		},
	}}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetDiscussionComments(10, "", 1, 1)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *discussionsSuite) TestGetCommentReplies() {
	responseBody := []byte(`{
		"data": {
			"commentReplies": [{
				"id": 123123,
				"pinned": false,
				"post": {
					"content": "this is a comment"
				}
			}]
		}
	}`)

	expected := []leetcodeapi.Comment{{
		Id:     123123,
		Pinned: false,
		Post: leetcodeapi.Post{
			Content: "this is a comment",
		},
	}}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetCommentReplies(10)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

//--------------------------------------problem---------------------------------------------

type problemsSuite struct {
	suite.Suite
}

func TestProblemsSuite(t *testing.T) {
	suite.Run(t, &problemsSuite{})
}

func (s *problemsSuite) SetupTest() {
	err := os.Setenv("LEETCODEAPI_ENV", "test")
	s.Assert().Nil(err)
}

func (s *problemsSuite) TearDownTest() {
	err := os.Unsetenv("LEETCODEAPI_ENV")
	s.Assert().Nil(err)
}

func (s *problemsSuite) TestGetAllProblems() {
	responseBody := []byte(`{
		"data": {
			"problemsetQuestionList": {
				"total": 1,
				"questions": [
					{
						"acRate": 23.21,
						"difficulty": "difficulty",
						"isFavor": false,
						"title": "title"
					}
				]
			}
		}
	}`)

	expected := leetcodeapi.ProblemList{
		Total: 1,
		Problems: []leetcodeapi.Problem{
			{
				AcRate:     23.21,
				Difficulty: "difficulty",
				IsFavor:    false,
				Title:      "title",
			},
		},
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetAllProblems(0, 20)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *problemsSuite) TestGetProblemContentByTitleSlug() {
	responseBody := []byte(`{
		"data": {
			"question": {
				"content": "problem content"
			}
		}
	}`)

	expected := leetcodeapi.ProblemContent{
		Content: "problem content",
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetProblemContentByTitleSlug("overthinking")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *problemsSuite) TestGetProblemsByTopic() {
	responseBody := []byte(`{
		"data": {
			"topicTag": {
				"name": "name",
				"slug": "slug",
				"questions": [{
					"acRate": 23.21,
					"difficulty": "difficulty",
					"isFavor": false,
					"title": "title"
				}]
			}
		}
	}`)

	expected := leetcodeapi.ProblemsByTopic{
		TopicName: "name",
		TopicSlug: "slug",
		Questions: []leetcodeapi.Problem{{
			AcRate:     23.21,
			Difficulty: "difficulty",
			IsFavor:    false,
			Title:      "title",
		}},
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetProblemsByTopic("no trust")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *problemsSuite) TestGetTopInterviewProblems() {
	responseBody := []byte(`{
		"data": {
			"problemsetQuestionList": {
				"total": 1,
				"questions": [
					{
						"acRate": 23.21,
						"difficulty": "difficulty",
						"isFavor": false,
						"title": "title"
					}
				]
			}
		}
	}`)

	expected := leetcodeapi.ProblemList{
		Total: 1,
		Problems: []leetcodeapi.Problem{
			{
				AcRate:     23.21,
				Difficulty: "difficulty",
				IsFavor:    false,
				Title:      "title",
			},
		},
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetTopInterviewProblems(0, 10)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

//--------------------------------------users---------------------------------------------

type usersSuite struct {
	suite.Suite
}

func TestUsersSuite(t *testing.T) {
	suite.Run(t, &usersSuite{})
}

func (s *usersSuite) SetupTest() {
	err := os.Setenv("LEETCODEAPI_ENV", "test")
	s.Assert().Nil(err)
}

func (s *usersSuite) TearDownTest() {
	err := os.Unsetenv("LEETCODEAPI_ENV")
	s.Assert().Nil(err)
}

func (s *usersSuite) TestGetUserPublicProfile() {
	responseBody := []byte(`{
		"data": {
			"matchedUser": {
				"githubUrl": "githubUrl",
				"linkedinUrl": "linkedinUrl",
				"profile": {
					"aboutMe": "aboutMe",
					"categoryDiscussCount": 4,
					"company": "company",
					"countryName": "countryName",
					"jobTitle": "jobTitle"
				},
				"twitterUrl": "twitterUrl",
				"username": "username"
			}
		}
	}`)

	expected := leetcodeapi.UserPublicProfile{
		GithubUrl:   "githubUrl",
		LinkedinUrl: "linkedinUrl",
		Profile: leetcodeapi.UserProfile{
			AboutMe:              "aboutMe",
			CategoryDiscussCount: 4,
			Company:              "company",
			CountryName:          "countryName",
			JobTitle:             "jobTitle",
		},
		TwitterUrl: "twitterUrl",
		Username:   "username",
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetUserPublicProfile("tourist")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *usersSuite) TestGetUserSolveCountByProblemTag() {
	responseBody := []byte(`{
		"data": {
			"matchedUser": {
				"tagProblemCounts": {
					"advanced": [{
						"problemsSolved": 1,
						"tagName": "tag1"
					}],
					"fundamental": [{
						"problemsSolved": 2,
						"tagName": "tag2"
					}]
				}
			}
		}
	}`)

	expected := leetcodeapi.TagProblemCounts{
		Advanced: []leetcodeapi.TagCount{{
			ProblemsSolved: 1,
			TagName:        "tag1",
		}},
		Fundamental: []leetcodeapi.TagCount{{
			ProblemsSolved: 2,
			TagName:        "tag2",
		}},
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetUserSolveCountByProblemTag("tourist")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *usersSuite) TestGetUserContestRankingHistory() {
	responseBody := []byte(`{
		"data": {
			"userContestRanking": {
				"globalRanking": 2,
				"rating": 23.5,
				"topPercentage": 21.5,
				"totalParticipants": 2
			},
			"userContestRankingHistory": [{
				"attended": false,
				"contest": {
					"title": "title"
				},
				"totalProblems": 2
			}]
		}
	}`)

	expected := leetcodeapi.UserContestRankingDetails{
		UserContestRanking: leetcodeapi.UserContestRanking{
			GlobalRanking:     2,
			Rating:            23.5,
			TopPercentage:     21.5,
			TotalParticipants: 2,
		},
		UserContestRankingHistory: []leetcodeapi.UserContestRankingHistory{{
			Attended:      false,
			TotalProblems: 2,
		}},
	}
	expected.UserContestRankingHistory[0].Contest.Title = "title"

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetUserContestRankingHistory("tourist")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *usersSuite) TestGetUserSolveCountByDifficulty() {
	responseBody := []byte(`{
		"data": {
			"allQuestionsCount": [{
				"count": 2,
				"difficulty": "hard"
			}],
			"matchedUser": {
				"problemsSolvedBeatsStats": [{
					"percentage": 20.2,
					"difficulty": "hard"
				}],
				"submitStatsGlobal": {
					"acSubmissionNum": [{
						"count": 2,
						"difficulty": "hard"
					}]
				}
			}
		}
	}`)

	expected := leetcodeapi.UserSolveCountByDifficultyDetails{
		AllQuestionsCount: []leetcodeapi.DifficultyCount{{
			Count:      2,
			Difficulty: "hard",
		}},
		SolveCount: leetcodeapi.UserSolveCountByDifficulty{
			BeatsStats: []leetcodeapi.DifficultyPercentage{{
				Percentage: 20.2,
				Difficulty: "hard",
			}},
		},
	}
	expected.SolveCount.SubmitStatsGlobal.AcSubmissionNum = []leetcodeapi.DifficultyCount{{
		Count:      2,
		Difficulty: "hard",
	}}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetUserSolveCountByDifficulty("tourist")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *usersSuite) TestGetUserProfileCalendar() {
	responseBody := []byte(`{
		"data": {
			"matchedUser": {
				"userCalendar": {
					"activeYears": [2021, 2022],
					"streak": 3,
					"submissionCalendar": "nice way to store it",
					"totalActiveDays": 3
				}
			}
		}
	}`)

	expected := leetcodeapi.UserCalendar{
		ActiveYears:        []int{2021, 2022},
		Streak:             3,
		SubmissionCalendar: "nice way to store it",
		TotalActiveDays:    3,
	}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetUserProfileCalendar("tourist")

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}

func (s *usersSuite) TestGetUserRecentAcSubmissions() {
	responseBody := []byte(`{
		"data": {
			"recentAcSubmissionList": [{
				"id": "id",
				"timestamp": "timestamp",
				"title": "title",
				"titleSlug": "titleSlug"
			}]
		}
	}`)

	expected := []leetcodeapi.AcSubmission{{
		Id:        "id",
		Timestamp: "timestamp",
		Title:     "title",
		TitleSlug: "titleSlug",
	}}

	server := leetcodeapi.GetMockedHttpServer(responseBody, 200)
	defer server.Close()

	result, err := leetcodeapi.GetUserRecentAcSubmissions("tourist", 2)

	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)
}
