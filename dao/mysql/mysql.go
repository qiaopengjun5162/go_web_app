package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql" // 匿名导入 自动执行 init()
)

var db *sqlx.DB

func Init() (err error) {
	//DSN (Data Source Name) Sprintf根据格式说明符进行格式化，并返回结果字符串。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// 连接到数据库并使用ping进行验证。
	// 也可以使用 MustConnect MustConnect连接到数据库，并在出现错误时恐慌 panic。
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) // 设置数据库的最大打开连接数。
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) // 设置空闲连接池中的最大连接数。
	return
}

func Close() {
	_ = db.Close()
}
