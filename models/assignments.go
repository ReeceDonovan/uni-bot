package models

type Assignments []AssignmentElement

type AssignmentElement struct {
	ID                              int64            `json:"id"`
	DueAt                           string           `json:"due_at"`
	UnlockAt                        *string          `json:"unlock_at"`
	LockAt                          string           `json:"lock_at"`
	PointsPossible                  float64          `json:"points_possible"`
	GradingType                     string           `json:"grading_type"`
	AssignmentGroupID               int64            `json:"assignment_group_id"`
	GradingStandardID               interface{}      `json:"grading_standard_id"`
	CreatedAt                       string           `json:"created_at"`
	UpdatedAt                       string           `json:"updated_at"`
	PeerReviews                     bool             `json:"peer_reviews"`
	AutomaticPeerReviews            bool             `json:"automatic_peer_reviews"`
	Position                        int64            `json:"position"`
	GradeGroupStudentsIndividually  bool             `json:"grade_group_students_individually"`
	AnonymousPeerReviews            bool             `json:"anonymous_peer_reviews"`
	GroupCategoryID                 interface{}      `json:"group_category_id"`
	PostToSis                       bool             `json:"post_to_sis"`
	ModeratedGrading                bool             `json:"moderated_grading"`
	OmitFromFinalGrade              bool             `json:"omit_from_final_grade"`
	IntraGroupPeerReviews           bool             `json:"intra_group_peer_reviews"`
	AnonymousInstructorAnnotations  bool             `json:"anonymous_instructor_annotations"`
	AnonymousGrading                bool             `json:"anonymous_grading"`
	GradersAnonymousToGraders       bool             `json:"graders_anonymous_to_graders"`
	GraderCount                     int64            `json:"grader_count"`
	GraderCommentsVisibleToGraders  bool             `json:"grader_comments_visible_to_graders"`
	FinalGraderID                   interface{}      `json:"final_grader_id"`
	GraderNamesVisibleToFinalGrader bool             `json:"grader_names_visible_to_final_grader"`
	AllowedAttempts                 int64            `json:"allowed_attempts"`
	LockInfo                        LockInfo         `json:"lock_info"`
	SecureParams                    string           `json:"secure_params"`
	CourseID                        int64            `json:"course_id"`
	Name                            string           `json:"name"`
	SubmissionTypes                 []string         `json:"submission_types"`
	HasSubmittedSubmissions         bool             `json:"has_submitted_submissions"`
	DueDateRequired                 bool             `json:"due_date_required"`
	MaxNameLength                   int64            `json:"max_name_length"`
	InClosedGradingPeriod           bool             `json:"in_closed_grading_period"`
	IsQuizAssignment                bool             `json:"is_quiz_assignment"`
	CanDuplicate                    bool             `json:"can_duplicate"`
	OriginalCourseID                interface{}      `json:"original_course_id"`
	OriginalAssignmentID            interface{}      `json:"original_assignment_id"`
	OriginalAssignmentName          interface{}      `json:"original_assignment_name"`
	OriginalQuizID                  interface{}      `json:"original_quiz_id"`
	WorkflowState                   string           `json:"workflow_state"`
	ImportantDates                  bool             `json:"important_dates"`
	Description                     *string          `json:"description"`
	Muted                           bool             `json:"muted"`
	HTMLURL                         string           `json:"html_url"`
	AllDates                        []AllDate        `json:"all_dates"`
	Published                       bool             `json:"published"`
	OnlyVisibleToOverrides          bool             `json:"only_visible_to_overrides"`
	Submission                      Submission       `json:"submission"`
	LockedForUser                   bool             `json:"locked_for_user"`
	LockExplanation                 string           `json:"lock_explanation"`
	SubmissionsDownloadURL          string           `json:"submissions_download_url"`
	PostManually                    bool             `json:"post_manually"`
	AnonymizeStudents               bool             `json:"anonymize_students"`
	RequireLockdownBrowser          bool             `json:"require_lockdown_browser"`
	AllowedExtensions               []string         `json:"allowed_extensions"`
	ScoreStatistics                 *ScoreStatistics `json:"score_statistics,omitempty"`
}

