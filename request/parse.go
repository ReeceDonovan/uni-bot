package request

import "time"

// TODO: Clean up struct, think of any other useful data we could use from the api

type CourseAssignment struct {
	Data struct {
		AllCourses []struct {
			ID         string `json:"_id"`
			CourseName string `json:"name"`
			State      string `json:"state"`
			CourseCode string `json:"courseCode"`
			Term       struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"term"`
			AssignmentsConnection struct {
				Nodes []struct {
					ID              string          `json:"_id"`
					Name            string          `json:"name"`
					DueAt           time.Time       `json:"dueAt"`
					HTMLURL         string          `json:"htmlUrl"`
					ScoreStatistics ScoreStatistics `json:"score_statistics,omitempty"`
				} `json:"nodes"`
			} `json:"assignmentsConnection"`
			EnrollmentsConnection interface{} `json:"enrollmentsConnection"`
		} `json:"allCourses"`
	} `json:"data"`
}

type ScoreStatistics struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Mean float64 `json:"mean"`
}

type ScoreRaw []struct {
	ScoreStatistics ScoreStatistics `json:"score_statistics,omitempty"`
}
