package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var Logger *logrus.Logger

var currentLogFile string

func Init() {
	f := &prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	}
	f.SetColorScheme(&prefixed.ColorScheme{
		TimestampStyle: "white+h",
	})
	Logger = &logrus.Logger{
		Out:       os.Stderr,
		Level:     logrus.DebugLevel,
		Formatter: f,
	}

	logFileNow := logFileNow()
	if currentLogFile != logFileNow {
		file, err := os.OpenFile(logFileNow, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Logger.Fatal(err)
		}
		writers := []io.Writer{file}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		Logger.SetOutput(fileAndStdoutWriter)
		currentLogFile = logFileNow
	}

}

func logFileNow() string {
	now := time.Now()
	// file := fmt.Sprintf("weibo-notify_%d_%d", now.Year(), now.Month())
	file := fmt.Sprintf("%d_%d_%d_dog.log", now.Year(), now.Month(), now.Day())
	return file
}
