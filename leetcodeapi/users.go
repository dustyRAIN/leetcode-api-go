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
		getGraphQLUserPublicProfile(username),
		&result,
	)

	if err != nil {
		log.Printf(err.Error())
		return UserPublicProfileReponseBody{}, err
	}

	return result, nil
}
