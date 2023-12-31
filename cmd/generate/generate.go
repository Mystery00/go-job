package main

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "dal/query",
		Mode:    gen.WithoutContext,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		//FieldWithIndexTag: true,
		//if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		WithUnitTest: true,
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessray or it will panic
	db, _ := gorm.Open(mysql.Open("root:12345678@tcp(127.0.0.1:3306)/db_go_job?charset=utf8mb4&parseTime=True&loc=Asia%2fShanghai"))
	g.UseDB(db)

	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"int": func(columnType gorm.ColumnType) (dataType string) {
			t, _ := columnType.ColumnType()
			if strings.HasPrefix(t, "int(6)") {
				return "int8"
			}
			return "int64"
		},
		// bool mapping
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			t, _ := columnType.ColumnType()
			if strings.HasPrefix(t, "tinyint(1)") {
				return "bool"
			}
			return "int8"
		},
	}
	g.WithDataTypeMap(dataMap)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	g.ApplyBasic(
		g.GenerateModelAs("job", "Job"),
	)

	// execute the action of code generation
	g.Execute()
}
