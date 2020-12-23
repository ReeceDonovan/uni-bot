package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// Req constructs a request and returns status code and response body
func Req(method, slug, token string, body []byte) (int, []byte) {

	r, err := http.NewRequest(method,
		fmt.Sprintf("%s/%s", viper.GetString("canvas.domain"), slug),
		bytes.NewReader(body),
	)
	if err != nil {
		log.Println("Error making http request <", method, ">", slug, "\n", err)
		return 0, []byte{}
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if method == "POST" {
		r.Header.Add("Content-Type", "application/json")
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Error sending http request <", method, ">", slug, "\n", err)
		return 0, []byte{}
	}
	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Couldn't read response \n", err)
		return 0, []byte{}
	}
	return resp.StatusCode, bd
}

// TODO: Might move everything from here down to different file for tidyness if we need other query commands, but ah sur. Also gotta change default struct variable names from the autogen

func QueryAssignments() (parsedAss Assignment) {

	_, res := Req("POST", "/api/graphql",
		fmt.Sprintf("%s", viper.GetString("canvas.token")),
		[]byte(`
			{"query": "query Assignments {
				allCourses {
					name
					assignmentsConnection {
						edges {
						node {
							name
							dueAt
							htmlUrl
							}
						}
					}
				}
			}"}`))

	jsonErr := json.Unmarshal(res, &parsedAss)
	if jsonErr != nil {
		log.Println("Error parsing response\n", jsonErr)
	}
	return parsedAss
}
