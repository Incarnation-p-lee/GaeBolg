package adb

import "os"
import "fmt"
import "os/exec"
import "strconv"

var adbBinary string = "adb"
var cmdScreenshot string = "screencap"
var cmdShell string = "shell"
var cmdLs string = "ls"
var cmdPull string = "pull"
var pngCount int = 0

func pngCountGet() int {
	return pngCount
}

func pngCountAdvance() {
	pngCount++
}

func AdbScreenshot() string {
	cmd := adbBinary
	args := make([]string, 0)

	tmpPng := "/sdcard/tmp.png"
	outPng := "./screenshot/screen" + "-" + strconv.Itoa(pngCountGet()) + ".png"
	cmdArgs := "-p"

	/* adb shell screencap -p > xxx.png */
	args = append(args, cmdShell);
	args = append(args, cmdScreenshot)
	args = append(args, cmdArgs)
	args = append(args, tmpPng)

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	args = args[:0]

	/* adb pull xxx.png yyy.png */
	args = append(args, cmdPull)
	args = append(args, tmpPng)
	args = append(args, outPng)

	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	pngCountAdvance()
	return outPng
}

