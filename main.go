package main

import (
	"fmt"

	"time"
)

func lazyProgramer() {
	// this will add lines of code but will not
	// effect anything ;)
}

func GetCurTime() string {
	date := time.Now()
	return date.String()
}

func GetCurZone() string {
	zone, _ := time.Now().Zone()
	return zone
}

func GetCurDate() string {
	date := time.DateOnly
	return date
}

func main() {
	fmt.Println("Time: ", GetCurTime())
	fmt.Println("Zone: ", GetCurZone())
	fmt.Println("Date: ", GetCurDate())
}
