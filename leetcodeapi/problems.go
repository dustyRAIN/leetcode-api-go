package leetcodeapi

import (
	"log"
)

func GetAllProblems() (ProblemList, error) {
	utils := &Util{}
	return (&problemsService{utils: utils, queries: &queryService{utils: utils}}).getAllProblems()
}

func GetProblemContentByTitleSlug(titleSlug string) (ProblemContent, error) {
	utils := &Util{}
	return (&problemsService{utils: utils, queries: &queryService{utils: utils}}).getProblemContentByTitleSlug(titleSlug)
}

func GetProblemsByTopic(topicSlug string) (ProblemsByTopic, error) {
	utils := &Util{}
	return (&problemsService{utils: utils, queries: &queryService{utils: utils}}).getProblemsByTopic(topicSlug)
}

func GetTopInterviewProblems() (ProblemList, error) {
	utils := &Util{}
	return (&problemsService{utils: utils, queries: &queryService{utils: utils}}).getTopInterviewProblems()
}

/*
---------------------------------------------------------------------------------------
*/

type problemsService struct {
	utils   IUtil
	queries IQuery
}

func (p *problemsService) getAllProblems() (ProblemList, error) {
	var result problemsetListResponseBody
	err := p.utils.MakeGraphQLRequest(p.queries.getGraphQLPayloadAllProblems(), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}

func (p *problemsService) getProblemContentByTitleSlug(titleSlug string) (ProblemContent, error) {
	var result problemContentResponseBody
	err := p.utils.MakeGraphQLRequest(p.queries.getGraphQLPayloadProblemContent(titleSlug), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemContent{}, err
	}

	return result.Data.Question, nil
}

func (p *problemsService) getProblemsByTopic(topicSlug string) (ProblemsByTopic, error) {
	var result problemsByTopicResponseBody
	err := p.utils.MakeGraphQLRequest(p.queries.getGraphQLPayloadProblemsByTopic(topicSlug), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemsByTopic{}, err
	}

	return result.Data.TopicTag, nil
}

func (p *problemsService) getTopInterviewProblems() (ProblemList, error) {
	var result problemsetListResponseBody
	err := p.utils.MakeGraphQLRequest(p.queries.getGraphQLPayloadTopInterviewProblems(), &result)

	if err != nil {
		log.Print(err.Error())
		return ProblemList{}, err
	}

	return result.Data.ProblemsetQuestionList, nil
}
