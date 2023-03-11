package leetcodeapi

import (
	"fmt"
	"log"
	"math"
)

func GetContestInfo(contestSlug string) (Contest, error) {
	return (&contestService{utils: &Util{}}).getContestInfo(contestSlug)
}

func GetContestRanking(contestSlug string, page int) (ContestRanking, error) {
	return (&contestService{utils: &Util{}}).getContestRanking(contestSlug, page)
}

/*
---------------------------------------------------------------------------------------
*/

type contestService struct {
	utils IUtil
}

func (c *contestService) getContestInfo(contestSlug string) (Contest, error) {
	var result Contest
	err := c.utils.makeHttpRequest(
		"GET",
		fmt.Sprintf("https://leetcode.com/contest/api/info/%v/", contestSlug),
		"",
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return Contest{}, err
	}

	return result, nil
}

func (c *contestService) getContestRanking(contestSlug string, page int) (ContestRanking, error) {
	var result ContestRanking
	err := c.utils.makeHttpRequest(
		"GET",
		fmt.Sprintf("https://leetcode.com/contest/api/ranking/%v/?pagination=%v&region=global", contestSlug, page),
		"",
		&result,
	)

	if err != nil {
		log.Print(err.Error())
		return ContestRanking{}, err
	}

	result.TotalPage = int(math.Ceil(float64(result.TotalUser) / 25.0))

	return result, nil
}
