package worker

import (
	"go-job/channel/webhook"
	"go-job/config"
	"go-job/dal"
	"go-job/dal/model"
	"go-job/worker"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var jobChan = make(chan model.Job)

func Init() {
	num := viper.GetInt(config.JobWorkerNum)
	//同时执行任务
	for i := 0; i < num; i++ {
		go waitJob()
	}
}

func Job(job model.Job) {
	jobChan <- job
}

func waitJob() {
	jobWaitTime := viper.GetString(config.JobWaitTime)
	duration, err := time.ParseDuration(jobWaitTime)
	if err != nil {
		logrus.Warnf("parse job wait time failed, err: %v", err)
	} else {
		time.Sleep(duration)
	}
	for {
		select {
		case job := <-jobChan:
			doExecute(job)
		}
	}
}

func doExecute(job model.Job) {
	j := dal.Query.Job
	defer func() {
		if err := recover(); err != nil {
			logrus.Warnf("execute job failed, err: %v", err)
			//更新执行状态为失败
			_, _ = j.Where(j.JobID.Eq(job.JobID)).Update(j.JobStatus, dal.FAILED)
			scheduledJob, _ := j.Where(j.JobID.Eq(job.JobID)).First()
			webhook.NewWebhook(*scheduledJob)
		}
	}()
	//更新执行状态为执行中
	now := time.Now()
	_, _ = j.Where(j.JobID.Eq(job.JobID)).Updates(model.Job{JobStatus: dal.EXECUTING, ExecuteTime: &now})
	//执行任务
	worker.Execute(job)
	//更新执行状态为成功
	_, _ = j.Where(j.JobID.Eq(job.JobID)).Update(j.JobStatus, dal.SUCCESS)
	scheduledJob, _ := j.Where(j.JobID.Eq(job.JobID)).First()
	webhook.NewWebhook(*scheduledJob)
}
