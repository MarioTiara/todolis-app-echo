package utils

import (
	"github.com/marioTiara/todolistapp/internal/api/dtos"
	"github.com/marioTiara/todolistapp/internal/api/models"
)

func ConverTaskToQueryModel(task models.Task) dtos.TaskQueryModel {
	queryModel := dtos.TaskQueryModel{}
	queryModel.ID = task.Title
	queryModel.Title = task.Title
	queryModel.Description = task.Description
	queryModel.Priority = uint(task.Priority)
	queryModel.CreatedAt = task.CreatedAt
	queryModel.Checked = task.Checked
	queryModel.IsActive = task.IsActive

	for _, file := range task.Files {
		queryFile := ConverFileToFileQueryModel(file)
		queryModel.Files = append(queryModel.Files, queryFile)
	}

	for _, subTask := range task.Children {
		subtaskQuery := ConvertSubTaskToSubtaskQueryModel(subTask)
		queryModel.SubTasks = append(queryModel.SubTasks, subtaskQuery)
	}

	return queryModel
}

func ConverFileToFileQueryModel(file models.Files) dtos.FileQueryModel {
	queryModel := dtos.FileQueryModel{
		ID:           file.ID,
		FileName:     file.FileName,
		FileSize:     file.FileSize,
		FileURL:      file.FileURL,
		UploadedTime: file.CreatedAt,
	}

	return queryModel
}

func ConvertSubTaskToSubtaskQueryModel(subTask models.Task) dtos.SubtaskQueryModel {
	subTaskQueryModel := dtos.SubtaskQueryModel{
		ID:          subTask.ID,
		Title:       subTask.Title,
		Description: subTask.Description,
		CreatedAt:   subTask.CreatedAt,
		UpdatedAt:   subTask.UpdatedAt,
		Priority:    uint(subTask.Priority),
		Checked:     subTask.Checked,
		IsActive:    subTask.IsActive,
		ParentID:    *subTask.ParentID,
	}

	for _, file := range subTask.Files {
		fileQuery := ConverFileToFileQueryModel(file)
		subTaskQueryModel.Files = append(subTaskQueryModel.Files, fileQuery)
	}

	return subTaskQueryModel
}