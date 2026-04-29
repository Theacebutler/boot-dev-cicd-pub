package main

import (
	"fmt"

	"time"
)

func GetCurTime() string {
	date := time.Now()
	return fmt.Sprintf("%s", date)
}

func GetCurZone() string {
	zone, _ := time.Now().Zone()
	return fmt.Sprintf("%s", zone)
}

func main() {
	out := GetCurTime()
	fmt.Println(out)
	fmt.Println(GetCurZone())
}
