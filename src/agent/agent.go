package agent

import (
	"context"
	"factorio-ai/src/provider"
)

type Agent struct {
	provider      *provider.LLMProvider
	mcp_client    string
	system_prompt string
	context       context.Context
	history       []string
}

func NewAgent(model *string, system_prompt string) *Agent {

	ctx := context.Background()
	provider := provider.NewLLMProvider(*model, ctx)

	return &Agent{
		provider:      provider,
		mcp_client:    "mock_mcp_client",
		system_prompt: system_prompt,
		context:       ctx,
		history:       []string{},
	}
}

func (agent *Agent) check_tool_call(message string) bool {

	return false
}

func (agent *Agent) handle_tool_call(message string) {
	if agent.check_tool_call(message) {

	}
}

func (agent *Agent) chat(message string) string {
	agent.history = append(agent.history, message)

	response := agent.provider.GenerateText(message)

	agent.history = append(agent.history, response)

	// handle mcp client handle tool
	agent.handle_tool_call(response)

	return response
}
