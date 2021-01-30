package log

import (
	"github.com/natefinch/lumberjack"
	"log"
)
// Logs 切割文件
func Logx ()  {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs.log",
		MaxSize:    500,// M
		MaxAge:     30,// days
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   true,// disabled by default
	})
}
