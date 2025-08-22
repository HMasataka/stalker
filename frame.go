package stalker

import "runtime"

type Frame struct {
	Locations []*Location
}

func (f *Frame) appendLocation(l *Location) *Frame {
	f.Locations = append(f.Locations, l)
	return f
}

func NewFrame(opts ...Option) *Frame {
	location := newLocation(opts...)

	return &Frame{
		Locations: []*Location{location},
	}
}

func Wrap(original *Frame, opts ...Option) *Frame {
	location := newLocation(opts...)

	return original.appendLocation(location)
}

func newLocation(opts ...Option) *Location {
	var fs [3]uintptr

	op := &Options{
		skipFrames: 1,
	}

	for _, opt := range opts {
		opt(op)
	}

	runtime.Callers(op.skipFrames, fs[:])

	return NewLocation(fs)
}
