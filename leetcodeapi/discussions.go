package leetcodeapi

import "log"

func GetDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int) (DiscussionList, error) {
	var result discussionListResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadDiscussionList(categories, tags, orderBy, searchQuery, offset),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return DiscussionList{}, err
	}

	return result.Data.CategoryTopicList, nil
}

func GetDiscussion(topicId int64) (Discussion, error) {
	var result discussionResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadDiscussion(topicId),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return Discussion{}, err
	}

	return result.Data.Topic, nil
}

func GetDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) ([]Comment, error) {
	var result discussionCommentsResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadDiscussionComments(topicId, orderBy, offset, pageSize),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return []Comment{}, err
	}

	return result.Data.TopicComments.Data, nil
}

func GetCommentReplies(commentId int64) ([]Comment, error) {
	var result commentRepliesResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadCommentReplies(commentId),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return []Comment{}, err
	}

	return result.Data.CommentReplies, nil
}
