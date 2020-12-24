package request

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
				Edges []struct {
					Node struct {
						Name    string      `json:"name"`
						DueAt   interface{} `json:"dueAt"`
						HTMLURL string      `json:"htmlUrl"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"assignmentsConnection"`
			EnrollmentsConnection interface{} `json:"enrollmentsConnection"`
		} `json:"allCourses"`
	} `json:"data"`
}


