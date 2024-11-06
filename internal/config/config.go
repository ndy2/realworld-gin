package config

import (
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap/zapcore"
)

var MysqlConfig = mysql.Config{
	User:   "root",
	Passwd: "password",
	Net:    "tcp",
	Addr:   "localhost:3306",
	DBName: "realworld",
}

var ZapConfig = struct {
	Level zapcore.Level
}{
	Level: zapcore.InfoLevel,
}
