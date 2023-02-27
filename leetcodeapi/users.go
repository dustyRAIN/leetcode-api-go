package leetcodeapi

import "log"

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
