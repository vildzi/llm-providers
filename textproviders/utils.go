package textproviders

import llmproviders "github.com/vildzi/llm-providers"

func withDefaultOptions(options llmproviders.TextCompletionOptions) llmproviders.TextCompletionOptions {
	if options.MaxTokens == 0 {
		options.MaxTokens = 512
	}
	if options.Temperature == 0 {
		options.Temperature = 1
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
