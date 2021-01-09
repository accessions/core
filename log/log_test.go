package log

import (
	"testing"
)

func TestLogs_AddLogs(t *testing.T) {

	ServiceLogs.New(func(options *Options) {
		options.Msg = "eros -- smoke"
		options.Level = LOG_LEVEL_ERROR
	}).Echo()

}
