package worker

import (
	"go-job/dal"
	"go-job/dal/model"
	"go-job/worker"

	"github.com/sirupsen/logrus"
)

var jobChan = make(chan model.Job)

func Init() {
	//同时执行10个任务
	for i := 0; i < 10; i++ {
		go waitJob()
	}
}

func Job(job model.Job) {
	jobChan <- job
}

func waitJob() {
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
		}
	}()
	//更新执行状态为执行中
	_, _ = j.Where(j.JobID.Eq(job.JobID)).Update(j.JobStatus, dal.EXECUTING)
	//执行任务
	worker.Execute(job)
	//更新执行状态为成功
	_, _ = j.Where(j.JobID.Eq(job.JobID)).Update(j.JobStatus, dal.SUCCESS)
}
