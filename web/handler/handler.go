package handler

import (
	"go-job/channel/save"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Job struct {
	PrepareExecuteTime int64             `json:"prepareExecuteTime" binding:"required"`
	Ext                map[string]string `json:"ext"`
}

func Handle(engine *gin.Engine) {
	//注册路由
	engine.GET(apiPath+"/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	engine.PUT(externalPath+"/job", func(context *gin.Context) {
		var job Job
		err := context.ShouldBindJSON(&job)
		if err != nil {
			panic(err)
		}
		t := time.UnixMilli(job.PrepareExecuteTime)
		if t.Before(time.Now()) {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "prepareExecuteTime must be in the future",
			})
			return
		}
		save.Job(t, job.Ext)
	})
	engine.PUT(externalPath+"/jobs", func(context *gin.Context) {
		var list []Job
		err := context.ShouldBindJSON(&list)
		if err != nil {
			panic(err)
		}
		now := time.Now()
		for _, job := range list {
			t := time.UnixMilli(job.PrepareExecuteTime)
			if t.Before(now) {
				logrus.Warnf("prepareExecuteTime must be in the future, job: %v", job)
				continue
			}
			save.Job(t, job.Ext)
		}
	})
}
