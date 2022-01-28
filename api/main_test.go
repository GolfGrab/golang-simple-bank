package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMain is the entry point for testing.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
