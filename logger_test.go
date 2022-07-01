package logger

import "testing"

func TestNewLogger(t *testing.T) {

	logerName := "test"

	l := NewLogger(logerName)

	if *l.logLevel != 1 {
		t.FailNow()
	}

	if v, ok := loggers[logerName]; !ok || v != l {
		t.FailNow()
	}
}

func TestSetRootLevel(t *testing.T) {
	logerName := "test"

	loglevel := 2

	SetRootLevel(loglevel)

	if *rootLevel != loglevel {
		t.FailNow()
	}

	l := NewLogger(logerName)

	if l.logLevel != rootLevel || *l.logLevel != loglevel {
		t.FailNow()
	}

}

func TestGetLogger(t *testing.T) {
	logerName := "test"
	logerName2 := "test2"

	l := GetLoger(logerName)

	if v, ok := loggers[logerName]; !ok || v != l {
		t.FailNow()
	}

	l = GetLoger(logerName2)

	if v, ok := loggers[logerName2]; !ok || v != l {
		t.FailNow()
	}

}

func TestSetLogLevel(t *testing.T) {
	logerName3 := "test3"
	logLevel := 3
	l := NewLogger(logerName3)

	if l.logLevel != rootLevel {
		t.Failed()
	}

	if l.SetLogLevel(logLevel); logLevel == *rootLevel {
		t.Failed()
	}

}
