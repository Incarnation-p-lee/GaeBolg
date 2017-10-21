package main

import "fmt"
import "time"
import "os/exec"
import "strconv"

func main() {
	data := []int{1, 2, 3, 4}

	for _, v := range data {
		go func(delay int) {
			time.Sleep(time.Duration(delay) * time.Second)

			cmd := exec.Command("./test.bin", strconv.Itoa(delay))
			msg, err := cmd.Output()

			if err != nil {
				fmt.Println("Cmd failed")
			}

			fmt.Printf(string(msg))
		} (v)
	}

	fmt.Println("Waiting")
	time.Sleep(15 * time.Second)
}

