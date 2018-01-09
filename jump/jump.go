package main

import "fmt"
import adb "./adb"

func main() {
	fmt.Println("Start Jump Game!")

	for i := 0; i < 10; i++ {
		adb.Screenshot()
		adb.Press(500);
	}
}

