package dal

import (
	"go-job/config"
	"go-job/dal/query"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 数据库配置
var (
	userName = viper.GetString(config.DbUserName)
	password = viper.GetString(config.DbPassword)
	ip       = viper.GetString(config.DbHost)
	port     = viper.GetString(config.DbPort)
	dbName   = viper.GetString(config.DbName)
)

var Query *query.Query

func InitDataBase(log *logrus.Entry) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True&loc=Asia%2fShanghai"}, "")
	client, err := gorm.Open(mysql.Open(path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: config.NewOrmLogger(logrus.StandardLogger()),
	})
	if err != nil {
		log.Fatal("Open database failed", err)
	}
	sqlDB, err := client.DB()
	if err != nil {
		log.Fatal("Get database failed", err)
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Info("Database connect success")
	Query = query.Use(client)
}
