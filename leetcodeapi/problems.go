package leetcodeapi

import (
	"log"
)

func GetAllProblems() (ProblemList, error) {
	return getAllProblems(Util{}, query{})
}

func GetProblemContentByTitleSlug(titleSlug string) (ProblemContent, error) {
	return getProblemContentByTitleSlug(titleSlug, Util{}, query{})
}

func GetProblemsByTopic(topicSlug string) (ProblemsByTopic, error) {
	return getProblemsByTopic(topicSlug, Util{}, query{})
}

func GetTopInterviewProblems() (ProblemList, error) {
	return getTopInterviewProblems(Util{}, query{})
}

/*
---------------------------------------------------------------------------------------
*/

func getAllProblems(utils IUtil, queries IQuery) (ProblemList, error) {
	var result problemsetListResponseBody
	err := utils.MakeGraphQLRequest(queries.getGraphQLPayloadAllProblems(), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}

func getProblemContentByTitleSlug(titleSlug string, utils IUtil, queries IQuery) (ProblemContent, error) {
	var result problemContentResponseBody
	err := utils.MakeGraphQLRequest(queries.getGraphQLPayloadProblemContent(titleSlug), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemContent{}, err
	}

	return result.Data.Question, nil
}

func getProblemsByTopic(topicSlug string, utils IUtil, queries IQuery) (ProblemsByTopic, error) {
	var result problemsByTopicResponseBody
	err := utils.MakeGraphQLRequest(queries.getGraphQLPayloadProblemsByTopic(topicSlug), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemsByTopic{}, err
	}

	return result.Data.TopicTag, nil
}

func getTopInterviewProblems(utils IUtil, queries IQuery) (ProblemList, error) {
	var result problemsetListResponseBody
	err := utils.MakeGraphQLRequest(queries.getGraphQLPayloadTopInterviewProblems(), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}
