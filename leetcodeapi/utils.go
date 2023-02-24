package leetcodeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func parseCookie(cookies []*http.Cookie, cookieName string) string {
	for _, cookie := range cookies {
		if cookie.Name == cookieName {
			return cookie.Value
		}
	}
	return ""
}

func makeHttpRequest(method string, url string, contentType string, body string, resultRef interface{}) error {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if len(leetcodeMeta.session) > 0 && len(leetcodeMeta.csrfToken) > 0 {
		request.Header.Add("Cookie", fmt.Sprintf("LEETCODE_SESSION=%v; csrftoken=%v", leetcodeMeta.session, leetcodeMeta.csrfToken))
	}
	if len(leetcodeMeta.csrfToken) > 0 {
		request.Header.Add("X-csrf-token", leetcodeMeta.csrfToken)
	}

	response, err := client.Do(request)

	if err != nil {
		return err
	}
	responseBodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return errors.New(fmt.Sprintf("statusCode: %v\nmessage: %v", response.StatusCode, string(responseBodyBytes)))
	}

	err = json.Unmarshal(responseBodyBytes, &resultRef)

	return err
}
