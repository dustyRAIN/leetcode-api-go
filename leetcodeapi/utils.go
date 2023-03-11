package leetcodeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type IUtil interface {
	parseCookie(cookies []*http.Cookie, cookieName string) string
	makeHttpRequest(method string, url string, body string, resultRef interface{}) error
	MakeGraphQLRequest(payload string, resultRef interface{}) error
	convertListToString(list []string) string
}

type Util struct{}

func (u *Util) parseCookie(cookies []*http.Cookie, cookieName string) string {
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			return cookie.Value
		}
	}
	return ""
}

func (u *Util) makeHttpRequest(method string, url string, body string, resultRef interface{}) error {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if len(credentials.session) > 0 && len(credentials.csrfToken) > 0 {
		request.Header.Add("Cookie", fmt.Sprintf("LEETCODE_SESSION=%v; csrftoken=%v", credentials.session, credentials.csrfToken))
	}
	if len(credentials.csrfToken) > 0 {
		request.Header.Add("X-csrf-token", credentials.csrfToken)
	}

	response, err := client.Do(request)

	if err != nil {
		return err
	}
	responseBodyBytes, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("statusCode: %v\nmessage: %v", response.StatusCode, string(responseBodyBytes))
	}

	err = json.Unmarshal(responseBodyBytes, &resultRef)

	return err
}

func (u *Util) MakeGraphQLRequest(payload string, resultRef interface{}) error {
	err := u.makeHttpRequest(
		"GET",
		"https://leetcode.com/graphql/",
		payload,
		&resultRef,
	)

	return err
}

func (u *Util) convertListToString(list []string) string {
	var listString string = "["
	for indx, item := range list {
		if indx > 0 {
			listString += ","
		}
		listString += fmt.Sprintf("\"%v\"", item)
	}
	listString += "]"
	return listString
}
