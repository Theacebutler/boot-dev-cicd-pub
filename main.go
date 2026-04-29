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

func GetCurDate() string {
	date := time.DateOnly
	return date
}

func main() {
	fmt.Println("Time: ", GetCurTime())
	fmt.Println("Zone: ", GetCurZone())
	fmt.Println("Date: ", GetCurDate())
}
