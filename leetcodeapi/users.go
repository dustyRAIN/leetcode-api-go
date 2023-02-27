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

type UserPublicProfileReponseBody struct {
	Data struct {
		MatchedUser struct {
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
		} `json:"matchedUser"`
	} `json:"data"`
}

func GetUserPublicProfile(username string) (UserPublicProfileReponseBody, error) {
	var result UserPublicProfileReponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserPublicProfile(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserPublicProfileReponseBody{}, err
	}

	return result, nil
}

type TagCount struct {
	ProblemsSolved int    `json:"problemsSolved"`
	TagName        string `json:"tagName"`
	TagSlug        string `json:"tagSlug"`
}

type UserSolveCountByTagResponseBody struct {
	Data struct {
		MatchedUser struct {
			TagProblemCounts struct {
				Advanced     []TagCount `json:"advanced"`
				Fundamental  []TagCount `json:"fundamental"`
				Intermediate []TagCount `json:"intermediate"`
			} `json:"tagProblemCounts"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func GetUserSolveCountByProblemTag(username string) (UserSolveCountByTagResponseBody, error) {
	var result UserSolveCountByTagResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserSolveCountByTag(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserSolveCountByTagResponseBody{}, err
	}

	return result, nil
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

type UserContestRankingHistoryResponseBody struct {
	Data struct {
		UserContestRanking struct {
			AttendedContestsCount int `json:"attendedContestsCount"`
			Badge                 struct {
				Name string `json:"name"`
			} `json:"badge"`
			GlobalRanking     int     `json:"globalRanking"`
			Rating            float32 `json:"rating"`
			TopPercentage     float32 `json:"topPercentage"`
			TotalParticipants int     `json:"totalParticipants"`
		} `json:"userContestRanking"`
		UserContestRankingHistory []UserContestRankingHistory `json:"userContestRankingHistory"`
	} `json:"data"`
}

func GetUserContestRankingHistory(username string) (UserContestRankingHistoryResponseBody, error) {
	var result UserContestRankingHistoryResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserContestRankingHistory(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserContestRankingHistoryResponseBody{}, err
	}

	return result, nil
}

type DifficultyCount struct {
	Count      int    `json:"count"`
	Difficulty string `json:"difficulty"`
}

type DifficultyPercentage struct {
	Percentage float32 `json:"percentage"`
	Difficulty string  `json:"difficulty"`
}

type UserSolveCountByDifficultyResponseBody struct {
	Data struct {
		AllQuestionsCount []DifficultyCount `json:"allQuestionsCount"`
		MatchedUser       struct {
			ProblemsSolvedBeatsStats []DifficultyPercentage `json:"problemsSolvedBeatsStats"`
			SubmitStatsGlobal        struct {
				AcSubmissionNum []DifficultyCount `json:"acSubmissionNum"`
			} `json:"submitStatsGlobal"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func GetUserSolveCountByDifficulty(username string) (UserSolveCountByDifficultyResponseBody, error) {
	var result UserSolveCountByDifficultyResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserSolveCountByDifficulty(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserSolveCountByDifficultyResponseBody{}, err
	}

	return result, nil
}

type UserProfileCalendarResponseBody struct {
	Data struct {
		MatchedUser struct {
			UserCalendar struct {
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
			} `json:"userCalendar"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func GetUserProfileCalendar(username string) (UserProfileCalendarResponseBody, error) {
	var result UserProfileCalendarResponseBody
	err := MakeGraphQLRequest(
		getGraphQLPayloadUserProfileCalendar(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserProfileCalendarResponseBody{}, err
	}

	return result, nil
}
