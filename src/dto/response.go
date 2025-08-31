package dto

// dto for response llm inference
type ChatResponse struct {
	Model              string                 `json:"model"`
	CreatedAt          string                 `json:"created_at"`
	Message            map[string]interface{} `json:"message"`
	Done               bool                   `json:"done"`
	TotalDuration      int64                  `json:"total_duration"`
	LoadDuration       int64                  `json:"load_duration"`
	PromptEvalCount    int64                  `json:"prompt_eval_count"`
	PromptEvalDuration int64                  `json:"prompt_eval_duration"`
	EvalCount          int64                  `json:"eval_count"`
	EvalDuration       int64                  `json:"eval_duration"`
}
