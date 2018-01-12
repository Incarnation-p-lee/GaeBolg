package main

import "fmt"
import "time"
import adb "./adb"
import screen "./screenshot"

var distanceTable [1024]int = [1024]int{
	0: 0,
}

var distanePerMs float64 = 0.619

func distanceToPressTime(distance int) int {
	t := distanceTable[distance]

	if t != 0 {
		return t
	}

	ms := float64(distance) / distanePerMs 

	return int(ms)
}

func main() {
	count := 0
	fmt.Println("Start Jump Game!")

	for {
		count++
		fmt.Printf("--- %d --------------------\n", count)

		path := adb.Screenshot()
		distance := screen.Distance(path)
		adb.Press(distanceToPressTime(distance))

		time.Sleep(400 * time.Millisecond)
	}
}

