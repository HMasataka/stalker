package stalker

import "runtime"

type Frame struct {
	Locations []*Location
}

func (f *Frame) appendLocation(l *Location) *Frame {
	f.Locations = append(f.Locations, l)
	return f
}

func NewFrame() *Frame {
	var fs [3]uintptr

	runtime.Callers(1, fs[:])

	location := NewLocation(fs)

	return &Frame{
		Locations: []*Location{location},
	}
}

func Wrap(original *Frame) *Frame {
	var fs [3]uintptr

	runtime.Callers(1, fs[:])

	return original.appendLocation(NewLocation(fs))
}
