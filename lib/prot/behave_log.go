package prot

import (
	"fmt"
	"log"
	"reflect"
)

func newLogBehavior() LogBehavior {
	return LogBehavior{
		name: "",
	}
}

type LogBehavior struct {
	name string
}

func (l *LogBehavior) Log(format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	if l.name != "" {
		text = fmt.Sprintf("[%s] %s", l.name, text)
	}
	log.Println(text)
}

func (l *LogBehavior) LogE(err error) {
	l.Log("err: %v", err)
}

func (l *LogBehavior) Use(i any) LogBehavior {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return LogBehavior{name: t.Name()}
}
