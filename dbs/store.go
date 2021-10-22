package dbs

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var mysqlDB *sql.DB
var mysqlDBErr error

type Config struct {
	USERNAME string
	PASSWORD string
	NETWORK string
	HOST string
	PORT string
	DATABASE string
	CHARSET string
}

func init() {
	var config Config
	vjson := viper.New()
	vjson.SetConfigName("mysql")
	vjson.SetConfigType("json")
	vjson.AddConfigPath("../config")

	if err := vjson.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	vjson.Unmarshal(&config)

	dbDSN := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s", config.USERNAME, config.PASSWORD, config.NETWORK, config.HOST, config.PORT, config.DATABASE, config.CHARSET);
	mysqlDB, mysqlDBErr = sql.Open("mysql", dbDSN);
	if mysqlDBErr != nil {
		log.Println("dbDSN: ", dbDSN);
		panic("数据源配置不正确: " + mysqlDBErr.Error());
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(100);
	// 闲置连接数
	mysqlDB.SetMaxIdleConns(20);
	// 最大连接周期
	mysqlDB.SetConnMaxLifetime(100 * time.Second);

	if mysqlDBErr = mysqlDB.Ping(); nil != mysqlDBErr {
		panic("数据库连接失败: " + mysqlDBErr.Error());
	}
}
