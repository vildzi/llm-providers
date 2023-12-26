package tests

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vildzi/llm-providers"
	"github.com/vildzi/llm-providers/textproviders"
	"os"
	"testing"
)

const (
	OpenAIProvider = "openai"
)

func TestTextProviders(t *testing.T) {
	oaiProvider, err := textproviders.NewOpenAIProvider(os.Getenv("OAI_KEY"), nil)
	if err != nil {
		t.Fatal("failed to init openai provider")
	}

	llmproviders.RegisterTextProvider(OpenAIProvider, oaiProvider)

	completions, err := llmproviders.WithTextProvider(OpenAIProvider).GetCompletion(context.TODO(), "Tell me a joke", llmproviders.TextCompletionOptions{
		Messages: textproviders.WithSystemPrompt("you are a helpful assistant", nil),
	})
	if err != nil {
		t.Fatal("failed to get completions", err)
	}

	t.Log(completions)
}
