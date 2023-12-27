# llm-providers

This is a package meant to give a unified API to access multiple LLM platforms. Currently only an OpenAI / Azure OpenAI provider is implemented.

> [!WARNING]  
> This package is a work in progress, expect breaking changes.

## Installation

`go get -u github.com/vildzi/llm-providers@master`

## Examples

### OpenAI

```go
import (
    llmproviders "github.com/vildzi/llm-providers"
    "github.com/vildzi/llm-providers/textproviders"
)

...

oaiProvider, err := textproviders.NewOpenAIProvider("<your openai api key>", nil)
if err != nil {
    panic("failed to init openai provider")
}

llmproviders.RegisterProvider("openai", oaiProvider)
// also becomes the default provider as the first registered provider, can be accessed with llmproviders.WithDefaultProvider()

completions, err := llmproviders.WithTextProvider("openai").GetCompletion(context.TODO(), "Tell me a joke", &llmproviders.TextCompletionOptions{
    Messages: textproviders.WithSystemPrompt("you are a helpful assistant", nil),
})
...
```
