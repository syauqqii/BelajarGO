package main

// Import library
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Initialization clearScreen() function
func clearScreen(){
  // Get OS name and save it into '_os' variable
	_os := runtime.GOOS

  // logic 1: windows
	if _os == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
  } // logic 2: linux
  else if _os == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
  } // logic 3: darwin, etc..
  else {
		fmt.Print()
	}
}

// Main program
func main(){
  // call clearScreen() function
  clearScreen()
}
