package leetcodeapi

import (
	"log"
)

func GetAllProblems() (ProblemList, error) {
	var result problemsetListResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadAllProblems(), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
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

func GetProblemsByTopic(topicSlug string) (ProblemsByTopic, error) {
	var result problemsByTopicResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadProblemContent(getGraphQLPayloadProblemsByTopic(topicSlug)), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemsByTopic{}, err
	}

	return result.Data.TopicTag, nil
}

func GetTopInterviewProblems() (ProblemList, error) {
	var result problemsetListResponseBody
	err := MakeGraphQLRequest(getGraphQLPayloadProblemContent(getGraphQLPayloadTopInterviewProblems()), &result)

	if err != nil {
		log.Printf(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}
