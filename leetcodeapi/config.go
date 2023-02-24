package leetcodeapi

import "fmt"

func getGraphQLPayloadAllProblems() string {
	return `{
		"query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        name\n        id\n        slug\n      }\n      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
		"variables": {
				"categorySlug": "",
				"skip": 0,
				"limit": 50,
				"filters": {}
		}
	}`
}

func getGraphQLPayloadProblemContent(titleSlug string) string {
	return fmt.Sprintf(`{
	    "query": "\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    mysqlSchemas\n  }\n}\n    ",
	    "variables": {
	        "titleSlug": "%v"
	    }
	}`, titleSlug)
}

func getGraphQLPayloadProblemsByTopic(topicStag string) string {
	return fmt.Sprintf(`{
	    "operationName": "getTopicTag",
	    "variables": {
	        "slug": "string"
	    },
	    "query": "query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    slug\n    questions {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        id\n        name\n        slug\n        }\n      companyTags {\n        name\n        slug\n        }\n      }\n    frequencies\n      }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n          }\n    privateFavorites {\n      ...favoriteFields\n          }\n      }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n      }\n  }\n"
	}`)
}
