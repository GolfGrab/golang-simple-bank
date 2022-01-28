package api

import (
	"os"
	"testing"
	"time"

	db "github.com/GolfGrab/journey-to-complete-backend/db/sqlc"
	"github.com/GolfGrab/journey-to-complete-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

// TestMain is the entry point for testing.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
