//go:build !windows
// +build !windows

package datura

import (
	"fmt"
	"syscall"

	"github.com/theapemachine/wrkspc/errnie"
)

const (
	// By increasing the minimum open files limit on Linux
	// based systems, we enable more network connections.
	minOpenFilesLimit = 65535
)

func Raise() error {
	var err errnie.Error
	var rLimit syscall.Rlimit

	errnie.Logs(fmt.Sprintf("ulimit max is %d", rLimit.Max)).With(errnie.INFO)

	if err = errnie.Handles(
		syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit),
	); err.Type != errnie.NIL {
		return err
	}

	if rLimit.Max < minOpenFilesLimit {
		return errnie.NewError(
			fmt.Errorf("max ulimit of %d reached", rLimit.Max),
		)
	}

	rLimit.Cur = minOpenFilesLimit

	if err = errnie.Handles(syscall.Setrlimit(
		syscall.RLIMIT_NOFILE, &rLimit,
	)); err.Type == errnie.NIL {
		errnie.Logs(fmt.Sprintf("new ulimit set of %d", rLimit.Cur))
	}

	errnie.Logs(fmt.Sprintf("ulimit set to %d", rLimit.Cur)).With(errnie.INFO)
	return err
}
