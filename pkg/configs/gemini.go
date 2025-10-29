package configs

import (
	"context"

	"google.golang.org/genai"
)

func NewGemini(ctx context.Context, apiKey string) *genai.Client {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		Log.Error().Err(err).Msg("Failed to create Gemini client")
		panic(err)
	}

	return client
}
