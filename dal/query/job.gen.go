// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"go-job/dal/model"
)

func newJob(db *gorm.DB, opts ...gen.DOOption) job {
	_job := job{}

	_job.jobDo.UseDB(db, opts...)
	_job.jobDo.UseModel(&model.Job{})

	tableName := _job.jobDo.TableName()
	_job.ALL = field.NewAsterisk(tableName)
	_job.JobID = field.NewInt64(tableName, "job_id")
	_job.PrepareExecuteTime = field.NewTime(tableName, "prepare_execute_time")
	_job.Scope = field.NewString(tableName, "scope")
	_job.JobStatus = field.NewInt8(tableName, "job_status")
	_job.ExecuteTime = field.NewTime(tableName, "execute_time")
	_job.Message = field.NewString(tableName, "message")
	_job.Tag = field.NewString(tableName, "tag")

	_job.fillFieldMap()

	return _job
}

type job struct {
	jobDo

	ALL                field.Asterisk
	JobID              field.Int64  // 任务id
	PrepareExecuteTime field.Time   // 预期执行时间
	Scope              field.String // 领域
	JobStatus          field.Int8   // 任务状态
	ExecuteTime        field.Time   // 最终执行时间
	Message            field.String // 任务执行详情
	Tag                field.String // 任务标签，用在回调

	fieldMap map[string]field.Expr
}

func (j job) Table(newTableName string) *job {
	j.jobDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j job) As(alias string) *job {
	j.jobDo.DO = *(j.jobDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *job) updateTableName(table string) *job {
	j.ALL = field.NewAsterisk(table)
	j.JobID = field.NewInt64(table, "job_id")
	j.PrepareExecuteTime = field.NewTime(table, "prepare_execute_time")
	j.Scope = field.NewString(table, "scope")
	j.JobStatus = field.NewInt8(table, "job_status")
	j.ExecuteTime = field.NewTime(table, "execute_time")
	j.Message = field.NewString(table, "message")
	j.Tag = field.NewString(table, "tag")

	j.fillFieldMap()

	return j
}

func (j *job) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *job) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 7)
	j.fieldMap["job_id"] = j.JobID
	j.fieldMap["prepare_execute_time"] = j.PrepareExecuteTime
	j.fieldMap["scope"] = j.Scope
	j.fieldMap["job_status"] = j.JobStatus
	j.fieldMap["execute_time"] = j.ExecuteTime
	j.fieldMap["message"] = j.Message
	j.fieldMap["tag"] = j.Tag
}

func (j job) clone(db *gorm.DB) job {
	j.jobDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j job) replaceDB(db *gorm.DB) job {
	j.jobDo.ReplaceDB(db)
	return j
}

type jobDo struct{ gen.DO }

func (j jobDo) Debug() *jobDo {
	return j.withDO(j.DO.Debug())
}

func (j jobDo) WithContext(ctx context.Context) *jobDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobDo) ReadDB() *jobDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobDo) WriteDB() *jobDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobDo) Session(config *gorm.Session) *jobDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobDo) Clauses(conds ...clause.Expression) *jobDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobDo) Returning(value interface{}, columns ...string) *jobDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobDo) Not(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobDo) Or(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobDo) Select(conds ...field.Expr) *jobDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobDo) Where(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *jobDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j jobDo) Order(conds ...field.Expr) *jobDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobDo) Distinct(cols ...field.Expr) *jobDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobDo) Omit(cols ...field.Expr) *jobDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobDo) Join(table schema.Tabler, on ...field.Expr) *jobDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobDo) LeftJoin(table schema.Tabler, on ...field.Expr) *jobDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobDo) RightJoin(table schema.Tabler, on ...field.Expr) *jobDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobDo) Group(cols ...field.Expr) *jobDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobDo) Having(conds ...gen.Condition) *jobDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobDo) Limit(limit int) *jobDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobDo) Offset(offset int) *jobDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *jobDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobDo) Unscoped() *jobDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobDo) Create(values ...*model.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobDo) CreateInBatches(values []*model.Job, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobDo) Save(values ...*model.Job) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobDo) First() (*model.Job, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Take() (*model.Job, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Last() (*model.Job, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) Find() ([]*model.Job, error) {
	result, err := j.DO.Find()
	return result.([]*model.Job), err
}

func (j jobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Job, err error) {
	buf := make([]*model.Job, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobDo) FindInBatches(result *[]*model.Job, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobDo) Attrs(attrs ...field.AssignExpr) *jobDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobDo) Assign(attrs ...field.AssignExpr) *jobDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobDo) Joins(fields ...field.RelationField) *jobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobDo) Preload(fields ...field.RelationField) *jobDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobDo) FirstOrInit() (*model.Job, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) FirstOrCreate() (*model.Job, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Job), nil
	}
}

func (j jobDo) FindByPage(offset int, limit int) (result []*model.Job, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobDo) Delete(models ...*model.Job) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobDo) withDO(do gen.Dao) *jobDo {
	j.DO = *do.(*gen.DO)
	return j
}
