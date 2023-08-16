package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-unittest/infrastructure/config"
	"go-unittest/pkg/shutdown"
	"log"
	"net/http"
	"time"
)

const (
	idleTimeout = 5 * time.Second // 서버가 유휴 상태로 유지되는 시간
)

func New() *gin.Engine {
	return gin.Default()
}

func Service(r *gin.Engine) {
	cfg := initConfig()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.PORT),
		Handler: r,
	}

	go func(s *http.Server) {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}(server)
	log.Printf("server started: %s", server.Addr)

	// 서버 종료 시그널을 받으면 서버 종료
	shutdown.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), idleTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
		log.Fatalf("server shutdown with timeout: %s", idleTimeout)
	} else {
		log.Printf("server shutdown gracefully")
	}
}

func initConfig() Config {
	return Config{
		PORT: config.GetInt("server.port"),
	}
}
