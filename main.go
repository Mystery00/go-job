package main

import (
	"context"
	"fmt"
	"go-job/channel/pull"
	"go-job/channel/worker"
	"go-job/config"
	"go-job/dal"
	"go-job/snowflake"
	"go-job/web/handler"
	"go-job/web/middleware"
	_ "go-job/worker/my-worker"
	_ "go-job/worker/simple"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.WithField("source", "main")

	runPort, exist := os.LookupEnv(config.EnvRunPort)
	if !exist {
		runPort = "9090"
	}
	config.InitLog()
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	dal.InitDataBase(log)
	snowflake.Init()

	router := gin.New()
	router.ForwardedByClientIP = true
	middleware.SetMiddleware(router)
	handler.Handle(router)

	pull.Init()
	worker.Init()

	srv := &http.Server{
		Addr:    fmt.Sprintf(`:%s`, runPort),
		Handler: router,
	}
	go func() {
		log.Infof(`Server is running at :%s`, runPort)
		_ = srv.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infoln("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Infoln("Server exit!")
}
