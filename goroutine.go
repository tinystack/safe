package safe

import (
	"log"
	"runtime/debug"
)

type Logger interface {
	Print(err, stack interface{})
}

type logger struct{}

func (l *logger) Print(err, stack interface{}) {
	log.Printf("Error in Go routine: %s\nStack: %s\n", err, stack)
}

var defaultLogger Logger = &logger{}

func SetLogger(l Logger) {
	defaultLogger = l
}

func Go(goroutine func()) {
	GoWithRecover(goroutine, func(err interface{}) {
		if defaultLogger != nil {
			defaultLogger.Print(err, debug.Stack())
		}
	})
}

func GoWithRecover(goroutine func(), customRecover func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				customRecover(err)
			}
		}()
		goroutine()
	}()
}
