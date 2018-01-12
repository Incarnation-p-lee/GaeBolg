package adb

import "os"
import "fmt"
import "os/exec"
import "strconv"

var adbBinary     string = "adb"
var cmdScreenshot string = "screencap"
var cmdShell      string = "shell"
var cmdLs         string = "ls"
var cmdPull       string = "pull"
var cmdInput      string = "input"
var cmdSwipe      string = "swipe"
var pngCount      int = 0

func pngCountGet() int {
	return pngCount
}

func pngCountAdvance() {
	pngCount++
}

func Screenshot() string {
	cmd := adbBinary
	args := make([]string, 0)

	tmpPng := "/sdcard/tmp.png"
	outPng := "./screenshot/screen" + "-" + strconv.Itoa(pngCountGet()) + ".png"
	cmdArgs := "-p"

	/* adb shell screencap -p > xxx.png */
	args = append(args, cmdShell, cmdScreenshot, cmdArgs, tmpPng);

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args = args[:0]

	/* adb pull xxx.png yyy.png */
	args = append(args, cmdPull, tmpPng, outPng)

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pngCountAdvance()
	return outPng
}

func Press(ms int) {
	if ms < 0 {
		fmt.Println("Cannot press negative ms!")
		return
	}

	fmt.Printf("Press %d ms.\n", ms)

	loc := "1"
	cmd := adbBinary
	args := make([]string, 0)

	/* adb shell input swipe 1 1 1 1 1000 */
	args = append(args, cmdShell, cmdInput, cmdSwipe)
	args = append(args, loc, loc, loc, loc, strconv.Itoa(ms))

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

