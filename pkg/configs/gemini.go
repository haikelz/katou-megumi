package configs

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/genai"
)

func NewGemini(ctx context.Context, apiKey string) *genai.Client {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		zap.L().Error("Failed to create Gemini client", zap.Error(err))
		panic(err)
	}

	return client
}
