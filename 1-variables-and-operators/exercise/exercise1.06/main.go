package main

import (
	"fmt"
	"time"
)

// # deklarasi multiple variable dari fungsi
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// # deklarasi multiple variable dengan shorthand
	// Debug, LogLevel, startUpTime := false, "info", time.Now()

	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
