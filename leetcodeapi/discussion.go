package leetcodeapi

import "log"

type Post struct {
	Id     int `json:"id"`
	Author struct {
		Username string `json:"username"`
	} `json:"author"`
	CreationDate int64 `json:"creationDate"`
}

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
