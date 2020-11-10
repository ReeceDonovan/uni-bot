package api

import "time"

type courseData []struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	AccountID            int         `json:"account_id"`
	UUID                 string      `json:"uuid"`
	StartAt              time.Time   `json:"start_at"`
	GradingStandardID    interface{} `json:"grading_standard_id"`
	IsPublic             bool        `json:"is_public"`
	CreatedAt            time.Time   `json:"created_at"`
	CourseCode           string      `json:"course_code"`
	DefaultView          string      `json:"default_view"`
	RootAccountID        int         `json:"root_account_id"`
	EnrollmentTermID     int         `json:"enrollment_term_id"`
	License              string      `json:"license"`
	GradePassbackSetting interface{} `json:"grade_passback_setting"`
	EndAt                interface{} `json:"end_at"`
	PublicSyllabus       bool        `json:"public_syllabus"`
	PublicSyllabusToAuth bool        `json:"public_syllabus_to_auth"`
	StorageQuotaMb       int         `json:"storage_quota_mb"`
	IsPublicToAuthUsers  bool        `json:"is_public_to_auth_users"`
	Term                 struct {
		ID                   int         `json:"id"`
		Name                 string      `json:"name"`
		StartAt              interface{} `json:"start_at"`
		EndAt                interface{} `json:"end_at"`
		CreatedAt            time.Time   `json:"created_at"`
		WorkflowState        string      `json:"workflow_state"`
		GradingPeriodGroupID interface{} `json:"grading_period_group_id"`
	} `json:"term"`
	ApplyAssignmentGroupWeights bool `json:"apply_assignment_group_weights"`
	Calendar                    struct {
		Ics string `json:"ics"`
	} `json:"calendar"`
	TimeZone    string `json:"time_zone"`
	Blueprint   bool   `json:"blueprint"`
	Enrollments []struct {
		Type                           string `json:"type"`
		Role                           string `json:"role"`
		RoleID                         int    `json:"role_id"`
		UserID                         int    `json:"user_id"`
		EnrollmentState                string `json:"enrollment_state"`
		LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
	} `json:"enrollments"`
	HideFinalGrades                  bool   `json:"hide_final_grades"`
	WorkflowState                    string `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates"`
	OverriddenCourseVisibility       string `json:"overridden_course_visibility"`
}

type assignmentData []struct {
	ID                              int         `json:"id"`
	Description                     string      `json:"description"`
	DueAt                           time.Time   `json:"due_at"`
	UnlockAt                        interface{} `json:"unlock_at"`
	LockAt                          interface{} `json:"lock_at"`
	PointsPossible                  float64     `json:"points_possible"`
	GradingType                     string      `json:"grading_type"`
	AssignmentGroupID               int         `json:"assignment_group_id"`
	GradingStandardID               interface{} `json:"grading_standard_id"`
	CreatedAt                       time.Time   `json:"created_at"`
	UpdatedAt                       time.Time   `json:"updated_at"`
	PeerReviews                     bool        `json:"peer_reviews"`
	AutomaticPeerReviews            bool        `json:"automatic_peer_reviews"`
	Position                        int         `json:"position"`
	GradeGroupStudentsIndividually  bool        `json:"grade_group_students_individually"`
	AnonymousPeerReviews            bool        `json:"anonymous_peer_reviews"`
	GroupCategoryID                 interface{} `json:"group_category_id"`
	PostToSis                       bool        `json:"post_to_sis"`
	ModeratedGrading                bool        `json:"moderated_grading"`
	OmitFromFinalGrade              bool        `json:"omit_from_final_grade"`
	IntraGroupPeerReviews           bool        `json:"intra_group_peer_reviews"`
	AnonymousInstructorAnnotations  bool        `json:"anonymous_instructor_annotations"`
	AnonymousGrading                bool        `json:"anonymous_grading"`
	GradersAnonymousToGraders       bool        `json:"graders_anonymous_to_graders"`
	GraderCount                     int         `json:"grader_count"`
	GraderCommentsVisibleToGraders  bool        `json:"grader_comments_visible_to_graders"`
	FinalGraderID                   interface{} `json:"final_grader_id"`
	GraderNamesVisibleToFinalGrader bool        `json:"grader_names_visible_to_final_grader"`
	AllowedAttempts                 int         `json:"allowed_attempts"`
	SecureParams                    string      `json:"secure_params"`
	CourseID                        int         `json:"course_id"`
	Name                            string      `json:"name"`
	SubmissionTypes                 []string    `json:"submission_types"`
	HasSubmittedSubmissions         bool        `json:"has_submitted_submissions"`
	DueDateRequired                 bool        `json:"due_date_required"`
	MaxNameLength                   int         `json:"max_name_length"`
	InClosedGradingPeriod           bool        `json:"in_closed_grading_period"`
	IsQuizAssignment                bool        `json:"is_quiz_assignment"`
	CanDuplicate                    bool        `json:"can_duplicate"`
	OriginalCourseID                interface{} `json:"original_course_id"`
	OriginalAssignmentID            interface{} `json:"original_assignment_id"`
	OriginalAssignmentName          interface{} `json:"original_assignment_name"`
	OriginalQuizID                  interface{} `json:"original_quiz_id"`
	WorkflowState                   string      `json:"workflow_state"`
	Muted                           bool        `json:"muted"`
	HTMLURL                         string      `json:"html_url"`
	Published                       bool        `json:"published"`
	OnlyVisibleToOverrides          bool        `json:"only_visible_to_overrides"`
	LockedForUser                   bool        `json:"locked_for_user"`
	SubmissionsDownloadURL          string      `json:"submissions_download_url"`
	PostManually                    bool        `json:"post_manually"`
	AnonymizeStudents               bool        `json:"anonymize_students"`
	RequireLockdownBrowser          bool        `json:"require_lockdown_browser"`
}

type ParsedCourse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CourseCode string `json:"course_code"`
}
type ParsedAssignment struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	HTMLURL     string    `json:"html_url"`
	DueAt       time.Time `json:"due_at"`
}
