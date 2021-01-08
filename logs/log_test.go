package logs

import (
	"testing"
)

func TestLogs_AddLogs(t *testing.T) {

	ServiceLogs.ActLogs(func(options *Options) {
		options.Msg = "eros -- smoke"
		options.Level = LOG_LEVEL_ERROR
	}).Echo()

}
