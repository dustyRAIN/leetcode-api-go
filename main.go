package main

import (
	"dustyRAIN/leetcode-api-go/leetcodeapi"
	"fmt"
)

func main() {
	//leetcodeapi.SetCredentials(session, token)
	//fmt.Println(leetcodeapi.GetAllProblems())
	//fmt.Println(leetcodeapi.GetProblemContentByTitleSlug("find-first-and-last-position-of-element-in-sorted-array"))
	//fmt.Println(leetcodeapi.GetProblemsByTopic("hash-table"))
	//fmt.Println(leetcodeapi.GetTopInterviewProblems())
	//fmt.Println(leetcodeapi.GetContestInfo("weekly-contest-333"))
	//fmt.Println(leetcodeapi.GetContestRanking("weekly-contest-333", 1))
	fmt.Println(leetcodeapi.GetDiscussions([]string{"interview-experience"}, []string{}, "", "", 0))
}
