package leetcodeapi

import "log"

func GetDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int) (DiscussionList, error) {
	return getDiscussions(categories, tags, orderBy, searchQuery, offset, Util{}, query{})
}

func GetDiscussion(topicId int64) (Discussion, error) {
	return getDiscussion(topicId, Util{}, query{})
}

func GetDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) ([]Comment, error) {
	return getDiscussionComments(topicId, orderBy, offset, pageSize, Util{}, query{})
}

func GetCommentReplies(commentId int64) ([]Comment, error) {
	return getCommentReplies(commentId, Util{}, query{})
}

/*
---------------------------------------------------------------------------------------
*/

func getDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int, utils IUtil, queries IQuery) (DiscussionList, error) {
	var result discussionListResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadDiscussionList(categories, tags, orderBy, searchQuery, offset, utils),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return DiscussionList{}, err
	}

	return result.Data.CategoryTopicList, nil
}

func getDiscussion(topicId int64, utils IUtil, queries IQuery) (Discussion, error) {
	var result discussionResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadDiscussion(topicId),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return Discussion{}, err
	}

	return result.Data.Topic, nil
}

func getDiscussionComments(topicId int64, orderBy string, offset int, pageSize int, utils IUtil, queries IQuery) ([]Comment, error) {
	var result discussionCommentsResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadDiscussionComments(topicId, orderBy, offset, pageSize),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return []Comment{}, err
	}

	return result.Data.TopicComments.Data, nil
}

func getCommentReplies(commentId int64, utils IUtil, queries IQuery) ([]Comment, error) {
	var result commentRepliesResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadCommentReplies(commentId),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return []Comment{}, err
	}

	return result.Data.CommentReplies, nil
}
