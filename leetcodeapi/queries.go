package leetcodeapi

import "fmt"

const problemCommonFields = `\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId: questionFrontendId\n      questionId\n      isFavor\n      paidOnly: isPaidOnly\n      status\n      title\n      titleSlug\n      stats\n      topicTags {\n        name\n        id\n        slug\n      }\n`

type IQuery interface {
	getGraphQLPayloadAllProblems(offset int, pageSize int) string
	getGraphQLPayloadProblemContent(titleSlug string) string
	getGraphQLPayloadProblemsByTopic(topicStag string) string
	getGraphQLPayloadTopInterviewProblems(offset int, pageSize int) string
	getGraphQLPayloadDiscussionList(categories []string, tags []string, orderBy string, searchQuery string, offset int) string
	getGraphQLPayloadDiscussion(topicId int64) string
	getGraphQLPayloadDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) string
	getGraphQLPayloadCommentReplies(commentId int64) string
	getGraphQLPayloadUserPublicProfile(username string) string
	getGraphQLPayloadUserSolveCountByTag(username string) string
	getGraphQLPayloadUserContestRankingHistory(username string) string
	getGraphQLPayloadUserSolveCountByDifficulty(username string) string
	getGraphQLPayloadUserProfileCalendar(username string) string
	getGraphQLPayloadUserRecentAcSubmissions(username string, pageSize int) string
}

type queryService struct {
	utils IUtil
}

func (q *queryService) getGraphQLPayloadAllProblems(offset int, pageSize int) string {
	return fmt.Sprintf(`{
		"query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {%v      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
		"variables": {
			"categorySlug": "",
			"skip": %v,
			"limit": %v,
			"filters": {}
		}
	}`, problemCommonFields, offset, pageSize)
}

func (q *queryService) getGraphQLPayloadProblemContent(titleSlug string) string {
	return fmt.Sprintf(`{
		"query": "\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    mysqlSchemas\n  }\n}\n    ",
		"variables": {
			"titleSlug": "%v"
		}
	}`, titleSlug)
}

func (q *queryService) getGraphQLPayloadProblemsByTopic(topicStag string) string {
	return fmt.Sprintf(`{
		"operationName": "getTopicTag",
		"variables": {
			"slug": "%v"
		},
		"query": "query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    slug\n    questions {%v     companyTags {\n        name\n        slug\n        }\n      }\n    frequencies\n      }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n          }\n    privateFavorites {\n      ...favoriteFields\n          }\n      }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n      }\n  }\n"
	}`, topicStag, problemCommonFields)
}

func (q *queryService) getGraphQLPayloadTopInterviewProblems(offset int, pageSize int) string {
	return fmt.Sprintf(`{
		"query": "\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList: questionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    total: totalNum\n    questions: data {%v      hasSolution\n      hasVideoSolution\n    }\n  }\n}\n    ",
		"variables": {
			"categorySlug": "",
			"skip": %v,
			"limit": %v,
			"filters": {
				"listId": "top-interview-questions"
			}
		}
	}`, problemCommonFields, offset, pageSize)
}

func (q *queryService) getGraphQLPayloadDiscussionList(categories []string, tags []string, orderBy string, searchQuery string, offset int) string {
	categoryListString := q.utils.convertListToString(categories)
	tagListString := q.utils.convertListToString(tags)

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

func (q *queryService) getGraphQLPayloadDiscussion(topicId int64) string {
	return fmt.Sprintf(`{
		"operationName": "DiscussTopic",
		"variables": {
			"topicId": %v
		},
		"query": "query DiscussTopic($topicId: Int!) {\n  topic(id: $topicId) {\n    id\n    viewCount\n    topLevelCommentCount\n    subscribed\n    title\n    pinned\n    tags\n    hideFromTrending\n    post {\n      ...DiscussPost\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`, topicId)
}

func (q *queryService) getGraphQLPayloadDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) string {
	return fmt.Sprintf(`{
		"operationName": "discussComments",
		"variables": {
			"orderBy": "%v",
			"pageNo": %v,
			"numPerPage": %v,
			"topicId": %v
		},
		"query": "query discussComments($topicId: Int!, $orderBy: String = \"newest_to_oldest\", $pageNo: Int = 1, $numPerPage: Int = 10) {\n  topicComments(topicId: $topicId, orderBy: $orderBy, pageNo: $pageNo, numPerPage: $numPerPage) {\n    data {\n      id\n      pinned\n      pinnedBy {\n        username\n        __typename\n      }\n      post {\n        ...DiscussPost\n        __typename\n      }\n      numChildren\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`, orderBy, offset, pageSize, topicId)
}

