package llmproviders

import (
	"context"
	"time"
)

type TextCompletionProvider interface {
	GetCompletion(ctx context.Context, query string, options *TextCompletionOptions) (TextCompletionResponse, error)
}

type TextCompletionOptions struct {
	Model       string
	Messages    []TextCompletionMessage
	MaxTokens   int32
	Temperature float32
}

type TextCompletionMessageRole string

const (
	TextCompletionRoleSystem    TextCompletionMessageRole = "system"
	TextCompletionRoleUser      TextCompletionMessageRole = "user"
	TextCompletionRoleAssistant TextCompletionMessageRole = "assistant"
)

type TextCompletionMessage struct {
	Role      TextCompletionMessageRole
	Message   string
	Timestamp time.Time
}

type TextCompletionChoice struct {
	Message TextCompletionMessage
}

type TextCompletionResponse struct {
	Choices []TextCompletionChoice
}
