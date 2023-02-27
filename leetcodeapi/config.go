package leetcodeapi

import "fmt"

const problemCommonFields = `\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        name\n        id\n        slug\n      }\n`

func getGraphQLPayloadAllProblems() string {
	return fmt.Sprintf(`{
		"query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {%v      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
		"variables": {
				"categorySlug": "",
				"skip": 0,
				"limit": 50,
				"filters": {}
		}
	}`, problemCommonFields)
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
	        "slug": "%v"
	    },
	    "query": "query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    slug\n    questions {%v     companyTags {\n        name\n        slug\n        }\n      }\n    frequencies\n      }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n          }\n    privateFavorites {\n      ...favoriteFields\n          }\n      }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n      }\n  }\n"
	}`, topicStag, problemCommonFields)
}

func getGraphQLPayloadTopInterviewProblems() string {
	return fmt.Sprintf(`{
	    "query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {%v      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
	    "variables": {
	        "categorySlug": "",
	        "skip": 0,
	        "limit": 50,
	        "filters": {
	            "listId": "top-interview-questions"
	        }
	    }
	}`, problemCommonFields)
}

func getGraphQLPayloadDiscussionList(categories []string, tags []string, orderBy string, searchQuery string, offset int) string {
	categoryListString := convertListToString(categories)
	tagListString := convertListToString(tags)

	if orderBy == "" {
		if len(searchQuery) > 0 {
			orderBy = "most_relevant"
		} else {
			orderBy = "hot"
		}
	}

	return fmt.Sprintf(`{
	    "operationName": "categoryTopicList",
	    "variables": {
	        "orderBy": "%v",
	        "query": "%v",
	        "skip": %v,
	        "first": 15,
	        "tags": %v,
	        "categories": %v
	    },
	    "query": "query categoryTopicList($categories: [String!]!, $first: Int!, $orderBy: TopicSortingOption, $skip: Int, $query: String, $tags: [String!]) {\n  categoryTopicList(categories: $categories, orderBy: $orderBy, skip: $skip, query: $query, first: $first, tags: $tags) {\n    ...TopicsList\n    __typename\n  }\n}\n\nfragment TopicsList on TopicConnection {\n  totalNum\n  edges {\n    node {\n      id\n      title\n      commentCount\n      viewCount\n      pinned\n      tags {\n        name\n        slug\n        __typename\n      }\n      post {\n        id\n        voteCount\n        creationDate\n        isHidden\n        author {\n          username\n          isActive\n          nameColor\n          activeBadge {\n            displayName\n            icon\n            __typename\n          }\n          profile {\n            userAvatar\n            __typename\n          }\n          __typename\n        }\n        status\n        coinRewards {\n          ...CoinReward\n          __typename\n        }\n        __typename\n      }\n      lastComment {\n        id\n        post {\n          id\n          author {\n            isActive\n            username\n            __typename\n          }\n          peek\n          creationDate\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    cursor\n    __typename\n  }\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`, orderBy, searchQuery, offset, tagListString, categoryListString)
}

func getGraphQLPayloadDiscussion(topicId int64) string {
	return fmt.Sprintf(`{
	    "operationName": "DiscussTopic",
	    "variables": {
	        "topicId": %v
	    },
	    "query": "query DiscussTopic($topicId: Int!) {\n  topic(id: $topicId) {\n    id\n    viewCount\n    topLevelCommentCount\n    subscribed\n    title\n    pinned\n    tags\n    hideFromTrending\n    post {\n      ...DiscussPost\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`, topicId)
}
