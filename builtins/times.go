package builtins

import (
	"fmt"
	"os/exec"
	"time"
)

func Times() {
	// recording start time
	begin := time.Now()

	// Run your command here
	cmd := exec.Command("enter shell command ")
	//if error ocurrs message will print out
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running command:", err)
		return
	}

	//record end time
	end := time.Now()

	//calculate the elapsed time (end time - begin time)
	elapsed := end.Sub(begin)

	//set Utime & Stime variables to the elapsed time
	UTime := elapsed
	STime := elapsed

	//print user and system times
	fmt.Printf("User time: %v\n", UTime)
	fmt.Printf("System time: %v\n", STime)
}
