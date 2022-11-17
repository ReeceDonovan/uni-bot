package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ReeceDonovan/uni-bot/models"
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

func GetCourses(token string) *models.Courses {
	_, res := Req("GET", "/api/v1/users/self/courses?enrollment_type=student&include[]=total_scores&include[]=current_grading_period_scores&include[]=total_students&include[]=teachers&include[]=term&per_page=1000", token, nil)
	tempCourses := &models.Courses{}
	jsonErr := json.Unmarshal(res, &tempCourses)
	if jsonErr != nil {
		log.Println("GetCourses - Error parsing response: ", jsonErr)
		log.Println(string(res))
	}
	courses := models.Courses{}
	for _, course := range *tempCourses {
		if course.AccountID == 1 {
			courses = append(courses, course)
		}
	}
	return &courses
}

func GetCourse(moduleID string, token string) (course *models.Course) {
	_, res := Req("GET", "/api/v1/courses/"+moduleID+"?include[]=total_scores&include[]=current_grading_period_scores&include[]=total_students&include[]=teachers", token, nil)

	jsonErr := json.Unmarshal(res, &course)
	if jsonErr != nil {
		log.Println("GetCourse - Error parsing response: ", jsonErr)
		log.Println(string(res))
	}
	return course
}

func GetAssignments(moduleID string, token string) (assignments *models.Assignments) {
	_, res := Req("GET", fmt.Sprintf("/api/v1/users/self/courses/%s/assignments?include[]=all_dates&include[]=submission&include[]=score_statistics&order_by=due_at&per_page=1000", moduleID), token, nil)

	jsonErr := json.Unmarshal(res, &assignments)
	if jsonErr != nil {
		log.Println("GetAssignments - Error parsing response: ", jsonErr)
		log.Println(string(res))
	}
	return assignments
}
