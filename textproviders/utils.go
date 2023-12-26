package textproviders

import llmproviders "github.com/vildzi/llm-providers"

func withDefaultOptions(options *llmproviders.TextCompletionOptions) *llmproviders.TextCompletionOptions {
	defaultOptions := &llmproviders.TextCompletionOptions{
		MaxTokens:   512,
		Temperature: 1,
		Messages:    []llmproviders.TextCompletionMessage{},
	}

	if options == nil {
		return defaultOptions
	}

	if options.MaxTokens == 0 {
		options.MaxTokens = defaultOptions.MaxTokens
	}
	if options.Temperature == 0 {
		options.Temperature = defaultOptions.Temperature
	}
	if options.Messages == nil {
		options.Messages = defaultOptions.Messages
	}

	return options
}

func WithSystemPrompt(prompt string, messages []llmproviders.TextCompletionMessage) []llmproviders.TextCompletionMessage {
	if messages == nil {
		messages = []llmproviders.TextCompletionMessage{}
	}
	newMessages := make([]llmproviders.TextCompletionMessage, 0, len(messages)+1)
	newMessages = append(newMessages, llmproviders.TextCompletionMessage{
		Role:    llmproviders.TextCompletionRoleSystem,
		Message: prompt,
	})
	newMessages = append(newMessages, messages...)
	return newMessages
}
