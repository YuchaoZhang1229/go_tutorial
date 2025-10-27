package server

import (
	"github.com/go-tutorial/08go-cli/config"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	db     *redis.Client
	logger zerolog.Logger
	config *config.Config
}

func New(db *redis.Client, logger zerolog.Logger, cfg *config.Config) *Server {
	return &Server{
		db:     db,
		logger: logger,
		config: cfg,
	}
}

func (s *Server) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", s.healthCheck)
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Service is healthy",
	})
}
