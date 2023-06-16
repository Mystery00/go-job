package webhook

import (
	"encoding/json"
	"go-job/config"
	"go-job/dal"
	"go-job/dal/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type JobWebHook struct {
	JobID              int64             `json:"jobId"`
	PrepareExecuteTime int64             `json:"prepareExecuteTime"`
	Scope              string            `json:"scope"`
	Success            bool              `json:"success"`
	ExecuteTime        int64             `json:"executeTime"`
	Ext                map[string]string `json:"ext"`
}

var (
	enable = viper.GetBool(config.WebHookEnable)
	url    = viper.GetString(config.WebHookUrl)
	method = viper.GetString(config.WebHookMethod)
)

func NewWebhook(job model.Job) {
	if !enable {
		return
	}
	var ext map[string]string
	err := json.Unmarshal([]byte(job.Tag), &ext)
	if err != nil {
		logrus.Errorf("unmarshal job tag failed, err: %v", err)
		return
	}
	jobWebHook := JobWebHook{
		JobID:              job.JobID,
		PrepareExecuteTime: job.PrepareExecuteTime.UnixMilli(),
		Scope:              job.Scope,
		Success:            job.JobStatus == dal.SUCCESS,
		ExecuteTime:        job.ExecuteTime.UnixMilli(),
		Ext:                ext,
	}
	sendJson(method, url, jobWebHook)
}
