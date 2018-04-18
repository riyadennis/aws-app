package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUpDB(t *testing.T) {
	err := SetUpDB("testDb.db")
	assert.NoError(t, err)
}
