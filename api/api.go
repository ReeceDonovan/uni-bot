package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
)

type returnAssignments struct {
}

var (
	qURL    string
	cachedC *cache.Cache
	cachedA *cache.Cache
	session *discordgo.Session
)

func Run(s *discordgo.Session) {
	cachedC = cache.New(6*time.Hour, 6*time.Hour)
	cachedA = cache.New(6*time.Hour, 6*time.Hour)

	session = s
}

func QueryCourse() []ParsedCourse {
	var cdata CourseData

	parsedC := []ParsedCourse{}

	cachedCourses, found := cachedC.Get("courses")

	if found {
		parsedC = cachedCourses.([]ParsedCourse)
		fmt.Println("Cache found")
	} else {
		fmt.Println("Cache created")
		qURL = viper.GetString("canvas.cURL") + viper.GetString("canvas.token")
		res, err := http.Get(qURL)
		if err != nil {
			log.Fatal(err)
		}
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		jsonErr := json.Unmarshal([]byte(body), &cdata)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		for _, c := range cdata {
			if c.CreatedAt.Unix() > time.Now().AddDate(-1, 0, 0).Unix() {
				parsedC = append(parsedC, ParsedCourse{
					c.ID,
					c.Name,
					c.CourseCode,
				})
			}
		}
		cachedC.Set("courses", parsedC, cache.DefaultExpiration)
	}
	return parsedC
}

func QueryAssign(c string) []ParsedAssignment {
	var adata AssignmentData

	parsedA := []ParsedAssignment{}

	cachedAssignments, found := cachedA.Get("assignments")

	if found {
		parsedA = cachedAssignments.([]ParsedAssignment)
		fmt.Println("Cache found")
	} else {
		qURL = viper.GetString("canvas.aURLs") + c + viper.GetString("canvas.aURLe") + viper.GetString("canvas.token")
		res, err := http.Get(qURL)
		if err != nil {
			log.Fatal(err)
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		jsonErr := json.Unmarshal([]byte(body), &adata)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for _, a := range adata {
			if a.DueAt.Unix() > time.Now().AddDate(0, 0, 0).Unix() {
				parsedA = append(parsedA, ParsedAssignment{
					a.ID,
					a.Name,
					a.Description,
					a.HTMLURL,
					a.DueAt,
				})
			}
		}
	}
	return parsedA
}
