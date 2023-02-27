package leetcodeapi

import (
	"fmt"
	"log"
)

func GetContestInfo(contestSlug string) (Contest, error) {
	var result Contest
	err := makeHttpRequest(
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

func GetContestRanking(contestSlug string, page int) (ContestRanking, error) {
	var result ContestRanking
	err := makeHttpRequest(
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
