package utils_test

import (
	"katou-megumi/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReusableStringVariables(t *testing.T) {
	assert.Equal(t, "Sedang memproses....", utils.WAIT_MESSAGE)
	assert.Equal(t, "Format yang dimasukkan Salah!", utils.WRONG_FORMAT)
	assert.Equal(t, "Error!", utils.ERROR_MESSAGE)
	assert.Equal(t, "Berhasil!", utils.SUCCESS_MESSAGE)
	assert.Equal(t, "user", utils.GEMINI_ROLE)
	assert.Equal(t, "gemini-2.5-pro", utils.GEMINI_MODEL)
	assert.Equal(t, "2006-01-02", utils.TIME_FORMAT)
}
