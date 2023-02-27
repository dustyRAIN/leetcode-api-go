package leetcodeapi

import "log"

type DiscussionNode struct {
	CommentCount int     `json:"commentCount"`
	Id           string  `json:"id"`
	LastComment  Comment `json:"lastComment"`
	Pinned       bool    `json:"pinned"`
	Post         Post    `json:"post"`
	Tags         []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Title     string `json:"title"`
	ViewCount int    `json:"viewCount"`
}

type DiscussionList struct {
	Data []struct {
		Cursor string         `json:"cursor"`
		Node   DiscussionNode `json:"node"`
	} `json:"edges"`
	TotalNum int `json:"totalNum"`
}

type discussionListResponseBody struct {
	Data struct {
		CategoryTopicList DiscussionList `josn:"categoryTopicList"`
	} `json:"data"`
}

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

type Badge struct {
	DisplayName string `json:"displayName,omitempty"`
	Icon        string `json:"icon,omitempty"`
}

type AuthorProfile struct {
	Reputation int    `json:"reputation,omitempty"`
	UserAvatar string `json:"userAvatar,omitempty"`
}

type Author struct {
	ActiveBadge    Badge         `json:"activeBadge,omitempty"`
	IsActive       bool          `json:"isActive,omitempty"`
	IsDiscussAdmin bool          `json:"isDiscussAdmin,omitempty"`
	IsDiscussStaff bool          `json:"isDiscussStaff,omitempty"`
	NameColor      string        `json:"nameColor,omitempty"`
	Profile        AuthorProfile `json:"profile,omitempty"`
	Username       string        `json:"username"`
}

type CoinReward struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Id          string `json:"id"`
	Score       int    `json:"score"`
}

type Post struct {
	Id                int          `json:"id"`
	Author            Author       `json:"author,omitempty"`
	AuthorIsModerator bool         `json:"authorIsModerator,omitempty"`
	CoinRewards       []CoinReward `json:"coinRewards,omitempty"`
	Content           string       `json:"content,omitempty"`
	CreationDate      int64        `json:"creationDate,omitempty"`
	IsHidden          bool         `json:"isHidden,omitempty"`
	IsOwnPost         bool         `json:"isOwnPost,omitempty"`
	Status            string       `json:"status,omitempty"`
	UpdationDate      int64        `json:"updationDate,omitempty"`
	VoteCount         int          `json:"voteCount,omitempty"`
	VoteStatus        int          `json:"voteStatus,omitempty"`
}

type Discussion struct {
	HideFromTrending     bool     `json:"hideFromTrending"`
	Id                   int64    `json:"id"`
	Pinned               bool     `json:"pinned"`
	Post                 Post     `json:"post"`
	Subscribed           bool     `json:"subscribed"`
	Tags                 []string `json:"tags"`
	Title                string   `json:"title"`
	TopLevelCommentCount int      `json:"topLevelCommentCount"`
	ViewCount            int      `json:"viewCount"`
}

type discussionResponseBody struct {
	Data struct {
		Topic Discussion `josn:"topic"`
	} `json:"data"`
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

type Comment struct {
	Id          int64 `json:"id,omitempty"`
	NumChildren int   `json:"numChildren,omitempty"`
	Pinned      bool  `json:"pinned,omitempty"`
	Post        Post  `json:"post"`
}

type discussionCommentsResponseBody struct {
	Data struct {
		TopicComments struct {
			Data []Comment `json:"data"`
		} `josn:"topicComments"`
	} `json:"data"`
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

type commentRepliesResponseBody struct {
	Data struct {
		CommentReplies []Comment `josn:"commentReplies"`
	} `json:"data"`
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
