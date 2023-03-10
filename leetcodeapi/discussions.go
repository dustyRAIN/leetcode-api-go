package leetcodeapi

import "log"

func GetDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int) (DiscussionList, error) {
	utils := &Util{}
	return (&discussionService{utils: utils, queries: &queryService{utils: utils}}).getDiscussions(categories, tags, orderBy, searchQuery, offset)
}

func GetDiscussion(topicId int64) (Discussion, error) {
	utils := &Util{}
	return (&discussionService{utils: utils, queries: &queryService{utils: utils}}).getDiscussion(topicId)
}

func GetDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) ([]Comment, error) {
	utils := &Util{}
	return (&discussionService{utils: utils, queries: &queryService{utils: utils}}).getDiscussionComments(topicId, orderBy, offset, pageSize)
}

func GetCommentReplies(commentId int64) ([]Comment, error) {
	utils := &Util{}
	return (&discussionService{utils: utils, queries: &queryService{utils: utils}}).getCommentReplies(commentId)
}

/*
---------------------------------------------------------------------------------------
*/

type discussionService struct {
	utils   IUtil
	queries IQuery
}

func (d *discussionService) getDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int) (DiscussionList, error) {
	var result discussionListResponseBody
	err := d.utils.MakeGraphQLRequest(
		d.queries.getGraphQLPayloadDiscussionList(categories, tags, orderBy, searchQuery, offset),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return DiscussionList{}, err
	}

	return result.Data.CategoryTopicList, nil
}

func (d *discussionService) getDiscussion(topicId int64) (Discussion, error) {
	var result discussionResponseBody
	err := d.utils.MakeGraphQLRequest(
		d.queries.getGraphQLPayloadDiscussion(topicId),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return Discussion{}, err
	}

	return result.Data.Topic, nil
}

func (d *discussionService) getDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) ([]Comment, error) {
	var result discussionCommentsResponseBody
	err := d.utils.MakeGraphQLRequest(
		d.queries.getGraphQLPayloadDiscussionComments(topicId, orderBy, offset, pageSize),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return []Comment{}, err
	}

	return result.Data.TopicComments.Data, nil
}

func (d *discussionService) getCommentReplies(commentId int64) ([]Comment, error) {
	var result commentRepliesResponseBody
	err := d.utils.MakeGraphQLRequest(
		d.queries.getGraphQLPayloadCommentReplies(commentId),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return []Comment{}, err
	}

	return result.Data.CommentReplies, nil
}
