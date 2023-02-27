package leetcodeapi

import "log"

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

func GetUserPublicProfile(username string) (UserPublicProfile, error) {
	var result userPublicProfileReponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserPublicProfile(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserPublicProfile{}, err
	}

	return result.Data.MatchedUser, nil
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

func GetUserSolveCountByProblemTag(username string) (TagProblemCounts, error) {
	var result userSolveCountByTagResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserSolveCountByTag(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return TagProblemCounts{}, err
	}

	return result.Data.MatchedUser.TagProblemCounts, nil
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

type UserContestRankingDetails struct {
	UserContestRanking        UserContestRanking          `json:"userContestRanking"`
	UserContestRankingHistory []UserContestRankingHistory `json:"userContestRankingHistory"`
}

type userContestRankingHistoryResponseBody struct {
	Data UserContestRankingDetails `json:"data"`
}

func GetUserContestRankingHistory(username string) (UserContestRankingDetails, error) {
	var result userContestRankingHistoryResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserContestRankingHistory(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserContestRankingDetails{}, err
	}

	return result.Data, nil
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
	ProblemsSolvedBeatsStats []DifficultyPercentage `json:"problemsSolvedBeatsStats"`
	SubmitStatsGlobal        struct {
		AcSubmissionNum []DifficultyCount `json:"acSubmissionNum"`
	} `json:"submitStatsGlobal"`
}

type UserSolveCountByDifficultyDetails struct {
	AllQuestionsCount     []DifficultyCount          `json:"allQuestionsCount"`
	MatchedUserSolveCount UserSolveCountByDifficulty `json:"matchedUser"`
}

type userSolveCountByDifficultyResponseBody struct {
	Data UserSolveCountByDifficultyDetails `json:"data"`
}

func GetUserSolveCountByDifficulty(username string) (UserSolveCountByDifficultyDetails, error) {
	var result userSolveCountByDifficultyResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserSolveCountByDifficulty(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserSolveCountByDifficultyDetails{}, err
	}

	return result.Data, nil
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

func GetUserProfileCalendar(username string) (UserCalendar, error) {
	var result userProfileCalendarResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserProfileCalendar(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserCalendar{}, err
	}

	return result.Data.MatchedUser.UserCalendar, nil
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

func GetUserRecentAcSubmissions(username string, pageSize int) ([]AcSubmission, error) {
	var result userRecentAcSubmissionsResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserRecentAcSubmissions(username, pageSize),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return []AcSubmission{}, err
	}

	return result.Data.RecentAcSubmissionList, nil
}
