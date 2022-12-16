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
	pt, file, line, _ := runtime.Caller(1)

	funcName := runtime.FuncForPC(pt).Name()

	return es{
		err:  errors.New(msg),
		file: file,
		fn:   funcName,
		line: line,
	}
}

func Wrap(err error) error {
	pt, file, line, _ := runtime.Caller(1)

	funcName := runtime.FuncForPC(pt).Name()

	return es{
		err:  err,
		file: file,
		fn:   funcName,
		line: line,
	}
}

type es struct {
	err  error
	file string
	fn   string
	line int
}

func (e es) Error() string {
	return strings.Join(e.Errors(), "\n")
}

func (e es) Unwrap() error {
	return e.err
}

func (e es) Errors() []string {
	var res []string
	var tmp es
	current := e

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
