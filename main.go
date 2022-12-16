package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/samber/lo"

	"errors"
)

func main() {
	err := New("error")
	fmt.Println(err)

	w := Wrap(err)
	fmt.Println(w)
}

func New(msg string) error {
	return newError(errors.New(msg))
}

func Wrap(err error) error {
	return newError(err)
}

func newError(err error) error {
	pt, file, line, _ := runtime.Caller(2)

	funcName := runtime.FuncForPC(pt).Name()

	return stalker{
		err:  err,
		file: file,
		fn:   funcName,
		line: line,
	}
}

type stalker struct {
	err  error
	file string
	fn   string
	line int
}

func (s stalker) Error() string {
	return strings.Join(s.Errors(), "\n")
}

func (s stalker) Unwrap() error {
	return s.err
}

func (s stalker) Errors() []string {
	var res []string
	var tmp stalker
	current := s

	for {
		res = append(res, fmt.Sprintf("%v %v %v", current.file, current.fn, current.line))
		if errors.As(current.err, &tmp) {
			current = tmp
		} else {
			res = append(res, current.err.Error())
			break
		}
	}

	return lo.Reverse(res)
}
