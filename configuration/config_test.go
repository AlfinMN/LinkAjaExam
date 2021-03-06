package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDBWithoutConfig(t *testing.T) {
	_, err, _, _ := Connection()
	assert.Nil(t, err)
}

func TestInitDBWithConfig(t *testing.T) {
	ReadEnv("DbUser", "root")
	ReadEnv("DbPass", "Apinchocs98")
	ReadEnv("DbHost", "localhost")
	ReadEnv("DbPort", "3306")
	ReadEnv("DbName", "linkaja")
	db, err, _, _ := Connection()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestCreateRouter(t *testing.T) {
	router := CreateRouter()
	assert.NotNil(t, router)
}
