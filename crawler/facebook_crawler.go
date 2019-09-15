package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// Comment facebook comment structure
type Comment struct {
	ID       string
	Name     string
	AuthorID string
	TargetID string
	Type     string
}

// FacebookComments json comment structure
type FacebookComments struct {
	CommentIds []string           `json:"commentIDs"`
	IDMap      map[string]Comment `json:"idMap"`
}

func fetch(url string) (string, error) {
	result, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return getBody(result)
}

func getBody(response *http.Response) (string, error) {
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func jsonComments(response string) string {
	r, _ := regexp.Compile(`{"comments":{(.*)},"meta"`)
	res := r.FindStringSubmatch(response)

	return fmt.Sprintf("{%s}", res[1])
}

func parseJSONComments(jsonString string, dest *FacebookComments) error {
	comments := jsonComments(jsonString)
	return json.Unmarshal([]byte(comments), dest)
}

func hasPendingComments(comments FacebookComments, officialUser string) bool {
	if len(comments.CommentIds) == 0 {
		return false
	}

	for _, id := range comments.CommentIds {
		answers := []string{}

		for answersID, comment := range comments.IDMap {
			if comment.Type != "comment" {
				continue
			}

			if comment.TargetID == id && comment.AuthorID == officialUser {
				answers = append(answers, answersID)
			}
		}

		if len(answers) == 0 {
			return true
		}
	}

	return false
}

// VerifyFacebookPage Fetch for url and check for pending comments
func VerifyFacebookPage(url string, officialUser string) (bool, error) {
	body, err := fetch(url)

	if err != nil {
		return false, err
	}

	var obj FacebookComments
	err = parseJSONComments(body, &obj)

	if err != nil {
		return false, err
	}

	return hasPendingComments(obj, officialUser), nil
}
