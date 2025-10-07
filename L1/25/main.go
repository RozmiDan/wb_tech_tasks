package main

import (
	"fmt"
	"time"
)

func CustomSleep(value time.Duration) {
	timer := time.NewTimer(value)
	defer timer.Stop()
	<-timer.C
}

func main() {
	fmt.Println("Start")
	CustomSleep(2 * time.Second)
	fmt.Println("End after 2 seconds")

}
