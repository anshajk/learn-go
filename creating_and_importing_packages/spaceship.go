package main

import "fmt"

func takeOff(countdown int) {
	for i := 0; i < countdown; i++ {
		fmt.Println("TakeOff in T-", i)
	}
}
