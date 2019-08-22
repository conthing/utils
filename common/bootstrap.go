package common

import (
	"fmt"
	"os"
	"time"
)

// BootFunc Boot函数的簽名
type BootFunc func(params interface{}) (needRetry bool, err error)

// Bootstrap 调用入口，timeout和retry单位ms
func Bootstrap(boot BootFunc, params interface{}, timeout int, interval int) (err error) {
	until := time.Now().Add(time.Millisecond * time.Duration(timeout))
	for time.Now().Before(until) {
		needRetry, e := boot(params)
		if e != nil {
			fmt.Fprintf(os.Stderr, "BOOTERR: %v\n", e)
			err = e
			if !needRetry {
				// no need to retry, break
				break
			}
		} else {
			// success, break
			err = nil
			break
		}
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
	return
}
