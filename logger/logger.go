package logger

import (
	"github.com/op/go-logging"
	"io"
	"os"
)

func Init(logFile string) *logging.Logger {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := io.Writer(file)
	logger := logging.MustGetLogger("changi")
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	be := logging.NewLogBackend(writer, "Changi: ", 0)
	formatter := logging.NewBackendFormatter(be, format)
	logging.SetBackend(be, formatter)
	return logger
}
