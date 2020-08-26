package main

import (
	"fmt"
	"time"
)

func main() {
	// # deklarasi variable dengan shorthand
	// harus berada didalam function
	// ini merupakan cara paling umum yang digunakan developer, sehingga code konsisten
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
