package guardpanic

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
)

type WorkerFunc func()
type ErrCallback func(err error)

type guardpanic struct {
	error
	restartCnt int
	worker WorkerFunc
	errCb ErrCallback
}

func (gp *guardpanic) Recover() {
	if r := recover(); r != nil {
		gp.error = errors.New(fmt.Sprintf("%s, %s", r, string(debug.Stack())))
		if gp.errCb != nil {
			gp.errCb(gp)
		}else {
			fmt.Fprint(os.Stderr, gp.error)
		}

		if gp.restartCnt > 0 {
			gp.restartCnt--
			go Run(gp.worker, gp.restartCnt, gp.errCb)
		}
	}
}

func (gp *guardpanic) Error() string {
	if gp.error != nil {
		return gp.error.Error()
	}
	return ""
}

func Run(worker WorkerFunc, cnt int, cb ErrCallback) {
	gp := &guardpanic{worker:worker, restartCnt:cnt, errCb:cb}
	defer gp.Recover()

	worker()
	return
}