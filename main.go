package main

import (
	"fmt"
	"runtime"

	"errors"
)

func main() {
	fmt.Println(e())
}

func e() error {
	return New("error")
}

func New(msg string) error {
	_, file, line, _ := runtime.Caller(1)

	return &Error{
		err:  errors.New(msg),
		file: file,
		line: line,
	}
}

type Error struct {
	err  error
	file string
	line int
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v %v %v", e.err, e.file, e.line)
}
