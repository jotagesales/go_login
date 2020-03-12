package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Server representation
type Server struct{}

// NewServer configure a httpserver
func NewServer(router *gin.Engine, port string) *http.Server {
	timeout := time.Second * 30
	return &http.Server{
		Addr:              port,
		Handler:           router,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
	}
}

// Runserver start the application server
func Runserver(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen port: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("server shutdown: ", err)
	}
}
