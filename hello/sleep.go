package main

import (
	"fmt"
	"time"
)

func SleepyTime() {
	fmt.Println("about to sleep")
	time.Sleep(10 * time.Second)
	fmt.Println("wakey wakey")
}