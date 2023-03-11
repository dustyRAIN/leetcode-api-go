package leetcodeapi_test

import (
	"os"
	"testing"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
	"github.com/stretchr/testify/suite"
)

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
