package Domain

type SuggestedTask struct {
	name string
}

func NewSuggestedTask(suggestedTask string) (*SuggestedTask, error) {
	return &SuggestedTask{
		name: suggestedTask,
	}, nil
}

func (t *SuggestedTask) GetName() string {
	return t.name
}
