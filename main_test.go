package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCurTime(t *testing.T) {
	out := GetCurTime()
	if out == "" {
		t.Errorf("out is empty")
	}
}

func TestGetCurZone(t *testing.T) {
	out := GetCurZone()
	if out == "" {
		t.Errorf("out is empty")
	}
	zone, _ := time.Now().Zone()
	expected := fmt.Sprintf("%s", zone)
	if expected != out {
		t.Errorf("out is not equal to expected")
	}
}

func TestGetCurDate(t *testing.T) {
	out := GetCurDate()
	if out == "" {
		t.Errorf("out is empty")
	}
	expected := time.DateOnly
	if expected != out {
		t.Errorf("out is not equal to expected")
	}
}
