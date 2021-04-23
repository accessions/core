package log

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	LOG_LEVEL_INFO    = byte(1)
	LOG_LEVEL_WARNING = byte(2)
	LOG_LEVEL_ERROR   = byte(3)
)



var ServiceLogs *Logs

func init() {
	ServiceLogs = &Logs{}
}

type Logs struct {
	option *Options
}

type Options struct {
	Level  byte
	Module string
	Action string
	Msg    interface{}
	Date   time.Time
}

type newOption func(options *Options) error

func (log *Logs) New (opts ...newOption) *Logs {
	opt := &Options{
		Level:  LOG_LEVEL_INFO,
		Module: "",
		Action: "",
		Msg:    "",
		Date:   time.Now(),
	}
	for _, o := range opts {
		o(opt)
	}
	log.option = opt
	return log
}

func SetLevel(level int) newOption {
	return func(options *Options) error {
		options.Level = level
		return nil
	}
}

func (log *Logs) Echo() {

	if str, ok := json.Marshal(log.option); ok == nil {
		fmt.Println("Fixme Eros -------- " + string(str))
	}
}

func (log *Logs) Back() []byte {
	back, _ := json.Marshal(log.option)
	return back
}

func (log *Logs) Push(inter ...interface{}) bool {
	return true
}
func (log *Logs) Add(inter ...interface{}) bool {
	return true
}
