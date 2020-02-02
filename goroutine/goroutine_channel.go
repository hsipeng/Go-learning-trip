package main

import (
	"time"
	"fmt"
)
var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1 // send 1 to c
}

func main() {
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(5* time.Second)
	<- c
	<- c
}

