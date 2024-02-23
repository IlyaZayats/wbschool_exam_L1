package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	fmt.Println("Hello ")
	mySleep(2 * time.Second)
	fmt.Print("World!")
}
