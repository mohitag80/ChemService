package models

// Question represents a chemistry exam question
type Question struct {
	ID           int      `json:"id"`
	QuestionText string   `json:"question_text"`
	Options      []string `json:"options"`
	CorrectAnswer string  `json:"correct_answer"`
	Topic        string   `json:"topic"`
	Grade        int      `json:"grade"`
	Complexity   string   `json:"complexity"` // easy | medium | hard
	Chapter      string   `json:"chapter"`
	Marks        int      `json:"marks"`
	Hint         string   `json:"hint,omitempty"`
}

// QuestionResponse wraps a list of questions with metadata
type QuestionResponse struct {
	Subject   string     `json:"subject"`
	Requested int        `json:"requested"`
	Returned  int        `json:"returned"`
	Questions []Question `json:"questions"`
}

// HealthResponse for the health check endpoint
type HealthResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

// ErrorResponse for error responses
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
