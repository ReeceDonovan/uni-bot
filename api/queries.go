package request

var AssignmentQuery = []byte(`
{"query": "query CourseAssignments {
	allCourses {
		_id
		name
		state
		courseCode
		term {
		_id
		name
		}
		assignmentsConnection {
		nodes {
			_id
			name
			dueAt
			htmlUrl
			pointsPossible
		}
		}
		enrollmentsConnection {
		nodes {
			type
			user {
			_id
			name
			}
		}
		}
	}
	}
	
"}`)
