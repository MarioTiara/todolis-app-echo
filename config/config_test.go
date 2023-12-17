package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_SuccessLoad(t *testing.T) {
	//Arrange
	config, err := LoadConfig(".")

	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.DBDriver)
	assert.NotEmpty(t, config.DbSource)
	assert.NotEmpty(t, config.ServerAddress)
}

// func TestLoadConfig_FailedLoad(t *testing.T) {
// 	//Arrange
// 	config, err := LoadConfig("../")
// 	assert.Error(t, err)
// 	assert.Empty(t, config.DBDriver)
// 	assert.Empty(t, config.DbSource)
// 	assert.Empty(t, config.ServerAddress)
// }
