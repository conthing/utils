package common

import (
	"fmt"
	"os"
	"time"
)

// BootFunc Boot函数的原型
type BootFunc func(params interface{}) (needRetry bool, err error)

// Bootstrap 调用入口，timeout和retry单位ms
func Bootstrap(boot BootFunc, params interface{}, timeout int, interval int) (err error) {
	until := time.Now().Add(time.Millisecond * time.Duration(timeout))
	for time.Now().Before(until) {
		needRetry, err := boot(params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "BOOTERR: %v\n", err)
			if !needRetry {
				// no need to retry, break
				break
			}
		} else {
			// success, break
			break
		}
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
	return
}
