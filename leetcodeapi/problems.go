package leetcodeapi

import (
	"log"
)

type TopicTag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Problem struct {
	AcRate             float32    `json:"acRate"`
	Difficulty         string     `json:"difficulty"`
	FreqBar            float32    `json:"freqBar"`
	FrontendQuestionId string     `json:"frontendQuestionId"`
	IsFavor            bool       `json:"isFavor"`
	PaidOnly           bool       `json:"paidOnly"`
	Status             string     `json:"status"`
	QuestionId         string     `json:"questionId"`
	Title              string     `json:"title"`
	TitleSlug          string     `json:"titleSlug"`
	Stats              string     `json:"stats"`
	TopicTags          []TopicTag `json:"topicTags"`
}

type ProblemsetListResponseBody struct {
	Data struct {
		ProblemsetQuestionList struct {
			Total     int       `json:"total"`
			Questions []Problem `json:"questions"`
		} `josn:"problemsetQuestionList"`
	} `json:"data"`
}

func GetAllProblems() (ProblemsetListResponseBody, error) {
	var result ProblemsetListResponseBody
	err := makeHttpRequest(
		"GET",
		"https://leetcode.com/graphql/",
		"application/json",
		getGraphQLPayloadAllProblems(),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return ProblemsetListResponseBody{}, err
	}

	return result, nil
}

/*
-----------------------------------------------------
*/

type ProblemContentResponseBody struct {
	Data struct {
		Question struct {
			Content string `json:"content"`
		} `json:"question"`
	} `json:"data"`
}

func GetProblemContentByTitleSlug(titleSlug string) (ProblemContentResponseBody, error) {
	var result ProblemContentResponseBody
	err := makeHttpRequest(
		"GET",
		"https://leetcode.com/graphql/",
		"application/json",
		getGraphQLPayloadProblemContent(titleSlug),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return ProblemContentResponseBody{}, err
	}

	return result, nil
}

/*
-----------------------------------------------------
*/

type ProblemsByTopicResponseBody struct {
	Data struct {
		TopicTag struct {
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
			Questions []Problem `json:"questions"`
		} `josn:"topicTag"`
	} `json:"data"`
}

func GetProblemsByTopic(topicSlug string) (ProblemsByTopicResponseBody, error) {
	var result ProblemsByTopicResponseBody
	err := makeHttpRequest(
		"GET",
		"https://leetcode.com/graphql/",
		"application/json",
		getGraphQLPayloadProblemsByTopic(topicSlug),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return ProblemsByTopicResponseBody{}, err
	}

	return result, nil
}
