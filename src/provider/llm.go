package provider

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

type LLMProvider struct {
	model   string
	client  *genai.Client
	context context.Context
	config  *genai.ClientConfig
}

func NewLLMProvider(model string, context context.Context) *LLMProvider {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
		APIKey:  os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	return &LLMProvider{
		model:   model,
		client:  client,
		context: context,
	}
}

func (p *LLMProvider) GenerateText(prompt string) string {
	config := &genai.GenerateContentConfig{}

	var contents = []*genai.Content{
		&genai.Content{
			Role:  "user",
			Parts: []*genai.Part{&genai.Part{Text: prompt}},
		},
	}

	resp, err := p.client.Models.GenerateContent(p.context, p.model, contents,
		config)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Text()
}
