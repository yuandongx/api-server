package log

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var (
	logger      *log.Logger
	level       int    = 20
	logFileName string = "auto-ping.log"
)

const ERROR = 40
const WARNING = 30
const INFO = 20
const DEBUG = 10

func init() {
	if home, err := os.UserHomeDir(); err == nil {
		now := time.Now()
		y := now.Year()
		m := now.Month()
		d := now.Day()
		h := now.Hour()
		mm := now.Minute()
		s := now.Second()
		lgp := path.Join(home, "auto-ping")
		if _, err := os.Stat(lgp); err != nil {
			os.Mkdir(lgp, os.ModePerm)
		}
		logFile := path.Join(lgp, fmt.Sprintf("%d%d%d_%d%d%d.log", y, m, d, h, mm, s))
		file, err1 := os.Create(logFile)
		if err1 != nil {
			return
		}
		logger = log.New(file, "INFO", log.Lshortfile|log.Ltime)
	}

}

func SetLevel(l int) {
	level = l
}

func Display(v ...interface{}) {
	logger.SetPrefix("NOTSET")
	logger.Println(v...)
}

func Info(v ...interface{}) {
	if level >= INFO {
		logger.SetPrefix("INFO")
		logger.Println(v...)
	}
}

func Debug(v ...interface{}) {
	if level >= DEBUG {
		logger.SetPrefix("DEBUG")
		logger.Println(v...)
	}
}
func Warn(v ...interface{}) {
	if level >= WARNING {
		logger.SetPrefix("WARNING")
		logger.Println(v...)
	}
}

func Err0r(v ...interface{}) {
	if level >= ERROR {
		logger.SetPrefix("ERROR")
		logger.Println(v...)
	}
}
