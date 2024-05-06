package shared

import (
	"mime/multipart"
)

type TranscriptionRequest struct {
	File  multipart.File `json:"file"`
	Model string         `json:"model"`
}

type ChatCompletion struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Segment struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	Text  string  `json:"text"`
}

type Transcription struct {
	Task     string    `json:"task"`
	Language string    `json:"language"`
	Duration float64   `json:"duration"`
	Text     string    `json:"text"`
	Segments []Segment `json:"segments"`
}

type Result struct {
	Error         string    `json:"error"`
	Transcription []Segment `json:"transcription"`
	Content       string    `json:"content"`
	HashTags      string    `json:"hashTags"`
}
