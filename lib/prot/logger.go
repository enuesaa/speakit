package prot

import (
	"fmt"
	"log"
)

type Logger struct {
	name string
}

func (l *Logger) Log(format string, v ...any) {
	text := fmt.Sprintf(format, v...)
	text = fmt.Sprintf("[%s] %s", l.name, text)

	log.Println(text)
}

func (l *Logger) LogE(err error) {
	l.Log("err: %v", err)
}
