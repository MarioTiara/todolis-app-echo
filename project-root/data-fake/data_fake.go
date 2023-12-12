package datafake

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
)

func GenerateFile() models.Files {
	var file models.Files
	rand, _ := faker.RandomInt(10, 23)
	file.ID = uint(rand[0])
	file.FileName = faker.Sentence()
	file.FileSize = uint(rand[1])
	file.FileURL = faker.Sentence()
	file.CreatedAt = time.Now()
	file.TaskID = uint(rand[2])
	return file
}

func GenerateFilesList(lenght int) []models.Files {
	var files []models.Files
	for i := 0; i < lenght; i++ {
		file := GenerateFile()
		files = append(files, file)
	}
	return files
}

func GenerateTask() models.Task {
	var task models.Task

	rand, _ := faker.RandomInt(50, 500)
	task.ID = uint(rand[0])
	task.Title = faker.Sentence()
	task.Description = faker.Sentence()
	task.CreatedAt = time.Now().UTC()
	task.UpdatedAt = time.Now().UTC()
	task.Priority = rand[1]
	task.Checked = false
	task.IsActive = true
	task.Files = GenerateFilesList(3)
	task.ParentID = nil
	task.Children = append(task.Children, models.Task{
		ID:          uint(rand[0]),
		Title:       faker.Sentence(),
		Description: faker.Sentence(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Priority:    rand[1],
		Checked:     false,
		IsActive:    true,
		ParentID:    &task.ID,
		Files:       GenerateFilesList(2),
	})
	return task
}

func GenerateSubtask() models.Task {
	intValue := 0
	uintValue := uint(intValue)
	var task models.Task

	rand, _ := faker.RandomInt(50, 500)
	task.ID = uint(rand[0])
	task.Title = faker.Sentence()
	task.Description = faker.Sentence()
	task.CreatedAt = time.Now().UTC()
	task.UpdatedAt = time.Now().UTC()
	task.Priority = rand[1]
	task.Checked = false
	task.IsActive = true
	task.Files = GenerateFilesList(3)
	task.ParentID = &uintValue

	return task
}

func GenerateSubtaskList(lenght int) []models.Task {
	var tasks []models.Task
	for i := 0; i < lenght; i++ {
		task := GenerateSubtask()
		tasks = append(tasks, task)
	}

	return tasks
}

func GenerateTasksList(lenght int) []models.Task {
	var tasks []models.Task
	for i := 0; i < lenght; i++ {
		task := GenerateTask()
		tasks = append(tasks, task)
	}

	return tasks
}

func GenerateAddTaskRequest(numberOfChild int) dtos.AddTaskRequest {
	request := dtos.AddTaskRequest{}
	request.Title = faker.Sentence()
	request.Description = faker.Sentence()
	request.Priority = 1
	for i := 0; i < numberOfChild; i++ {
		child := dtos.AddTaskRequest{
			Title:       faker.Sentence(),
			Description: faker.Sentence(),
			Priority:    1,
		}
		request.Children = append(request.Children, child)
	}
	return request
}

func GenerateAddSubTaskRequest(parentID uint) dtos.AddSubTaskRequest {
	request := dtos.AddSubTaskRequest{}
	request.Title = faker.Sentence()
	request.Description = faker.Sentence()
	request.Priority = 1
	request.ParentID = parentID
	return request
}

func GenerateUpdateTaskRequest() dtos.UpdateTaskRequest {
	parentId := uint(0)
	request := dtos.UpdateTaskRequest{
		Title:       faker.Sentence(),
		Description: faker.Sentence(),
		ID:          1,
		Priority:    1,
		Checked:     false,
		ParentID:    &parentId,
	}

	return request
}
