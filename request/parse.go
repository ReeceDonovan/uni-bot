package request

import "time"

// TODO: Clean up struct, think of any other useful data we could use from the api

type CourseAssignment struct {
	Data struct {
		AllCourses []struct {
			AssignmentsConnection struct {
				Nodes []struct {
					DueAt           time.Time       `json:"dueAt"`
					HTMLURL         string          `json:"htmlUrl"`
					ID              string          `json:"_id"`
					Name            string          `json:"name"`
					PointsPossible  float64         `json:"pointsPossible"`
					ScoreStatistics ScoreStatistics `json:"score_statistics,omitempty"`
				} `json:"nodes"`
			} `json:"assignmentsConnection"`
			CourseCode            string `json:"courseCode"`
			EnrollmentsConnection struct {
				Nodes []struct {
					Type string `json:"type"`
					User struct {
						ID   string `json:"_id"`
						Name string `json:"name"`
					} `json:"user"`
				} `json:"nodes"`
			} `json:"enrollmentsConnection"`
			ID         string `json:"_id"`
			CourseName string `json:"name"`
			State      string `json:"state"`
			Term       struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"term"`
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
