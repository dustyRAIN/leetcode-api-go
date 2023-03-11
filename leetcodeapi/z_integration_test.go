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
