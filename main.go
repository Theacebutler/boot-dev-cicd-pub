package main

import (
	"fmt"

	"time"
)

func GetCurTime() string {
	date := time.Now()
	return fmt.Sprintf("%s", date)
}

func main() {
	out := GetCurTime()
	fmt.Println(out)
}
