package elog_test

import (
	"errors"
	"testing"

	"github.com/wjhdec/echo-ext/pkg/elog"
)

func TestLog(t *testing.T) {
	log := elog.NewConsoleLogger()
	log.Debug("info")
}

func testError1() error {
	return testError2()
}

func testError2() error {
	return errors.New("error2")
}

func TestErrLog(t *testing.T) {
	log := elog.NewConsoleLogger()
	err := testError1()
	log.Error(err)
}
