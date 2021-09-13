package models

type Courses []Course

type Course struct {
	ID                               int64         `json:"id"`
	RootAccountID                    int64         `json:"root_account_id"`
	AccountID                        int64         `json:"account_id"`
	Name                             string        `json:"name"`
	GradingStandardID                interface{}   `json:"grading_standard_id"`
	EnrollmentTermID                 int64         `json:"enrollment_term_id"`
	UUID                             string        `json:"uuid"`
	StartAt                          string        `json:"start_at"`
	IsPublic                         bool          `json:"is_public"`
	CreatedAt                        string        `json:"created_at"`
	CourseCode                       string        `json:"course_code"`
	DefaultView                      DefaultView   `json:"default_view"`
	License                          License       `json:"license"`
	GradePassbackSetting             interface{}   `json:"grade_passback_setting"`
	EndAt                            interface{}   `json:"end_at"`
	PublicSyllabus                   bool          `json:"public_syllabus"`
	PublicSyllabusToAuth             bool          `json:"public_syllabus_to_auth"`
	StorageQuotaMB                   int64         `json:"storage_quota_mb"`
	IsPublicToAuthUsers              bool          `json:"is_public_to_auth_users"`
	HomeroomCourse                   bool          `json:"homeroom_course"`
	CourseColor                      interface{}   `json:"course_color"`
	FriendlyName                     interface{}   `json:"friendly_name"`
	Term                             Term          `json:"term"`
	ApplyAssignmentGroupWeights      bool          `json:"apply_assignment_group_weights"`
	Teachers                         []Teacher     `json:"teachers"`
	Calendar                         Calendar      `json:"calendar"`
	TimeZone                         TimeZone      `json:"time_zone"`
	Blueprint                        bool          `json:"blueprint"`
	Template                         bool          `json:"template"`
	Enrollments                      []Enrollment  `json:"enrollments"`
	HideFinalGrades                  bool          `json:"hide_final_grades"`
	WorkflowState                    WorkflowState `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool          `json:"restrict_enrollments_to_course_dates"`
	OverriddenCourseVisibility       string        `json:"overridden_course_visibility"`
	HasGradingPeriods                bool          `json:"has_grading_periods"`
	MultipleGradingPeriodsEnabled    bool          `json:"multiple_grading_periods_enabled"`
	HasWeightedGradingPeriods        bool          `json:"has_weighted_grading_periods"`
}

type Calendar struct {
	ICS string `json:"ics"`
}

type Enrollment struct {
	Type                              Type        `json:"type"`
	Role                              Role        `json:"role"`
	RoleID                            int64       `json:"role_id"`
	UserID                            int64       `json:"user_id"`
	EnrollmentState                   State       `json:"enrollment_state"`
	LimitPrivilegesToCourseSection    bool        `json:"limit_privileges_to_course_section"`
	CurrentGradingPeriodID            interface{} `json:"current_grading_period_id"`
	CurrentGradingPeriodTitle         interface{} `json:"current_grading_period_title"`
	HasGradingPeriods                 bool        `json:"has_grading_periods"`
	MultipleGradingPeriodsEnabled     bool        `json:"multiple_grading_periods_enabled"`
	ComputedCurrentGrade              interface{} `json:"computed_current_grade"`
	ComputedCurrentScore              *float64    `json:"computed_current_score"`
	ComputedFinalGrade                interface{} `json:"computed_final_grade"`
	ComputedFinalScore                *float64    `json:"computed_final_score"`
	TotalsForAllGradingPeriodsOption  *bool       `json:"totals_for_all_grading_periods_option,omitempty"`
	CurrentPeriodComputedCurrentScore interface{} `json:"current_period_computed_current_score"`
	CurrentPeriodComputedFinalScore   interface{} `json:"current_period_computed_final_score"`
	CurrentPeriodComputedCurrentGrade interface{} `json:"current_period_computed_current_grade"`
	CurrentPeriodComputedFinalGrade   interface{} `json:"current_period_computed_final_grade"`
}

type Teacher struct {
	ID             int64       `json:"id"`
	DisplayName    string      `json:"display_name"`
	AvatarImageURL string      `json:"avatar_image_url"`
	HTMLURL        string      `json:"html_url"`
	Pronouns       interface{} `json:"pronouns"`
}

type Term struct {
	ID                   int64       `json:"id"`
	GradingPeriodGroupID interface{} `json:"grading_period_group_id"`
	Name                 string      `json:"name"`
	StartAt              interface{} `json:"start_at"`
	EndAt                interface{} `json:"end_at"`
	CreatedAt            string      `json:"created_at"`
	WorkflowState        State       `json:"workflow_state"`
}

type DefaultView string

type State string

type Role string

type Type string

type License string

type Name string

type TimeZone string

type WorkflowState string
