package simple

import (
	"go-job/dal/model"

	"github.com/sirupsen/logrus"
)

func init() {
	//worker.Register(&simpleWorker{})
}

type simpleWorker struct {
}

func (w *simpleWorker) Execute(job model.Job) {
	logrus.Infof("job scheduled and execute success, job_id: %v", job.JobID)
}
