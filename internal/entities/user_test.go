package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jhon Doe", "j@j.com", "12345678")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jhon Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Jhon Doe", "j@j.com", "12345678")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12345678"))
	assert.False(t, user.ValidatePassword("123456789"))
	assert.NotEqual(t, user.Password, "12345678")
}
