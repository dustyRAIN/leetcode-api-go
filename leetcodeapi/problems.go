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

type ProblemList struct {
	Total    int       `json:"total"`
	Problems []Problem `json:"questions"`
}

type problemsetListResponseBody struct {
	Data struct {
		ProblemsetQuestionList ProblemList `josn:"problemsetQuestionList"`
	} `json:"data"`
}

func GetAllProblems() (ProblemList, error) {
	var result problemsetListResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadAllProblems(), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}

/*
-----------------------------------------------------
*/

type ProblemContent struct {
	Content string `json:"content"`
}

type problemContentResponseBody struct {
	Data struct {
		Question ProblemContent `json:"question"`
	} `json:"data"`
}

func GetProblemContentByTitleSlug(titleSlug string) (ProblemContent, error) {
	var result problemContentResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadProblemContent(titleSlug), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemContent{}, err
	}

	return result.Data.Question, nil
}

/*
-----------------------------------------------------
*/

type ProblemsByTopic struct {
	TopicName string    `json:"name"`
	TopicSlug string    `json:"slug"`
	Questions []Problem `json:"questions"`
}

type problemsByTopicResponseBody struct {
	Data struct {
		TopicTag ProblemsByTopic `josn:"topicTag"`
	} `json:"data"`
}

func GetProblemsByTopic(topicSlug string) (ProblemsByTopic, error) {
	var result problemsByTopicResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadProblemContent(getGraphQLPayloadProblemsByTopic(topicSlug)), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemsByTopic{}, err
	}

	return result.Data.TopicTag, nil
}

/*
-----------------------------------------------------
*/

func GetTopInterviewProblems() (ProblemList, error) {
	var result problemsetListResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadProblemContent(getGraphQLPayloadTopInterviewProblems()), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}
