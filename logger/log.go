package logger

import (
	logging "github.com/op/go-logging"
	"os"
	"sync"
)

var fmt = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{level:.4s} %{shortfunc:12s} - %{color:reset} %{message}`,
)

/*var fmt = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{level:.4s} [%{longfunc}]  %{message}`,▷
)
*/
/*var fmt = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
*/
var Log = logging.MustGetLogger("gannet")
var once sync.Once

func LogInit() *(logging.Logger) {
	onceLog := func() {
		backend := logging.NewLogBackend(os.Stderr, "", 0)
		backendFormatter := logging.NewBackendFormatter(backend, fmt)
		logging.SetBackend(backendFormatter)
	}
	once.Do(onceLog)
	return Log
}

/*// the other init function in this go source file
func init() {
	onceLog := func() {
		backend := logging.NewLogBackend(os.Stderr, "", 0)
		backendFormatter := logging.NewBackendFormatter(backend, fmt)
		logging.SetBackend(backendFormatter)
		logging.SetLevel(logging.INFO, "")
	}
	once.Do(onceLog)
}
*/
