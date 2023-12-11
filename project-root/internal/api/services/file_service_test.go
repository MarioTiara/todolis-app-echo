package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	repository "github.com/marioTiara/todolistapp/internal/repository/test"
	"github.com/stretchr/testify/assert"
)

func TestNewFileSevice_ReturnsNewInstanceWithDependencies(t *testing.T) {
	ctr := gomock.NewController(t)
	uow := repository.NewMockUnitOfWork(ctr)
	store := storages.NewMockStorage(ctr)

	// Act
	fileService := NewFileSevice(uow, store)

	// Assert
	assert.NotNil(t, fileService)
}
