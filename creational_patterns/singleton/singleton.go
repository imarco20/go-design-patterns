package main

// Purpose:
//		- The Singleton pattern restricts the instantiation of a class
//		  to a single instance and provides global access
//		- Allows for lazy initialization of the class
//
// Scenarios:
//		- Situations where you want to ensure there is only one instance
//		  of a class: logging, configuration, etc...

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	loglevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

var logger *MyLogger

var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("creating the logger instance")
		logger = &MyLogger{}
	})

	fmt.Println("returning logger instance")
	return logger
}
