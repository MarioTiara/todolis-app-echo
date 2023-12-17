package database

import (
	"testing"

	"github.com/marioTiara/todolistapp/config"
	"github.com/stretchr/testify/assert"
)

func TestNewPostGressDB(t *testing.T) {
	//Arrange
	configuration, _ := config.LoadConfig("../../../config")
	//Act
	postgress := NewPostGressDB(configuration)

	assert.NotNil(t, postgress)
}

func TestGetDB(t *testing.T) {
	//Arrange
	configuration, _ := config.LoadConfig("../../../config")
	//Act
	postgress := NewPostGressDB(configuration)
	db := postgress.GetDB()

	assert.NotNil(t, postgress)
	assert.NotNil(t, db)
}
