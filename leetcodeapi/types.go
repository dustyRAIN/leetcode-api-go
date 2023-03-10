package leetcodeapi

type leetcodeMeta struct {
	session   string
	csrfToken string
}

type ContestMeta struct {
	Description     string `json:"description"`
	DiscussTopicId  int    `json:"discuss_topic_id"`
	Duration        int    `json:"duration"`
	Id              int    `json:"id"`
	IsPrivate       bool   `json:"is_private"`
	IsVirtual       bool   `json:"is_virtual"`
	OriginStartTime int64  `json:"origin_start_time"`
	StartTime       int64  `json:"start_time"`
	Title           string `json:"title"`
	TitleSlug       string `json:"title_slug"`
}

type ContestProblemInfo struct {
	Credit     int    `json:"credit"`
	Id         int    `json:"id"`
	QuestionId int    `json:"question_id"`
	Title      string `json:"title"`
	TitleSlug  string `json:"title_slug"`
}

type Contest struct {
	Company struct {
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Name        string `json:"name"`
	} `json:"company"`
	ContainsPremium bool                 `json:"containsPremium"`
	ContestMeta     ContestMeta          `json:"contest"`
	Questions       []ContestProblemInfo `json:"questions"`
	Registered      bool                 `json:"registered"`
}

type ParticipantDetails struct {
	ContestId     int    `json:"contest_id"`
	CountryCode   string `json:"country_code"`
	CountryName   string `json:"country_name"`
	DataRegion    string `json:"data_region"`
	FinishTime    int64  `json:"finish_time"`
	GlobalRanking int    `json:"global_ranking"`
	Rank          int    `json:"rank"`
	Score         int    `json:"score"`
	UserBadge     struct {
		DisplayName string `json:"display_name"`
		Icon        string `json:"icon"`
	} `json:"user_badge"`
	UserSlug      string `json:"user_slug"`
	Username      string `json:"username"`
	UsernameColor string `json:"username_color"`
}

type SubmissionInContest struct {
	ContestId    int    `json:"contest_id"`
	DataRegion   string `json:"data_region"`
	Date         int64  `json:"date"`
	FailCount    int    `json:"fail_count"`
	Id           int    `json:"id"`
	QuestionId   int    `json:"question_id"`
	Status       int    `json:"status"`
	SubmissionId int64  `json:"submission_id"`
}

type TopicTag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Problem struct {
	AcRate             float32    `json:"acRate"`
	Difficulty         string     `json:"difficulty"`
	FreqBar            float32    `json:"freqBar"`
	FrontendQuestionId string     `json:"frontendQuestionId"`
	IsFavor            bool       `json:"isFavor"`
	PaidOnly           bool       `json:"paidOnly"`
	Status             string     `json:"status"`
	QuestionId         string     `json:"questionId"`
	Title              string     `json:"title"`
	TitleSlug          string     `json:"titleSlug"`
	Stats              string     `json:"stats"`
	TopicTags          []TopicTag `json:"topicTags"`
}

type ProblemList struct {
	Total    int       `json:"total"`
	Problems []Problem `json:"questions"`
}

type problemsetListResponseBody struct {
	Data struct {
		ProblemsetQuestionList ProblemList `josn:"problemsetQuestionList"`
	} `json:"data"`
}

type ContestRanking struct {
	IsPast      bool                             `json:"is_past"`
	Questions   []ContestProblemInfo             `json:"questions"`
	Ranks       []ParticipantDetails             `json:"total_rank"`
	Submissions []map[string]SubmissionInContest `json:"submissions"`
	Time        float64                          `json:"time"`
	TotalUser   int                              `json:"user_num"`
	TotalPage   int
}

type ProblemContent struct {
	Content string `json:"content"`
}

type problemContentResponseBody struct {
	Data struct {
		Question ProblemContent `json:"question"`
	} `json:"data"`
}

type ProblemsByTopic struct {
	TopicName string    `json:"name"`
	TopicSlug string    `json:"slug"`
	Questions []Problem `json:"questions"`
}

type problemsByTopicResponseBody struct {
	Data struct {
		TopicTag ProblemsByTopic `josn:"topicTag"`
	} `json:"data"`
}

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

type DiscussionListItem struct {
	Cursor string         `json:"cursor"`
	Node   DiscussionNode `json:"node"`
}

type DiscussionList struct {
	Data     []DiscussionListItem `json:"edges"`
	TotalNum int                  `json:"totalNum"`
}

