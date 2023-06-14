package pull

import (
	"context"
	"go-job/channel/worker"
	"go-job/dal"
	"time"

	"github.com/sirupsen/logrus"
)

var intChan = make(chan int64)

func Init() {
	go func() {
		intChan <- time.Now().UnixMilli()
		// 每隔 2 分钟，向 intChan 发送一个数据
		for {
			time.Sleep(2 * time.Minute)
			intChan <- time.Now().UnixMilli()
		}
	}()

	go func() {
		ctx := context.Background()
		for {
			select {
			case <-intChan:
				doPull(ctx)
			}
		}
	}()
}

func doPull(ctx context.Context) {
	j := dal.Query.Job
	jobs, err := j.WithContext(ctx).Where(j.PrepareExecuteTime.Lte(time.Now().Add(time.Minute * 2))).Where(j.JobStatus.Eq(dal.WAIT)).Find()
	if err != nil {
		logrus.Warnf("pull job failed, err: %v", err)
		return
	}
	if len(jobs) == 0 {
		logrus.Infof("no job need to pull")
		return
	}
	for _, job := range jobs {
		//更新状态并且获取锁
		update, err := j.WithContext(ctx).Where(j.JobID.Eq(job.JobID)).Where(j.JobStatus.Eq(dal.WAIT)).Update(j.JobStatus, dal.SCHEDULED)
		if err != nil {
			logrus.Warnf("update job status failed, err: %v", err)
			return
		}
		if update.RowsAffected == 0 {
			logrus.Debugf("job has been pulled, job_id: %v", job.JobID)
			continue
		}
		scheduledJob, _ := j.WithContext(ctx).Where(j.JobID.Eq(job.JobID)).First()
		logrus.Debugf("pull job success, job_id: %v", scheduledJob.JobID)
		//将任务放到调度器中
		worker.Job(*scheduledJob)
	}
}
