package leetcodeapi

import (
	"fmt"
	"log"
)

func GetContestInfo(contestSlug string) (Contest, error) {
	return getContestInfo(contestSlug, Util{})
}

func GetContestRanking(contestSlug string, page int) (ContestRanking, error) {
	return getContestRanking(contestSlug, page, Util{})
}

/*
---------------------------------------------------------------------------------------
*/

func getContestInfo(contestSlug string, utils IUtil) (Contest, error) {
	var result Contest
	err := utils.makeHttpRequest(
		"GET",
		fmt.Sprintf("https://leetcode.com/contest/api/info/%v/", contestSlug),
		"application/json",
		"",
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return Contest{}, err
	}

	return result, nil
}

func getContestRanking(contestSlug string, page int, utils IUtil) (ContestRanking, error) {
	var result ContestRanking
	err := utils.makeHttpRequest(
		"GET",
		fmt.Sprintf("https://leetcode.com/contest/api/ranking/%v/?pagination=%v&region=global", contestSlug, page),
		"application/json",
		"",
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return ContestRanking{}, err
	}

	return result, nil
}
