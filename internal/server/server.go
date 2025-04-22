package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SRV struct {
	server *http.Server
}

func (s *SRV) Run(port string, handler *gin.Engine) error {
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s *SRV) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