func (q *queryService) getGraphQLPayloadCommentReplies(commentId int64) string {
	return fmt.Sprintf(`{
		"operationName": "fetchCommentReplies",
		"variables": {
			"commentId": %v
		},
		"query": "query fetchCommentReplies($commentId: Int!) {\n  commentReplies(commentId: $commentId) {\n    id\n    pinned\n    pinnedBy {\n      username\n      __typename\n    }\n    post {\n      ...DiscussPost\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DiscussPost on PostNode {\n  id\n  voteCount\n  voteStatus\n  content\n  updationDate\n  creationDate\n  status\n  isHidden\n  coinRewards {\n    ...CoinReward\n    __typename\n  }\n  author {\n    isDiscussAdmin\n    isDiscussStaff\n    username\n    nameColor\n    activeBadge {\n      displayName\n      icon\n      __typename\n    }\n    profile {\n      userAvatar\n      reputation\n      __typename\n    }\n    isActive\n    __typename\n  }\n  authorIsModerator\n  isOwnPost\n  __typename\n}\n\nfragment CoinReward on ScoreNode {\n  id\n  score\n  description\n  date\n  __typename\n}\n"
	}`, commentId)
}

func (q *queryService) getGraphQLPayloadUserPublicProfile(username string) string {
	return fmt.Sprintf(`{
		"query": "\n    query userPublicProfile($username: String!) {\n  matchedUser(username: $username) {\n    contestBadge {\n      name\n      expired\n      hoverText\n      icon\n    }\n    username\n    githubUrl\n    twitterUrl\n    linkedinUrl\n    profile {\n      ranking\n      userAvatar\n      realName\n      aboutMe\n      school\n      websites\n      countryName\n      company\n      jobTitle\n      skillTags\n      postViewCount\n      postViewCountDiff\n      reputation\n      reputationDiff\n      solutionCount\n      solutionCountDiff\n      categoryDiscussCount\n      categoryDiscussCountDiff\n    }\n  }\n}\n    ",
		"variables": {
			"username": "%v"
		}
	}`, username)
}

func (q *queryService) getGraphQLPayloadUserSolveCountByTag(username string) string {
	return fmt.Sprintf(`{
		"query": "\n    query skillStats($username: String!) {\n  matchedUser(username: $username) {\n    tagProblemCounts {\n      advanced {\n        tagName\n        tagSlug\n        problemsSolved\n      }\n      intermediate {\n        tagName\n        tagSlug\n        problemsSolved\n      }\n      fundamental {\n        tagName\n        tagSlug\n        problemsSolved\n      }\n    }\n  }\n}\n    ",
		"variables": {
			"username": "%v"
		}
	}`, username)
}

func (q *queryService) getGraphQLPayloadUserContestRankingHistory(username string) string {
	return fmt.Sprintf(`{
		"query": "\n    query userContestRankingInfo($username: String!) {\n  userContestRanking(username: $username) {\n    attendedContestsCount\n    rating\n    globalRanking\n    totalParticipants\n    topPercentage\n    badge {\n      name\n    }\n  }\n  userContestRankingHistory(username: $username) {\n    attended\n    trendDirection\n    problemsSolved\n    totalProblems\n    finishTimeInSeconds\n    rating\n    ranking\n    contest {\n      title\n      startTime\n    }\n  }\n}\n    ",
		"variables": {
			"username": "%v"
		}
	}`, username)
}

func (q *queryService) getGraphQLPayloadUserSolveCountByDifficulty(username string) string {
	return fmt.Sprintf(`{
		"query": "\n    query userProblemsSolved($username: String!) {\n  allQuestionsCount {\n    difficulty\n    count\n  }\n  matchedUser(username: $username) {\n    problemsSolvedBeatsStats {\n      difficulty\n      percentage\n    }\n    submitStatsGlobal {\n      acSubmissionNum {\n        difficulty\n        count\n      }\n    }\n  }\n}\n    ",
		"variables": {
			"username": "%v"
		}
	}`, username)
}

func (q *queryService) getGraphQLPayloadUserProfileCalendar(username string) string {
	return fmt.Sprintf(`{
		"query": "\n    query userProfileCalendar($username: String!, $year: Int) {\n  matchedUser(username: $username) {\n    userCalendar(year: $year) {\n      activeYears\n      streak\n      totalActiveDays\n      dccBadges {\n        timestamp\n        badge {\n          name\n          icon\n        }\n      }\n      submissionCalendar\n    }\n  }\n}\n    ",
		"variables": {
			"username": "%v"
		}
	}`, username)
}

func (q *queryService) getGraphQLPayloadUserRecentAcSubmissions(username string, pageSize int) string {
	return fmt.Sprintf(`{
		"query": "\n    query recentAcSubmissions($username: String!, $limit: Int!) {\n  recentAcSubmissionList(username: $username, limit: $limit) {\n    id\n    title\n    titleSlug\n    timestamp\n  }\n}\n    ",
		"variables": {
			"username": "%v",
			"limit": %v
		}
	}`, username, pageSize)
}
