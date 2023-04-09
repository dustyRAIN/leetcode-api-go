# leetcode-api-go
Some of the APIs from [Leetcode](https://leetcode.com), written in our favorite language. Let's Goooo.... (sorry if you're already tired of it, I am)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/dustyRAIN/leetcode-api-go.svg)](https://github.com/gomods/athens) ![Go Report Card](https://goreportcard.com/badge/github.com/dustyRAIN/leetcode-api-go)   

[![forthebadge](https://forthebadge.com/images/badges/works-on-my-machine.svg)](https://forthebadge.com)  


## Get Started
All the APIs written here can be accessed without any user context, hence there is no need for authentication. 
Although, user context will provide additional data such as the user's solve count, submission status, premium data (if applicable to the user), etc. 
In order to inject or remove the user context, the following methods can be called.

```go
package your_package

import (
  "github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

var session string = "your_leetcode_session_from_cookie"
var csrfToken string = "your_leetcode_csrfToken_from_cookie"

func main() {
  //  to inject user context
  leetcodeapi.SetCredentials(session, csrfToken)

  //  any api call will now use the user context

  //  to remove the context
  leetcodeapi.RemoveCredentials()
}
```

Available APIs can be called in the following way,

```go
package your_package

import (
  "github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func DoSomething() {
  var allProblemList leetcodeapi.ProblemList
  allProblemList, err := leetcodeapi.GetAllProblems(0, 50)
}
```

## Available APIs Related to Leetcode Contests Page

```go
  leetcodeapi.GetContestInfo(contestSlug string) (leetcodeapi.Contest, error)
  leetcodeapi.GetContestRanking(contestSlug string, page int) (leetcodeapi.ContestRanking, error)
```

## Available APIs Related to Leetcode Discussions Page

```go
  leetcodeapi.GetDiscussions(categories []string, tags []string, orderBy string, searchQuery string, offset int) (leetcodeapi.DiscussionList, error)
  leetcodeapi.GetDiscussion(topicId int64) (leetcodeapi.Discussion, error)
  leetcodeapi.GetDiscussionComments(topicId int64, orderBy string, offset int, pageSize int) ([]leetcodeapi.Comment, error)
  leetcodeapi.GetCommentReplies(commentId int64) ([]leetcodeapi.Comment, error)
```

## Available APIs Related to Leetcode Problems Page

```go
  leetcodeapi.GetAllProblems(offset int, pageSize int) (leetcodeapi.ProblemList, error)
  leetcodeapi.GetProblemContentByTitleSlug(titleSlug string) (leetcodeapi.ProblemContent, error)
  leetcodeapi.GetProblemsByTopic(topicSlug string) (leetcodeapi.ProblemsByTopic, error)
  leetcodeapi.GetTopInterviewProblems(offset int, pageSize int) (leetcodeapi.ProblemList, error)
```

## Available APIs Related to Leetcode Users Page

```go
  leetcodeapi.GetUserPublicProfile(username string) (leetcodeapi.UserPublicProfile, error)
  leetcodeapi.GetUserSolveCountByProblemTag(username string) (leetcodeapi.TagProblemCounts, error)
  leetcodeapi.GetUserContestRankingHistory(username string) (leetcodeapi.UserContestRankingDetails, error)
  leetcodeapi.GetUserSolveCountByDifficulty(username string) (leetcodeapi.UserSolveCountByDifficultyDetails, error)
  leetcodeapi.GetUserProfileCalendar(username string) (leetcodeapi.UserCalendar, error)
  leetcodeapi.GetUserRecentAcSubmissions(username string, pageSize int) ([]leetcodeapi.AcSubmission, error)
```

## Make Your Own Call to Leetcode

Indeed the available APIs may not be sufficient, hence a GraphQL request can be made directly. The following method will only allow us to make GET calls
to prevent the misusage of the API. The method takes two parameters, GraphQL request payload as a `string` and the reference object as an `interface{}`
where the response body will be translated.

```go
package your_package

import (
  "github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

type Something struct {
  Data struct {
    Name      string  `json:"name"`
    IsBlocked bool    `json:"is_blocked"`
  } `json:"data"`
}

func DidYouDoDat() {
  var responseBody Something
  payload := `{
    "query": "\n query Something() {\n name \nis_blocked \n} \n",
    "variables": {}
  }`
  err := (&leetcodeapi.Util{}).MakeGraphQLRequest(payload, &responseBody)
}
```
