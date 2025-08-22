package stalker

import (
	"runtime"
	"strconv"
)

type Location struct {
	function string
	file     string
	line     int
}

func NewLocation(fs [3]uintptr) *Location {
	rfs := runtime.CallersFrames(fs[:])

	var frames []runtime.Frame
	for {
		fr, ok := rfs.Next()
		if !ok {
			break
		}

		frames = append(frames, fr)
	}

	return &Location{
		function: frames[len(frames)-1].Function,
		file:     frames[len(frames)-1].File,
		line:     frames[len(frames)-1].Line,
	}
}

func (l *Location) toStr() string {
	return l.function + " at " + l.file + ":" + strconv.Itoa(l.line)
}

func (l *Location) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.toStr() + `"`), nil
}