type discussionListResponseBody struct {
	Data struct {
		CategoryTopicList DiscussionList `josn:"categoryTopicList"`
	} `json:"data"`
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

type commentRepliesResponseBody struct {
	Data struct {
		CommentReplies []Comment `josn:"commentReplies"`
	} `json:"data"`
}

type UserProfile struct {
	AboutMe                  string   `json:"aboutMe"`
	CategoryDiscussCount     int      `json:"categoryDiscussCount"`
	CategoryDiscussCountDiff int      `json:"categoryDiscussCountDiff"`
	Company                  string   `json:"company"`
	CountryName              string   `json:"countryName"`
	JobTitle                 string   `json:"jobTitle"`
	PostViewCount            int      `json:"postViewCount"`
	PostViewCountDiff        int      `json:"postViewCountDiff"`
	Ranking                  int      `json:"ranking"`
	RealName                 string   `json:"realName"`
	Reputation               int      `json:"reputation"`
	ReputationDiff           int      `json:"reputationDiff"`
	School                   string   `json:"school"`
	SkillTags                []string `json:"skillTags"`
	SolutionCount            int      `json:"solutionCount"`
	SolutionCountDiff        int      `json:"solutionCountDiff"`
	UserAvatar               string   `json:"userAvatar"`
	Websites                 []string `json:"websites"`
}

type UserPublicProfile struct {
	ContestBadge struct {
		Expired   bool   `json:"expired"`
		HoverText string `json:"hoverText"`
		Icon      string `json:"icon"`
		Name      string `json:"name"`
	} `json:"contestBadge"`
	GithubUrl   string      `json:"githubUrl"`
	LinkedinUrl string      `json:"linkedinUrl"`
	Profile     UserProfile `json:"profile"`
	TwitterUrl  string      `json:"twitterUrl"`
	Username    string      `json:"username"`
}

type userPublicProfileReponseBody struct {
	Data struct {
		MatchedUser UserPublicProfile `json:"matchedUser"`
	} `json:"data"`
}

type TagCount struct {
	ProblemsSolved int    `json:"problemsSolved"`
	TagName        string `json:"tagName"`
	TagSlug        string `json:"tagSlug"`
}

type TagProblemCounts struct {
	Advanced     []TagCount `json:"advanced"`
	Fundamental  []TagCount `json:"fundamental"`
	Intermediate []TagCount `json:"intermediate"`
}

type userSolveCountByTagResponseBody struct {
	Data struct {
		MatchedUser struct {
			TagProblemCounts TagProblemCounts `json:"tagProblemCounts"`
		} `json:"matchedUser"`
	} `json:"data"`
}

type UserContestRanking struct {
	AttendedContestsCount int `json:"attendedContestsCount"`
	Badge                 struct {
		Name string `json:"name"`
	} `json:"badge"`
	GlobalRanking     int     `json:"globalRanking"`
	Rating            float32 `json:"rating"`
	TopPercentage     float32 `json:"topPercentage"`
	TotalParticipants int     `json:"totalParticipants"`
}

type UserContestRankingHistory struct {
	Attended bool `json:"attended"`
	Contest  struct {
		Title     string `json:"title"`
		StartTime int64  `json:"startTime"`
	} `json:"contest"`
	FinishTimeInSeconds int     `json:"finishTimeInSeconds"`
	ProblemsSolved      int     `json:"problemsSolved"`
	Ranking             int     `json:"ranking"`
	Rating              float32 `json:"rating"`
	TotalProblems       int     `json:"totalProblems"`
	TrendDirection      string  `json:"trendDirection"`
}

type UserContestRankingDetails struct {
	UserContestRanking        UserContestRanking          `json:"userContestRanking"`
	UserContestRankingHistory []UserContestRankingHistory `json:"userContestRankingHistory"`
}

type userContestRankingHistoryResponseBody struct {
	Data UserContestRankingDetails `json:"data"`
}

type DifficultyCount struct {
	Count      int    `json:"count"`
	Difficulty string `json:"difficulty"`
}

type DifficultyPercentage struct {
	Percentage float32 `json:"percentage"`
	Difficulty string  `json:"difficulty"`
}

type UserSolveCountByDifficulty struct {
	BeatsStats        []DifficultyPercentage `json:"problemsSolvedBeatsStats"`
	SubmitStatsGlobal struct {
		AcSubmissionNum []DifficultyCount `json:"acSubmissionNum"`
	} `json:"submitStatsGlobal"`
}

type UserSolveCountByDifficultyDetails struct {
	AllQuestionsCount []DifficultyCount          `json:"allQuestionsCount"`
	SolveCount        UserSolveCountByDifficulty `json:"matchedUser"`
}

type userSolveCountByDifficultyResponseBody struct {
	Data UserSolveCountByDifficultyDetails `json:"data"`
}

type AcSubmission struct {
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Title     string `json:"title"`
	TitleSlug string `json:"titleSlug"`
}

type userRecentAcSubmissionsResponseBody struct {
	Data struct {
		RecentAcSubmissionList []AcSubmission `json:"recentAcSubmissionList"`
	} `json:"data"`
}

type UserCalendar struct {
	ActiveYears []int `json:"activeYears"`
	DccBadges   []struct {
		Badge struct {
			Icon string `json:"icon"`
			Name string `json:"name"`
		} `json:"badge"`
		Timestamp int64 `json:"timestamp"`
	} `json:"dccBadges"`
	Streak             int    `json:"streak"`
	SubmissionCalendar string `json:"submissionCalendar"`
	TotalActiveDays    int    `json:"totalActiveDays"`
}

type userProfileCalendarResponseBody struct {
	Data struct {
		MatchedUser struct {
			UserCalendar UserCalendar `json:"userCalendar"`
		} `json:"matchedUser"`
	} `json:"data"`
}
