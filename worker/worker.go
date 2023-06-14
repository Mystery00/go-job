package worker

import (
	"go-job/dal/model"
)

type Worker interface {
	Execute(model.Job)
}

var worker *Worker

func Register(w Worker) {
	worker = &w
}

func Execute(job model.Job) {
	if worker == nil {
		panic("worker is nil")
	}
	(*worker).Execute(job)
}
