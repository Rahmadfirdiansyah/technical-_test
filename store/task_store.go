package store

import "stability-test-task-api/models"

var Tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}

func GetAllTasks() []models.Task {
	return Tasks
}

func GetTaskByID(id int) *models.Task {
	for _, t := range Tasks {
		if t.ID == id {
			return &t
		}
	}
	return nil
}

func AddTask(task models.Task) {
	task.ID = len(Tasks) + 1
	Tasks = append(Tasks, task)
}

func DeleteTask(id int) bool {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true
		}
	}
}
