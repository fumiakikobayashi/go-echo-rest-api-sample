package Domain

type SuggestedTasks struct {
	tasks []SuggestedTask
}

func NewSuggestedTasks() SuggestedTasks {
	return SuggestedTasks{
		tasks: []SuggestedTask{},
	}
}

func (st SuggestedTasks) Push(suggestedTask SuggestedTask) SuggestedTasks {
	st.tasks = append(st.tasks, suggestedTask)
	return SuggestedTasks{
		tasks: st.tasks,
	}
}

func (st SuggestedTasks) GetSuggestedTasks() []SuggestedTask {
	return st.tasks
}
