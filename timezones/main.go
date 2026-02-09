package main

import (
	"fmt"
	"time"
)

func main() {
	dt := "2025-01-29T23:00:00Z"
	luxLocation, err := time.LoadLocation("Europe/Luxembourg")
	if err != nil {
		fmt.Println("_______________ %#v", err)
	}
	luxVisitDate := dt
	if utcTime, err := time.Parse(time.RFC3339, dt); err == nil {
		luxVisitDate = utcTime.In(luxLocation).Format("2006-01-02T15:04:05Z")
	}
	fmt.Println(luxVisitDate)
}
