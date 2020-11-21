package secberus

import (
	"fmt"
	"runtime"
	"strings"
)

func handleError(err error) error {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		file, line := details.FileLine(pc)
		spl := strings.Split(file, "/")
		return fmt.Errorf("error [%s:%d]: %s", spl[len(spl)-1], line, err)
	}

	return fmt.Errorf("error [unable to get location]: %s", err)
}
