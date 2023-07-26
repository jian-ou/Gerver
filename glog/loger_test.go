package gloger_test

import (
	gloger "Gerver/glog"
	"testing"
)

func TestLogger(t *testing.T) {
	l := gloger.NewLogger(gloger.LogDebug)
	l.Debug()
}
