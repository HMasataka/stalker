package main

import (
	"fmt"
	"runtime"

	"errors"
)

func main() {
	fmt.Println(Wrap(e()))
}

func e() error {
	return New("error")
}

func New(msg string) error {
	_, file, line, _ := runtime.Caller(1)

	return &errs{
		single{
			err:  errors.New(msg),
			file: file,
			line: line,
		},
	}
}

func Wrap(err error) error {
	_, file, line, _ := runtime.Caller(1)

	return &errs{
		single{
			err:  err,
			file: file,
			line: line,
		},
	}
}

type single struct {
	err  error
	file string
	line int
}

type errs []single

func (es errs) Error() string {
	var res string
	for _, e := range es {
		res += fmt.Sprintf("%v %v %v\n", e.err, e.file, e.line)
	}
	return res
}

func (es errs) Unwrap() error {
	if len(es) == 0 {
		return nil
	}
	return es[1:]
}
