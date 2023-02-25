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
