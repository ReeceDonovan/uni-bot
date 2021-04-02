package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ReeceDonovan/uni-bot/config"
	"github.com/spf13/viper"
)

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

func GetAssignments(sID string) (parsedData *CourseAssignment) {
	var token string

	sr := viper.Get("servers.active").([]config.ServerData)
	for _, s := range sr {
		if sID == s.ServerID {
			token = s.CanvasToken
		}
	}
	_, res := Req("POST", "/api/graphql", token, AssignmentQuery)
	jsonErr := json.Unmarshal(res, &parsedData)
	if jsonErr != nil {
		log.Println("Error parsing response: ", jsonErr)
	}
	return parsedData
}
