package leetcodeapi

import "log"

func GetUserPublicProfile(username string) (UserPublicProfile, error) {
	return getUserPublicProfile(username, &Util{}, &query{})
}

func GetUserSolveCountByProblemTag(username string) (TagProblemCounts, error) {
	return getUserSolveCountByProblemTag(username, &Util{}, &query{})
}

func GetUserContestRankingHistory(username string) (UserContestRankingDetails, error) {
	return getUserContestRankingHistory(username, &Util{}, &query{})
}

func GetUserSolveCountByDifficulty(username string) (UserSolveCountByDifficultyDetails, error) {
	return getUserSolveCountByDifficulty(username, &Util{}, &query{})
}

func GetUserProfileCalendar(username string) (UserCalendar, error) {
	return getUserProfileCalendar(username, &Util{}, &query{})
}

func GetUserRecentAcSubmissions(username string, pageSize int) ([]AcSubmission, error) {
	return getUserRecentAcSubmissions(username, pageSize, &Util{}, &query{})
}

/*
---------------------------------------------------------------------------------------
*/

func getUserPublicProfile(username string, utils IUtil, queries IQuery) (UserPublicProfile, error) {
	var result userPublicProfileReponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadUserPublicProfile(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserPublicProfile{}, err
	}

	return result.Data.MatchedUser, nil
}

func getUserSolveCountByProblemTag(username string, utils IUtil, queries IQuery) (TagProblemCounts, error) {
	var result userSolveCountByTagResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadUserSolveCountByTag(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return TagProblemCounts{}, err
	}

	return result.Data.MatchedUser.TagProblemCounts, nil
}

func getUserContestRankingHistory(username string, utils IUtil, queries IQuery) (UserContestRankingDetails, error) {
	var result userContestRankingHistoryResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadUserContestRankingHistory(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserContestRankingDetails{}, err
	}

	return result.Data, nil
}

func getUserSolveCountByDifficulty(username string, utils IUtil, queries IQuery) (UserSolveCountByDifficultyDetails, error) {
	var result userSolveCountByDifficultyResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadUserSolveCountByDifficulty(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserSolveCountByDifficultyDetails{}, err
	}

	return result.Data, nil
}

func getUserProfileCalendar(username string, utils IUtil, queries IQuery) (UserCalendar, error) {
	var result userProfileCalendarResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadUserProfileCalendar(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserCalendar{}, err
	}

	return result.Data.MatchedUser.UserCalendar, nil
}

func getUserRecentAcSubmissions(username string, pageSize int, utils IUtil, queries IQuery) ([]AcSubmission, error) {
	var result userRecentAcSubmissionsResponseBody
	err := utils.MakeGraphQLRequest(
		queries.getGraphQLPayloadUserRecentAcSubmissions(username, pageSize),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return []AcSubmission{}, err
	}

	return result.Data.RecentAcSubmissionList, nil
}
