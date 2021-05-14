package log

import (
	"errors"
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

const BT = 512
var (
	ERROR_A = errors.New("error A")
	ERROR_B = errors.New("error B")
)

type logInterceptor interface {
	errorTest(int) error
}
type A struct {}
func (A) errorTest(b int) error {
	return [BT]error{
		ERROR_A,
		ERROR_B,
	}[b]
}

func smoke () logInterceptor {
	return A{}
}
func test() error {
 return smoke().errorTest(1)
}
