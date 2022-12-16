package main

import (
	"fmt"
	"runtime"

	"errors"
)

func main() {
	w := Wrap(e())
	var e es
	fmt.Println(errors.As(w, &e))
	fmt.Println(e)

	fmt.Println(w)
}

func e() error {
	return New("error")
}

func New(msg string) error {
	_, file, line, _ := runtime.Caller(1)

	return es{
		err:  errors.New(msg),
		file: file,
		line: line,
	}
}

func Wrap(err error) error {
	_, file, line, _ := runtime.Caller(1)

	return es{
		err:  err,
		file: file,
		line: line,
	}
}

type es struct {
	err  error
	file string
	line int
}

func (e es) Error() string {
	var res string
	res += fmt.Sprintf("%v %v %v\n", e.err, e.file, e.line)
	return res
}

func (e es) Unwrap() error {
	return e.err
}
