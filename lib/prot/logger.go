package prot

import (
	"fmt"
	"log"
	"reflect"
)

type Logger struct {
	name string
}

func (l *Logger) Log(format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	if l.name != "" {
		text = fmt.Sprintf("[%s] %s", l.name, text)
	}
	log.Println(text)
}

func (l *Logger) LogE(err error) {
	l.Log("err: %v", err)
}

func (l *Logger) Use(i any) Logger {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return Logger{name: t.Name()}
}
