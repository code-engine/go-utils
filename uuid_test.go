package utils_test

import (
	"testing"

	utils "github.com/code-engine/go-utils"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID(t *testing.T) {
	subject := utils.NewUUID()
	assert.IsType(t, "string", subject)
	assert.Len(t, subject, 36)
}
