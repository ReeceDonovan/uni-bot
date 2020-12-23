package request

// TODO: Clean up struct, think of any other useful data we could use from the api

type Assignment struct {
	Data struct {
		AllCourses []struct {
			CourseName            string `json:"name"`
			AssignmentsConnection struct {
				Edges []struct {
					Node struct {
						Name string `json:"name"`
						DueAt   interface{} `json:"dueAt"`
						HTMLURL string      `json:"htmlUrl"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"assignmentsConnection"`
		} `json:"allCourses"`
	} `json:"data"`
}
