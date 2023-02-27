package leetcodeapi

import (
	"fmt"
	"log"
)

type ContestInfo struct {
	Description     string `json:"description"`
	DiscussTopicId  int    `json:"discuss_topic_id"`
	Duration        int    `json:"duration"`
	Id              int    `json:"id"`
	IsPrivate       bool   `json:"is_private"`
	IsVirtual       bool   `json:"is_virtual"`
	OriginStartTime int64  `json:"origin_start_time"`
	StartTime       int64  `json:"start_time"`
	Title           string `json:"title"`
	TitleSlug       string `json:"title_slug"`
}

type ContestProblemInfo struct {
	Credit     int    `json:"credit"`
	Id         int    `json:"id"`
	QuestionId int    `json:"question_id"`
	Title      string `json:"title"`
	TitleSlug  string `json:"title_slug"`
}

type ContestInfoResponseBody struct {
	Company struct {
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Name        string `json:"name"`
	} `json:"company"`
	ContainsPremium bool                 `json:"containsPremium"`
	Contest         ContestInfo          `json:"contest"`
	Questions       []ContestProblemInfo `json:"questions"`
	Registered      bool                 `json:"registered"`
}

func GetContestInfo(contestSlug string) (ContestInfoResponseBody, error) {
	var result ContestInfoResponseBody
	err := makeHttpRequest(
		"GET",
		fmt.Sprintf("https://leetcode.com/contest/api/info/%v/", contestSlug),
		"application/json",
		"",
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return ContestInfoResponseBody{}, err
	}

	return result, nil
}

/*
-----------------------------------------------------
*/

type ParticipantDetails struct {
	ContestId     int    `json:"contest_id"`
	CountryCode   string `json:"country_code"`
	CountryName   string `json:"country_name"`
	DataRegion    string `json:"data_region"`
	FinishTime    int64  `json:"finish_time"`
	GlobalRanking int    `json:"global_ranking"`
	Rank          int    `json:"rank"`
	Score         int    `json:"score"`
	UserBadge     struct {
		DisplayName string `json:"display_name"`
		Icon        string `json:"icon"`
	} `json:"user_badge"`
	UserSlug      string `json:"user_slug"`
	Username      string `json:"username"`
	UsernameColor string `json:"username_color"`
}

type SubmissionInContest struct {
	ContestId    int    `json:"contest_id"`
	DataRegion   string `json:"data_region"`
	Date         int64  `json:"date"`
	FailCount    int    `json:"fail_count"`
	Id           int    `json:"id"`
	QuestionId   int    `json:"question_id"`
	Status       int    `json:"status"`
	SubmissionId int64  `json:"submission_id"`
}

type ContestRankingResponseBody struct {
	IsPast      bool                             `json:"is_past"`
	Questions   []ContestProblemInfo             `json:"questions"`
	Ranks       []ParticipantDetails             `json:"total_rank"`
	Submissions []map[string]SubmissionInContest `json:"submissions"`
	Time        float64                          `json:"time"`
	TotalUser   int                              `json:"user_num"`
	TotalPage   int
}

func GetContestRanking(contestSlug string, page int) (ContestRankingResponseBody, error) {
	var result ContestRankingResponseBody
	err := makeHttpRequest(
		"GET",
		fmt.Sprintf("https://leetcode.com/contest/api/ranking/%v/?pagination=%v&region=global", contestSlug, page),
		"application/json",
		"",
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return ContestRankingResponseBody{}, err
	}

	return result, nil
}
