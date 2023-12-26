package textproviders

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/pkg/errors"
	llmproviders "github.com/vildzi/llm-providers"
)

type OpenAIProvider struct {
	llmproviders.TextCompletionProvider

	client *azopenai.Client
}

const defaultOpenAIEndpoint = "https://api.openai.com/v1"

type OpenAIProviderOptions struct {
	Endpoint string
}

type OpenAIModel string

const (
	OpenAIModelGPT35Turbo    OpenAIModel = "gpt-3.5-turbo"
	OpenAIModelGPT35Turbo16k OpenAIModel = "gpt-3.5-turbo-16k"
	OpenAIModelGPT4          OpenAIModel = "gpt-4"
	OpenAIModelGPT4128k      OpenAIModel = "gpt-4-1106-preview"
)

func NewOpenAIProvider(apiKey string, options *OpenAIProviderOptions) (*OpenAIProvider, error) {
	keyCredential := azcore.NewKeyCredential(apiKey)

	endpoint := defaultOpenAIEndpoint
	if options != nil {
		if options.Endpoint != "" {
			endpoint = options.Endpoint
		}
	}

	oaiClient, err := azopenai.NewClientForOpenAI(endpoint, keyCredential, nil)

	if err != nil {
		return nil, errors.Wrap(err, "init openai client")
	}

	return &OpenAIProvider{
		client: oaiClient,
	}, nil
}

func (p *OpenAIProvider) convertMessages(messages []llmproviders.TextCompletionMessage) []azopenai.ChatRequestMessageClassification {
	convertedMessages := make([]azopenai.ChatRequestMessageClassification, len(messages))

	for i, genericMessage := range messages {
		switch genericMessage.Role {
		case llmproviders.TextCompletionRoleSystem:
			convertedMessages[i] = &azopenai.ChatRequestSystemMessage{Content: to.Ptr(genericMessage.Message)}
		case llmproviders.TextCompletionRoleAssistant:
			convertedMessages[i] = &azopenai.ChatRequestAssistantMessage{Content: to.Ptr(genericMessage.Message)}
		case llmproviders.TextCompletionRoleUser:
			convertedMessages[i] = &azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent(genericMessage.Message)}
		}
	}

	return convertedMessages
}

func (p *OpenAIProvider) withDefaultOptions(options *llmproviders.TextCompletionOptions) *llmproviders.TextCompletionOptions {
	baseDefaultOptions := withDefaultOptions(options)
	if baseDefaultOptions.Model == "" {
		baseDefaultOptions.Model = string(OpenAIModelGPT35Turbo)
	}
	return baseDefaultOptions
}

func (p *OpenAIProvider) GetCompletion(ctx context.Context, input string, options *llmproviders.TextCompletionOptions) (llmproviders.TextCompletionResponse, error) {
	// options will never be nil after this call
	options = p.withDefaultOptions(options)

	messages := append(options.Messages, llmproviders.TextCompletionMessage{
		Role:    llmproviders.TextCompletionRoleUser,
		Message: input,
	})

	completions, err := p.client.GetChatCompletions(ctx, azopenai.ChatCompletionsOptions{
		MaxTokens:      &options.MaxTokens,
		Temperature:    &options.Temperature,
		Messages:       p.convertMessages(messages),
		DeploymentName: to.Ptr(options.Model),
	}, nil)

	if err != nil {
		return llmproviders.TextCompletionResponse{}, errors.Wrap(err, "get text completion")
	}

	choices := make([]llmproviders.TextCompletionChoice, len(completions.Choices))
	for i, azChoice := range completions.Choices {
		message := llmproviders.TextCompletionMessage{
			Role:    llmproviders.TextCompletionRoleAssistant,
			Message: "",
		}

		if azChoice.Message != nil && azChoice.Message.Content != nil {
			message.Message = *azChoice.Message.Content
		}

		choices[i] = llmproviders.TextCompletionChoice{
			Message: message,
		}
	}

	return llmproviders.TextCompletionResponse{
		Choices: choices,
	}, nil
}
