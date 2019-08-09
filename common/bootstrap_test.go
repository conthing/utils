package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func boot(_ interface{}) (needRetry bool, err error) {
	return false, fmt.Errorf("fatal error")
}

func TestBootstrap1(t *testing.T) {
	check := 0
	if Bootstrap(boot, nil, 1000, 500) != nil {
		check = 1
	}
	assert.Equal(t, 1, check, "Bootstrap should return error")
}
