package leetcodeapi

import "log"

type DiscussionNode struct {
	CommentCount int    `json:"commentCount"`
	Id           string `json:"id"`
	LastComment  struct {
		Id   int  `json:"id"`
		Post Post `json:"post"`
	} `json:"lastComment"`
	Pinned bool `json:"pinned"`
	Post   Post `json:"post"`
	Tags   []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Title     string `json:"title"`
	ViewCount int    `json:"viewCount"`
}

type DiscussionListResponseBody struct {
	Data struct {
		CategoryTopicList struct {
			Edges []struct {
				Cursor string         `json:"cursor"`
				Node   DiscussionNode `json:"node"`
			} `json:"edges"`
			TotalNum int `json:"totalNum"`
		} `josn:"categoryTopicList"`
	} `json:"data"`
}

func GetDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int) (DiscussionListResponseBody, error) {
	var result DiscussionListResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadDiscussionList(categories, tags, orderBy, searchQuery, offset),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return DiscussionListResponseBody{}, err
	}

	return result, nil
}

type Author struct {
	ActiveBadge struct {
		DisplayName string `json:"displayName,omitempty"`
		Icon        string `json:"icon,omitempty"`
	} `json:"activeBadge,omitempty"`
	IsActive       bool   `json:"isActive,omitempty"`
	IsDiscussAdmin bool   `json:"isDiscussAdmin,omitempty"`
	IsDiscussStaff bool   `json:"isDiscussStaff,omitempty"`
	NameColor      string `json:"nameColor,omitempty"`
	Profile        struct {
		Reputation int    `json:"reputation,omitempty"`
		UserAvatar string `json:"userAvatar,omitempty"`
	} `json:"profile,omitempty"`
	Username string `json:"username"`
}

type Post struct {
	Id                int      `json:"id"`
	Author            Author   `json:"author,omitempty"`
	AuthorIsModerator bool     `json:"authorIsModerator,omitempty"`
	CoinRewards       []string `json:"coinRewards,omitempty"`
	Content           string   `json:"content,omitempty"`
	CreationDate      int64    `json:"creationDate,omitempty"`
	IsHidden          bool     `json:"isHidden,omitempty"`
	IsOwnPost         bool     `json:"isOwnPost,omitempty"`
	Status            string   `json:"status,omitempty"`
	UpdationDate      int64    `json:"updationDate,omitempty"`
	VoteCount         int      `json:"voteCount,omitempty"`
	VoteStatus        int      `json:"voteStatus,omitempty"`
}

type DiscussionResponseBody struct {
	Data struct {
		Topic struct {
			HideFromTrending     bool     `json:"hideFromTrending"`
			Id                   int64    `json:"id"`
			Pinned               bool     `json:"pinned"`
			Post                 Post     `json:"post"`
			Subscribed           bool     `json:"subscribed"`
			Tags                 []string `json:"tags"`
			Title                string   `json:"title"`
			TopLevelCommentCount int      `json:"topLevelCommentCount"`
			ViewCount            int      `json:"viewCount"`
		} `josn:"topic"`
	} `json:"data"`
}

func GetDiscussion(topicId int64) (DiscussionResponseBody, error) {
	var result DiscussionResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadDiscussion(topicId),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return DiscussionResponseBody{}, err
	}

	return result, nil
}
