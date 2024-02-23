package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	ticker := time.NewTicker(d)
	<-ticker.C
}

func main() {
	fmt.Println("Hello ")
	mySleep(2 * time.Second)
	fmt.Print("World!")
}
