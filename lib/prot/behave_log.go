package prot

import (
	"fmt"
	"log"
	"reflect"
)

type LogBehavior struct {
	name string
}

func (l *LogBehavior) log(format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	if l.name != "" {
		text = fmt.Sprintf("[%s] %s", l.name, text)
	}
	log.Println(text)
}

func (l *LogBehavior) Info(format string, v ...any) {
	l.log(format, v...)
}

func (l *LogBehavior) Head(format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	runes := []rune(text)
	n := min(20, len(runes))
	head := string(runes[:n])
	l.log("%s", head)
}

func (l *LogBehavior) Err(err error) {
	l.log("err: %v", err)
}

func (l *LogBehavior) Use(i any) *LogBehavior {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return &LogBehavior{name: t.Name()}
}
