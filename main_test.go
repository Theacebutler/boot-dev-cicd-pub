package main

import (
	"testing"
)

func TestGetCurTime(t *testing.T) {
	out := GetCurTime()
	if out == "" {
		t.Errorf("out is empty")
	}
}
