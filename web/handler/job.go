package handler

import (
	"encoding/json"
	"go-job/channel/save"
	"go-job/dal"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Job struct {
	Scope              string            `json:"scope" binding:"required"`
	PrepareExecuteTime int64             `json:"prepareExecuteTime" binding:"required"`
	Ext                map[string]string `json:"ext"`
}

var job = func(context *gin.Context) {
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
	save.Job(job.Scope, t, job.Ext)
}

var jobs = func(context *gin.Context) {
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
		save.Job(job.Scope, t, job.Ext)
	}
}

type History struct {
	Scope string `form:"scope" binding:"required"`
	Size  int    `form:"size" binding:"required"`
}

type JobHistory struct {
	JobId              int64             `json:"jobId"`
	PrepareExecuteTime int64             `json:"prepareExecuteTime"`
	Scope              string            `json:"scope"`
	Success            bool              `json:"success"`
	ExecuteTime        int64             `json:"executeTime"`
	Message            string            `json:"message"`
	Ext                map[string]string `json:"ext"`
}

var history = func(context *gin.Context) {
	var h History
	err := context.ShouldBindQuery(&h)
	if err != nil {
		panic(err)
	}
	j := dal.Query.Job
	list, err := j.WithContext(context.Request.Context()).Where(j.Scope.Eq(h.Scope)).Order(j.PrepareExecuteTime.Desc()).Limit(h.Size).Find()
	if err != nil {
		panic(err)
	}
	r := make([]JobHistory, 0)
	for _, job := range list {
		var ext map[string]string
		err := json.Unmarshal([]byte(job.Tag), &ext)
		if err != nil {
			logrus.Warnf("unmarshal ext failed, job: %v, err: %v", job, err)
		}
		r = append(r, JobHistory{
			JobId:              job.JobID,
			PrepareExecuteTime: job.PrepareExecuteTime.UnixMilli(),
			Scope:              job.Scope,
			Success:            job.JobStatus == dal.SUCCESS,
			ExecuteTime:        job.ExecuteTime.UnixMilli(),
			Message:            job.Message,
			Ext:                ext,
		})
	}
	context.JSON(http.StatusOK, r)
}
