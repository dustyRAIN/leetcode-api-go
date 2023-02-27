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
	//fmt.Println(leetcodeapi.GetDiscussions([]string{"interview-experience"}, []string{}, "", "google", 0))
	//fmt.Println(leetcodeapi.GetDiscussion(1674246))
	//fmt.Println(leetcodeapi.GetDiscussionComments(2069641, "best", 0, 10))
	//fmt.Println(leetcodeapi.GetCommentReplies(1404906))
	//fmt.Println(leetcodeapi.GetUserPublicProfile("dustyRAIN"))
	//fmt.Println(leetcodeapi.GetUserSolveCountByProblemTag("dustyRAIN"))
	fmt.Println(leetcodeapi.GetUserContestRankingHistory("dustyRAIN"))
}
