package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-tutorial/08go-cli/config"
	"github.com/go-tutorial/08go-cli/internal/database"
	"github.com/go-tutorial/08go-cli/internal/logger"
	"github.com/go-tutorial/08go-cli/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the HTTP server",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().String("host", "0.0.0.0", "Server host")
	serverCmd.Flags().Int("port", 8080, "Server port")

	viper.BindPFlag("server.host", serverCmd.Flags().Lookup("host"))
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
}

func runServer(cmd *cobra.Command, args []string) {
	// 初始化配置
	cfg := config.Load()

	// 初始化日志
	logger.Init(cfg.Log)
	log := logger.GetLogger()

	// 初始化数据库
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	log.Info().Msg("Database connection established")

	// 初始化Gin服务器
	gin.SetMode(gin.ReleaseMode)
	if cfg.Log.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// 初始化业务服务器
	appServer := server.New(db, log, cfg)
	appServer.RegisterRoutes(router)

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	// 启动服务器
	go func() {
		log.Info().Str("addr", srv.Addr).Msg("Starting HTTP server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited")
}
