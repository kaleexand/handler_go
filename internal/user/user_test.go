package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {

	err := NewUser("Anton", -24, 1)
	assert.Error(t, err)
	err = NewUser("Anton", 24, 1)
	assert.NoError(t, err)
}

func TestGetUserById(t *testing.T) {
	err := NewUser("Anton", 24, 1)
	assert.NoError(t, err)
	_, err = GetUserById(100)
	assert.Error(t, err)
	user, err := GetUserById(0)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

}

//func remove(slice []int, s int) []int {
//	return append(slice[:s], slice[s+1:]...)
//}