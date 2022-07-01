package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	DebugLogLevel int = iota
	InfoLogLevel  int = iota
	WarnLogLevel  int = iota
	ErrorLogLevel int = iota
	OffLogLevel   int = iota
	otput             = log.Ldate | log.Ltime | log.Lmsgprefix
)

var (
	rootLevel *int = new(int)
	loggers        = make(map[string]*logger)
)

type logger struct {
	logLevel    *int
	infoLogger  *log.Logger
	debubLogger *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

func NewLogger(name string) *logger {
	l := logger{}
	l.logLevel = rootLevel
	l.infoLogger = log.New(os.Stdout, fmt.Sprint("INFO ", name, "\t"), otput)
	l.debubLogger = log.New(os.Stdout, fmt.Sprint("DEBUG ", name, "\t"), otput)
	l.errorLogger = log.New(os.Stdout, fmt.Sprint("ERROR ", name, "\t"), otput)
	l.warnLogger = log.New(os.Stdout, fmt.Sprint("WARN ", name, "\t"), otput)

	loggers[name] = &l
	return &l
}

func (l *logger) Debug(v ...interface{}) {
	if *l.logLevel >= DebugLogLevel {
		l.debubLogger.Println(v...)
	}
}

func (l *logger) Info(v ...interface{}) {
	if *l.logLevel >= InfoLogLevel {
		l.infoLogger.Println(v...)
	}
}
func (l *logger) Warn(v ...interface{}) {
	if *l.logLevel >= WarnLogLevel {
		l.warnLogger.Println(v...)
	}
}
func (l *logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)

}
func (l *logger) SetLogLevel(logLevel int) {
	l.logLevel = &logLevel
}

func GetLoger(name string) (logger *logger) {
	if v, ok := loggers[name]; ok {
		return v
	} else {
		logger = NewLogger(name)
	}
	return
}

func SetRootLevel(logLevel int) {
	*rootLevel = logLevel
}

func init() {
	*rootLevel = 1
}
