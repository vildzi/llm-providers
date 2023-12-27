package llmproviders

import "sync"

var textProviders = make(map[string]TextCompletionProvider)
var textProvidersMux sync.RWMutex
var DefaultTextProvider *string

func RegisterTextProvider(key string, provider TextCompletionProvider) {
	textProvidersMux.Lock()
	defer textProvidersMux.Unlock()

	textProviders[key] = provider
	if DefaultTextProvider == nil {
		DefaultTextProvider = &key
	}
}

func WithTextProvider(key string) TextCompletionProvider {
	textProvidersMux.RLock()
	defer textProvidersMux.RUnlock()

	return textProviders[key]
}

func SetDefaultTextProvider(key string) error {
	textProvidersMux.Lock()
	defer textProvidersMux.Unlock()

	if textProviders[key] == nil {

	}
	return nil
}

func WithDefaultTextProvider() TextCompletionProvider {
	textProvidersMux.RLock()
	defer textProvidersMux.RUnlock()

	if DefaultTextProvider == nil {
		return nil
	}
	return textProviders[*DefaultTextProvider]
}
