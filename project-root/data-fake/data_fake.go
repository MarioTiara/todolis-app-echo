package datafake

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/marioTiara/todolistapp/internal/api/models"
)

func GenerateFile() models.Files {
	var file models.Files

	err := faker.FakeData(&file)
	if err != nil {
		fmt.Println("Error:", err)
		return models.Files{}
	}
	return file
}

func GenerateFilesList() []models.Files {
	var files []models.Files
	for i := 0; 1 <= 5; i++ {
		file := GenerateFile()
		files = append(files, file)
	}
	return files
}

func GenerateTask(depth int) models.Task {
	if depth <= 0 {
		return models.Task{} // Return an empty task when depth is exhausted
	}

	var task models.Task

	// Generate fake data for the Task struct
	err := faker.FakeData(&task)
	if err != nil {
		fmt.Println("Error generating fake Task data:", err)
	}

	// Manually customize or exclude specific fields if needed
	task.Title = "CustomTitle"
	task.Description = "CustomDescription"

	// Generate fake data for the Files struct
	for i := 0; i < 3; i++ {
		var file models.Files
		err := faker.FakeData(&file)
		if err != nil {
			fmt.Println("Error generating fake Files data:", err)
		}

		// Manually set the TaskID for the Files struct
		file.TaskID = task.ID

		// Append the generated file to the task's Files slice
		task.Files = append(task.Files, file)
	}

	// Recursively generate fake data for children with reduced depth
	for i := 0; i < 2; i++ {
		childTask := GenerateTask(depth - 1)
		if childTask.ID != 0 {
			task.Children = append(task.Children, childTask)
		}
	}

	return task
}

func GenerateTasksList() []models.Task {
	var tasks []models.Task
	for i := 0; i < 10; i++ {
		task := GenerateTask(5)
		tasks = append(tasks, task)
	}

	return tasks
}
