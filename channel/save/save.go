package save

import (
	"encoding/json"
	"go-job/dal"
	"go-job/dal/model"
	"go-job/snowflake"
	"time"

	"github.com/sirupsen/logrus"
)

type CreateJob struct {
	Scope              string
	PrepareExecuteTime time.Time
	Ext                map[string]string
}

var c chan CreateJob

func init() {
	c = make(chan CreateJob, 50)

	go func() {
		for {
			select {
			case job := <-c:
				doSave(job)
			}
		}
	}()
}

func Job(scope string, prepareExecuteTime time.Time, ext map[string]string) {
	job := CreateJob{
		Scope:              scope,
		PrepareExecuteTime: prepareExecuteTime,
		Ext:                ext,
	}
	c <- job
}

func doSave(c CreateJob) {
	j := dal.Query.Job
	tag, err := json.Marshal(c.Ext)
	if err != nil {
		logrus.Warnf("json marshal failed, err: %v", err)
		return
	}
	job := &model.Job{
		JobID:              snowflake.NextID(),
		PrepareExecuteTime: c.PrepareExecuteTime,
		Scope:              c.Scope,
		JobStatus:          dal.WAIT,
		ExecuteTime:        nil,
		Tag:                string(tag),
	}
	err = j.Create(job)
	if err != nil {
		logrus.Warnf("create job failed, err: %v", err)
		return
	}
}
