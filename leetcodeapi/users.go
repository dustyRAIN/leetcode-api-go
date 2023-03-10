package leetcodeapi

import "log"

func GetUserPublicProfile(username string) (UserPublicProfile, error) {
	utils := &Util{}
	return (&usersService{utils: utils, queries: &queryService{utils: utils}}).getUserPublicProfile(username)
}

func GetUserSolveCountByProblemTag(username string) (TagProblemCounts, error) {
	utils := &Util{}
	return (&usersService{utils: utils, queries: &queryService{utils: utils}}).getUserSolveCountByProblemTag(username)
}

func GetUserContestRankingHistory(username string) (UserContestRankingDetails, error) {
	utils := &Util{}
	return (&usersService{utils: utils, queries: &queryService{utils: utils}}).getUserContestRankingHistory(username)
}

func GetUserSolveCountByDifficulty(username string) (UserSolveCountByDifficultyDetails, error) {
	utils := &Util{}
	return (&usersService{utils: utils, queries: &queryService{utils: utils}}).getUserSolveCountByDifficulty(username)
}

func GetUserProfileCalendar(username string) (UserCalendar, error) {
	utils := &Util{}
	return (&usersService{utils: utils, queries: &queryService{utils: utils}}).getUserProfileCalendar(username)
}

func GetUserRecentAcSubmissions(username string, pageSize int) ([]AcSubmission, error) {
	utils := &Util{}
	return (&usersService{utils: utils, queries: &queryService{utils: utils}}).getUserRecentAcSubmissions(username, pageSize)
}

/*
---------------------------------------------------------------------------------------
*/

type usersService struct {
	utils   IUtil
	queries IQuery
}

func (u *usersService) getUserPublicProfile(username string) (UserPublicProfile, error) {
	var result userPublicProfileReponseBody
	err := u.utils.MakeGraphQLRequest(
		u.queries.getGraphQLPayloadUserPublicProfile(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserPublicProfile{}, err
	}

	return result.Data.MatchedUser, nil
}

func (u *usersService) getUserSolveCountByProblemTag(username string) (TagProblemCounts, error) {
	var result userSolveCountByTagResponseBody
	err := u.utils.MakeGraphQLRequest(
		u.queries.getGraphQLPayloadUserSolveCountByTag(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return TagProblemCounts{}, err
	}

	return result.Data.MatchedUser.TagProblemCounts, nil
}

func (u *usersService) getUserContestRankingHistory(username string) (UserContestRankingDetails, error) {
	var result userContestRankingHistoryResponseBody
	err := u.utils.MakeGraphQLRequest(
		u.queries.getGraphQLPayloadUserContestRankingHistory(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserContestRankingDetails{}, err
	}

	return result.Data, nil
}

func (u *usersService) getUserSolveCountByDifficulty(username string) (UserSolveCountByDifficultyDetails, error) {
	var result userSolveCountByDifficultyResponseBody
	err := u.utils.MakeGraphQLRequest(
		u.queries.getGraphQLPayloadUserSolveCountByDifficulty(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserSolveCountByDifficultyDetails{}, err
	}

	return result.Data, nil
}

func (u *usersService) getUserProfileCalendar(username string) (UserCalendar, error) {
	var result userProfileCalendarResponseBody
	err := u.utils.MakeGraphQLRequest(
		u.queries.getGraphQLPayloadUserProfileCalendar(username),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return UserCalendar{}, err
	}

	return result.Data.MatchedUser.UserCalendar, nil
}

func (u *usersService) getUserRecentAcSubmissions(username string, pageSize int) ([]AcSubmission, error) {
	var result userRecentAcSubmissionsResponseBody
	err := u.utils.MakeGraphQLRequest(
		u.queries.getGraphQLPayloadUserRecentAcSubmissions(username, pageSize),
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return []AcSubmission{}, err
	}

	return result.Data.RecentAcSubmissionList, nil
}
