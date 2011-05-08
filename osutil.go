// Utility functions for functionality you would expect to find in the os package, such as running programs and testing if files exist.
package osutil

import (
	"os"
	"exec"
	"log"
)

// Checks if a file exists, returns true if the file exists and can be opened, false otherwise. If the file exists but cannot be opened, it still returns false.
// NOTE: Assumes that any error opening the file for reading means that the file does not exist. This is not always true, but is probably the best behaviour anyway. There is rarely a usage case where you want to check if a file exists, but not if you can open it.
func FileExists(name string) bool {
	_, err := os.Open(name)
	
	if err != nil {
		return false
	}
	return true
}

// Simple way to run most programs. Searches for the program in PATH, and runs the first found program. Args need not contain the program name as the zeroth argument, it is prepended automatically.
func Run(command string, args []string) (proc *exec.Cmd, err os.Error) {
	log.Println(command, args)
	hho := exec.PassThrough
	args = prepend(args, command)
	
	binpath, err := exec.LookPath(command)
	if err != nil {
		return nil, err
	}
	
	proc, err = exec.Run(binpath, args, os.Environ(), "", hho, hho, hho)
	if err != nil {
		log.Println("Error running command ", command, ": ", err)
		return nil, err
	}
	return
}

// Runs a command using Run(), but waits for it to complete before returning.
func WaitRun(command string, args []string) (proc *exec.Cmd, err os.Error) {
	proc, err = Run(command, args)
	if err != nil {
		return proc, err
	}
	proc.Close()
	return
}

// Doesn't really belong here. Oh well...
func prepend(arr []string, prep string) []string {
	arr = append(arr, "")
	for i := len(arr)-1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = prep
	return arr
}