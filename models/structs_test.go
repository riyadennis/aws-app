package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestNewUser(t *testing.T) {
	user := NewUser("test@test.com", "test", "password", "test user")
	assert.IsType(t, &User{}, user)
}

func TestNewFollower(t *testing.T) {
	follower := NewFollower(123, 123)
	assert.IsType(t, &Follower{}, follower)
}
func TestNewComment(t *testing.T) {
	comment := NewComment(123, 123, "Good photo", time.Now())
	assert.IsType(t, &Comment{}, comment)
}
func TestNewPhoto(t *testing.T) {
	photo := NewPhoto(123, 1, "photo.jpg", "My Photo", time.Now())
	assert.IsType(t,&Photo{}, photo)
}