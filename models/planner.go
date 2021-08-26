package models

type Planner []PlannerElement

type PlannerElement struct {
	ContextType     ContextType       `json:"context_type"`
	CourseID        int64             `json:"course_id"`
	PlannableID     int64             `json:"plannable_id"`
	PlannerOverride interface{}       `json:"planner_override"`
	PlannableType   PlannableType     `json:"plannable_type"`
	NewActivity     bool              `json:"new_activity"`
	Submissions     *SubmissionsUnion `json:"submissions"`
	PlannableDate   string            `json:"plannable_date"`
	Plannable       Plannable         `json:"plannable"`
	HTMLURL         string            `json:"html_url"`
	ContextName     ContextName       `json:"context_name"`
	ContextImage    *string           `json:"context_image"`
}

type Plannable struct {
	ID              int64      `json:"id"`
	Title           string     `json:"title"`
	UnreadCount     *int64     `json:"unread_count,omitempty"`
	ReadState       *ReadState `json:"read_state,omitempty"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	LocationName    *string    `json:"location_name,omitempty"`
	AllDay          *bool      `json:"all_day,omitempty"`
	LocationAddress *string    `json:"location_address,omitempty"`
	Description     *string    `json:"description,omitempty"`
	StartAt         *string    `json:"start_at,omitempty"`
	EndAt           *string    `json:"end_at,omitempty"`
	PointsPossible  *float64   `json:"points_possible,omitempty"`
	DueAt           *string    `json:"due_at,omitempty"`
}

type SubmissionsClass struct {
	Submitted    bool      `json:"submitted"`
	Excused      bool      `json:"excused"`
	Graded       bool      `json:"graded"`
	Late         bool      `json:"late"`
	Missing      bool      `json:"missing"`
	NeedsGrading bool      `json:"needs_grading"`
	HasFeedback  bool      `json:"has_feedback"`
	RedoRequest  bool      `json:"redo_request"`
	Feedback     *Feedback `json:"feedback,omitempty"`
}

type Feedback struct {
	Comment         string `json:"comment"`
	IsMedia         bool   `json:"is_media"`
	AuthorName      string `json:"author_name"`
	AuthorAvatarURL string `json:"author_avatar_url"`
}

type ContextName string

type ContextType string

type ReadState string

type PlannableType string

type SubmissionsUnion struct {
	Bool             *bool
	SubmissionsClass *SubmissionsClass
}
