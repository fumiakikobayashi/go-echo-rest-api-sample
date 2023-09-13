package Infrastructures

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"go-echo-rest-api-sample/src/Domains/SuggestedTask"
	"go-echo-rest-api-sample/src/Shared"
	"strings"
)

type SuggestionTask struct {
	client *openai.Client
}

func NewSuggestionTaskClient(client *openai.Client) SuggestionTask {
	return SuggestionTask{
		client: client,
	}
}

func (s SuggestionTask) SuggestTasksBy(target string) (Domain.SuggestedTasks, error) {
	res, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: target + "を実現するために必要なタスクを箇条書きで洗い出してください。タスクの上限は10個まで。1行につき30文字以内。いきなり箇条書きから出力を始め、余計な文章を含まず最後の箇条書きで終了してください。",
				},
			},
		},
	)

	if err != nil {
		return Domain.SuggestedTasks{}, Shared.NewSampleError("001-001", "doSomethingでエラー発生")
	}

	suggestedTasks := Domain.NewSuggestedTasks()
	lines := strings.Split(res.Choices[0].Message.Content, "\n")
	for _, line := range lines {
		suggestedTask, err := Domain.NewSuggestedTask(line)
		if err != nil {
			return Domain.SuggestedTasks{}, err
		}
		suggestedTasks = suggestedTasks.Push(*suggestedTask)
	}

	return suggestedTasks, nil
}
