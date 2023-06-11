package config

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
)

type OrmLogger struct {
	*logrus.Logger
	logger.Config
}

func NewOrmLogger(log *logrus.Logger) logger.Interface {
	return &OrmLogger{
		Logger: log,
		Config: logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Warn,
		},
	}
}

// LogMode log mode
func (l *OrmLogger) LogMode(level logger.LogLevel) logger.Interface {
	ormLogger := *l
	switch level {
	case logger.Error:
		ormLogger.SetLevel(logrus.ErrorLevel)
	case logger.Warn:
		ormLogger.SetLevel(logrus.WarnLevel)
	case logger.Info:
		ormLogger.SetLevel(logrus.InfoLevel)
	}
	return &ormLogger
}

// Info print info
func (l *OrmLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	log := checkLogger(l, ctx)
	if l.LogLevel >= logger.Info {
		//去掉换行符
		m := strings.Replace(msg, "\n", "", -1)
		log.Infof(m, data...)
	}
}

// Warn print warn messages
func (l *OrmLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	log := checkLogger(l, ctx)
	if l.LogLevel >= logger.Warn {
		//去掉换行符
		m := strings.Replace(msg, "\n", "", -1)
		log.Warnf(m, data...)
	}
}

// Error print error messages
func (l *OrmLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	log := checkLogger(l, ctx)
	if l.LogLevel >= logger.Error {
		//去掉换行符
		m := strings.Replace(msg, "\n", "", -1)
		log.Errorf(m, data...)
	}
}

// Trace print sql message
func (l *OrmLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	log := checkLogger(l, ctx)
	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil && err != gorm.ErrRecordNotFound:
		if strings.Contains(sql, `t_snowflake`) {
			return
		}
		log.Errorf(`%s %s`, utils.FileWithLineNum(), err)
		if rows == -1 {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": "-",
			}).Error(sql)
		} else {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": rows,
			}).Error(sql)
		}
	default:
		if strings.Contains(sql, `t_snowflake`) {
			return
		}
		if rows == -1 {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": "-",
			}).Debug(sql)
		} else {
			log.WithFields(logrus.Fields{
				"cost": elapsed,
				"rows": rows,
			}).Debug(sql)
		}
	}
}

func checkLogger(l *OrmLogger, ctx context.Context) *logrus.Entry {
	return l.WithContext(ctx).WithField("source", "gorm")
}
