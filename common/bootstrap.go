package common

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// BootFunc Boot函数的原型
type BootFunc func(param interface{}) error

// BootstrapError 带是否需要retry信息的error，由BootFunc函数返回
type BootstrapError struct {
	DontRetry bool
	ErrMsg    string
}

// NewBootstrapError 创建BootstrapError
func NewBootstrapError(dontRetry bool, err string) error {
	e := BootstrapError{DontRetry: dontRetry, ErrMsg: err}
	return e
}

// Error 标准接口
func (e BootstrapError) Error() string {
	return e.ErrMsg
}

// Bootstrap 调用入口
func Bootstrap(params interface{}, boot BootFunc, timeout int) {
	deps := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go retry(params, boot, timeout, &wg, deps)
	go func(ch chan error) {
		for {
			select {
			case e, ok := <-ch:
				if ok {
					fmt.Fprintf(os.Stderr, "Boot error:%v\n", e)
				} else {
					return
				}
			}
		}
	}(deps)

	wg.Wait()
}

func retry(params interface{}, boot BootFunc, timeout int, wait *sync.WaitGroup, ch chan error) {
	until := time.Now().Add(time.Millisecond * time.Duration(timeout))
	for time.Now().Before(until) {
		var err error

		err = boot(params)
		if err != nil {
			ch <- err
			if bootstrapErr, ok := err.(BootstrapError); ok {
				if bootstrapErr.DontRetry {
					// no need to retry, break
					break
				}
			}
		} else {
			// success, break
			break
		}
		time.Sleep(time.Second * time.Duration(1))
	}
	close(ch)
	wait.Done()

	return
}
