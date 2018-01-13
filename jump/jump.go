package main

import "fmt"
import "time"
import adb "./adb"
import screen "./screenshot"

func main() {
	count := 0
	fmt.Println("Start Jump Game!")

	// screen.Distance("screenshot/screen-5.png")

	for {
		count++
		fmt.Printf("--- %d --------------------\n", count)

		path := adb.Screenshot()
		distance := screen.Distance(path)
		adb.Press(distanceToPressTime(distance))

		time.Sleep(1500 * time.Millisecond)
		fmt.Println()
	}
}
