package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Simple GO App to restart COntrol Strip when frozen on Macs with TOuchbar
// Has been tested on 2018 Macbook Pro with Touchbar Running Catalina 64 bit

// excaute function sends command "killall ControlStrip" to terminal and responds with result
func execute() {
	out, err := exec.Command("killall", "ControlStrip").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed, Control Strip has been Restarted")
	output := string(out[:])
	fmt.Println(output)
}

// main func checks for proper OS support takes in result of execute command
func main() {
	if runtime.GOOS == "darwin/amd64" {
		fmt.Println("Can't Execute this on a MacOS machine")
	} else {
		execute()
	}
}
