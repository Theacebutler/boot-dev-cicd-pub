package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetApiKey(t *testing.T) {
	os.Setenv("APY_KEY", "test")
	key, err := GetApiKey()
	assert.Nil(t, err)
	require.NotEmpty(t, key)
}
