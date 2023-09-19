package Task

import (
	Domains "go-echo-rest-api-sample/src/Domains/Task"
	"testing"
	"time"
)

func TestCreateNewTask(t *testing.T) {
	// given
	name := "Test Task"
	deadline, pErr := time.Parse(Domains.DeadlineFormat, "2023-09-20")
	if pErr != nil {
		t.Errorf("Error parsing deadline: %v", pErr)
	}

	// when
	task, err := Domains.CreateNewTask(name, deadline)

	// then
	if err != nil {
		t.Fatalf("Error creating new task: %v", err)
	}

	if task.GetName() != name {
		t.Errorf("Expected task name to be %s, but got %s", name, task.GetName())
	}

	if !task.GetDeadline().Equal(deadline) {
		t.Errorf("Expected task deadline to be %v, but got %v", deadline, task.GetDeadline())
	}

	if task.GetIsFavorite() {
		t.Error("Expected task to not be favorite by default")
	}

	if task.GetIsCompleted() {
		t.Error("Expected task to not be completed by default")
	}
}

func TestReconstructTask(t *testing.T) {
	// given
	taskId, tErr := Domains.NewTaskId(1)
	if tErr != nil {
		t.Errorf("Error creating taskId: %v", tErr)
	}
	name := "Test Task"
	deadline, pErr := time.Parse(Domains.DeadlineFormat, "2023-09-20")
	if pErr != nil {
		t.Errorf("Error parsing deadline: %v", pErr)
	}
	isFavorite := true
	isCompleted := false

	// when
	task := Domains.ReconstructTask(taskId, name, deadline, isFavorite, isCompleted)

	// then
	if task.GetTaskId() != taskId {
		t.Errorf("Task ID should be %d, but got %d", taskId, task.GetTaskId())
	}

	if task.GetName() != name {
		t.Errorf("Task name should be '%s', but got '%s'", name, task.GetName())
	}

	if task.GetDeadline() != deadline {
		t.Errorf("Task deadline should be '%s', but got '%s'", deadline, task.GetDeadline())
	}

	if task.GetIsFavorite() != isFavorite {
		t.Errorf("Task should be favorite")
	}

	if task.GetIsCompleted() != isCompleted {
		t.Errorf("Task should not be completed")
	}
}

func TestUpdateTask(t *testing.T) {
	// given
	name := "Test Task"
	deadline, err := time.Parse(Domains.DeadlineFormat, "2023-09-20")
	if err != nil {
		t.Errorf("Error parsing deadline: %v", err)
	}

	task, err := Domains.CreateNewTask(name, deadline)
	if err != nil {
		t.Errorf("Error creating new task: %v", err)
	}

	newTaskName := "New Test Task"
	newDeadline, err := time.Parse(Domains.DeadlineFormat, "2023-09-21")
	if err != nil {
		t.Errorf("Error parsing deadline: %v", err)
	}

	// when
	task.UpdateTask(newTaskName, newDeadline)

	// then
	if task.GetName() != newTaskName {
		t.Errorf("Task name should be '%s', but got '%s'", newTaskName, task.GetName())
	}

	if task.GetDeadline() != newDeadline {
		t.Errorf("Task deadline should be '%s', but got '%s'", newDeadline, task.GetDeadline())
	}
}

func TestUpdateTaskFavorite(t *testing.T) {
	// given
	name := "Test Task"
	deadline, err := time.Parse(Domains.DeadlineFormat, "2023-09-20")
	if err != nil {
		t.Errorf("Error parsing deadline: %v", err)
	}

	task, err := Domains.CreateNewTask(name, deadline)
	if err != nil {
		t.Errorf("Error creating new task: %v", err)
	}

	// when
	task.UpdateTaskFavorite()

	// then
	if !task.GetIsFavorite() {
		t.Errorf("Task should be favorite")
	}
}

func TestUpdateTaskComplete(t *testing.T) {
	// given
	name := "Test Task"
	deadline, err := time.Parse(Domains.DeadlineFormat, "2023-09-20")
	if err != nil {
		t.Errorf("Error parsing deadline: %v", err)
	}

	task, err := Domains.CreateNewTask(name, deadline)
	if err != nil {
		t.Errorf("Error creating new task: %v", err)
	}

	// when
	task.UpdateTaskComplete()

	// then
	if !task.GetIsCompleted() {
		t.Errorf("Task should be completed")
	}
}

func TestGetters(t *testing.T) {
	// given
	name := "Test Task"
	deadline, err := time.Parse(Domains.DeadlineFormat, "2023-09-20")
	if err != nil {
		t.Errorf("Error parsing deadline: %v", err)
	}

	// when
	task, err := Domains.CreateNewTask(name, deadline)

	// then
	//if task.GetTaskId().GetValue() != 1 {
	//	t.Errorf("Task ID should be 1, but got %d", task.GetTaskId())
	//}

	if task.GetName() != name {
		t.Errorf("Task name should be '%s', but got '%s'", name, task.GetName())
	}

	if task.GetDeadline() != deadline {
		t.Errorf("Task deadline should be '%s', but got '%s'", deadline, task.GetDeadline())
	}

	if task.GetIsFavorite() {
		t.Errorf("Task should be favorite")
	}

	if task.GetIsCompleted() {
		t.Errorf("Task should not be completed")
	}
}