type AllDate struct {
	DueAt    string  `json:"due_at"`
	UnlockAt *string `json:"unlock_at"`
	LockAt   string  `json:"lock_at"`
	Base     *bool   `json:"base,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	Title    *string `json:"title,omitempty"`
	SetType  *string `json:"set_type,omitempty"`
	SetID    *int64  `json:"set_id,omitempty"`
}

type LockInfo struct {
	AssetString   string         `json:"asset_string"`
	ContextModule *ContextModule `json:"context_module,omitempty"`
	LockAt        *string        `json:"lock_at,omitempty"`
	CanView       *bool          `json:"can_view,omitempty"`
}

type ContextModule struct {
	ID                        int64         `json:"id"`
	ContextID                 int64         `json:"context_id"`
	ContextType               string        `json:"context_type"`
	Name                      string        `json:"name"`
	Position                  int64         `json:"position"`
	Prerequisites             []interface{} `json:"prerequisites"`
	CompletionRequirements    []interface{} `json:"completion_requirements"`
	CreatedAt                 string        `json:"created_at"`
	UpdatedAt                 string        `json:"updated_at"`
	WorkflowState             string        `json:"workflow_state"`
	DeletedAt                 interface{}   `json:"deleted_at"`
	UnlockAt                  *string       `json:"unlock_at"`
	MigrationID               interface{}   `json:"migration_id"`
	RequireSequentialProgress bool          `json:"require_sequential_progress"`
	ClonedItemID              interface{}   `json:"cloned_item_id"`
	CompletionEvents          interface{}   `json:"completion_events"`
	RequirementCount          interface{}   `json:"requirement_count"`
	RootAccountID             int64         `json:"root_account_id"`
}

type ScoreStatistics struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Mean float64 `json:"mean"`
}

type Submission struct {
	ID                            int64        `json:"id"`
	Body                          *string      `json:"body"`
	URL                           interface{}  `json:"url"`
	SubmittedAt                   *string      `json:"submitted_at"`
	AssignmentID                  int64        `json:"assignment_id"`
	UserID                        int64        `json:"user_id"`
	SubmissionType                *string      `json:"submission_type"`
	WorkflowState                 string       `json:"workflow_state"`
	GradeMatchesCurrentSubmission bool         `json:"grade_matches_current_submission"`
	GradedAt                      *string      `json:"graded_at"`
	GraderID                      *int64       `json:"grader_id"`
	Attempt                       *int64       `json:"attempt"`
	CachedDueDate                 string       `json:"cached_due_date"`
	Excused                       *bool        `json:"excused"`
	LatePolicyStatus              interface{}  `json:"late_policy_status"`
	PointsDeducted                interface{}  `json:"points_deducted"`
	GradingPeriodID               interface{}  `json:"grading_period_id"`
	ExtraAttempts                 interface{}  `json:"extra_attempts"`
	PostedAt                      *string      `json:"posted_at"`
	Late                          bool         `json:"late"`
	Missing                       bool         `json:"missing"`
	SecondsLate                   int64        `json:"seconds_late"`
	PreviewURL                    string       `json:"preview_url"`
	Grade                         *string      `json:"grade,omitempty"`
	Score                         *float64     `json:"score,omitempty"`
	EnteredGrade                  *string      `json:"entered_grade,omitempty"`
	EnteredScore                  *float64     `json:"entered_score,omitempty"`
	Attachments                   []Attachment `json:"attachments"`
}

type Attachment struct {
	ID            int64       `json:"id"`
	UUID          string      `json:"uuid"`
	FolderID      int64       `json:"folder_id"`
	DisplayName   string      `json:"display_name"`
	Filename      string      `json:"filename"`
	UploadStatus  string      `json:"upload_status"`
	ContentType   string      `json:"content-type"`
	URL           string      `json:"url"`
	Size          int64       `json:"size"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	UnlockAt      interface{} `json:"unlock_at"`
	Locked        bool        `json:"locked"`
	Hidden        bool        `json:"hidden"`
	LockAt        interface{} `json:"lock_at"`
	HiddenForUser bool        `json:"hidden_for_user"`
	ThumbnailURL  interface{} `json:"thumbnail_url"`
	ModifiedAt    string      `json:"modified_at"`
	MIMEClass     string      `json:"mime_class"`
	MediaEntryID  interface{} `json:"media_entry_id"`
	LockedForUser bool        `json:"locked_for_user"`
	PreviewURL    string      `json:"preview_url"`
}
