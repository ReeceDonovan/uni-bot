package api

import (
	"encoding/json"
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

func QueryCourse() courseData {
	var cdata courseData
	// cachedCourses, found := cachedC.Get("courses")
	// if found {
	// 	cdata = cachedCourses.(courseData)
	// } else {
	// assignments := []*Assignment{}
	qURL = viper.GetString("canvas.cURL") + viper.GetString("canvas.token")
	// qURL = "https://ucc.instructure.com/api/v1/users/self/courses?enrollment_state=active&state[]=available&include[]=term&exclude[]=enrollments&sort=nickname&access_token=" + "13518~I85HRvFXEgBQSlWzClrmG4VnKaEZwAqgwjYBvgXbkE2fAMy32PiRZk1sQhcUmwqU"

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
	// fmt.Println(cdata)
	// cachedC.Set("courses", cdata, cache.DefaultExpiration)
	// }
	return cdata
}

func QueryAssign(c string) []parsedAssignment {
	var adata assignmentData
	parsedData := []parsedAssignment{}
	// time := time.Now()
	// fmt.Println(c)
	// cachedAssignments, found := cachedA.Get(c)
	// if found {
	// 	adata = cachedAssignments.(assignmentData)
	// } else {
	// 	var adata assignmentData
	qURL = ("https://ucc.instructure.com/api/v1/users/self/courses/" + c + "/assignments?&order_by=due_at&access_token=" + viper.GetString("canvas.token"))
	res, err := http.Get(qURL)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(res)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// fmt.Println(body)

	jsonErr := json.Unmarshal([]byte(body), &adata)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// fmt.Println(adata)
	// cachedA.Set(c, adata, cache.DefaultExpiration)
	// }
	for _, a := range adata {
		if a.DueAt.Unix() > time.Now().AddDate(0, 0, 0).Unix() {
			parsedData = append(parsedData, parsedAssignment{
				a.ID,
				a.Name,
				a.Description,
				a.HTMLURL,
				a.DueAt,
			})
		}
		// fmt.Println(parsedData)
	}
	return parsedData
}
